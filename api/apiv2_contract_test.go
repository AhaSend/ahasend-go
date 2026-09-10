package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func newContractTestClient(t *testing.T, handler http.HandlerFunc) (*APIClient, func()) {
	t.Helper()

	server := httptest.NewServer(handler)
	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	cfg := NewConfiguration()
	cfg.Host = serverURL.Host
	cfg.Scheme = serverURL.Scheme
	cfg.APIKey = "test-key"
	cfg.RetryConfig.Enabled = false

	return NewAPIClientWithConfig(cfg), server.Close
}

const contactResponseFixture = `{
	"object":"contact",
	"id":"11111111-1111-4111-8111-111111111111",
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-10T11:00:00Z",
	"email":"User+Tag@Example.COM",
	"first_name":"Pat",
	"last_name":"Example",
	"status":"enabled",
	"status_reason":"",
	"unsubscribed":false,
	"unsubscribed_at":null,
	"attributes":{"legacy":{"nested":[true,12.5,null]}},
	"validation_status":"unvalidated",
	"last_validated_at":null
}`

func TestContactsAPIEmitsDeclaredTransport(t *testing.T) {
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	rawEmail := "User+Tag@Example.COM"
	fromTime := time.Date(2026, 9, 9, 10, 0, 0, 0, time.UTC)
	toTime := fromTime.Add(24 * time.Hour)

	type contactCall func(*APIClient) (any, *http.Response, error)
	tests := []struct {
		name              string
		method            string
		path              string
		status            int
		response          string
		idempotencyKey    string
		call              contactCall
		assertRequestBody func(*testing.T, []byte)
	}{
		{
			name:     "get contacts",
			method:   http.MethodGet,
			path:     "/v2/accounts/" + accountID.String() + "/contacts",
			status:   http.StatusOK,
			response: `{"object":"list","data":[` + contactResponseFixture + `],"pagination":{"has_more":true,"next_cursor":"next","previous_cursor":null}}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ContactsAPI.GetContacts(context.Background(), accountID, requests.GetContactsParams{
					Limit:      ahasend.Int32(25),
					After:      ahasend.String("after-cursor"),
					Before:     ahasend.String("before-cursor"),
					Email:      ahasend.String("person@example.com"),
					Status:     ahasend.String("enabled"),
					Subscribed: ahasend.Bool(true),
					FromTime:   &fromTime,
					ToTime:     &toTime,
				})
			},
		},
		{
			name:     "get contact",
			method:   http.MethodGet,
			path:     "/v2/accounts/" + accountID.String() + "/contacts/" + rawEmail,
			status:   http.StatusOK,
			response: contactResponseFixture,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ContactsAPI.GetContact(context.Background(), accountID, rawEmail)
			},
		},
		{
			name:           "create contact",
			method:         http.MethodPost,
			path:           "/v2/accounts/" + accountID.String() + "/contacts",
			status:         http.StatusCreated,
			response:       contactResponseFixture,
			idempotencyKey: "create-contact",
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ContactsAPI.CreateContact(context.Background(), accountID, requests.CreateContactRequest{
					Email:      rawEmail,
					Attributes: map[string]any{"customer": true},
				}, WithIdempotencyKey("create-contact"))
			},
			assertRequestBody: func(t *testing.T, body []byte) {
				var request requests.CreateContactRequest
				require.NoError(t, json.Unmarshal(body, &request))
				assert.Equal(t, rawEmail, request.Email)
				assert.Equal(t, true, request.Attributes["customer"])
			},
		},
		{
			name:     "update contact",
			method:   http.MethodPut,
			path:     "/v2/accounts/" + accountID.String() + "/contacts/" + rawEmail,
			status:   http.StatusOK,
			response: contactResponseFixture,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ContactsAPI.UpdateContact(context.Background(), accountID, rawEmail, requests.UpdateContactRequest{
					Attributes: map[string]any{"obsolete": nil},
				})
			},
			assertRequestBody: func(t *testing.T, body []byte) {
				var request map[string]any
				require.NoError(t, json.Unmarshal(body, &request))
				attributes, ok := request["attributes"].(map[string]any)
				require.True(t, ok)
				assert.Contains(t, attributes, "obsolete")
				assert.Nil(t, attributes["obsolete"])
			},
		},
		{
			name:     "delete contact",
			method:   http.MethodDelete,
			path:     "/v2/accounts/" + accountID.String() + "/contacts/" + rawEmail,
			status:   http.StatusOK,
			response: `{"message":"contact deleted"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ContactsAPI.DeleteContact(context.Background(), accountID, rawEmail)
			},
		},
		{
			name:           "batch upsert contacts",
			method:         http.MethodPost,
			path:           "/v2/accounts/" + accountID.String() + "/contacts/batch",
			status:         http.StatusOK,
			response:       `{"object":"list","created":0,"updated":0,"failed":1,"data":[{"position":0,"email":"User+Tag@Example.COM","outcome":"failed","reason":"invalid attribute"}]}`,
			idempotencyKey: "batch-contacts",
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ContactsAPI.BatchUpsertContacts(context.Background(), accountID, requests.BatchUpsertContactsRequest{
					Data: []requests.BatchUpsertContactInput{{Email: rawEmail, Attributes: map[string]any{"obsolete": nil}}},
				}, WithIdempotencyKey("batch-contacts"))
			},
			assertRequestBody: func(t *testing.T, body []byte) {
				var request requests.BatchUpsertContactsRequest
				require.NoError(t, json.Unmarshal(body, &request))
				require.Len(t, request.Data, 1)
				assert.Contains(t, request.Data[0].Attributes, "obsolete")
				assert.Nil(t, request.Data[0].Attributes["obsolete"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedRequest *http.Request
			var capturedBody []byte
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				capturedRequest = r
				capturedBody, _ = io.ReadAll(r.Body)
				w.Header().Set("Content-Type", "application/json")
				if tt.idempotencyKey != "" {
					w.Header().Set("Idempotent-Replayed", "true")
				}
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.response))
			})
			defer cleanup()

			result, httpResponse, err := tt.call(client)

			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, httpResponse)
			require.NotNil(t, capturedRequest)
			assert.Equal(t, tt.method, capturedRequest.Method)
			assert.Equal(t, tt.path, capturedRequest.URL.Path)
			assert.Equal(t, tt.idempotencyKey, capturedRequest.Header.Get("Idempotency-Key"))
			if tt.idempotencyKey != "" {
				assert.Equal(t, "true", httpResponse.Header.Get("Idempotent-Replayed"))
			}
			if tt.assertRequestBody != nil {
				tt.assertRequestBody(t, capturedBody)
			}

			if tt.name == "get contacts" {
				query := capturedRequest.URL.Query()
				assert.Equal(t, "25", query.Get("limit"))
				assert.Equal(t, "after-cursor", query.Get("after"))
				assert.Equal(t, "before-cursor", query.Get("before"))
				assert.Equal(t, "person@example.com", query.Get("email"))
				assert.Equal(t, "enabled", query.Get("status"))
				assert.Equal(t, "true", query.Get("subscribed"))
				assert.Equal(t, fromTime.Format(time.RFC3339), query.Get("from_time"))
				assert.Equal(t, toTime.Format(time.RFC3339), query.Get("to_time"))
			}
			if tt.name == "batch upsert contacts" {
				batch := result.(*responses.BatchUpsertContactsResponse)
				assert.Equal(t, 1, batch.Failed)
				require.Len(t, batch.Data, 1)
				assert.Equal(t, 0, batch.Data[0].Position)
				assert.Equal(t, "failed", batch.Data[0].Outcome)
				assert.Equal(t, "invalid attribute", batch.Data[0].Reason)
			}
		})
	}
}

