package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// allClassifications is every value the SDK ships as a constant. It is
// hand-maintained alongside the const block and checked against the
// specification by TestDeliveryAttemptClassificationMatchesSpec.
var allClassifications = []DeliveryAttemptClassification{
	ClassificationInvalidRecipient,
	ClassificationBadDomain,
	ClassificationInactiveMailbox,
	ClassificationInvalidSender,
	ClassificationQuotaIssues,
	ClassificationNoAnswerFromHost,
	ClassificationBadConnection,
	ClassificationDNSFailure,
	ClassificationRoutingErrors,
	ClassificationTransientFailure,
	ClassificationMessageExpired,
	ClassificationProtocolErrors,
	ClassificationAuthenticationFailed,
	ClassificationPolicyRelated,
	ClassificationUncategorized,
}

// webhookSpec is the fragment of openapi/webhooks.yaml these tests read.
type webhookSpec struct {
	// Webhooks is the OpenAPI 3.1 webhooks section, keyed by event type. Only
	// each operation's request example is read.
	Webhooks map[string]struct {
		Post struct {
			RequestBody struct {
				Content map[string]struct {
					Example map[string]interface{} `yaml:"example"`
				} `yaml:"content"`
			} `yaml:"requestBody"`
		} `yaml:"post"`
	} `yaml:"webhooks"`
	Components struct {
		Schemas struct {
			DeliveryAttempt struct {
				Properties struct {
					Classification struct {
						Enum        []string `yaml:"enum"`
						KnownValues []string `yaml:"x-known-values"`
					} `yaml:"classification"`
				} `yaml:"properties"`
			} `yaml:"DeliveryAttempt"`
		} `yaml:"schemas"`
	} `yaml:"components"`
}

// findWebhookSpec walks up from the working directory and returns the absolute
// path of the webhook specification, so that the test's package directory does
// not matter. internal/prismmock does the same for openapi/openapi.yaml, but
// its finder is unexported and hardcodes that filename.
func findWebhookSpec() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		candidate := filepath.Join(dir, "openapi", "webhooks.yaml")
		if _, statErr := os.Stat(candidate); statErr == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("openapi/webhooks.yaml not found above the working directory")
		}
		dir = parent
	}
}

// loadWebhookSpec reads and decodes openapi/webhooks.yaml.
func loadWebhookSpec(t *testing.T) webhookSpec {
	t.Helper()

	path, err := findWebhookSpec()
	require.NoError(t, err)

	raw, err := os.ReadFile(path)
	require.NoError(t, err)

	var spec webhookSpec
	require.NoError(t, yaml.Unmarshal(raw, &spec))
	return spec
}

// TestDeliveryAttemptClassificationMatchesSpec guards the one seam in this
// feature that is genuinely loose: the constants above are hand-copied out of
// x-known-values in a specification that is refreshed by hand, so a spec
// refresh that adds or renames a bucket must not leave the SDK silently behind.
func TestDeliveryAttemptClassificationMatchesSpec(t *testing.T) {
	spec := loadWebhookSpec(t)

	classification := spec.Components.Schemas.DeliveryAttempt.Properties.Classification
	require.NotEmpty(t, classification.KnownValues,
		"x-known-values is missing from the DeliveryAttempt classification schema")

	t.Run("the spec list has no duplicates", func(t *testing.T) {
		seen := make(map[string]bool, len(classification.KnownValues))
		for _, value := range classification.KnownValues {
			assert.False(t, seen[value], "%q appears twice in x-known-values", value)
			seen[value] = true
		}
	})

	t.Run("the constants match the spec exactly", func(t *testing.T) {
		declared := make([]string, 0, len(allClassifications))
		for _, value := range allClassifications {
			declared = append(declared, string(value))
		}
		// A set comparison, not an ordered one: membership is the property
		// under test, so asserting order would only raise failures whose fix
		// is a no-op reordering.
		assert.ElementsMatch(t, classification.KnownValues, declared)
	})

	t.Run("classification is not a closed set", func(t *testing.T) {
		// An enum in the spec would mean generators and validators start
		// rejecting values, which is the failure mode this field exists to
		// avoid: a rejected delivery becomes a 400, and 100 consecutive
		// failures disable the endpoint.
		assert.Empty(t, classification.Enum,
			"classification gained an enum; the bucket set is open and must stay open")
	})
}

