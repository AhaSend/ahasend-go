package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test webhook secret for testing
const testWebhookSecret = "MfKQ9r8GKYqrTwjUPD8ILPZIo2LaLaSw"

func TestNewWebhookVerifier(t *testing.T) {
	t.Run("valid secret", func(t *testing.T) {
		verifier, err := NewWebhookVerifier(testWebhookSecret)
		assert.NoError(t, err)
		assert.NotNil(t, verifier)
	})

	t.Run("valid secret with aha-whsec prefix", func(t *testing.T) {
		verifier, err := NewWebhookVerifier("aha-whsec-" + testWebhookSecret)
		assert.NoError(t, err)
		assert.NotNil(t, verifier)
	})
}

func TestWebhookVerification(t *testing.T) {
	verifier, err := NewWebhookVerifier(testWebhookSecret)
	require.NoError(t, err)

	t.Run("valid webhook", func(t *testing.T) {
		payload := `{"type":"message.delivered","timestamp":"2024-05-06T09:50:16.687031577Z","data":{"account_id":"4cdd7bdd-294e-4762-892f-83d40abf5a87","event":"on_delivered","from":"sender@example.com","recipient":"recipient@example.com","subject":"Test Email","message_id_header":"<message-id-12345@localhost>","id":"407926766d2711f09b30960002cafe7c"}}`

		headers := http.Header{}
		headers.Set("webhook-id", "msg_2Ej8Gx5VCOPKUhbMr9Zw7qvxPtt")
		headers.Set("webhook-timestamp", fmt.Sprintf("%d", time.Now().Unix()))
		headers.Set("webhook-signature", generateSignature(t, verifier, headers.Get("webhook-id"), headers.Get("webhook-timestamp"), payload))

		err := verifier.Verify([]byte(payload), headers)
		assert.NoError(t, err)
	})

	t.Run("invalid signature", func(t *testing.T) {
		payload := `{"type":"message.delivered","timestamp":"2024-05-06T09:50:16.687031577Z","data":{}}`

		headers := http.Header{}
		headers.Set("webhook-id", "msg_2Ej8Gx5VCOPKUhbMr9Zw7qvxPtt")
		headers.Set("webhook-timestamp", fmt.Sprintf("%d", time.Now().Unix()))
		headers.Set("webhook-signature", "v1=invalid_signature")

		err := verifier.Verify([]byte(payload), headers)
		assert.ErrorIs(t, err, ErrInvalidSignature)
	})

	t.Run("expired timestamp", func(t *testing.T) {
		payload := `{"type":"message.delivered","timestamp":"2024-05-06T09:50:16.687031577Z","data":{}}`

		// Use timestamp from 10 minutes ago
		oldTimestamp := fmt.Sprintf("%d", time.Now().Add(-10*time.Minute).Unix())

		headers := http.Header{}
		headers.Set("webhook-id", "msg_2Ej8Gx5VCOPKUhbMr9Zw7qvxPtt")
		headers.Set("webhook-timestamp", oldTimestamp)
		headers.Set("webhook-signature", generateSignature(t, verifier, headers.Get("webhook-id"), oldTimestamp, payload))

		err := verifier.Verify([]byte(payload), headers)
		assert.ErrorIs(t, err, ErrExpiredTimestamp)
	})

	t.Run("missing headers", func(t *testing.T) {
		payload := `{"type":"message.delivered","timestamp":"2024-05-06T09:50:16.687031577Z","data":{}}`

		testCases := []struct {
			name    string
			headers http.Header
		}{
			{
				name: "missing webhook-id",
				headers: http.Header{
					"webhook-timestamp": []string{fmt.Sprintf("%d", time.Now().Unix())},
					"webhook-signature": []string{"v1=test"},
				},
			},
			{
				name: "missing webhook-timestamp",
				headers: http.Header{
					"webhook-id":        []string{"msg_123"},
					"webhook-signature": []string{"v1=test"},
				},
			},
			{
				name: "missing webhook-signature",
				headers: http.Header{
					"webhook-id":        []string{"msg_123"},
					"webhook-timestamp": []string{fmt.Sprintf("%d", time.Now().Unix())},
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				err := verifier.Verify([]byte(payload), tc.headers)
				assert.ErrorIs(t, err, ErrMissingHeaders)
			})
		}
	})

	t.Run("multiple signatures with one valid", func(t *testing.T) {
		payload := `{"type":"message.delivered","timestamp":"2024-05-06T09:50:16.687031577Z","data":{}}`

		headers := http.Header{}
		headers.Set("webhook-id", "msg_2Ej8Gx5VCOPKUhbMr9Zw7qvxPtt")
		headers.Set("webhook-timestamp", fmt.Sprintf("%d", time.Now().Unix()))

		validSig := generateSignature(t, verifier, headers.Get("webhook-id"), headers.Get("webhook-timestamp"), payload)
		// Include multiple signatures (one valid, others invalid)
		headers.Set("webhook-signature", fmt.Sprintf("v1=invalid_sig1 %s v1=invalid_sig2", validSig))

		err := verifier.Verify([]byte(payload), headers)
		assert.NoError(t, err)
	})

	t.Run("custom tolerance", func(t *testing.T) {
		// Set tolerance to 1 minute
		verifier.SetTolerance(1 * time.Minute)

		payload := `{"type":"message.delivered","timestamp":"2024-05-06T09:50:16.687031577Z","data":{}}`

		// Use timestamp from 2 minutes ago (should fail with 1 minute tolerance)
		oldTimestamp := fmt.Sprintf("%d", time.Now().Add(-2*time.Minute).Unix())

		headers := http.Header{}
		headers.Set("webhook-id", "msg_2Ej8Gx5VCOPKUhbMr9Zw7qvxPtt")
		headers.Set("webhook-timestamp", oldTimestamp)
		headers.Set("webhook-signature", generateSignature(t, verifier, headers.Get("webhook-id"), oldTimestamp, payload))

		err := verifier.Verify([]byte(payload), headers)
		assert.ErrorIs(t, err, ErrExpiredTimestamp)

		// Reset tolerance
		verifier.SetTolerance(5 * time.Minute)
	})
}