func TestContactsAPIRawEmailIsEncodedExactlyOnce(t *testing.T) {
	accountID := uuid.New()
	rawEmail := "User+Tag@Example.COM"
	wantPath := "/v2/accounts/" + accountID.String() + "/contacts/" + rawEmail

	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, wantPath, r.URL.Path)
		assert.Equal(t, wantPath, r.URL.EscapedPath())
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(contactResponseFixture))
	})
	defer cleanup()

	_, _, err := client.ContactsAPI.GetContact(context.Background(), accountID, rawEmail)
	require.NoError(t, err)
}

func TestContactsAPIPreservesDocumentedErrorHeaders(t *testing.T) {
	tests := []struct {
		name           string
		status         int
		headers        map[string]string
		wantType       ErrorType
		wantRetryAfter int
	}{
		{name: "in-progress conflict", status: http.StatusConflict, headers: map[string]string{"Idempotent-Replayed": "false", "Retry-After": "7"}, wantType: ErrorTypeIdempotencyConflict, wantRetryAfter: 7},
		{name: "payload mismatch", status: http.StatusUnprocessableEntity, wantType: ErrorTypeIdempotency},
		{name: "rate limit", status: http.StatusTooManyRequests, headers: map[string]string{"Retry-After": "11"}, wantType: ErrorTypeRateLimit, wantRetryAfter: 11},
		{name: "server error", status: http.StatusInternalServerError, wantType: ErrorTypeServer},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
				for key, value := range tt.headers {
					w.Header().Set(key, value)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(`{"message":"contact request failed"}`))
			})
			defer cleanup()

			_, httpResponse, err := client.ContactsAPI.BatchUpsertContacts(context.Background(), uuid.New(), requests.BatchUpsertContactsRequest{
				Data: []requests.BatchUpsertContactInput{{Email: "person@example.com"}},
			}, WithIdempotencyKey("contact-error"))

			require.Error(t, err)
			require.NotNil(t, httpResponse)
			assert.Equal(t, tt.status, httpResponse.StatusCode)
			for key, value := range tt.headers {
				assert.Equal(t, value, httpResponse.Header.Get(key))
			}
			var apiErr *APIError
			require.ErrorAs(t, err, &apiErr)
			assert.Equal(t, tt.wantType, apiErr.Type)
			assert.Equal(t, tt.wantRetryAfter, apiErr.RetryAfter)
		})
	}
}

