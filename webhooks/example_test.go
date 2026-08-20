package webhooks_test

import (
	"fmt"
	"net/http"

	"github.com/AhaSend/ahasend-go/webhooks"
)

// ExampleWebhookVerifier_Parse shows how to read the diagnostics for the
// delivery attempt a message event reports on.
//
// There is no Output comment, so this example is compiled and vetted but never
// run: verification is over the current time, and pinning a signature to make
// the output deterministic would test the fixture rather than the code.
func ExampleWebhookVerifier_Parse() {
	verifier, err := webhooks.NewWebhookVerifier("aha-whsec-your-webhook-secret")
	if err != nil {
		return
	}

	// In a real handler these are the request body and its headers. Use
	// ParseRequest to let the SDK read the body for you.
	var (
		payload []byte
		headers http.Header
	)

	event, err := verifier.Parse(payload, headers)
	if err != nil {
		// Parse verifies the HMAC signature before it decodes anything, so this
		// covers both an unsigned payload and a malformed one. Answer 400.
		return
	}

	bounced, ok := event.(*webhooks.MessageBouncedEvent)
	if !ok {
		return
	}

	// The attempt is optional and absent more often than present. A nil pointer
	// covers both a missing field and an explicit null, which mean the same
	// thing, so this is never an error.
	attempt := bounced.Data.DeliveryAttempt
	if attempt == nil {
		fmt.Println("bounced with no recorded delivery attempt")
		return
	}

	// 0 is a real code, meaning response content was recorded without an SMTP
	// code, so report it rather than testing for truthiness.
	fmt.Printf("bounced with SMTP code %d\n", attempt.SMTPCode)

	// The set of classifications is open: branch on the buckets you handle and
	// keep a fallback for values this SDK predates. Rejecting an unfamiliar
	// value would answer 400 to a valid delivery, and 100 consecutive failures
	// disable the endpoint.
	if attempt.Classification != nil {
		switch *attempt.Classification {
		case webhooks.ClassificationInvalidRecipient, webhooks.ClassificationBadDomain:
			fmt.Println("permanent addressing failure")
		case webhooks.ClassificationQuotaIssues:
			fmt.Println("mailbox full, worth retrying later")
		default:
			fmt.Printf("unhandled classification %s\n", *attempt.Classification)
		}
	}

	// Response and Description are free-form diagnostic text and routinely embed
	// the recipient and parts of the message, so they are read for display and
	// never logged. Branch on the codes and the classification instead.
}
