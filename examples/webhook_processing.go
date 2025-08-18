package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AhaSend/ahasend-go/webhooks"
)

func main() {
	// Get webhook secret from environment variable
	webhookSecret := os.Getenv("AHASEND_WEBHOOK_SECRET")
	if webhookSecret == "" {
		log.Fatal("AHASEND_WEBHOOK_SECRET environment variable is required")
	}

	// Create webhook verifier
	verifier, err := webhooks.NewWebhookVerifier(webhookSecret)
	if err != nil {
		log.Fatalf("Failed to create webhook verifier: %v", err)
	}

	// Create HTTP handler for webhook endpoint
	http.HandleFunc("/webhooks/ahasend", webhookHandler(verifier))

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting webhook server on port %s", port)
	log.Printf("Webhook endpoint: http://localhost:%s/webhooks/ahasend", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// webhookHandler creates an HTTP handler function for processing AhaSend webhooks
func webhookHandler(verifier *webhooks.WebhookVerifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only accept POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse and verify the webhook
		event, err := verifier.ParseRequest(r)
		if err != nil {
			log.Printf("Failed to verify webhook: %v", err)

			// Determine appropriate error code
			switch err {
			case webhooks.ErrMissingHeaders, webhooks.ErrInvalidSignature, webhooks.ErrExpiredTimestamp:
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			case webhooks.ErrInvalidPayload, webhooks.ErrUnknownEventType:
				http.Error(w, "Bad Request", http.StatusBadRequest)
			default:
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		// Process the event based on its type
		if err := processWebhookEvent(event); err != nil {
			log.Printf("Failed to process webhook event: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Return success response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}

// processWebhookEvent handles different types of webhook events
func processWebhookEvent(event webhooks.WebhookEvent) error {
	// Log the event type and timestamp
	log.Printf("Received webhook event: %s at %s", event.GetType(), event.GetTimestamp())

	// Handle different event types
	switch e := event.(type) {
	case *webhooks.MessageReceptionEvent:
		return handleMessageReception(e)

	case *webhooks.MessageDeliveredEvent:
		return handleMessageDelivered(e)

	case *webhooks.MessageTransientErrorEvent:
		return handleMessageTransientError(e)

	case *webhooks.MessageFailedEvent:
		return handleMessageFailed(e)

	case *webhooks.MessageBouncedEvent:
		return handleMessageBounced(e)

	case *webhooks.MessageSuppressedEvent:
		return handleMessageSuppressed(e)

	case *webhooks.MessageOpenedEvent:
		return handleMessageOpened(e)

	case *webhooks.MessageClickedEvent:
		return handleMessageClicked(e)

	case *webhooks.SuppressionCreatedEvent:
		return handleSuppressionCreated(e)

	case *webhooks.DomainDNSErrorEvent:
		return handleDomainDNSError(e)

	case *webhooks.RouteMessageEvent:
		return handleRouteMessage(e)

	default:
		log.Printf("Unhandled event type: %s", event.GetType())
		return nil
	}
}

// Message event handlers

func handleMessageReception(event *webhooks.MessageReceptionEvent) error {
	log.Printf("📬 Email queued for delivery:")
	log.Printf("  From: %s", event.Data.From)
	log.Printf("  To: %s", event.Data.Recipient)
	log.Printf("  Subject: %s", event.Data.Subject)
	log.Printf("  Message ID: %s", event.Data.ID)

	// TODO: Add your business logic here
	// For example, update message status in database

	return nil
}

func handleMessageDelivered(event *webhooks.MessageDeliveredEvent) error {
	log.Printf("✅ Email delivered successfully:")
	log.Printf("  To: %s", event.Data.Recipient)
	log.Printf("  Subject: %s", event.Data.Subject)
	log.Printf("  Message ID: %s", event.Data.ID)

	// TODO: Add your business logic here
	// For example, mark message as delivered in your system

	return nil
}

func handleMessageTransientError(event *webhooks.MessageTransientErrorEvent) error {
	log.Printf("⚠️ Temporary delivery issue (will retry):")
	log.Printf("  To: %s", event.Data.Recipient)
	log.Printf("  Subject: %s", event.Data.Subject)

	// TODO: Add your business logic here
	// For example, log the temporary failure for monitoring

	return nil
}

func handleMessageFailed(event *webhooks.MessageFailedEvent) error {
	log.Printf("❌ Email delivery failed permanently:")
	log.Printf("  To: %s", event.Data.Recipient)
	log.Printf("  Subject: %s", event.Data.Subject)
	log.Printf("  Message ID: %s", event.Data.ID)

	// TODO: Add your business logic here
	// For example, notify the sender or update customer records

	return nil
}

func handleMessageBounced(event *webhooks.MessageBouncedEvent) error {
	log.Printf("🔙 Email bounced:")
	log.Printf("  To: %s", event.Data.Recipient)
	log.Printf("  Subject: %s", event.Data.Subject)

	// TODO: Add your business logic here
	// For example, mark email address as invalid

	return nil
}

func handleMessageSuppressed(event *webhooks.MessageSuppressedEvent) error {
	log.Printf("🚫 Email suppressed (not sent):")
	log.Printf("  To: %s", event.Data.Recipient)
	log.Printf("  Subject: %s", event.Data.Subject)

	// TODO: Add your business logic here
	// For example, log why the email was suppressed

	return nil
}

func handleMessageOpened(event *webhooks.MessageOpenedEvent) error {
	log.Printf("👀 Email opened:")
	log.Printf("  By: %s", event.Data.Recipient)
	log.Printf("  Subject: %s", event.Data.Subject)

	if event.Data.UserAgent != nil {
		log.Printf("  User Agent: %s", *event.Data.UserAgent)
	}
	if event.Data.IP != nil {
		log.Printf("  IP: %s", *event.Data.IP)
	}
	if event.Data.IsBot != nil {
		log.Printf("  Is Bot: %s", *event.Data.IsBot)
	}

	// TODO: Add your business logic here
	// For example, track engagement metrics

	return nil
}

func handleMessageClicked(event *webhooks.MessageClickedEvent) error {
	log.Printf("🔗 Link clicked in email:")
	log.Printf("  By: %s", event.Data.Recipient)
	log.Printf("  URL: %s", event.Data.URL)
	log.Printf("  Subject: %s", event.Data.Subject)
	log.Printf("  User Agent: %s", event.Data.UserAgent)
	log.Printf("  IP: %s", event.Data.IP)
	log.Printf("  Is Bot: %v", event.Data.IsBot)

	// TODO: Add your business logic here
	// For example, track click-through rates

	return nil
}

// Suppression event handler

func handleSuppressionCreated(event *webhooks.SuppressionCreatedEvent) error {
	log.Printf("⛔ Email address suppressed:")
	log.Printf("  Email: %s", event.Data.Recipient)
	log.Printf("  Reason: %s", event.Data.Reason)
	log.Printf("  Domain: %s", event.Data.SendingDomain)
	log.Printf("  Expires: %s", event.Data.ExpiresAt)

	// TODO: Add your business logic here
	// For example, update your email list to avoid sending to this address

	return nil
}

// Domain event handler

func handleDomainDNSError(event *webhooks.DomainDNSErrorEvent) error {
	log.Printf("🔴 Domain DNS configuration error:")
	log.Printf("  Domain: %s", event.Data.Domain)
	log.Printf("  SPF Valid: %v", event.Data.SPFValid)
	log.Printf("  DKIM Valid: %v", event.Data.DKIMValid)
	log.Printf("  DMARC Valid: %v", event.Data.DMARCValid)
	log.Printf("  Last Checked: %s", event.Data.DNSLastCheckedAt)

	// TODO: Add your business logic here
	// For example, alert your operations team

	return nil
}

// Route event handler (for inbound emails)

func handleRouteMessage(event *webhooks.RouteMessageEvent) error {
	log.Printf("📨 Inbound email received:")
	log.Printf("  From: %s", event.Data.From)
	log.Printf("  To: %s", event.Data.To)
	log.Printf("  Subject: %s", event.Data.Subject)
	log.Printf("  Size: %d bytes", event.Data.Size)

	if event.Data.SpamScore != nil {
		log.Printf("  Spam Score: %.2f", *event.Data.SpamScore)
	}

	log.Printf("  Bounce: %v", event.Data.Bounce)

	// Process attachments if any
	if len(event.Data.Attachments) > 0 {
		log.Printf("  Attachments: %d", len(event.Data.Attachments))
		for i, attachment := range event.Data.Attachments {
			log.Printf("    [%d] %s (%s)", i+1, attachment.Filename, attachment.ContentType)
		}
	}

	// TODO: Add your business logic here
	// For example, create a support ticket, process an auto-reply, etc.

	return nil
}

// Example: Advanced webhook processing with database integration
func exampleEventTypeChecks(event webhooks.WebhookEvent) {
	// This is a placeholder to show how you might integrate with a database

	// Check if it's a message event
	if webhooks.IsMessageEvent(event) {
		// Extract common message data
		if messageData := webhooks.GetMessageEventData(event); messageData != nil {
			// Update your database
			// db.UpdateMessageStatus(messageData.ID, event.GetType(), event.GetTimestamp())
			fmt.Printf("Would update message %s with status %s\n", messageData.ID, event.GetType())
		}
	}

	// Handle suppression events specially
	if webhooks.IsSuppressionEvent(event) {
		if suppEvent, ok := event.(*webhooks.SuppressionCreatedEvent); ok {
			// Update your email list
			// db.MarkEmailAsSuppressed(suppEvent.Data.Recipient, suppEvent.Data.Reason)
			fmt.Printf("Would mark %s as suppressed\n", suppEvent.Data.Recipient)
		}
	}
}

// Example: Webhook endpoint with custom error handling and logging
func exampleAdvancedWebhookHandler(verifier *webhooks.WebhookVerifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add request ID for tracing
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("webhook-%d", time.Now().UnixNano())
		}

		log.Printf("[%s] Received webhook request from %s", requestID, r.RemoteAddr)

		// Parse and verify
		event, err := verifier.ParseRequest(r)
		if err != nil {
			log.Printf("[%s] Webhook verification failed: %v", requestID, err)

			// Return appropriate status codes
			switch err {
			case webhooks.ErrExpiredTimestamp:
				// Still process but log the delay
				log.Printf("[%s] WARNING: Processing expired webhook (timestamp too old)", requestID)
				// You might still want to process it depending on your use case
			case webhooks.ErrInvalidSignature, webhooks.ErrMissingHeaders:
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			default:
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
		}

		// Log successful verification
		log.Printf("[%s] Webhook verified: type=%s timestamp=%s",
			requestID, event.GetType(), event.GetTimestamp())

		// Process asynchronously to return quickly
		go func() {
			if err := processWebhookEvent(event); err != nil {
				log.Printf("[%s] Failed to process event: %v", requestID, err)
				// Could add retry logic or dead letter queue here
			}
		}()

		// Return success immediately
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     "accepted",
			"request_id": requestID,
		})
	}
}
