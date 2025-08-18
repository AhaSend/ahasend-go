package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// HeaderWebhookID is the header key for the webhook ID
	HeaderWebhookID = "webhook-id"
	// HeaderWebhookTimestamp is the header key for the webhook timestamp
	HeaderWebhookTimestamp = "webhook-timestamp"
	// HeaderWebhookSignature is the header key for the webhook signature
	HeaderWebhookSignature = "webhook-signature"

	// DefaultTolerance is the default time tolerance for webhook verification (5 minutes)
	DefaultTolerance = 5 * time.Minute

	// SignatureVersion is the version prefix for webhook signatures
	SignatureVersion = "v1"
)

var (
	// ErrMissingHeaders is returned when required webhook headers are missing
	ErrMissingHeaders = errors.New("missing required webhook headers")
	// ErrInvalidSignature is returned when the webhook signature is invalid
	ErrInvalidSignature = errors.New("invalid webhook signature")
	// ErrExpiredTimestamp is returned when the webhook timestamp is too old
	ErrExpiredTimestamp = errors.New("webhook timestamp expired")
	// ErrInvalidPayload is returned when the webhook payload cannot be parsed
	ErrInvalidPayload = errors.New("invalid webhook payload")
	// ErrUnknownEventType is returned when the webhook event type is not recognized
	ErrUnknownEventType = errors.New("unknown webhook event type")
)

// WebhookVerifier verifies webhook signatures according to the Standard Webhooks specification
type WebhookVerifier struct {
	secret    []byte
	tolerance time.Duration
}

// NewWebhookVerifier creates a new webhook verifier with the given secret
func NewWebhookVerifier(secret string) (*WebhookVerifier, error) {
	return &WebhookVerifier{
		secret:    []byte(secret),
		tolerance: DefaultTolerance,
	}, nil
}

// SetTolerance sets the time tolerance for webhook verification
func (v *WebhookVerifier) SetTolerance(tolerance time.Duration) {
	v.tolerance = tolerance
}

// Verify verifies a webhook payload with the given headers
func (v *WebhookVerifier) Verify(payload []byte, headers http.Header) error {
	// Get required headers
	msgID := headers.Get(HeaderWebhookID)
	msgTimestamp := headers.Get(HeaderWebhookTimestamp)
	msgSignature := headers.Get(HeaderWebhookSignature)

	if msgID == "" || msgTimestamp == "" || msgSignature == "" {
		return ErrMissingHeaders
	}

	// Parse and validate timestamp
	timestamp, err := strconv.ParseInt(msgTimestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %w", err)
	}

	webhookTime := time.Unix(timestamp, 0)
	if time.Since(webhookTime) > v.tolerance {
		return ErrExpiredTimestamp
	}

	// Build the signed content
	signedContent := fmt.Sprintf("%s.%s.%s", msgID, msgTimestamp, string(payload))

	// Calculate expected signature
	expectedSignature := v.sign([]byte(signedContent))

	// Parse signatures from header (space-delimited list)
	signatures := strings.Split(msgSignature, " ")

	// Check if any signature matches
	for _, sig := range signatures {
		// Remove version prefix if present
		sig = strings.TrimPrefix(sig, SignatureVersion+"=")

		// Constant-time comparison
		if hmac.Equal([]byte(sig), []byte(expectedSignature)) {
			return nil
		}
	}

	return ErrInvalidSignature
}

// VerifyRequest verifies a webhook from an HTTP request
func (v *WebhookVerifier) VerifyRequest(r *http.Request) error {
	// Read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	// Verify the webhook
	return v.Verify(body, r.Header)
}

// Parse verifies and parses a webhook payload into the appropriate event type
func (v *WebhookVerifier) Parse(payload []byte, headers http.Header) (WebhookEvent, error) {
	// First verify the webhook
	if err := v.Verify(payload, headers); err != nil {
		return nil, err
	}

	// Parse the base event to determine the type
	var baseEvent baseWebhookEvent
	if err := json.Unmarshal(payload, &baseEvent); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
	}

	// Parse into specific event type based on the type field
	switch baseEvent.Type {
	// Message events
	case "message.reception":
		var event MessageReceptionEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	case "message.delivered":
		var event MessageDeliveredEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	case "message.transient_error":
		var event MessageTransientErrorEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	case "message.failed":
		var event MessageFailedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	case "message.bounced":
		var event MessageBouncedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	case "message.suppressed":
		var event MessageSuppressedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	case "message.opened":
		var event MessageOpenedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	case "message.clicked":
		var event MessageClickedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	// Suppression events
	case "suppression.created":
		var event SuppressionCreatedEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	// Domain events
	case "domain.dns_error":
		var event DomainDNSErrorEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	// Route events
	case "route.message":
		var event RouteMessageEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidPayload, err)
		}
		return &event, nil

	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownEventType, baseEvent.Type)
	}
}

// ParseRequest verifies and parses a webhook from an HTTP request
func (v *WebhookVerifier) ParseRequest(r *http.Request) (WebhookEvent, error) {
	// Read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	// Parse the webhook
	return v.Parse(body, r.Header)
}

// sign calculates the HMAC-SHA256 signature for the given data
func (v *WebhookVerifier) sign(data []byte) string {
	h := hmac.New(sha256.New, v.secret)
	h.Write(data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// WebhookEvent is the interface for all webhook events
type WebhookEvent interface {
	GetType() string
	GetTimestamp() time.Time
}

// baseWebhookEvent is the base structure for all webhook events
type baseWebhookEvent struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *baseWebhookEvent) GetType() string {
	return e.Type
}

func (e *baseWebhookEvent) GetTimestamp() time.Time {
	return e.Timestamp
}