func TestWebhookParsing(t *testing.T) {
	verifier, err := NewWebhookVerifier(testWebhookSecret)
	require.NoError(t, err)

	createValidHeaders := func(payload string) http.Header {
		headers := http.Header{}
		headers.Set("webhook-id", "msg_2Ej8Gx5VCOPKUhbMr9Zw7qvxPtt")
		headers.Set("webhook-timestamp", fmt.Sprintf("%d", time.Now().Unix()))
		headers.Set("webhook-signature", generateSignature(t, verifier, headers.Get("webhook-id"), headers.Get("webhook-timestamp"), payload))
		return headers
	}

	t.Run("parse message.delivered event", func(t *testing.T) {
		payload := `{
			"type": "message.delivered",
			"webhook_id": "abe11757-2886-4b55-96f1-0e0afc95795a",
			"timestamp": "2024-05-06T09:50:16.687031577Z",
			"data": {
				"account_id": "4cdd7bdd-294e-4762-892f-83d40abf5a87",
				"event": "on_delivered",
				"from": "sender@example.com",
				"recipient": "recipient@example.com",
				"subject": "Welcome to our service",
				"message_id_header": "<message-id-12345@localhost>",
				"id": "407926766d2711f09b30960002cafe7c"
			}
		}`

		event, err := verifier.Parse([]byte(payload), createValidHeaders(payload))
		require.NoError(t, err)
		require.NotNil(t, event)

		deliveredEvent, ok := event.(*MessageDeliveredEvent)
		require.True(t, ok)
		assert.Equal(t, "message.delivered", deliveredEvent.Type)
		assert.Equal(t, "4cdd7bdd-294e-4762-892f-83d40abf5a87", deliveredEvent.Data.AccountID)
		assert.Equal(t, "sender@example.com", deliveredEvent.Data.From)
		assert.Equal(t, "recipient@example.com", deliveredEvent.Data.Recipient)
		assert.Equal(t, "Welcome to our service", deliveredEvent.Data.Subject)
	})

	t.Run("parse message.opened event with extra fields", func(t *testing.T) {
		payload := `{
			"type": "message.opened",
			"timestamp": "2024-05-06T10:15:16.687031577Z",
			"data": {
				"account_id": "4cdd7bdd-294e-4762-892f-83d40abf5a87",
				"event": "on_opened",
				"from": "sender@example.com",
				"recipient": "recipient@example.com",
				"subject": "Welcome to our service",
				"message_id_header": "<message-id-12345@localhost>",
				"id": "407926766d2711f09b30960002cafe7c",
				"user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
				"ip": "192.168.1.100",
				"is_bot": "false"
			}
		}`

		event, err := verifier.Parse([]byte(payload), createValidHeaders(payload))
		require.NoError(t, err)
		require.NotNil(t, event)

		openedEvent, ok := event.(*MessageOpenedEvent)
		require.True(t, ok)
		assert.Equal(t, "message.opened", openedEvent.Type)
		assert.NotNil(t, openedEvent.Data.UserAgent)
		assert.Equal(t, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36", *openedEvent.Data.UserAgent)
		assert.NotNil(t, openedEvent.Data.IP)
		assert.Equal(t, "192.168.1.100", *openedEvent.Data.IP)
		assert.NotNil(t, openedEvent.Data.IsBot)
		assert.Equal(t, "false", *openedEvent.Data.IsBot)
	})

	t.Run("parse message.clicked event", func(t *testing.T) {
		payload := `{
			"type": "message.clicked",
			"timestamp": "2024-05-06T10:20:16.687031577Z",
			"data": {
				"account_id": "4cdd7bdd-294e-4762-892f-83d40abf5a87",
				"event": "on_clicked",
				"from": "sender@example.com",
				"recipient": "recipient@example.com",
				"subject": "Welcome to our service",
				"message_id_header": "<message-id-12345@localhost>",
				"id": "407926766d2711f09b30960002cafe7c",
				"url": "https://example.com/link",
				"user_agent": "Mozilla/5.0",
				"ip": "192.168.1.100",
				"is_bot": false
			}
		}`

		event, err := verifier.Parse([]byte(payload), createValidHeaders(payload))
		require.NoError(t, err)
		require.NotNil(t, event)

		clickedEvent, ok := event.(*MessageClickedEvent)
		require.True(t, ok)
		assert.Equal(t, "message.clicked", clickedEvent.Type)
		assert.Equal(t, "https://example.com/link", clickedEvent.Data.URL)
		assert.Equal(t, "Mozilla/5.0", clickedEvent.Data.UserAgent)
		assert.Equal(t, false, clickedEvent.Data.IsBot)
	})

	t.Run("parse suppression.created event", func(t *testing.T) {
		payload := `{
			"type": "suppression.created",
			"timestamp": "2024-05-06T12:57:06.451529527Z",
			"data": {
				"account_id": "4cdd7bdd-294e-4762-892f-83d40abf5a87",
				"recipient": "bounced@example.com",
				"created_at": "2024-05-06T12:57:06.451529617Z",
				"expires_at": "2024-06-05T12:57:06.451529617Z",
				"reason": "Too many hard bounces",
				"sending_domain": "your-domain.com"
			}
		}`

		event, err := verifier.Parse([]byte(payload), createValidHeaders(payload))
		require.NoError(t, err)
		require.NotNil(t, event)

		suppressionEvent, ok := event.(*SuppressionCreatedEvent)
		require.True(t, ok)
		assert.Equal(t, "suppression.created", suppressionEvent.Type)
		assert.Equal(t, "bounced@example.com", suppressionEvent.Data.Recipient)
		assert.Equal(t, "Too many hard bounces", suppressionEvent.Data.Reason)
		assert.Equal(t, "your-domain.com", suppressionEvent.Data.SendingDomain)
	})

	t.Run("parse domain.dns_error event", func(t *testing.T) {
		payload := `{
			"type": "domain.dns_error",
			"webhook_id": "abe11757-2886-4b55-96f1-0e0afc95795a",
			"timestamp": "2024-05-06T12:59:46.404433272Z",
			"data": {
				"domain": "example.com",
				"account_id": "4cdd7bdd-294e-4762-892f-83d40abf5a87",
				"spf_valid": false,
				"dkim_valid": false,
				"dmarc_valid": false,
				"dns_last_checked_at": "2024-05-06T12:59:46.404433312Z"
			}
		}`

		event, err := verifier.Parse([]byte(payload), createValidHeaders(payload))
		require.NoError(t, err)
		require.NotNil(t, event)

		domainEvent, ok := event.(*DomainDNSErrorEvent)
		require.True(t, ok)
		assert.Equal(t, "domain.dns_error", domainEvent.Type)
		assert.Equal(t, "example.com", domainEvent.Data.Domain)
		assert.Equal(t, false, domainEvent.Data.SPFValid)
		assert.Equal(t, false, domainEvent.Data.DKIMValid)
		assert.Equal(t, false, domainEvent.Data.DMARCValid)
	})

	t.Run("parse route.message event", func(t *testing.T) {
		payload := `{
			"type": "route.message",
			"route_id": "550e8400-e29b-41d4-a716-446655440000",
			"timestamp": "2024-05-06T13:15:46.404433272Z",
			"data": {
				"id": "route-msg-12345",
				"from": "customer@gmail.com",
				"reply_to": "customer@gmail.com",
				"to": "support@yourdomain.com",
				"subject": "Help with my account",
				"message_id": "<unique-message-id@gmail.com>",
				"size": 2048,
				"spam_score": 0.1,
				"bounce": false,
				"cc": "",
				"date": "Mon, 06 May 2024 13:15:46 +0000",
				"in_reply_to": "",
				"references": "",
				"auto_submitted": "",
				"html_body": "<p>I need help with my account settings.</p>",
				"plain_body": "I need help with my account settings.",
				"reply_from_plain_body": "I need help with my account settings.",
				"attachments": [
					{
						"filename": "document.pdf",
						"content_type": "application/pdf",
						"data": "base64encodeddata=="
					}
				],
				"headers": {
					"X-Mailer": "Gmail",
					"X-Priority": "3"
				}
			}
		}`

		event, err := verifier.Parse([]byte(payload), createValidHeaders(payload))
		require.NoError(t, err)
		require.NotNil(t, event)

		routeEvent, ok := event.(*RouteMessageEvent)
		require.True(t, ok)
		assert.Equal(t, "route.message", routeEvent.Type)
		assert.Equal(t, "route-msg-12345", routeEvent.Data.ID)
		assert.Equal(t, "customer@gmail.com", routeEvent.Data.From)
		assert.Equal(t, "support@yourdomain.com", routeEvent.Data.To)
		assert.Equal(t, "Help with my account", routeEvent.Data.Subject)
		assert.Equal(t, 1, len(routeEvent.Data.Attachments))
		assert.Equal(t, "document.pdf", routeEvent.Data.Attachments[0].Filename)
		assert.Equal(t, "Gmail", routeEvent.Data.Headers["X-Mailer"])
	})

	t.Run("parse unknown event type", func(t *testing.T) {
		payload := `{
			"type": "unknown.event",
			"timestamp": "2024-05-06T09:50:16.687031577Z",
			"data": {}
		}`

		event, err := verifier.Parse([]byte(payload), createValidHeaders(payload))
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownEventType)
		assert.Nil(t, event)
	})

	t.Run("parse invalid JSON", func(t *testing.T) {
		payload := `{invalid json}`

		event, err := verifier.Parse([]byte(payload), createValidHeaders(payload))
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidPayload)
		assert.Nil(t, event)
	})

	t.Run("verify fails before parsing", func(t *testing.T) {
		payload := `{"type": "message.delivered", "timestamp": "2024-05-06T09:50:16.687031577Z", "data": {}}`

		headers := http.Header{}
		headers.Set("webhook-id", "msg_123")
		headers.Set("webhook-timestamp", fmt.Sprintf("%d", time.Now().Unix()))
		headers.Set("webhook-signature", "invalid_signature")

		event, err := verifier.Parse([]byte(payload), headers)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidSignature)
		assert.Nil(t, event)
	})
}