func TestOpenAPIContactOperationsAndResponses(t *testing.T) {
	type responseContract struct {
		Ref     string         `yaml:"$ref"`
		Headers map[string]any `yaml:"headers"`
	}
	type operationContract struct {
		OperationID string                      `yaml:"operationId"`
		CodeSamples []struct{ Lang string }     `yaml:"x-code-samples"`
		Responses   map[string]responseContract `yaml:"responses"`
	}
	type openAPIContract struct {
		Paths map[string]map[string]operationContract `yaml:"paths"`
	}

	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	var spec openAPIContract
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	tests := []struct {
		path        string
		method      string
		operationID string
		statuses    []string
	}{
		{path: "/v2/accounts/{account_id}/contacts", method: "get", operationID: "getContacts", statuses: []string{"200", "400", "401", "403", "429", "500"}},
		{path: "/v2/accounts/{account_id}/contacts", method: "post", operationID: "createContact", statuses: []string{"201", "400", "401", "403", "409", "422", "429", "500"}},
		{path: "/v2/accounts/{account_id}/contacts/batch", method: "post", operationID: "batchUpsertContacts", statuses: []string{"200", "400", "401", "403", "409", "422", "429", "500"}},
		{path: "/v2/accounts/{account_id}/contacts/{id_or_email}", method: "get", operationID: "getContact", statuses: []string{"200", "401", "403", "404", "429", "500"}},
		{path: "/v2/accounts/{account_id}/contacts/{id_or_email}", method: "put", operationID: "updateContact", statuses: []string{"200", "400", "401", "403", "404", "409", "429", "500"}},
		{path: "/v2/accounts/{account_id}/contacts/{id_or_email}", method: "delete", operationID: "deleteContact", statuses: []string{"200", "401", "403", "404", "429", "500"}},
	}

	for _, tt := range tests {
		t.Run(tt.operationID, func(t *testing.T) {
			path, ok := spec.Paths[tt.path]
			require.True(t, ok)
			operation, ok := path[tt.method]
			require.True(t, ok)
			assert.Equal(t, tt.operationID, operation.OperationID)

			var goSamples int
			for _, sample := range operation.CodeSamples {
				if sample.Lang == "go" {
					goSamples++
				}
			}
			assert.Equal(t, 1, goSamples)

			actualStatuses := make([]string, 0, len(operation.Responses))
			for status := range operation.Responses {
				actualStatuses = append(actualStatuses, status)
			}
			sort.Strings(actualStatuses)
			assert.Equal(t, tt.statuses, actualStatuses)
		})
	}

	create := spec.Paths["/v2/accounts/{account_id}/contacts"]["post"].Responses
	assert.Contains(t, create["201"].Headers, "Idempotent-Replayed")
	assert.Contains(t, create["400"].Headers, "Idempotent-Replayed")
	assert.Contains(t, create["409"].Headers, "Idempotent-Replayed")
	assert.Contains(t, create["409"].Headers, "Retry-After")
	assert.Equal(t, "#/components/responses/IdempotencyPayloadMismatch", create["422"].Ref)

	batch := spec.Paths["/v2/accounts/{account_id}/contacts/batch"]["post"].Responses
	assert.Contains(t, batch["200"].Headers, "Idempotent-Replayed")
	assert.Contains(t, batch["400"].Headers, "Idempotent-Replayed")
	assert.Contains(t, batch["409"].Headers, "Idempotent-Replayed")
	assert.Contains(t, batch["409"].Headers, "Retry-After")
	assert.Equal(t, "#/components/responses/IdempotencyPayloadMismatch", batch["422"].Ref)
}