func TestDeliveryAttemptParsing(t *testing.T) {
	verifier, err := NewWebhookVerifier(testWebhookSecret)
	require.NoError(t, err)

	// Only shapes a delivery can actually carry. The contract in
	// openapi/webhooks.yaml lists smtp_code as required and typed integer
	// within an attempt that is itself an object or null, so a missing code, a
	// non-integer code, and an attempt that is neither object nor null are all
	// outside it — and Parse verifies the HMAC signature before it decodes
	// anything, so nothing unsigned reaches the decoder to begin with. A payload
	// that broke the contract anyway would be a fault at the source, and
	// rejecting it as ErrInvalidPayload is the loud, correct response.
	testCases := []struct {
		name     string
		fragment string
		assert   func(t *testing.T, attempt *DeliveryAttempt)
	}{
		{
			name:     "absent",
			fragment: "",
			assert: func(t *testing.T, attempt *DeliveryAttempt) {
				assert.Nil(t, attempt)
			},
		},
		{
			// An explicit null carries exactly the meaning of a missing field,
			// so it must reach the consumer as the same nil rather than as an
			// error.
			name:     "explicit null",
			fragment: `,"delivery_attempt": null`,
			assert: func(t *testing.T, attempt *DeliveryAttempt) {
				assert.Nil(t, attempt)
			},
		},
		{
			name: "every field populated",
			fragment: `,"delivery_attempt": {
				"classification": "InvalidRecipient",
				"smtp_code": 550,
				"enhanced_status_code": "5.1.1",
				"response": "The email account that you tried to reach does not exist",
				"description": "The recipient address does not exist",
				"command": "RCPT TO"
			}`,
			assert: func(t *testing.T, attempt *DeliveryAttempt) {
				require.NotNil(t, attempt)
				require.NotNil(t, attempt.Classification)
				assert.Equal(t, ClassificationInvalidRecipient, *attempt.Classification)
				assert.Equal(t, 550, attempt.SMTPCode)
				require.NotNil(t, attempt.EnhancedStatusCode)
				assert.Equal(t, "5.1.1", *attempt.EnhancedStatusCode)
				require.NotNil(t, attempt.Response)
				assert.Equal(t, "The email account that you tried to reach does not exist", *attempt.Response)
				require.NotNil(t, attempt.Description)
				assert.Equal(t, "The recipient address does not exist", *attempt.Description)
				require.NotNil(t, attempt.Command)
				assert.Equal(t, "RCPT TO", *attempt.Command)
			},
		},
		{
			// A successful delivery carries a code and nothing the classifier
			// produced, which is the common shape on message.delivered.
			name: "smtp code only",
			fragment: `,"delivery_attempt": {
				"smtp_code": 250
			}`,
			assert: func(t *testing.T, attempt *DeliveryAttempt) {
				require.NotNil(t, attempt)
				assert.Equal(t, 250, attempt.SMTPCode)
				assert.Nil(t, attempt.Classification)
				assert.Nil(t, attempt.EnhancedStatusCode)
				assert.Nil(t, attempt.Response)
				assert.Nil(t, attempt.Description)
				assert.Nil(t, attempt.Command)
			},
		},
		{
			// 0 is a real code, meaning response content was recorded without
			// an SMTP code. It must be distinguishable from an absent attempt.
			name: "zero smtp code",
			fragment: `,"delivery_attempt": {
				"smtp_code": 0,
				"response": "connection reset before greeting"
			}`,
			assert: func(t *testing.T, attempt *DeliveryAttempt) {
				require.NotNil(t, attempt)
				assert.Equal(t, 0, attempt.SMTPCode)
				// Asserted so that the row cannot pass on an attempt whose fields
				// were all silently dropped, which would leave SMTPCode at 0 too.
				require.NotNil(t, attempt.Response)
				assert.Equal(t, "connection reset before greeting", *attempt.Response)
			},
		},
		{
			// The bucket set is open, so a value this release predates must
			// decode and reach the consumer verbatim rather than being rejected
			// or flattened to a fallback.
			name: "unknown classification",
			fragment: `,"delivery_attempt": {
				"classification": "SomeFutureBucket",
				"smtp_code": 550
			}`,
			assert: func(t *testing.T, attempt *DeliveryAttempt) {
				require.NotNil(t, attempt)
				require.NotNil(t, attempt.Classification)
				assert.Equal(t, DeliveryAttemptClassification("SomeFutureBucket"), *attempt.Classification)
			},
		},
		{
			// A field added to the object by a later server release must not
			// break an already-published client.
			name: "unknown field on the object",
			fragment: `,"delivery_attempt": {
				"smtp_code": 550,
				"classification": "BadDomain",
				"some_future_field": {"nested": [1, 2, 3]}
			}`,
			assert: func(t *testing.T, attempt *DeliveryAttempt) {
				require.NotNil(t, attempt)
				assert.Equal(t, 550, attempt.SMTPCode)
				require.NotNil(t, attempt.Classification)
				assert.Equal(t, ClassificationBadDomain, *attempt.Classification)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			payload := messageEventPayload("message.bounced", "on_bounced", tc.fragment)

			event, err := verifier.Parse([]byte(payload), signedHeaders(t, verifier, payload))
			require.NoError(t, err)

			bounced, ok := event.(*MessageBouncedEvent)
			require.True(t, ok)
			tc.assert(t, bounced.Data.DeliveryAttempt)
		})
	}

	t.Run("every event sharing MessageEventData carries the field", func(t *testing.T) {
		fragment := `,"delivery_attempt": {"smtp_code": 421, "classification": "NoAnswerFromHost"}`

		events := []struct {
			eventType string
			eventName string
		}{
			{"message.reception", "on_reception"},
			{"message.delivered", "on_delivered"},
			{"message.transient_error", "on_transient_error"},
			{"message.failed", "on_failed"},
			{"message.bounced", "on_bounced"},
			{"message.suppressed", "on_suppressed"},
			{"message.opened", "on_opened"},
		}

		for _, e := range events {
			e := e
			t.Run(e.eventType, func(t *testing.T) {
				payload := messageEventPayload(e.eventType, e.eventName, fragment)

				event, err := verifier.Parse([]byte(payload), signedHeaders(t, verifier, payload))
				require.NoError(t, err)

				// Reaching the attempt through the accessor rather than the
				// concrete type is what proves the field is on the shared data
				// shape and not on one event's copy of it.
				data := GetMessageEventData(event)
				require.NotNil(t, data, "GetMessageEventData did not recognize %s", e.eventType)
				require.NotNil(t, data.DeliveryAttempt)
				assert.Equal(t, 421, data.DeliveryAttempt.SMTPCode)
				require.NotNil(t, data.DeliveryAttempt.Classification)
				assert.Equal(t, ClassificationNoAnswerFromHost, *data.DeliveryAttempt.Classification)
			})
		}
	})

	t.Run("message.clicked carrying an attempt still parses", func(t *testing.T) {
		// message.clicked has its own data shape and does not expose the
		// field, but the specification warns against treating an attempt on an
		// unexpected event as impossible, so the payload must still decode.
		payload := `{
			"type": "message.clicked",
			"webhook_id": "abe11757-2886-4b55-96f1-0e0afc95795a",
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
				"is_bot": false,
				"delivery_attempt": {"smtp_code": 250}
			}
		}`

		event, err := verifier.Parse([]byte(payload), signedHeaders(t, verifier, payload))
		require.NoError(t, err)

		clicked, ok := event.(*MessageClickedEvent)
		require.True(t, ok)
		assert.Equal(t, "https://example.com/link", clicked.Data.URL)
	})
}