func TestWebhookHelperFunctions(t *testing.T) {
	t.Run("IsMessageEvent", func(t *testing.T) {
		assert.True(t, IsMessageEvent(&MessageDeliveredEvent{}))
		assert.True(t, IsMessageEvent(&MessageOpenedEvent{}))
		assert.True(t, IsMessageEvent(&MessageClickedEvent{}))
		assert.False(t, IsMessageEvent(&SuppressionCreatedEvent{}))
		assert.False(t, IsMessageEvent(&DomainDNSErrorEvent{}))
	})

	t.Run("IsSuppressionEvent", func(t *testing.T) {
		assert.True(t, IsSuppressionEvent(&SuppressionCreatedEvent{}))
		assert.False(t, IsSuppressionEvent(&MessageDeliveredEvent{}))
	})

	t.Run("IsDomainEvent", func(t *testing.T) {
		assert.True(t, IsDomainEvent(&DomainDNSErrorEvent{}))
		assert.False(t, IsDomainEvent(&MessageDeliveredEvent{}))
	})

	t.Run("IsRouteEvent", func(t *testing.T) {
		assert.True(t, IsRouteEvent(&RouteMessageEvent{}))
		assert.False(t, IsRouteEvent(&MessageDeliveredEvent{}))
	})

	t.Run("GetMessageEventData", func(t *testing.T) {
		messageData := MessageEventData{
			AccountID: "test-account",
			From:      "test@example.com",
			Recipient: "recipient@example.com",
			Subject:   "Test Subject",
		}

		deliveredEvent := &MessageDeliveredEvent{Data: messageData}
		data := GetMessageEventData(deliveredEvent)
		require.NotNil(t, data)
		assert.Equal(t, "test-account", data.AccountID)
		assert.Equal(t, "test@example.com", data.From)

		// Test with non-message event
		suppressionEvent := &SuppressionCreatedEvent{}
		data = GetMessageEventData(suppressionEvent)
		assert.Nil(t, data)

		// Test with clicked event (which has different data structure)
		clickedEvent := &MessageClickedEvent{}
		data = GetMessageEventData(clickedEvent)
		assert.Nil(t, data)
	})
}

// Helper function to generate a valid signature for testing
func generateSignature(t *testing.T, verifier *WebhookVerifier, msgID, timestamp, payload string) string {
	t.Helper()

	// Use the secret as-is (no base64 decoding) - matches the new webhook implementation
	secret := []byte(testWebhookSecret)
	signedContent := fmt.Sprintf("%s.%s.%s", msgID, timestamp, payload)

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(signedContent))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("v1=%s", signature)
}