func TestSuppressionsAPIGetSuppressionsUsesTimeQueryNames(t *testing.T) {
	var lastRequest *http.Request
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	})
	defer cleanup()

	fromTime := time.Date(2026, 4, 18, 10, 0, 0, 0, time.UTC)
	toTime := fromTime.Add(time.Hour)
	email := "user@example.com"
	domain := "example.com"

	_, _, err := client.SuppressionsAPI.GetSuppressions(context.Background(), uuid.New(), requests.GetSuppressionsParams{
		Email:    &email,
		Domain:   &domain,
		FromTime: &fromTime,
		ToTime:   &toTime,
	})

	require.NoError(t, err)
	require.NotNil(t, lastRequest)
	query := lastRequest.URL.Query()
	assert.Equal(t, email, query.Get("email"))
	assert.Equal(t, domain, query.Get("domain"))
	assert.Equal(t, fromTime.Format(time.RFC3339), query.Get("from_time"))
	assert.Equal(t, toTime.Format(time.RFC3339), query.Get("to_time"))
	assert.Empty(t, query.Get("from_date"))
	assert.Empty(t, query.Get("to_date"))
}

func TestRoutesAPIGetRoutesWithParamsSerializesDomain(t *testing.T) {
	var lastRequest *http.Request
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	})
	defer cleanup()

	_, _, err := client.RoutesAPI.GetRoutesWithParams(context.Background(), uuid.New(), requests.GetRoutesParams{
		Domain: ahasend.String("example.com"),
		PaginationParams: common.PaginationParams{
			Limit: ahasend.Int32(25),
		},
	})

	require.NoError(t, err)
	require.NotNil(t, lastRequest)
	assert.Equal(t, "example.com", lastRequest.URL.Query().Get("domain"))
	assert.Equal(t, "25", lastRequest.URL.Query().Get("limit"))
}

func TestSMTPCredentialsAPIUsesUUIDPathIDs(t *testing.T) {
	accountID := uuid.New()
	credentialID := uuid.New()
	var lastPath string

	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"credential_smtp","id":"` + credentialID.String() + `","created_at":"2026-04-18T10:00:00Z","updated_at":"2026-04-18T10:00:00Z","name":"prod","username":"user","sandbox":false,"scope":"global","domains":[]}`))
	})
	defer cleanup()

	resp, _, err := client.SMTPCredentialsAPI.GetSMTPCredential(context.Background(), accountID, credentialID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, credentialID, resp.ID)
	assert.Equal(t, "/v2/accounts/"+accountID.String()+"/smtp-credentials/"+credentialID.String(), lastPath)
}

func TestMessagesAPICreateConversationMessageUsesToField(t *testing.T) {
	accountID := uuid.New()
	var body map[string]interface{}
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/v2/accounts/"+accountID.String()+"/messages/conversation", r.URL.Path)
		err := json.NewDecoder(r.Body).Decode(&body)
		require.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"object":"list","data":[]}`))
	})
	defer cleanup()

	_, _, err := client.MessagesAPI.CreateConversationMessage(context.Background(), accountID, requests.CreateConversationMessageRequest{
		From:    common.SenderAddress{Email: "sender@example.com"},
		To:      []common.SenderAddress{{Email: "recipient@example.com"}},
		Subject: "Hello",
	})

	require.NoError(t, err)
	require.Contains(t, body, "to")
	assert.NotContains(t, body, "recipients")
}