// TestSpecExamplesDecodeStrictly decodes every example payload the
// specification publishes, twice: once through the real Parse, and once through
// a decoder with DisallowUnknownFields set. The second pass is the guard — it
// fails if a key the example declares maps to no field on the type Parse chose.
// Keys inside a map-typed field are exempt, correctly: Headers accepts whatever
// the sending server wrote.
//
// This is the guard the classification drift test cannot provide. That test
// watches the classification *values*; this one watches the *keys*. A spec
// refresh that renames a property — response to response_text, say — leaves the
// Go struct tag pointing at a key nobody sends, and every consumer then reads
// nil forever with no test failing anywhere: the drift test still passes
// because the values did not change, and the hand-written parse fixtures still
// pass because they were written against the old key. Using the specification's
// own examples as the fixtures closes that gap, since they are refreshed with
// the spec rather than alongside it.
//
// Two limits are worth knowing. A field no example carries cannot be checked
// this way, and description is currently the only such field on
// DeliveryAttempt; it is also where a silent nil costs least, being display
// prose that must never be branched on. And this cannot see a consistent
// mismapping — swap the json tags of two same-typed fields and every key still
// maps to a field, so only a fixture naming an expected value against a named
// Go field catches it. That is why the hand-written tests here and in
// TestWebhookParsing are not made redundant by this one; verified by swapping
// Subject and MessageIDHeader, which this test ignores and TestWebhookParsing
// fails on.
//
// A mis-dispatch between two events that share a data shape is invisible here
// for the same reason: the seven events carrying MessageEventData are
// interchangeable to a key check, and GetType reads the payload's own type
// string rather than anything derived from the Go type. TestWebhookParsing is
// again what catches that.
//
// It also fails when a spec refresh adds a property to an example, reported as
// an unknown field. Usually that is a true positive — the SDK is dropping a
// field the specification chose to publish an example of — and the fix is to
// model the field. If instead the field is one this SDK deliberately declines
// to surface, skip it here by name with a comment recording that decision,
// rather than by loosening the check. Note all of this is about what the SDK
// *surfaces*: the production decoder in Parse stays tolerant, and an unknown
// field on the wire is still ignored rather than rejected. Only this test is
// strict, and only against the specification's own examples.
func TestSpecExamplesDecodeStrictly(t *testing.T) {
	spec := loadWebhookSpec(t)
	require.NotEmpty(t, spec.Webhooks, "no webhooks section found in the specification")

	verifier, err := NewWebhookVerifier(testWebhookSecret)
	require.NoError(t, err)

	examined := 0
	for eventType, operation := range spec.Webhooks {
		example := operation.Post.RequestBody.Content["application/json"].Example
		if example == nil {
			continue
		}
		examined++

		t.Run(eventType, func(t *testing.T) {
			payload, err := json.Marshal(example)
			require.NoError(t, err)

			// Through the real entry point rather than a second dispatch table
			// that could drift from it. This catches a mis-dispatch only across
			// differing shapes — see the limits above.
			event, err := verifier.Parse(payload, signedHeaders(t, verifier, string(payload)))
			require.NoError(t, err)
			require.Equal(t, eventType, event.GetType())

			// A fresh value of whatever concrete type Parse chose, so this
			// needs no second copy of Parse's type dispatch to drift from it.
			// Asserted rather than assumed: Elem panics on a non-pointer, and a
			// clear failure beats a stack trace if Parse ever returns a value.
			require.Equal(t, reflect.Ptr, reflect.TypeOf(event).Kind(),
				"Parse returned %T by value; this test needs a pointer to allocate a fresh one", event)
			strict := reflect.New(reflect.TypeOf(event).Elem()).Interface()
			decoder := json.NewDecoder(bytes.NewReader(payload))
			decoder.DisallowUnknownFields()
			require.NoError(t, decoder.Decode(strict),
				"a key in the specification's example does not map to a field on %T. "+
					"Either a json tag no longer matches the key the producer sends, or "+
					"the specification published a field this SDK does not model yet", event)
		})
	}

	// Without this the loop could silently examine nothing — an empty spec
	// section, or a content type that is no longer application/json.
	assert.Equal(t, len(spec.Webhooks), examined,
		"every published webhook should carry a request example")
}
