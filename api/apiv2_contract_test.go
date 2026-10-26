package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strings"
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
			if tt.assertRequestBody != nil {
				tt.assertRequestBody(t, capturedBody)
			}

			if tt.name == "get contacts" {
				query := capturedRequest.URL.Query()
				assert.Equal(t, "25", query.Get("limit"))
				assert.Equal(t, "after-cursor", query.Get("after"))
				assert.False(t, query.Has("before"))
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
	contactsPath := "/v2/accounts/" + accountID.String() + "/contacts/"

	tests := []struct {
		name        string
		rawEmail    string
		wantEscaped string
	}{
		{name: "plus and uppercase", rawEmail: "User+Tag@Example.COM", wantEscaped: "User+Tag@Example.COM"},
		{name: "slash", rawEmail: "a/b@example.com", wantEscaped: "a%2Fb@example.com"},
		{name: "double slash", rawEmail: "a//b@example.com", wantEscaped: "a%2F%2Fb@example.com"},
		{name: "percent", rawEmail: "100%off@example.com", wantEscaped: "100%25off@example.com"},
		{name: "unicode local part", rawEmail: "jöhn@example.com", wantEscaped: "j%C3%B6hn@example.com"},
		{name: "query and fragment delimiters", rawEmail: "a?b#c@example.com", wantEscaped: "a%3Fb%23c@example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestURI, escapedPath, decodedPath string
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				requestURI = r.RequestURI
				escapedPath = r.URL.EscapedPath()
				decodedPath = r.URL.Path
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(contactResponseFixture))
			})
			defer cleanup()

			_, _, err := client.ContactsAPI.GetContact(context.Background(), accountID, tt.rawEmail)

			require.NoError(t, err)
			assert.Equal(t, contactsPath+tt.wantEscaped, requestURI)
			assert.Equal(t, contactsPath+tt.wantEscaped, escapedPath)
			assert.Equal(t, contactsPath+tt.rawEmail, decodedPath)
		})
	}
}

func TestContactsAPIRejectsEncodedNullCreateAttribute(t *testing.T) {
	var nilString *string
	requestSent := make(chan struct{}, 1)
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		requestSent <- struct{}{}
		w.WriteHeader(http.StatusCreated)
	})
	defer cleanup()

	result, response, err := client.ContactsAPI.CreateContact(
		context.Background(),
		uuid.New(),
		requests.CreateContactRequest{
			Email:      "person@example.com",
			Attributes: map[string]any{"obsolete": nilString},
		},
	)

	require.Error(t, err)
	assert.Nil(t, response)
	assert.NotNil(t, result)
	select {
	case <-requestSent:
		t.Fatal("CreateContact sent a request with a null attribute")
	default:
	}
	var apiError *APIError
	require.ErrorAs(t, err, &apiError)
	assert.Equal(t, ErrorTypeValidation, apiError.Type)
	assert.Contains(t, apiError.Message, `attribute "obsolete" must not be null`)
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

// responseContract, operationContract and openAPIContract read just enough of
// the bundled specification for the tests below to assert what an operation
// documents.
type responseContract struct {
	Ref     string         `yaml:"$ref"`
	Headers map[string]any `yaml:"headers"`
}

type operationContract struct {
	OperationID string                      `yaml:"operationId"`
	CodeSamples []struct{ Lang string }     `yaml:"x-code-samples"`
	Security    []map[string][]string       `yaml:"security"`
	Responses   map[string]responseContract `yaml:"responses"`
}

type openAPIContract struct {
	Paths map[string]map[string]operationContract `yaml:"paths"`
}

func loadOpenAPIContract(t *testing.T) openAPIContract {
	t.Helper()

	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	var spec openAPIContract
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	return spec
}

func TestOpenAPIContactOperationsAndResponses(t *testing.T) {
	spec := loadOpenAPIContract(t)

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

const listResponseFixture = `{
	"object":"contact_list",
	"id":"22222222-2222-4222-8222-222222222222",
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-10T11:00:00Z",
	"name":"Newsletter",
	"description":"",
	"tags":[],
	"contact_count":7
}`

const listContactResponseFixture = `{
	"object":"list_contact",
	"list_id":"22222222-2222-4222-8222-222222222222",
	"contact_id":"11111111-1111-4111-8111-111111111111",
	"email":"user+tag@example.com",
	"subscription_status":"unsubscribed",
	"subscribed_at":"2026-09-10T10:00:00Z",
	"unsubscribed_at":"2026-09-11T10:00:00Z",
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-11T10:00:00Z"
}`

func TestListsAPIEmitsDeclaredTransport(t *testing.T) {
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	listID := uuid.MustParse("22222222-2222-4222-8222-222222222222")
	contactID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	rawEmail := "a/b+Tag@Example.COM"
	escapedEmail := "a%2Fb+Tag@Example.COM"
	listsPath := "/v2/accounts/" + accountID.String() + "/lists"
	listPath := listsPath + "/" + listID.String()

	type listCall func(*APIClient) (any, *http.Response, error)
	tests := []struct {
		name           string
		method         string
		path           string
		status         int
		response       string
		idempotencyKey string
		wantQuery      url.Values
		wantBody       string
		call           listCall
		assertResult   func(*testing.T, any)
	}{
		{
			name:      "get lists",
			method:    http.MethodGet,
			path:      listsPath,
			status:    http.StatusOK,
			response:  `{"object":"list","data":[` + listResponseFixture + `],"pagination":{"has_more":true,"next_cursor":"next","previous_cursor":null}}`,
			wantQuery: url.Values{"limit": {"25"}, "after": {"after-cursor"}, "name": {"News"}},
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.GetLists(context.Background(), accountID, requests.GetListsParams{
					Limit: ahasend.Int32(25),
					After: ahasend.String("after-cursor"),
					Name:  ahasend.String("News"),
				})
			},
			assertResult: func(t *testing.T, result any) {
				page := result.(*responses.PaginatedContactListsResponse)
				require.Len(t, page.Data, 1)
				assert.Equal(t, listID, page.Data[0].ID)
				assert.Equal(t, 7, page.Data[0].ContactCount)
				assert.Equal(t, []string{}, page.Data[0].Tags)
				assert.True(t, page.Pagination.HasMore)
				require.NotNil(t, page.Pagination.NextCursor)
				assert.Equal(t, "next", *page.Pagination.NextCursor)
			},
		},
		{
			name:           "create list",
			method:         http.MethodPost,
			path:           listsPath,
			status:         http.StatusCreated,
			response:       listResponseFixture,
			idempotencyKey: "create-list",
			wantBody:       `{"name":"Newsletter","description":"Monthly","tags":["news"]}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.CreateList(context.Background(), accountID, requests.CreateListRequest{
					Name:        "Newsletter",
					Description: ahasend.String("Monthly"),
					Tags:        []string{"news"},
				}, WithIdempotencyKey("create-list"))
			},
			assertResult: func(t *testing.T, result any) {
				assert.Equal(t, "contact_list", result.(*responses.ContactList).Object)
			},
		},
		{
			name:     "get list",
			method:   http.MethodGet,
			path:     listPath,
			status:   http.StatusOK,
			response: listResponseFixture,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.GetList(context.Background(), accountID, listID)
			},
			assertResult: func(t *testing.T, result any) {
				assert.Equal(t, "Newsletter", result.(*responses.ContactList).Name)
			},
		},
		{
			name:     "update list",
			method:   http.MethodPut,
			path:     listPath,
			status:   http.StatusOK,
			response: listResponseFixture,
			wantBody: `{"description":"","tags":[]}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.UpdateList(context.Background(), accountID, listID, requests.UpdateListRequest{
					Description: ahasend.String(""),
					Tags:        &[]string{},
				})
			},
		},
		{
			name:     "delete list",
			method:   http.MethodDelete,
			path:     listPath,
			status:   http.StatusOK,
			response: `{"message":"list deleted"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.DeleteList(context.Background(), accountID, listID)
			},
			assertResult: func(t *testing.T, result any) {
				assert.Equal(t, "list deleted", result.(*common.SuccessResponse).Message)
			},
		},
		{
			name:      "get list contacts",
			method:    http.MethodGet,
			path:      listPath + "/contacts",
			status:    http.StatusOK,
			response:  `{"object":"list","data":[` + listContactResponseFixture[:len(listContactResponseFixture)-1] + `,"contact":` + contactResponseFixture + `}],"pagination":{"has_more":false,"next_cursor":null,"previous_cursor":null}}`,
			wantQuery: url.Values{"limit": {"100"}, "before": {"before-cursor"}, "subscription_status": {"unsubscribed"}, "email": {"person@example.com"}, "include_contacts": {"true"}},
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.GetListContacts(context.Background(), accountID, listID, requests.GetListContactsParams{
					Before:             ahasend.String("before-cursor"),
					SubscriptionStatus: ahasend.String(requests.ListContactStatusUnsubscribed),
					Email:              ahasend.String("person@example.com"),
					IncludeContacts:    ahasend.Bool(true),
				})
			},
			assertResult: func(t *testing.T, result any) {
				page := result.(*responses.PaginatedListContactsResponse)
				require.Len(t, page.Data, 1)
				membership := page.Data[0]
				assert.Equal(t, contactID, membership.ContactID)
				assert.Equal(t, requests.ListContactStatusUnsubscribed, membership.SubscriptionStatus)
				require.NotNil(t, membership.UnsubscribedAt)
				require.NotNil(t, membership.Contact)
				assert.Equal(t, contactID, membership.Contact.ID)
				assert.Nil(t, membership.List)
			},
		},
		{
			name:           "batch add list contacts",
			method:         http.MethodPost,
			path:           listPath + "/contacts/batch",
			status:         http.StatusOK,
			response:       `{"object":"list","added":0,"skipped":1,"failed":2,"data":[{"position":0,"id":"11111111-1111-4111-8111-111111111111","outcome":"already_member","membership":` + listContactResponseFixture + `},{"position":1,"email":"nobody@example.com","outcome":"not_found","reason":"contact not found"},{"position":2,"outcome":"invalid","reason":"entry must name exactly one of email or id"}]}`,
			idempotencyKey: "batch-list-contacts",
			wantBody:       `{"data":[{"id":"11111111-1111-4111-8111-111111111111"},{"email":"nobody@example.com"}]}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.BatchAddListContacts(context.Background(), accountID, listID, requests.BatchAddListContactsRequest{
					Data: []requests.BatchAddListContactInput{{ID: &contactID}, {Email: ahasend.String("nobody@example.com")}},
				}, WithIdempotencyKey("batch-list-contacts"))
			},
			assertResult: func(t *testing.T, result any) {
				batch := result.(*responses.BatchAddListContactsResponse)
				assert.Equal(t, 0, batch.Added)
				assert.Equal(t, 1, batch.Skipped)
				assert.Equal(t, 2, batch.Failed)
				require.Len(t, batch.Data, 3)
				assert.Equal(t, responses.BatchListContactOutcomeAlreadyMember, batch.Data[0].Outcome)
				require.NotNil(t, batch.Data[0].ID)
				assert.Equal(t, contactID, *batch.Data[0].ID)
				require.NotNil(t, batch.Data[0].Membership)
				assert.Equal(t, requests.ListContactStatusUnsubscribed, batch.Data[0].Membership.SubscriptionStatus)
				assert.Equal(t, responses.BatchListContactOutcomeNotFound, batch.Data[1].Outcome)
				assert.Equal(t, "nobody@example.com", batch.Data[1].Email)
				assert.Equal(t, 2, batch.Data[2].Position)
				assert.Equal(t, responses.BatchListContactOutcomeInvalid, batch.Data[2].Outcome)
				assert.Empty(t, batch.Data[2].Email)
				assert.Nil(t, batch.Data[2].ID)
				assert.NotEmpty(t, batch.Data[2].Reason)
			},
		},
		{
			name:     "upsert list contact",
			method:   http.MethodPut,
			path:     listPath + "/contacts/" + escapedEmail,
			status:   http.StatusOK,
			response: listContactResponseFixture,
			wantBody: `{"subscription_status":"unsubscribed"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.UpsertListContact(context.Background(), accountID, listID, rawEmail, requests.UpsertListContactRequest{
					SubscriptionStatus: ahasend.String(requests.ListContactStatusUnsubscribed),
				})
			},
		},
		{
			name:     "upsert list contact naming no status",
			method:   http.MethodPut,
			path:     listPath + "/contacts/" + contactID.String(),
			status:   http.StatusOK,
			response: listContactResponseFixture,
			wantBody: `{}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.UpsertListContact(context.Background(), accountID, listID, contactID.String(), requests.UpsertListContactRequest{})
			},
		},
		{
			name:     "delete list contact",
			method:   http.MethodDelete,
			path:     listPath + "/contacts/" + escapedEmail,
			status:   http.StatusOK,
			response: `{"message":"contact removed from list"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.DeleteListContact(context.Background(), accountID, listID, rawEmail)
			},
		},
		{
			name:      "get contact lists",
			method:    http.MethodGet,
			path:      "/v2/accounts/" + accountID.String() + "/contacts/" + escapedEmail + "/lists",
			status:    http.StatusOK,
			response:  `{"object":"list","data":[` + listContactResponseFixture[:len(listContactResponseFixture)-1] + `,"list":{"object":"contact_list","id":"22222222-2222-4222-8222-222222222222","created_at":"2026-09-10T10:00:00Z","updated_at":"2026-09-10T11:00:00Z","name":"Newsletter","description":"","tags":["news"]}}],"pagination":{"has_more":false,"next_cursor":null,"previous_cursor":null}}`,
			wantQuery: url.Values{"limit": {"10"}, "subscription_status": {"complained"}},
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.ListsAPI.GetContactLists(context.Background(), accountID, rawEmail, requests.GetContactListsParams{
					Limit:              ahasend.Int32(10),
					SubscriptionStatus: ahasend.String(requests.ListContactStatusComplained),
				})
			},
			assertResult: func(t *testing.T, result any) {
				page := result.(*responses.PaginatedListContactsResponse)
				require.Len(t, page.Data, 1)
				require.NotNil(t, page.Data[0].List)
				assert.Equal(t, listID, page.Data[0].List.ID)
				assert.Equal(t, []string{"news"}, page.Data[0].List.Tags)
				assert.Nil(t, page.Data[0].Contact)
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
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.response))
			})
			defer cleanup()

			result, httpResponse, err := tt.call(client)

			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, httpResponse)
			require.NotNil(t, capturedRequest)
			assert.Equal(t, tt.status, httpResponse.StatusCode)
			assert.Equal(t, tt.method, capturedRequest.Method)
			assert.Equal(t, tt.path, capturedRequest.URL.EscapedPath())
			assert.Equal(t, tt.idempotencyKey, capturedRequest.Header.Get("Idempotency-Key"))
			if tt.wantQuery != nil {
				assert.Equal(t, tt.wantQuery, capturedRequest.URL.Query())
			} else {
				assert.Empty(t, capturedRequest.URL.RawQuery)
			}
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, string(capturedBody))
			} else {
				assert.Empty(t, capturedBody)
			}
			if tt.assertResult != nil {
				tt.assertResult(t, result)
			}
		})
	}
}

func TestListsAPIPreservesComplainedConflict(t *testing.T) {
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"message":"contact reported this list as spam and their membership cannot be changed"}`))
	})
	defer cleanup()

	calls := map[string]func() (*http.Response, error){
		"upsert": func() (*http.Response, error) {
			_, httpResponse, err := client.ListsAPI.UpsertListContact(context.Background(), uuid.New(), uuid.New(), "person@example.com", requests.UpsertListContactRequest{
				SubscriptionStatus: ahasend.String(requests.ListContactStatusConfirmed),
			})
			return httpResponse, err
		},
		"delete": func() (*http.Response, error) {
			_, httpResponse, err := client.ListsAPI.DeleteListContact(context.Background(), uuid.New(), uuid.New(), "person@example.com")
			return httpResponse, err
		},
	}

	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			httpResponse, err := call()

			require.Error(t, err)
			require.NotNil(t, httpResponse)
			assert.Equal(t, http.StatusConflict, httpResponse.StatusCode)
			var apiErr *APIError
			require.ErrorAs(t, err, &apiErr)
			assert.Equal(t, http.StatusConflict, apiErr.StatusCode)
			assert.Contains(t, apiErr.Message, "cannot be changed")
		})
	}
}

func TestOpenAPIListOperationsAndResponses(t *testing.T) {
	type responseContract struct {
		Ref     string         `yaml:"$ref"`
		Headers map[string]any `yaml:"headers"`
	}
	type operationContract struct {
		OperationID string                      `yaml:"operationId"`
		Security    []map[string][]string       `yaml:"security"`
		Responses   map[string]responseContract `yaml:"responses"`
	}
	type openAPIContract struct {
		Paths map[string]map[string]operationContract `yaml:"paths"`
	}

	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	var spec openAPIContract
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	const (
		lists        = "/v2/accounts/{account_id}/lists"
		list         = "/v2/accounts/{account_id}/lists/{list_id}"
		listContacts = "/v2/accounts/{account_id}/lists/{list_id}/contacts"
		listBatch    = "/v2/accounts/{account_id}/lists/{list_id}/contacts/batch"
		listContact  = "/v2/accounts/{account_id}/lists/{list_id}/contacts/{id_or_email}"
		contactLists = "/v2/accounts/{account_id}/contacts/{id_or_email}/lists"
	)
	tests := []struct {
		path        string
		method      string
		operationID string
		scope       string
		statuses    []string
	}{
		{path: lists, method: "get", operationID: "getLists", scope: "lists:read", statuses: []string{"200", "400", "401", "403", "429", "500"}},
		{path: lists, method: "post", operationID: "createList", scope: "lists:write", statuses: []string{"201", "400", "401", "403", "409", "422", "429", "500"}},
		{path: list, method: "get", operationID: "getList", scope: "lists:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: list, method: "put", operationID: "updateList", scope: "lists:write", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: list, method: "delete", operationID: "deleteList", scope: "lists:delete", statuses: []string{"200", "400", "401", "403", "404", "409", "429", "500"}},
		{path: listContacts, method: "get", operationID: "getListContacts", scope: "lists:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: listBatch, method: "post", operationID: "batchAddListContacts", scope: "lists:write", statuses: []string{"200", "400", "401", "403", "404", "409", "422", "429", "500"}},
		{path: listContact, method: "put", operationID: "upsertListContact", scope: "lists:write", statuses: []string{"200", "400", "401", "403", "404", "409", "429", "500"}},
		{path: listContact, method: "delete", operationID: "deleteListContact", scope: "lists:write", statuses: []string{"200", "400", "401", "403", "404", "409", "429", "500"}},
		{path: contactLists, method: "get", operationID: "getContactLists", scope: "lists:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
	}

	for _, tt := range tests {
		t.Run(tt.operationID, func(t *testing.T) {
			path, ok := spec.Paths[tt.path]
			require.True(t, ok)
			operation, ok := path[tt.method]
			require.True(t, ok)
			assert.Equal(t, tt.operationID, operation.OperationID)
			assert.Equal(t, []map[string][]string{{"BearerAuth": {tt.scope}}}, operation.Security)

			actualStatuses := make([]string, 0, len(operation.Responses))
			for status := range operation.Responses {
				actualStatuses = append(actualStatuses, status)
			}
			sort.Strings(actualStatuses)
			assert.Equal(t, tt.statuses, actualStatuses)
		})
	}

	create := spec.Paths[lists]["post"].Responses
	assert.Contains(t, create["201"].Headers, "Idempotent-Replayed")
	assert.Equal(t, "#/components/responses/IdempotencyConflict", create["409"].Ref)
	assert.Equal(t, "#/components/responses/IdempotencyPayloadMismatch", create["422"].Ref)

	batch := spec.Paths[listBatch]["post"].Responses
	assert.Contains(t, batch["200"].Headers, "Idempotent-Replayed")
	assert.Equal(t, "#/components/responses/IdempotencyConflict", batch["409"].Ref)
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

func TestSuppressionsAPIGetSuppressionsFallsBackToDeprecatedDates(t *testing.T) {
	var lastRequest *http.Request
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	})
	defer cleanup()

	fromDate := time.Date(2026, 4, 18, 10, 0, 0, 0, time.UTC)
	toDate := fromDate.Add(time.Hour)

	t.Run("deprecated fields alone are sent as the time filters", func(t *testing.T) {
		_, _, err := client.SuppressionsAPI.GetSuppressions(context.Background(), uuid.New(), requests.GetSuppressionsParams{
			FromDate: &fromDate,
			ToDate:   &toDate,
		})

		require.NoError(t, err)
		require.NotNil(t, lastRequest)
		query := lastRequest.URL.Query()
		assert.Equal(t, fromDate.Format(time.RFC3339), query.Get("from_time"))
		assert.Equal(t, toDate.Format(time.RFC3339), query.Get("to_time"))
		assert.Empty(t, query.Get("from_date"))
		assert.Empty(t, query.Get("to_date"))
	})

	t.Run("FromTime and ToTime win over the deprecated fields", func(t *testing.T) {
		fromTime := fromDate.Add(24 * time.Hour)
		toTime := fromTime.Add(time.Hour)

		_, _, err := client.SuppressionsAPI.GetSuppressions(context.Background(), uuid.New(), requests.GetSuppressionsParams{
			FromTime: &fromTime,
			ToTime:   &toTime,
			FromDate: &fromDate,
			ToDate:   &toDate,
		})

		require.NoError(t, err)
		require.NotNil(t, lastRequest)
		query := lastRequest.URL.Query()
		assert.Equal(t, fromTime.Format(time.RFC3339), query.Get("from_time"))
		assert.Equal(t, toTime.Format(time.RFC3339), query.Get("to_time"))
	})
}

func TestDomainsAPIGetDomainsWithParamsSerializesFilters(t *testing.T) {
	var lastRequest *http.Request
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	})
	defer cleanup()

	t.Run("sending type and DNS validity", func(t *testing.T) {
		_, _, err := client.DomainsAPI.GetDomainsWithParams(context.Background(), uuid.New(), requests.GetDomainsParams{
			DNSValid:    ahasend.Bool(true),
			SendingType: ahasend.String(common.DomainSendingTypeMarketing),
			PaginationParams: common.PaginationParams{
				Limit: ahasend.Int32(25),
				After: ahasend.String("after-cursor"),
			},
		})

		require.NoError(t, err)
		require.NotNil(t, lastRequest)
		query := lastRequest.URL.Query()
		assert.Equal(t, "marketing", query.Get("sending_type"))
		assert.Equal(t, "true", query.Get("dns_valid"))
		assert.Equal(t, "25", query.Get("limit"))
		assert.Equal(t, "after-cursor", query.Get("after"))
	})

	t.Run("no filters", func(t *testing.T) {
		_, _, err := client.DomainsAPI.GetDomainsWithParams(context.Background(), uuid.New(), requests.GetDomainsParams{})

		require.NoError(t, err)
		require.NotNil(t, lastRequest)
		query := lastRequest.URL.Query()
		assert.NotContains(t, query, "sending_type")
		assert.NotContains(t, query, "dns_valid")
		assert.Equal(t, "100", query.Get("limit"))
	})

	t.Run("GetDomains keeps its DNS validity filter and pagination", func(t *testing.T) {
		_, _, err := client.DomainsAPI.GetDomains(context.Background(), uuid.New(), ahasend.Bool(false), &common.PaginationParams{
			Before: ahasend.String("before-cursor"),
		})

		require.NoError(t, err)
		require.NotNil(t, lastRequest)
		query := lastRequest.URL.Query()
		assert.Equal(t, "false", query.Get("dns_valid"))
		assert.Equal(t, "100", query.Get("limit"))
		assert.Equal(t, "before-cursor", query.Get("before"))
		assert.NotContains(t, query, "sending_type")
	})
}

const unpausedMarketingDomainResponseFixture = `{
	"object":"domain",
	"id":"33333333-3333-4333-8333-333333333333",
	"created_at":"2026-10-01T10:00:00Z",
	"updated_at":"2026-10-01T11:00:00Z",
	"domain":"mail.example.com",
	"account_id":"44444444-4444-4444-8444-444444444444",
	"dns_records":[],
	"last_dns_check_at":null,
	"dns_valid":true,
	"tracking_subdomain":null,
	"return_path_subdomain":null,
	"subscription_subdomain":null,
	"media_subdomain":null,
	"dkim_rotation_interval_days":null,
	"dkim_selector":null,
	"rotation_ready":false,
	"dsn_recipient":null,
	"sending_type":"marketing",
	"paused":false,
	"paused_at":null,
	"pause_reason":null
}`

func TestSubAccountsAPIUnpauseSubAccountDomain(t *testing.T) {
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	subAccountID := uuid.MustParse("bbbbbbbb-bbbb-4bbb-8bbb-bbbbbbbbbbbb")
	basePath := "/v2/accounts/" + accountID.String() + "/sub-accounts/" + subAccountID.String() + "/domains/"

	tests := []struct {
		name        string
		domain      string
		escapedPath string
	}{
		{name: "hostname", domain: "mail.example.com", escapedPath: basePath + "mail.example.com/unpause"},
		{name: "a slash stays inside the domain segment", domain: "mail.example.com/x", escapedPath: basePath + "mail.example.com%2Fx/unpause"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotMethod, gotPath, gotIdempotencyKey string
			var gotBody []byte
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.EscapedPath()
				gotIdempotencyKey = r.Header.Get("Idempotency-Key")
				gotBody, _ = io.ReadAll(r.Body)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(unpausedMarketingDomainResponseFixture))
			})
			defer cleanup()

			resp, httpResp, err := client.SubAccountsAPI.UnpauseSubAccountDomain(context.Background(), accountID, subAccountID, tt.domain)

			require.NoError(t, err)
			require.NotNil(t, httpResp)
			assert.Equal(t, http.StatusOK, httpResp.StatusCode)
			assert.Equal(t, http.MethodPost, gotMethod)
			assert.Equal(t, tt.escapedPath, gotPath)
			assert.NotEmpty(t, gotIdempotencyKey)
			assert.Empty(t, gotBody)

			require.NotNil(t, resp)
			assert.Equal(t, "mail.example.com", resp.Domain)
			assert.Equal(t, common.DomainSendingTypeMarketing, resp.SendingType)
			assert.False(t, resp.Paused)
			assert.Nil(t, resp.PausedAt)
			assert.Nil(t, resp.PauseReason)
		})
	}

	t.Run("an empty domain is refused before the request", func(t *testing.T) {
		var called bool
		client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			called = true
		})
		defer cleanup()

		_, _, err := client.SubAccountsAPI.UnpauseSubAccountDomain(context.Background(), accountID, subAccountID, "")

		require.Error(t, err)
		assert.False(t, called)
	})

	t.Run("a refused unpause is an APIError", func(t *testing.T) {
		client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"domain not found"}`))
		})
		defer cleanup()

		_, httpResp, err := client.SubAccountsAPI.UnpauseSubAccountDomain(context.Background(), accountID, subAccountID, "mail.example.com")

		require.Error(t, err)
		require.NotNil(t, httpResp)
		var apiErr *APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, ErrorTypeNotFound, apiErr.Type)
		assert.Equal(t, "domain not found", apiErr.Message)
	})
}

func TestOpenAPIUnpauseSubAccountDomainOperation(t *testing.T) {
	spec := loadOpenAPIContract(t)

	path, ok := spec.Paths["/v2/accounts/{account_id}/sub-accounts/{sub_account_id}/domains/{domain}/unpause"]
	require.True(t, ok)
	operation, ok := path["post"]
	require.True(t, ok)
	assert.Equal(t, "unpauseSubAccountDomain", operation.OperationID)

	var goSamples int
	for _, sample := range operation.CodeSamples {
		if sample.Lang == "go" {
			goSamples++
		}
	}
	assert.Equal(t, 1, goSamples)

	assert.Equal(t, []map[string][]string{{"BearerAuth": {"sub-accounts:suspend"}}}, operation.Security)

	actualStatuses := make([]string, 0, len(operation.Responses))
	for status := range operation.Responses {
		actualStatuses = append(actualStatuses, status)
	}
	sort.Strings(actualStatuses)
	assert.Equal(t, []string{"200", "400", "401", "403", "404", "500"}, actualStatuses)
}

// openAPISchemaProperties returns the sorted property names of a schema in
// the bundled specification, including those of its inline allOf branches.
// A $ref branch adds nothing: name the referenced schema as well.
func openAPISchemaProperties(t *testing.T, schemaNames ...string) []string {
	t.Helper()

	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	type schemaProperties struct {
		Properties map[string]any `yaml:"properties"`
	}
	var spec struct {
		Components struct {
			Schemas map[string]struct {
				schemaProperties `yaml:",inline"`
				AllOf            []schemaProperties `yaml:"allOf"`
			} `yaml:"schemas"`
		} `yaml:"components"`
	}
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	var names []string
	for _, schemaName := range schemaNames {
		schema, ok := spec.Components.Schemas[schemaName]
		require.True(t, ok, "schema %s", schemaName)
		for _, part := range append([]schemaProperties{schema.schemaProperties}, schema.AllOf...) {
			for name := range part.Properties {
				names = append(names, name)
			}
		}
	}
	sort.Strings(names)
	return names
}

// modelJSONFields returns the sorted JSON names of a struct's fields,
// including those of the structs it embeds.
func modelJSONFields(model any) []string {
	var names []string
	var collect func(reflect.Type)
	collect = func(modelType reflect.Type) {
		for i := 0; i < modelType.NumField(); i++ {
			field := modelType.Field(i)
			if field.Anonymous {
				collect(field.Type)
				continue
			}
			name, _, _ := strings.Cut(field.Tag.Get("json"), ",")
			names = append(names, name)
		}
	}
	collect(reflect.TypeOf(model))
	sort.Strings(names)
	return names
}

// TestOpenAPIDomainSchemaMatchesModel fails when the bundled specification
// gives Domain a property that responses.Domain cannot decode.
func TestOpenAPIDomainSchemaMatchesModel(t *testing.T) {
	assert.Equal(t, openAPISchemaProperties(t, "Domain"), modelJSONFields(responses.Domain{}))
}

func TestOpenAPIMessageSchemaMatchesModel(t *testing.T) {
	// Message is an allOf of a $ref to MessageSummary and the content fields.
	assert.Equal(t, openAPISchemaProperties(t, "Message", "MessageSummary"), modelJSONFields(responses.Message{}))
}

func TestOpenAPITemplateSchemaMatchesModel(t *testing.T) {
	assert.Equal(t, openAPISchemaProperties(t, "Template"), modelJSONFields(responses.Template{}))
}

// TestOpenAPIWebhookSchemasMatchModels fails when the bundled specification
// gives a webhook or a webhook request a property its model cannot carry.
// responses.Webhook also decodes the secret that only CreatedWebhook adds.
func TestOpenAPIWebhookSchemasMatchModels(t *testing.T) {
	assert.Equal(t, openAPISchemaProperties(t, "CreatedWebhook", "Webhook"), modelJSONFields(responses.Webhook{}))
	assert.Equal(t, openAPISchemaProperties(t, "CreateWebhookRequest"), modelJSONFields(requests.CreateWebhookRequest{}))
	assert.Equal(t, openAPISchemaProperties(t, "UpdateWebhookRequest"), modelJSONFields(requests.UpdateWebhookRequest{}))
}

// TestWebhooksAPIGetWebhooksSendsEveryEventFilter sets every event filter and
// checks that the request carries each event filter the specification lists.
func TestWebhooksAPIGetWebhooksSendsEveryEventFilter(t *testing.T) {
	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)
	var spec struct {
		Paths map[string]map[string]struct {
			Parameters []struct {
				Name string `yaml:"name"`
				In   string `yaml:"in"`
			} `yaml:"parameters"`
		} `yaml:"paths"`
	}
	require.NoError(t, yaml.Unmarshal(contents, &spec))
	var filters []string
	for _, parameter := range spec.Paths["/v2/accounts/{account_id}/webhooks"]["get"].Parameters {
		if parameter.In == "query" && strings.HasPrefix(parameter.Name, "on_") {
			filters = append(filters, parameter.Name)
		}
	}
	require.NotEmpty(t, filters)

	var lastRequest *http.Request
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
	})
	defer cleanup()

	params := GetWebhooksParams{}
	value := reflect.ValueOf(&params).Elem()
	for i := 0; i < value.NumField(); i++ {
		if field := value.Field(i); strings.HasPrefix(value.Type().Field(i).Name, "On") {
			field.Set(reflect.ValueOf(ahasend.Bool(true)))
		}
	}
	_, _, err = client.WebhooksAPI.GetWebhooks(context.Background(), uuid.New(), params)
	require.NoError(t, err)
	require.NotNil(t, lastRequest)

	var sent []string
	for name := range lastRequest.URL.Query() {
		if strings.HasPrefix(name, "on_") {
			assert.Equal(t, "true", lastRequest.URL.Query().Get(name), name)
			sent = append(sent, name)
		}
	}
	assert.ElementsMatch(t, filters, sent)
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

func TestMessagesAPICreateTemplateMessage(t *testing.T) {
	accountID := uuid.New()
	templateID := uuid.New()
	var gotMethod, gotPath, gotIdempotencyKey string
	var body map[string]interface{}
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotIdempotencyKey = r.Header.Get("Idempotency-Key")
		err := json.NewDecoder(r.Body).Decode(&body)
		require.NoError(t, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte(`{"object":"list","data":[]}`))
	})
	defer cleanup()

	_, _, err := client.MessagesAPI.CreateTemplateMessage(context.Background(), accountID, requests.CreateTemplateMessageRequest{
		TemplateID: templateID,
		Recipients: []common.Recipient{{Email: "recipient@example.com"}},
	}, WithIdempotencyKey("template-send"))

	require.NoError(t, err)
	assert.Equal(t, http.MethodPost, gotMethod)
	assert.Equal(t, "/v2/accounts/"+accountID.String()+"/messages/template", gotPath)
	assert.Equal(t, "template-send", gotIdempotencyKey)
	assert.Equal(t, templateID.String(), body["template_id"])
	// A request with no from sends from the template's default sender.
	assert.NotContains(t, body, "from")
}

// The send operations check every message before it is sent. When the check
// cannot run, the API answers 503 with a Retry-After and asks for the same
// key on the retry; the client retries it as it retries any 503.
func TestSendsRetryABusyTemplateCheckWithTheirIdempotencyKey(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		response string
		call     func(client *APIClient) (*http.Response, error)
	}{
		{
			name:     "create message",
			status:   http.StatusAccepted,
			response: `{"object":"list","data":[]}`,
			call: func(client *APIClient) (*http.Response, error) {
				_, httpResponse, err := client.MessagesAPI.CreateMessage(context.Background(), uuid.New(), requests.CreateMessageRequest{
					From:        common.SenderAddress{Email: "sender@example.com"},
					Recipients:  []common.Recipient{{Email: "recipient@example.com"}},
					Subject:     "Hello",
					TextContent: ahasend.String("Hello"),
				})
				return httpResponse, err
			},
		},
		{
			name:     "create template message",
			status:   http.StatusAccepted,
			response: `{"object":"list","data":[]}`,
			call: func(client *APIClient) (*http.Response, error) {
				_, httpResponse, err := client.MessagesAPI.CreateTemplateMessage(context.Background(), uuid.New(), requests.CreateTemplateMessageRequest{
					TemplateID: uuid.New(),
					Recipients: []common.Recipient{{Email: "recipient@example.com"}},
				})
				return httpResponse, err
			},
		},
		{
			name:     "create conversation message",
			status:   http.StatusAccepted,
			response: `{"object":"list","data":[]}`,
			call: func(client *APIClient) (*http.Response, error) {
				_, httpResponse, err := client.MessagesAPI.CreateConversationMessage(context.Background(), uuid.New(), requests.CreateConversationMessageRequest{
					From:    common.SenderAddress{Email: "sender@example.com"},
					To:      []common.SenderAddress{{Email: "recipient@example.com"}},
					Subject: "Hello",
				})
				return httpResponse, err
			},
		},
		{
			name:     "test campaign",
			status:   http.StatusOK,
			response: `{"object":"campaign_test","sent_to":["ann@example.com"],"failed":[],"result_unknown":false}`,
			call: func(client *APIClient) (*http.Response, error) {
				_, httpResponse, err := client.CampaignsAPI.TestCampaign(context.Background(), uuid.New(), uuid.New(), requests.TestCampaignRequest{
					To: []string{"ann@example.com"},
				})
				return httpResponse, err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var attempts int
			var keys []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempts++
				keys = append(keys, r.Header.Get("Idempotency-Key"))
				w.Header().Set("Content-Type", "application/json")
				if attempts == 1 {
					w.Header().Set("Retry-After", "1")
					w.WriteHeader(http.StatusServiceUnavailable)
					_, _ = w.Write([]byte(`{"message":"the message could not be checked right now; try again in a moment"}`))
					return
				}
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()
			serverURL, err := url.Parse(server.URL)
			require.NoError(t, err)

			cfg := NewConfiguration()
			cfg.Host = serverURL.Host
			cfg.Scheme = serverURL.Scheme
			cfg.APIKey = "test-key"
			cfg.RetryConfig.Enabled = true
			cfg.RetryConfig.MaxRetries = 2
			cfg.RetryConfig.BaseDelay = time.Millisecond
			// Above the server's Retry-After of 1s, so the 503 is retried
			cfg.RetryConfig.MaxDelay = 2 * time.Second
			client := NewAPIClientWithConfig(cfg)

			httpResponse, err := tt.call(client)

			require.NoError(t, err)
			require.NotNil(t, httpResponse)
			assert.Equal(t, tt.status, httpResponse.StatusCode)
			require.Len(t, keys, 2)
			assert.NotEmpty(t, keys[0])
			assert.Equal(t, keys[0], keys[1])
		})
	}
}

func TestOpenAPISendRequestSchemasMatchModels(t *testing.T) {
	assert.Equal(t, openAPISchemaProperties(t, "CreateMessageRequest"), modelJSONFields(requests.CreateMessageRequest{}))
	assert.Equal(t, openAPISchemaProperties(t, "CreateTemplateMessageRequest"), modelJSONFields(requests.CreateTemplateMessageRequest{}))
}

func TestOpenAPISendRequestSender(t *testing.T) {
	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	type sendRequestSchema struct {
		Required   []string `yaml:"required"`
		Properties struct {
			From struct {
				Ref string `yaml:"$ref"`
			} `yaml:"from"`
		} `yaml:"properties"`
	}
	var spec struct {
		Components struct {
			Schemas struct {
				CreateMessageRequest         sendRequestSchema `yaml:"CreateMessageRequest"`
				CreateTemplateMessageRequest sendRequestSchema `yaml:"CreateTemplateMessageRequest"`
			} `yaml:"schemas"`
		} `yaml:"components"`
	}
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	inline := spec.Components.Schemas.CreateMessageRequest
	assert.Contains(t, inline.Required, "from")
	assert.Equal(t, "#/components/schemas/Address", inline.Properties.From.Ref)

	template := spec.Components.Schemas.CreateTemplateMessageRequest
	assert.NotContains(t, template.Required, "from")
	assert.Equal(t, "#/components/schemas/Address", template.Properties.From.Ref)
}

func TestOpenAPICreateTemplateMessageOperation(t *testing.T) {
	spec := loadOpenAPIContract(t)

	path, ok := spec.Paths["/v2/accounts/{account_id}/messages/template"]
	require.True(t, ok)
	operation, ok := path["post"]
	require.True(t, ok)
	assert.Equal(t, "createTemplateMessage", operation.OperationID)

	var goSamples int
	for _, sample := range operation.CodeSamples {
		if sample.Lang == "go" {
			goSamples++
		}
	}
	assert.Equal(t, 1, goSamples)

	assert.Equal(t, []map[string][]string{
		{"BearerAuth": {"messages:send:all"}},
		{"BearerAuth": {"messages:send:{domain}"}},
	}, operation.Security)

	actualStatuses := make([]string, 0, len(operation.Responses))
	for status := range operation.Responses {
		actualStatuses = append(actualStatuses, status)
	}
	sort.Strings(actualStatuses)
	assert.Equal(t, []string{"202", "400", "401", "403", "404", "409", "422", "500", "503"}, actualStatuses)
	assert.Contains(t, operation.Responses["202"].Headers, "Idempotent-Replayed")
}

func TestOpenAPITemplateCheckBusyResponses(t *testing.T) {
	spec := loadOpenAPIContract(t)

	tests := []struct {
		path        string
		operationID string
	}{
		{path: "/v2/accounts/{account_id}/messages", operationID: "createMessage"},
		{path: "/v2/accounts/{account_id}/messages/template", operationID: "createTemplateMessage"},
		{path: "/v2/accounts/{account_id}/messages/conversation", operationID: "createConversationMessage"},
		{path: "/v2/accounts/{account_id}/campaigns/{campaign_id}/test", operationID: "testCampaign"},
	}
	for _, tt := range tests {
		t.Run(tt.operationID, func(t *testing.T) {
			operation, ok := spec.Paths[tt.path]["post"]
			require.True(t, ok)
			assert.Equal(t, tt.operationID, operation.OperationID)
			assert.Equal(t, "#/components/responses/TemplateCheckBusy", operation.Responses["503"].Ref)
		})
	}

	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)
	var components struct {
		Components struct {
			Responses map[string]responseContract `yaml:"responses"`
		} `yaml:"components"`
	}
	require.NoError(t, yaml.Unmarshal(contents, &components))
	assert.Contains(t, components.Components.Responses["TemplateCheckBusy"].Headers, "Retry-After")
}

func TestOpenAPITemplateOperationsAndResponses(t *testing.T) {
	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	type templateOperationContract struct {
		operationContract `yaml:",inline"`
		Parameters        []struct {
			Ref string `yaml:"$ref"`
		} `yaml:"parameters"`
	}
	var spec struct {
		Paths map[string]map[string]templateOperationContract `yaml:"paths"`
	}
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	const (
		templates = "/v2/accounts/{account_id}/templates"
		template  = templates + "/{template_id}"
		draft     = template + "/draft"
		publish   = template + "/publish"
		versions  = template + "/versions"
		version   = versions + "/{version_id}"
		restore   = version + "/restore"
	)
	tests := []struct {
		path        string
		method      string
		operationID string
		scope       string
		idempotent  bool
		raced       bool
		statuses    []string
	}{
		{path: templates, method: "get", operationID: "listTemplates", scope: "templates:read", statuses: []string{"200", "400", "401", "403", "429", "500"}},
		{path: templates, method: "post", operationID: "createTemplate", scope: "templates:write", idempotent: true, statuses: []string{"201", "400", "401", "403", "409", "422", "429", "500"}},
		{path: template, method: "get", operationID: "getTemplate", scope: "templates:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: template, method: "put", operationID: "updateTemplate", scope: "templates:write", idempotent: true, raced: true, statuses: []string{"200", "400", "401", "403", "404", "409", "422", "429", "500", "503"}},
		{path: template, method: "delete", operationID: "deleteTemplate", scope: "templates:delete", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: draft, method: "get", operationID: "getTemplateDraft", scope: "templates:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: draft, method: "delete", operationID: "discardTemplateDraft", scope: "templates:write", raced: true, statuses: []string{"200", "400", "401", "403", "404", "429", "500", "503"}},
		{path: publish, method: "post", operationID: "publishTemplate", scope: "templates:write", idempotent: true, raced: true, statuses: []string{"200", "400", "401", "403", "404", "409", "422", "429", "500", "503"}},
		{path: versions, method: "get", operationID: "listTemplateVersions", scope: "templates:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: version, method: "get", operationID: "getTemplateVersion", scope: "templates:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: restore, method: "post", operationID: "restoreTemplateVersion", scope: "templates:write", idempotent: true, raced: true, statuses: []string{"200", "400", "401", "403", "404", "409", "422", "429", "500", "503"}},
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

			assert.Equal(t, []map[string][]string{{"BearerAuth": {tt.scope}}}, operation.Security)

			actualStatuses := make([]string, 0, len(operation.Responses))
			for status := range operation.Responses {
				actualStatuses = append(actualStatuses, status)
			}
			sort.Strings(actualStatuses)
			assert.Equal(t, tt.statuses, actualStatuses)

			var takesIdempotencyKey bool
			for _, parameter := range operation.Parameters {
				if parameter.Ref == "#/components/parameters/IdempotencyKey" {
					takesIdempotencyKey = true
				}
			}
			assert.Equal(t, tt.idempotent, takesIdempotencyKey)
			if tt.idempotent {
				for status, response := range operation.Responses {
					if strings.HasPrefix(status, "2") {
						assert.Contains(t, response.Headers, "Idempotent-Replayed")
					}
				}
				assert.Equal(t, "#/components/responses/IdempotencyConflict", operation.Responses["409"].Ref)
				assert.Equal(t, "#/components/responses/IdempotencyPayloadMismatch", operation.Responses["422"].Ref)
			}
			if tt.raced {
				assert.Equal(t, "#/components/responses/TemplateBeingChanged", operation.Responses["503"].Ref)
			}
		})
	}

	// A template_id naming no template of this account is the template
	// send's 404; an inline send names no stored template.
	assert.Contains(t, spec.Paths["/v2/accounts/{account_id}/messages/template"]["post"].Responses, "404")
	assert.NotContains(t, spec.Paths["/v2/accounts/{account_id}/messages"]["post"].Responses, "404")
}

func TestOpenAPITemplateSchemasMatchModels(t *testing.T) {
	tests := []struct {
		schemas []string
		model   any
	}{
		{schemas: []string{"TemplateContent"}, model: responses.TemplateContent{}},
		{schemas: []string{"TemplateDraft"}, model: responses.TemplateDraft{}},
		{schemas: []string{"TemplatePublisher"}, model: responses.TemplatePublisher{}},
		{schemas: []string{"TemplateVersion"}, model: responses.TemplateVersion{}},
		// TemplateVersionDetail is an allOf of a $ref to TemplateVersion and
		// the content fields.
		{schemas: []string{"TemplateVersionDetail", "TemplateVersion"}, model: responses.TemplateVersionDetail{}},
		{schemas: []string{"TemplateVersionsResponse"}, model: responses.TemplateVersionsResponse{}},
		{schemas: []string{"CreateTemplateRequest"}, model: requests.CreateTemplateRequest{}},
		{schemas: []string{"UpdateTemplateRequest"}, model: requests.UpdateTemplateRequest{}},
		{schemas: []string{"RestoreTemplateVersionRequest"}, model: requests.RestoreTemplateVersionRequest{}},
	}

	for _, tt := range tests {
		t.Run(tt.schemas[0], func(t *testing.T) {
			assert.Equal(t, openAPISchemaProperties(t, tt.schemas...), modelJSONFields(tt.model))
		})
	}

	// The API accepts and ignores text_is_custom in a write, so a read content
	// can be sent back; the SDK leaves it out.
	input := openAPISchemaProperties(t, "TemplateContentInput")
	assert.Equal(t, []string{"html", "mjml", "text", "text_is_custom"}, input)
	assert.Equal(t, []string{"html", "mjml", "text"}, modelJSONFields(requests.TemplateContentInput{}))
}

const templateWriteResponseFixture = `{
	"object":"template",
	"id":"11111111-1111-4111-8111-111111111111",
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-10T11:00:00Z",
	"name":"Password reset",
	"subject":"Reset your password",
	"preheader":"",
	"variables":[],
	"from":null,
	"reply_to":null,
	"editor":"html",
	"has_draft":true,
	"content":{"html":"<p>Hi</p>","text":"Hi","text_is_custom":false}
}`

func TestTemplatesAPIEmitsDeclaredTransport(t *testing.T) {
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	versionID := uuid.MustParse("22222222-2222-4222-8222-222222222222")
	keyID := uuid.MustParse("33333333-3333-4333-8333-333333333333")
	templatesPath := "/v2/accounts/" + accountID.String() + "/templates"
	templatePath := templatesPath + "/" + templateID.String()
	versionPath := templatePath + "/versions/" + versionID.String()
	versionFixture := `{"object":"template_version","id":"` + versionID.String() + `","version":2,"published_at":"2026-09-12T10:00:00Z","published_by":{"type":"api_key","id":"` + keyID.String() + `"}}`

	assertTemplate := func(t *testing.T, result any) {
		template := result.(*responses.Template)
		assert.Equal(t, templateID, template.ID)
		assert.Equal(t, common.TemplateEditorHTML, template.Editor)
		assert.True(t, template.HasDraft)
		require.NotNil(t, template.Content)
		assert.Equal(t, "<p>Hi</p>", template.Content.HTML)
	}

	type templateCall func(*APIClient) (any, *http.Response, error)
	tests := []struct {
		name           string
		method         string
		path           string
		status         int
		response       string
		idempotencyKey string
		wantBody       string
		call           templateCall
		assertResult   func(*testing.T, any)
	}{
		{
			name:           "create template",
			method:         http.MethodPost,
			path:           templatesPath,
			status:         http.StatusCreated,
			response:       templateWriteResponseFixture,
			idempotencyKey: "create-template",
			wantBody:       `{"name":"Password reset","subject":"Reset your password","content":{"html":"<p>Hi</p>"}}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.CreateTemplate(context.Background(), accountID, requests.CreateTemplateRequest{
					Name:    "Password reset",
					Subject: ahasend.String("Reset your password"),
					Content: &requests.TemplateContentInput{HTML: ahasend.String("<p>Hi</p>")},
				}, WithIdempotencyKey("create-template"))
			},
			assertResult: assertTemplate,
		},
		{
			name:           "update template",
			method:         http.MethodPut,
			path:           templatePath,
			status:         http.StatusOK,
			response:       templateWriteResponseFixture,
			idempotencyKey: "update-template",
			wantBody:       `{"preheader":"","from":null,"content":{"text":null},"publish":true}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.UpdateTemplate(context.Background(), accountID, templateID, requests.UpdateTemplateRequest{
					Preheader: ahasend.String(""),
					From:      &common.SenderAddress{},
					Content:   &requests.TemplateContentInput{Text: ahasend.String("")},
					Publish:   true,
				}, WithIdempotencyKey("update-template"))
			},
			assertResult: assertTemplate,
		},
		{
			name:     "delete template",
			method:   http.MethodDelete,
			path:     templatePath,
			status:   http.StatusOK,
			response: `{"message":"template deleted"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.DeleteTemplate(context.Background(), accountID, templateID)
			},
			assertResult: func(t *testing.T, result any) {
				assert.Equal(t, "template deleted", result.(*common.SuccessResponse).Message)
			},
		},
		{
			name:     "get template draft",
			method:   http.MethodGet,
			path:     templatePath + "/draft",
			status:   http.StatusOK,
			response: `{"object":"template_draft","template_id":"` + templateID.String() + `","updated_at":"2026-09-11T10:00:00Z","subject":"New subject","preheader":"","variables":[],"from":null,"reply_to":null,"content":null}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.GetTemplateDraft(context.Background(), accountID, templateID)
			},
			assertResult: func(t *testing.T, result any) {
				draft := result.(*responses.TemplateDraft)
				assert.Equal(t, templateID, draft.TemplateID)
				assert.Equal(t, "New subject", draft.Subject)
				assert.Nil(t, draft.Content)
			},
		},
		{
			name:     "discard template draft",
			method:   http.MethodDelete,
			path:     templatePath + "/draft",
			status:   http.StatusOK,
			response: templateWriteResponseFixture,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.DiscardTemplateDraft(context.Background(), accountID, templateID)
			},
			assertResult: assertTemplate,
		},
		{
			name:           "publish template",
			method:         http.MethodPost,
			path:           templatePath + "/publish",
			status:         http.StatusOK,
			response:       templateWriteResponseFixture,
			idempotencyKey: "publish-template",
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.PublishTemplate(context.Background(), accountID, templateID, WithIdempotencyKey("publish-template"))
			},
			assertResult: assertTemplate,
		},
		{
			name:     "get template versions",
			method:   http.MethodGet,
			path:     templatePath + "/versions",
			status:   http.StatusOK,
			response: `{"object":"list","data":[` + versionFixture + `]}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.GetTemplateVersions(context.Background(), accountID, templateID)
			},
			assertResult: func(t *testing.T, result any) {
				versions := result.(*responses.TemplateVersionsResponse)
				require.Len(t, versions.Data, 1)
				assert.Equal(t, versionID, versions.Data[0].ID)
				assert.Equal(t, 2, versions.Data[0].Version)
				require.NotNil(t, versions.Data[0].PublishedBy)
				assert.Equal(t, responses.TemplatePublisherTypeAPIKey, versions.Data[0].PublishedBy.Type)
				assert.Equal(t, keyID, versions.Data[0].PublishedBy.ID)
			},
		},
		{
			name:     "get template version",
			method:   http.MethodGet,
			path:     versionPath,
			status:   http.StatusOK,
			response: versionFixture[:len(versionFixture)-1] + `,"subject":"Old subject","preheader":"","variables":[],"from":null,"reply_to":null,"content":{"text":"Hi","text_is_custom":true}}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.GetTemplateVersion(context.Background(), accountID, templateID, versionID)
			},
			assertResult: func(t *testing.T, result any) {
				version := result.(*responses.TemplateVersionDetail)
				assert.Equal(t, versionID, version.ID)
				assert.Equal(t, "Old subject", version.Subject)
				require.NotNil(t, version.Content)
				assert.Equal(t, "Hi", version.Content.Text)
				assert.True(t, version.Content.TextIsCustom)
			},
		},
		{
			name:           "restore template version",
			method:         http.MethodPost,
			path:           versionPath + "/restore",
			status:         http.StatusOK,
			response:       templateWriteResponseFixture,
			idempotencyKey: "restore-template",
			wantBody:       `{"publish":true}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.RestoreTemplateVersion(context.Background(), accountID, templateID, versionID, requests.RestoreTemplateVersionRequest{Publish: true}, WithIdempotencyKey("restore-template"))
			},
			assertResult: assertTemplate,
		},
		{
			name:           "restore template version without publishing",
			method:         http.MethodPost,
			path:           versionPath + "/restore",
			status:         http.StatusOK,
			response:       templateWriteResponseFixture,
			idempotencyKey: "restore-template",
			wantBody:       `{}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.TemplatesAPI.RestoreTemplateVersion(context.Background(), accountID, templateID, versionID, requests.RestoreTemplateVersionRequest{}, WithIdempotencyKey("restore-template"))
			},
			assertResult: assertTemplate,
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
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.response))
			})
			defer cleanup()

			result, httpResponse, err := tt.call(client)

			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, httpResponse)
			require.NotNil(t, capturedRequest)
			assert.Equal(t, tt.status, httpResponse.StatusCode)
			assert.Equal(t, tt.method, capturedRequest.Method)
			assert.Equal(t, tt.path, capturedRequest.URL.EscapedPath())
			assert.Empty(t, capturedRequest.URL.RawQuery)
			if tt.idempotencyKey != "" {
				assert.Equal(t, tt.idempotencyKey, capturedRequest.Header.Get("Idempotency-Key"))
			}
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, string(capturedBody))
			} else {
				assert.Empty(t, capturedBody)
			}
			if tt.assertResult != nil {
				tt.assertResult(t, result)
			}
		})
	}
}

func TestTemplatesAPIRetriesARacedWriteWithItsIdempotencyKey(t *testing.T) {
	tests := []struct {
		name string
		call func(*APIClient) (*responses.Template, *http.Response, error)
	}{
		{
			name: "publish template",
			call: func(client *APIClient) (*responses.Template, *http.Response, error) {
				return client.TemplatesAPI.PublishTemplate(context.Background(), uuid.New(), uuid.New())
			},
		},
		{
			name: "update template",
			call: func(client *APIClient) (*responses.Template, *http.Response, error) {
				return client.TemplatesAPI.UpdateTemplate(context.Background(), uuid.New(), uuid.New(), requests.UpdateTemplateRequest{
					Subject: ahasend.String("Reset your password"),
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var attempts int
			var keys []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempts++
				keys = append(keys, r.Header.Get("Idempotency-Key"))
				w.Header().Set("Content-Type", "application/json")
				if attempts == 1 {
					w.Header().Set("Retry-After", "1")
					w.WriteHeader(http.StatusServiceUnavailable)
					_, _ = w.Write([]byte(`{"message":"template is being changed by another request, try again"}`))
					return
				}
				_, _ = w.Write([]byte(templateWriteResponseFixture))
			}))
			defer server.Close()
			serverURL, err := url.Parse(server.URL)
			require.NoError(t, err)

			cfg := NewConfiguration()
			cfg.Host = serverURL.Host
			cfg.Scheme = serverURL.Scheme
			cfg.APIKey = "test-key"
			cfg.RetryConfig.Enabled = true
			cfg.RetryConfig.MaxRetries = 2
			cfg.RetryConfig.BaseDelay = time.Millisecond
			// Above the server's Retry-After of 1s, so the 503 is retried
			cfg.RetryConfig.MaxDelay = 2 * time.Second
			client := NewAPIClientWithConfig(cfg)

			template, httpResponse, err := tt.call(client)

			require.NoError(t, err)
			require.NotNil(t, httpResponse)
			assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
			assert.Equal(t, "Password reset", template.Name)
			require.Len(t, keys, 2)
			assert.NotEmpty(t, keys[0])
			assert.Equal(t, keys[0], keys[1])
		})
	}
}

func TestOpenAPICampaignOperationsAndResponses(t *testing.T) {
	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	type campaignResponseContract struct {
		responseContract `yaml:",inline"`
		Content          map[string]struct {
			Schema struct {
				Ref string `yaml:"$ref"`
			} `yaml:"schema"`
		} `yaml:"content"`
	}
	type campaignOperationContract struct {
		OperationID string                              `yaml:"operationId"`
		CodeSamples []struct{ Lang string }             `yaml:"x-code-samples"`
		Security    []map[string][]string               `yaml:"security"`
		Responses   map[string]campaignResponseContract `yaml:"responses"`
		Parameters  []struct {
			Ref string `yaml:"$ref"`
		} `yaml:"parameters"`
	}
	var spec struct {
		Paths map[string]map[string]campaignOperationContract `yaml:"paths"`
	}
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	const (
		idempotencyConflict    = "#/components/responses/IdempotencyConflict"
		campaignConflict       = "#/components/responses/CampaignConflict"
		campaignCreateConflict = "#/components/responses/CampaignCreateConflict"
		campaignNotReady       = "#/components/schemas/CampaignNotReady"

		campaignTemplates = "/v2/accounts/{account_id}/campaign-templates"
		campaignTemplate  = campaignTemplates + "/{template_id}"
		campaigns         = "/v2/accounts/{account_id}/campaigns"
		campaign          = campaigns + "/{campaign_id}"
	)
	tests := []struct {
		path        string
		method      string
		operationID string
		scope       string
		idempotent  bool
		// conflict is the response a 409 of an idempotent operation refers
		// to.
		conflict string
		// badRequest is the schema of a 400 that is not an ErrorResponse.
		badRequest string
		statuses   []string
	}{
		{path: campaignTemplates, method: "get", operationID: "listCampaignTemplates", scope: "templates:read", statuses: []string{"200", "400", "401", "403", "429", "500"}},
		{path: campaignTemplates, method: "post", operationID: "createCampaignTemplate", scope: "templates:write", idempotent: true, conflict: idempotencyConflict, statuses: []string{"201", "400", "401", "403", "409", "422", "429", "500"}},
		{path: campaignTemplate, method: "get", operationID: "getCampaignTemplate", scope: "templates:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: campaignTemplate, method: "put", operationID: "updateCampaignTemplate", scope: "templates:write", idempotent: true, conflict: idempotencyConflict, statuses: []string{"200", "400", "401", "403", "404", "409", "422", "429", "500"}},
		{path: campaignTemplate, method: "delete", operationID: "deleteCampaignTemplate", scope: "templates:delete", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: campaigns, method: "get", operationID: "listCampaigns", scope: "campaigns:read", statuses: []string{"200", "400", "401", "403", "429", "500"}},
		{path: campaigns, method: "post", operationID: "createCampaign", scope: "campaigns:write", idempotent: true, conflict: campaignCreateConflict, statuses: []string{"201", "400", "401", "403", "409", "422", "429", "500"}},
		{path: campaign, method: "get", operationID: "getCampaign", scope: "campaigns:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: campaign, method: "put", operationID: "updateCampaign", scope: "campaigns:write", idempotent: true, conflict: campaignConflict, statuses: []string{"200", "400", "401", "403", "404", "409", "422", "429", "500"}},
		{path: campaign, method: "delete", operationID: "deleteCampaign", scope: "campaigns:delete", statuses: []string{"200", "400", "401", "403", "404", "409", "429", "500"}},
		{path: campaign + "/duplicate", method: "post", operationID: "duplicateCampaign", scope: "campaigns:write", idempotent: true, conflict: campaignCreateConflict, statuses: []string{"201", "400", "401", "403", "404", "409", "422", "429", "500"}},
		{path: campaign + "/send", method: "post", operationID: "sendCampaign", scope: "campaigns:send", idempotent: true, conflict: campaignConflict, badRequest: campaignNotReady, statuses: []string{"202", "400", "401", "403", "404", "409", "422", "429", "500"}},
		{path: campaign + "/unschedule", method: "post", operationID: "unscheduleCampaign", scope: "campaigns:send", idempotent: true, conflict: campaignConflict, statuses: []string{"200", "400", "401", "403", "404", "409", "422", "429", "500"}},
		{path: campaign + "/test", method: "post", operationID: "testCampaign", scope: "campaigns:send", idempotent: true, conflict: campaignConflict, badRequest: campaignNotReady, statuses: []string{"200", "400", "401", "403", "404", "409", "422", "429", "500", "502", "503"}},
		{path: campaign + "/stats", method: "get", operationID: "getCampaignStats", scope: "campaigns:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: campaign + "/recipients", method: "get", operationID: "listCampaignRecipients", scope: "campaigns:read", statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
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

			assert.Equal(t, []map[string][]string{{"BearerAuth": {tt.scope}}}, operation.Security)

			actualStatuses := make([]string, 0, len(operation.Responses))
			for status := range operation.Responses {
				actualStatuses = append(actualStatuses, status)
			}
			sort.Strings(actualStatuses)
			assert.Equal(t, tt.statuses, actualStatuses)

			var takesIdempotencyKey bool
			for _, parameter := range operation.Parameters {
				if parameter.Ref == "#/components/parameters/IdempotencyKey" {
					takesIdempotencyKey = true
				}
			}
			assert.Equal(t, tt.idempotent, takesIdempotencyKey)
			if tt.idempotent {
				for status, response := range operation.Responses {
					if strings.HasPrefix(status, "2") {
						assert.Contains(t, response.Headers, "Idempotent-Replayed")
					}
				}
				assert.Equal(t, tt.conflict, operation.Responses["409"].Ref)
				assert.Equal(t, "#/components/responses/IdempotencyPayloadMismatch", operation.Responses["422"].Ref)
			}

			wantBadRequest := "#/components/schemas/ErrorResponse"
			if tt.badRequest != "" {
				wantBadRequest = tt.badRequest
			}
			assert.Equal(t, wantBadRequest, operation.Responses["400"].Content["application/json"].Schema.Ref)
		})
	}
}

func TestOpenAPICampaignSchemasMatchModels(t *testing.T) {
	tests := []struct {
		schemas []string
		model   any
	}{
		{schemas: []string{"CampaignTemplate"}, model: responses.CampaignTemplate{}},
		{schemas: []string{"PaginatedCampaignTemplatesResponse"}, model: responses.PaginatedCampaignTemplatesResponse{}},
		{schemas: []string{"CampaignIssue"}, model: responses.CampaignIssue{}},
		{schemas: []string{"CampaignListItem"}, model: responses.CampaignListItem{}},
		// Campaign is an allOf of a $ref to CampaignListItem and the draft
		// fields, and CampaignSend an allOf of a $ref to Campaign and the
		// warnings.
		{schemas: []string{"Campaign", "CampaignListItem"}, model: responses.Campaign{}},
		{schemas: []string{"CampaignSend", "Campaign", "CampaignListItem"}, model: responses.CampaignSend{}},
		{schemas: []string{"PaginatedCampaignsResponse"}, model: responses.PaginatedCampaignsResponse{}},
		{schemas: []string{"CampaignNotReady"}, model: responses.CampaignNotReady{}},
		{schemas: []string{"CampaignTestFailure"}, model: responses.CampaignTestFailure{}},
		{schemas: []string{"CampaignTest"}, model: responses.CampaignTest{}},
		{schemas: []string{"CampaignStatsReport"}, model: responses.CampaignStatsReport{}},
		{schemas: []string{"CampaignStatsRates"}, model: responses.CampaignStatsRates{}},
		{schemas: []string{"CampaignStats"}, model: responses.CampaignStats{}},
		{schemas: []string{"CampaignRecipient"}, model: responses.CampaignRecipient{}},
		{schemas: []string{"PaginatedCampaignRecipientsResponse"}, model: responses.PaginatedCampaignRecipientsResponse{}},
		{schemas: []string{"CreateCampaignTemplateRequest"}, model: requests.CreateCampaignTemplateRequest{}},
		{schemas: []string{"UpdateCampaignTemplateRequest"}, model: requests.UpdateCampaignTemplateRequest{}},
		{schemas: []string{"CreateCampaignRequest"}, model: requests.CreateCampaignRequest{}},
		{schemas: []string{"UpdateCampaignRequest"}, model: requests.UpdateCampaignRequest{}},
		{schemas: []string{"DuplicateCampaignRequest"}, model: requests.DuplicateCampaignRequest{}},
		{schemas: []string{"SendCampaignRequest"}, model: requests.SendCampaignRequest{}},
		{schemas: []string{"TestCampaignRequest"}, model: requests.TestCampaignRequest{}},
	}

	for _, tt := range tests {
		t.Run(tt.schemas[0], func(t *testing.T) {
			assert.Equal(t, openAPISchemaProperties(t, tt.schemas...), modelJSONFields(tt.model))
		})
	}
}

const campaignTemplateResponseFixture = `{
	"object":"campaign_template",
	"id":"44444444-4444-4444-8444-444444444444",
	"created_at":"2026-10-01T10:00:00Z",
	"updated_at":"2026-10-01T11:00:00Z",
	"name":"Monthly newsletter",
	"editor":"html",
	"content":{"html":"<p>News</p>","text":"News","text_is_custom":false}
}`

func TestCampaignTemplatesAPIEmitsDeclaredTransport(t *testing.T) {
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	templateID := uuid.MustParse("44444444-4444-4444-8444-444444444444")
	templatesPath := "/v2/accounts/" + accountID.String() + "/campaign-templates"
	templatePath := templatesPath + "/" + templateID.String()

	assertCampaignTemplate := func(t *testing.T, result any) {
		template := result.(*responses.CampaignTemplate)
		assert.Equal(t, templateID, template.ID)
		assert.Equal(t, common.TemplateEditorHTML, template.Editor)
		require.NotNil(t, template.Content)
		assert.Equal(t, "<p>News</p>", template.Content.HTML)
	}

	tests := []struct {
		name           string
		method         string
		path           string
		status         int
		response       string
		idempotencyKey string
		wantBody       string
		call           func(*APIClient) (any, *http.Response, error)
		assertResult   func(*testing.T, any)
	}{
		{
			name:           "create campaign template",
			method:         http.MethodPost,
			path:           templatesPath,
			status:         http.StatusCreated,
			response:       campaignTemplateResponseFixture,
			idempotencyKey: "create-campaign-template",
			wantBody:       `{"name":"Monthly newsletter","content":{"html":"<p>News</p>"}}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignTemplatesAPI.CreateCampaignTemplate(context.Background(), accountID, requests.CreateCampaignTemplateRequest{
					Name:    "Monthly newsletter",
					Content: &requests.TemplateContentInput{HTML: ahasend.String("<p>News</p>")},
				}, WithIdempotencyKey("create-campaign-template"))
			},
			assertResult: assertCampaignTemplate,
		},
		{
			name:     "get campaign template",
			method:   http.MethodGet,
			path:     templatePath,
			status:   http.StatusOK,
			response: campaignTemplateResponseFixture,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignTemplatesAPI.GetCampaignTemplate(context.Background(), accountID, templateID)
			},
			assertResult: assertCampaignTemplate,
		},
		{
			name:           "update campaign template",
			method:         http.MethodPut,
			path:           templatePath,
			status:         http.StatusOK,
			response:       campaignTemplateResponseFixture,
			idempotencyKey: "update-campaign-template",
			wantBody:       `{"name":"Monthly newsletter","content":{"text":null}}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignTemplatesAPI.UpdateCampaignTemplate(context.Background(), accountID, templateID, requests.UpdateCampaignTemplateRequest{
					Name:    ahasend.String("Monthly newsletter"),
					Content: &requests.TemplateContentInput{Text: ahasend.String("")},
				}, WithIdempotencyKey("update-campaign-template"))
			},
			assertResult: assertCampaignTemplate,
		},
		{
			name:     "delete campaign template",
			method:   http.MethodDelete,
			path:     templatePath,
			status:   http.StatusOK,
			response: `{"message":"campaign template deleted"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignTemplatesAPI.DeleteCampaignTemplate(context.Background(), accountID, templateID)
			},
			assertResult: func(t *testing.T, result any) {
				assert.Equal(t, "campaign template deleted", result.(*common.SuccessResponse).Message)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertDeclaredTransport(t, tt.method, tt.path, tt.status, tt.response, tt.idempotencyKey, tt.wantBody, tt.call, tt.assertResult)
		})
	}
}

const campaignResponseFixture = `{
	"object":"campaign",
	"id":"55555555-5555-4555-8555-555555555555",
	"created_at":"2026-10-01T10:00:00Z",
	"updated_at":"2026-10-01T11:00:00Z",
	"name":"October newsletter",
	"subject":"Our news",
	"preheader":null,
	"from":{"email":"news@example.com","name":""},
	"reply_to":null,
	"status":"draft",
	"status_reason":null,
	"send_at":null,
	"send_resumes_at":null,
	"finished_at":null,
	"list_ids":["66666666-6666-4666-8666-666666666666"],
	"segment_ids":["33333333-3333-4333-8333-333333333333"],
	"exclude_list_ids":[],
	"exclude_segment_ids":[],
	"editor":"html",
	"content":{"html":"<p>News</p>","text":"News","text_is_custom":false},
	"ready":false,
	"issues":[{"code":"preheader_missing","message":"Add a preview text before you send the campaign."}]
}`

func TestCampaignsAPIEmitsDeclaredTransport(t *testing.T) {
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	campaignID := uuid.MustParse("55555555-5555-4555-8555-555555555555")
	listID := uuid.MustParse("66666666-6666-4666-8666-666666666666")
	templateID := uuid.MustParse("44444444-4444-4444-8444-444444444444")
	segmentID := uuid.MustParse("33333333-3333-4333-8333-333333333333")
	campaignsPath := "/v2/accounts/" + accountID.String() + "/campaigns"
	campaignPath := campaignsPath + "/" + campaignID.String()
	sendAt := time.Date(2026, 10, 20, 9, 0, 0, 0, time.UTC)

	assertDraft := func(t *testing.T, result any) {
		campaign := result.(*responses.Campaign)
		assert.Equal(t, campaignID, campaign.ID)
		assert.Equal(t, common.CampaignStatusDraft, campaign.Status)
		assert.Equal(t, []uuid.UUID{listID}, campaign.ListIDs)
		assert.Equal(t, []uuid.UUID{segmentID}, campaign.SegmentIDs)
		require.NotNil(t, campaign.ExcludeListIDs)
		assert.Empty(t, campaign.ExcludeListIDs)
		require.NotNil(t, campaign.Ready)
		assert.False(t, *campaign.Ready)
		require.Len(t, campaign.Issues, 1)
		assert.Equal(t, responses.CampaignIssuePreheaderMissing, campaign.Issues[0].Code)
	}

	tests := []struct {
		name           string
		method         string
		path           string
		status         int
		response       string
		idempotencyKey string
		wantBody       string
		call           func(*APIClient) (any, *http.Response, error)
		assertResult   func(*testing.T, any)
	}{
		{
			name:           "create campaign",
			method:         http.MethodPost,
			path:           campaignsPath,
			status:         http.StatusCreated,
			response:       campaignResponseFixture,
			idempotencyKey: "create-campaign",
			wantBody:       `{"name":"October newsletter","subject":"Our news","from":{"email":"news@example.com"},"list_ids":["66666666-6666-4666-8666-666666666666"],"segment_ids":["33333333-3333-4333-8333-333333333333"],"template_id":"44444444-4444-4444-8444-444444444444"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.CreateCampaign(context.Background(), accountID, requests.CreateCampaignRequest{
					Name:       "October newsletter",
					Subject:    ahasend.String("Our news"),
					From:       &common.SenderAddress{Email: "news@example.com"},
					ListIDs:    []uuid.UUID{listID},
					SegmentIDs: []uuid.UUID{segmentID},
					TemplateID: &templateID,
				}, WithIdempotencyKey("create-campaign"))
			},
			assertResult: assertDraft,
		},
		{
			name:     "get campaign",
			method:   http.MethodGet,
			path:     campaignPath,
			status:   http.StatusOK,
			response: campaignResponseFixture,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.GetCampaign(context.Background(), accountID, campaignID)
			},
			assertResult: assertDraft,
		},
		{
			name:           "update campaign",
			method:         http.MethodPut,
			path:           campaignPath,
			status:         http.StatusOK,
			response:       campaignResponseFixture,
			idempotencyKey: "update-campaign",
			wantBody:       `{"subject":null,"preheader":"See what is new","from":null,"reply_to":null,"list_ids":["66666666-6666-4666-8666-666666666666"],"exclude_segment_ids":[]}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.UpdateCampaign(context.Background(), accountID, campaignID, requests.UpdateCampaignRequest{
					Subject:           ahasend.String(""),
					Preheader:         ahasend.String("See what is new"),
					From:              &common.SenderAddress{},
					ReplyTo:           ahasend.String(""),
					ListIDs:           []uuid.UUID{listID},
					ExcludeSegmentIDs: []uuid.UUID{},
				}, WithIdempotencyKey("update-campaign"))
			},
			assertResult: assertDraft,
		},
		{
			name:     "delete campaign",
			method:   http.MethodDelete,
			path:     campaignPath,
			status:   http.StatusOK,
			response: `{"message":"campaign deleted"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.DeleteCampaign(context.Background(), accountID, campaignID)
			},
			assertResult: func(t *testing.T, result any) {
				assert.Equal(t, "campaign deleted", result.(*common.SuccessResponse).Message)
			},
		},
		{
			name:           "duplicate campaign",
			method:         http.MethodPost,
			path:           campaignPath + "/duplicate",
			status:         http.StatusCreated,
			response:       campaignResponseFixture,
			idempotencyKey: "duplicate-campaign",
			wantBody:       `{"name":"November newsletter"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.DuplicateCampaign(context.Background(), accountID, campaignID, requests.DuplicateCampaignRequest{
					Name: ahasend.String("November newsletter"),
				}, WithIdempotencyKey("duplicate-campaign"))
			},
			assertResult: assertDraft,
		},
		{
			name:           "duplicate campaign with the default name",
			method:         http.MethodPost,
			path:           campaignPath + "/duplicate",
			status:         http.StatusCreated,
			response:       campaignResponseFixture,
			idempotencyKey: "duplicate-campaign",
			wantBody:       `{}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.DuplicateCampaign(context.Background(), accountID, campaignID, requests.DuplicateCampaignRequest{}, WithIdempotencyKey("duplicate-campaign"))
			},
			assertResult: assertDraft,
		},
		{
			name:           "send campaign later",
			method:         http.MethodPost,
			path:           campaignPath + "/send",
			status:         http.StatusAccepted,
			response:       `{"object":"campaign","id":"` + campaignID.String() + `","status":"validating","send_at":"2026-10-20T09:00:00Z","ready":null,"issues":null,"warnings":["the account cannot send right now"]}`,
			idempotencyKey: "send-campaign",
			wantBody:       `{"send_at":"2026-10-20T09:00:00Z"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.SendCampaign(context.Background(), accountID, campaignID, requests.SendCampaignRequest{SendAt: &sendAt}, WithIdempotencyKey("send-campaign"))
			},
			assertResult: func(t *testing.T, result any) {
				send := result.(*responses.CampaignSend)
				assert.Equal(t, campaignID, send.ID)
				assert.Equal(t, common.CampaignStatusValidating, send.Status)
				require.NotNil(t, send.SendAt)
				assert.Equal(t, sendAt, send.SendAt.UTC())
				assert.Nil(t, send.Ready)
				assert.Nil(t, send.Issues)
				assert.Equal(t, []string{"the account cannot send right now"}, send.Warnings)
			},
		},
		{
			name:           "send campaign now",
			method:         http.MethodPost,
			path:           campaignPath + "/send",
			status:         http.StatusAccepted,
			response:       `{"object":"campaign","id":"` + campaignID.String() + `","status":"validating","warnings":[]}`,
			idempotencyKey: "send-campaign",
			wantBody:       `{}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.SendCampaign(context.Background(), accountID, campaignID, requests.SendCampaignRequest{}, WithIdempotencyKey("send-campaign"))
			},
			assertResult: func(t *testing.T, result any) {
				send := result.(*responses.CampaignSend)
				require.NotNil(t, send.Warnings)
				assert.Empty(t, send.Warnings)
			},
		},
		{
			name:           "unschedule campaign",
			method:         http.MethodPost,
			path:           campaignPath + "/unschedule",
			status:         http.StatusOK,
			response:       campaignResponseFixture,
			idempotencyKey: "unschedule-campaign",
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.UnscheduleCampaign(context.Background(), accountID, campaignID, WithIdempotencyKey("unschedule-campaign"))
			},
			assertResult: assertDraft,
		},
		{
			name:           "test campaign",
			method:         http.MethodPost,
			path:           campaignPath + "/test",
			status:         http.StatusOK,
			response:       `{"object":"campaign_test","sent_to":["ann@example.com"],"failed":[{"email":"bob@example.com","message":"mailbox unavailable"}],"result_unknown":false}`,
			idempotencyKey: "test-campaign",
			wantBody:       `{"to":["Ann <ann@example.com>","bob@example.com"]}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.TestCampaign(context.Background(), accountID, campaignID, requests.TestCampaignRequest{
					To: []string{"Ann <ann@example.com>", "bob@example.com"},
				}, WithIdempotencyKey("test-campaign"))
			},
			assertResult: func(t *testing.T, result any) {
				test := result.(*responses.CampaignTest)
				assert.Equal(t, []string{"ann@example.com"}, test.SentTo)
				assert.Equal(t, []responses.CampaignTestFailure{{Email: "bob@example.com", Message: "mailbox unavailable"}}, test.Failed)
				assert.False(t, test.ResultUnknown)
			},
		},
		{
			name:     "get campaign stats",
			method:   http.MethodGet,
			path:     campaignPath + "/stats",
			status:   http.StatusOK,
			response: `{"object":"campaign_stats","campaign_id":"` + campaignID.String() + `","recipients":10,"delivered":9,"opens":4,"clicks":1,"unsubscribes":0,"complaints":0,"report":null,"rates":{"open_rate":null,"click_rate":null,"bounce_rate":null,"deferral_rate":null},"recount_failed":false}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.CampaignsAPI.GetCampaignStats(context.Background(), accountID, campaignID)
			},
			assertResult: func(t *testing.T, result any) {
				stats := result.(*responses.CampaignStats)
				assert.Equal(t, campaignID, stats.CampaignID)
				assert.Equal(t, 9, stats.Delivered)
				assert.Nil(t, stats.Report)
				assert.Nil(t, stats.Rates.OpenRate)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertDeclaredTransport(t, tt.method, tt.path, tt.status, tt.response, tt.idempotencyKey, tt.wantBody, tt.call, tt.assertResult)
		})
	}
}

// assertDeclaredTransport runs call against a server that answers status and
// response, and asserts the request had the method, path, idempotency key
// and body given, and no query.
func assertDeclaredTransport(
	t *testing.T,
	method, path string,
	status int,
	response, idempotencyKey, wantBody string,
	call func(*APIClient) (any, *http.Response, error),
	assertResult func(*testing.T, any),
) {
	t.Helper()

	var capturedRequest *http.Request
	var capturedBody []byte
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		capturedRequest = r
		capturedBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(response))
	})
	defer cleanup()

	result, httpResponse, err := call(client)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, httpResponse)
	require.NotNil(t, capturedRequest)
	assert.Equal(t, status, httpResponse.StatusCode)
	assert.Equal(t, method, capturedRequest.Method)
	assert.Equal(t, path, capturedRequest.URL.EscapedPath())
	assert.Empty(t, capturedRequest.URL.RawQuery)
	if idempotencyKey != "" {
		assert.Equal(t, idempotencyKey, capturedRequest.Header.Get("Idempotency-Key"))
	}
	if wantBody != "" {
		assert.JSONEq(t, wantBody, string(capturedBody))
	} else {
		assert.Empty(t, capturedBody)
	}
	if assertResult != nil {
		assertResult(t, result)
	}
}

func TestCampaignWritesRetryWithTheirIdempotencyKey(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		response string
		call     func(*APIClient) (*http.Response, error)
	}{
		{
			name:     "update campaign template",
			status:   http.StatusInternalServerError,
			response: campaignTemplateResponseFixture,
			call: func(client *APIClient) (*http.Response, error) {
				_, httpResponse, err := client.CampaignTemplatesAPI.UpdateCampaignTemplate(context.Background(), uuid.New(), uuid.New(), requests.UpdateCampaignTemplateRequest{
					Name: ahasend.String("Monthly newsletter"),
				})
				return httpResponse, err
			},
		},
		{
			name:     "update campaign",
			status:   http.StatusInternalServerError,
			response: campaignResponseFixture,
			call: func(client *APIClient) (*http.Response, error) {
				_, httpResponse, err := client.CampaignsAPI.UpdateCampaign(context.Background(), uuid.New(), uuid.New(), requests.UpdateCampaignRequest{
					Subject: ahasend.String("Our news"),
				})
				return httpResponse, err
			},
		},
		{
			// A 502 means the mail server took nothing; the API asks for a
			// retry with the same key.
			name:     "test campaign",
			status:   http.StatusBadGateway,
			response: `{"object":"campaign_test","sent_to":["ann@example.com"],"failed":[],"result_unknown":false}`,
			call: func(client *APIClient) (*http.Response, error) {
				_, httpResponse, err := client.CampaignsAPI.TestCampaign(context.Background(), uuid.New(), uuid.New(), requests.TestCampaignRequest{
					To: []string{"ann@example.com"},
				})
				return httpResponse, err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var attempts int
			var keys []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				attempts++
				keys = append(keys, r.Header.Get("Idempotency-Key"))
				w.Header().Set("Content-Type", "application/json")
				if attempts == 1 {
					w.WriteHeader(tt.status)
					_, _ = w.Write([]byte(`{"message":"try again"}`))
					return
				}
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()
			serverURL, err := url.Parse(server.URL)
			require.NoError(t, err)

			cfg := NewConfiguration()
			cfg.Host = serverURL.Host
			cfg.Scheme = serverURL.Scheme
			cfg.APIKey = "test-key"
			cfg.RetryConfig.Enabled = true
			cfg.RetryConfig.MaxRetries = 2
			cfg.RetryConfig.BaseDelay = time.Millisecond
			cfg.RetryConfig.MaxDelay = time.Millisecond
			client := NewAPIClientWithConfig(cfg)

			httpResponse, err := tt.call(client)

			require.NoError(t, err)
			require.NotNil(t, httpResponse)
			assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
			require.Len(t, keys, 2)
			assert.NotEmpty(t, keys[0])
			assert.Equal(t, keys[0], keys[1])
		})
	}
}

func TestCampaignsAPIClassifiesCampaignConflicts(t *testing.T) {
	tests := []struct {
		name           string
		headers        map[string]string
		wantType       ErrorType
		wantRetryAfter int
	}{
		{name: "fresh status conflict", wantType: ErrorTypeConflict},
		{name: "replayed status conflict", headers: map[string]string{"Idempotent-Replayed": "true"}, wantType: ErrorTypeConflict},
		{name: "execution in progress", headers: map[string]string{"Idempotent-Replayed": "false", "Retry-After": "5"}, wantType: ErrorTypeIdempotencyConflict, wantRetryAfter: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
				for key, value := range tt.headers {
					w.Header().Set(key, value)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				_, _ = w.Write([]byte(`{"message":"campaign is not a draft"}`))
			})
			defer cleanup()

			_, _, err := client.CampaignsAPI.SendCampaign(context.Background(), uuid.New(), uuid.New(), requests.SendCampaignRequest{}, WithIdempotencyKey("send-campaign"))

			var apiErr *APIError
			require.ErrorAs(t, err, &apiErr)
			assert.Equal(t, tt.wantType, apiErr.Type)
			assert.Equal(t, tt.wantRetryAfter, apiErr.RetryAfter)
			assert.Equal(t, "campaign is not a draft", apiErr.Message)
			_, notReady := AsCampaignNotReady(err)
			assert.False(t, notReady)
		})
	}
}

func TestOpenAPISegmentOperationsAndResponses(t *testing.T) {
	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	type segmentOperationContract struct {
		OperationID string                      `yaml:"operationId"`
		CodeSamples []struct{ Lang string }     `yaml:"x-code-samples"`
		Security    []map[string][]string       `yaml:"security"`
		Responses   map[string]responseContract `yaml:"responses"`
		Parameters  []struct {
			Ref string `yaml:"$ref"`
		} `yaml:"parameters"`
	}
	var spec struct {
		Paths map[string]map[string]segmentOperationContract `yaml:"paths"`
	}
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	const (
		segmentCreateConflict = "#/components/responses/SegmentCreateConflict"

		segments        = "/v2/accounts/{account_id}/segments"
		segment         = segments + "/{segment_id}"
		contactSegments = "/v2/accounts/{account_id}/contacts/{id_or_email}/segments"
	)
	tests := []struct {
		path        string
		method      string
		operationID string
		scopes      []string
		idempotent  bool
		// conflict is the response a 409 of an idempotent operation refers
		// to.
		conflict string
		statuses []string
	}{
		{path: segments, method: "get", operationID: "getSegments", scopes: []string{"segments:read"}, statuses: []string{"200", "400", "401", "403", "429", "500"}},
		{path: segments, method: "post", operationID: "createSegment", scopes: []string{"segments:write"}, idempotent: true, conflict: segmentCreateConflict, statuses: []string{"201", "400", "401", "403", "409", "422", "429", "500"}},
		{path: segments + "/preview", method: "post", operationID: "previewSegment", scopes: []string{"segments:read", "contacts:read"}, statuses: []string{"200", "400", "401", "403", "429", "500"}},
		{path: segment, method: "get", operationID: "getSegment", scopes: []string{"segments:read"}, statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: segment, method: "put", operationID: "updateSegment", scopes: []string{"segments:write"}, statuses: []string{"200", "400", "401", "403", "404", "409", "429", "500"}},
		{path: segment, method: "delete", operationID: "deleteSegment", scopes: []string{"segments:delete"}, statuses: []string{"200", "400", "401", "403", "404", "409", "429", "500"}},
		{path: segment + "/contacts", method: "get", operationID: "getSegmentContacts", scopes: []string{"segments:read", "contacts:read"}, statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
		{path: contactSegments, method: "get", operationID: "getContactSegments", scopes: []string{"segments:read", "contacts:read"}, statuses: []string{"200", "400", "401", "403", "404", "429", "500"}},
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

			assert.Equal(t, []map[string][]string{{"BearerAuth": tt.scopes}}, operation.Security)

			actualStatuses := make([]string, 0, len(operation.Responses))
			for status := range operation.Responses {
				actualStatuses = append(actualStatuses, status)
			}
			sort.Strings(actualStatuses)
			assert.Equal(t, tt.statuses, actualStatuses)

			var takesIdempotencyKey bool
			for _, parameter := range operation.Parameters {
				if parameter.Ref == "#/components/parameters/IdempotencyKey" {
					takesIdempotencyKey = true
				}
			}
			assert.Equal(t, tt.idempotent, takesIdempotencyKey)
			if tt.idempotent {
				for status, response := range operation.Responses {
					if strings.HasPrefix(status, "2") {
						assert.Contains(t, response.Headers, "Idempotent-Replayed")
					}
				}
				assert.Equal(t, tt.conflict, operation.Responses["409"].Ref)
				assert.Equal(t, "#/components/responses/IdempotencyPayloadMismatch", operation.Responses["422"].Ref)
			} else if conflict, ok := operation.Responses["409"]; ok {
				assert.Empty(t, conflict.Ref)
				assert.Empty(t, conflict.Headers)
			}
		})
	}
}

func TestOpenAPISegmentSchemasMatchModels(t *testing.T) {
	tests := []struct {
		schemas []string
		model   any
	}{
		{schemas: []string{"Segment"}, model: responses.Segment{}},
		// SegmentDetail is an allOf of a $ref to Segment and the marketable
		// count.
		{schemas: []string{"SegmentDetail", "Segment"}, model: responses.SegmentDetail{}},
		{schemas: []string{"SegmentDefinition"}, model: common.SegmentDefinition{}},
		{schemas: []string{"SegmentGroup"}, model: common.SegmentGroup{}},
		{schemas: []string{"SegmentCondition"}, model: common.SegmentCondition{}},
		{schemas: []string{"SegmentCampaignCount"}, model: common.SegmentCampaignCount{}},
		{schemas: []string{"CreateSegmentRequest"}, model: requests.CreateSegmentRequest{}},
		{schemas: []string{"UpdateSegmentRequest"}, model: requests.UpdateSegmentRequest{}},
		{schemas: []string{"PreviewSegmentRequest"}, model: requests.PreviewSegmentRequest{}},
		{schemas: []string{"SegmentPreview"}, model: responses.SegmentPreview{}},
		{schemas: []string{"SegmentPreviewContact"}, model: responses.SegmentPreviewContact{}},
		{schemas: []string{"SegmentContact"}, model: responses.SegmentContact{}},
		{schemas: []string{"PaginatedSegmentsResponse"}, model: responses.PaginatedSegmentsResponse{}},
		{schemas: []string{"PaginatedSegmentContactsResponse"}, model: responses.PaginatedSegmentContactsResponse{}},
	}

	for _, tt := range tests {
		t.Run(tt.schemas[0], func(t *testing.T) {
			assert.Equal(t, openAPISchemaProperties(t, tt.schemas...), modelJSONFields(tt.model))
		})
	}
}

func TestOpenAPIDeliveryEventSchemaMatchesModel(t *testing.T) {
	assert.Equal(t, openAPISchemaProperties(t, "DeliveryEvent"), modelJSONFields(responses.DeliveryEvent{}))
}

const segmentResponseFixture = `{
	"object":"segment",
	"id":"33333333-3333-4333-8333-333333333333",
	"created_at":"2026-10-20T10:00:00Z",
	"updated_at":"2026-10-20T11:00:00Z",
	"name":"Engaged customers",
	"description":"",
	"definition":{"schema_version":1,"match":"all","conditions":[{"field":"list","op":"in","value":"22222222-2222-4222-8222-222222222222"},{"match":"any","conditions":[{"field":"campaign_count","op":"opened_at_least","value":{"count":3,"since":"90 days ago"}}]}]},
	"timezone":"Europe/Berlin",
	"state":"building",
	"failure_reason":null,
	"member_count":0,
	"last_evaluated_at":null
}`

const segmentContactResponseFixture = `{
	"object":"segment_contact",
	"segment_id":"33333333-3333-4333-8333-333333333333",
	"contact_id":"11111111-1111-4111-8111-111111111111",
	"email":"user+tag@example.com",
	"entered_at":"2026-10-20T11:00:00Z"
}`

// segmentTestDefinition is the definition segmentResponseFixture holds.
func segmentTestDefinition() common.SegmentDefinition {
	return common.SegmentDefinition{
		SchemaVersion: common.SegmentSchemaVersion,
		Match:         common.SegmentMatchAll,
		Conditions: []common.SegmentRule{
			{Condition: &common.SegmentCondition{
				Field: common.SegmentFieldList,
				Op:    common.SegmentOpIn,
				Value: &common.SegmentValue{String: ahasend.String("22222222-2222-4222-8222-222222222222")},
			}},
			{Group: &common.SegmentGroup{
				Match: common.SegmentMatchAny,
				Conditions: []common.SegmentCondition{{
					Field: common.SegmentFieldCampaignCount,
					Op:    common.SegmentOpOpenedAtLeast,
					Value: &common.SegmentValue{Count: &common.SegmentCampaignCount{Count: 3, Since: "90 days ago"}},
				}},
			}},
		},
	}
}

func TestSegmentsAPIEmitsDeclaredTransport(t *testing.T) {
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	segmentID := uuid.MustParse("33333333-3333-4333-8333-333333333333")
	contactID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	rawEmail := "a/b+Tag@Example.COM"
	escapedEmail := "a%2Fb+Tag@Example.COM"
	segmentsPath := "/v2/accounts/" + accountID.String() + "/segments"
	segmentPath := segmentsPath + "/" + segmentID.String()
	definitionJSON := `{"schema_version":1,"match":"all","conditions":[{"field":"list","op":"in","value":"22222222-2222-4222-8222-222222222222"},{"match":"any","conditions":[{"field":"campaign_count","op":"opened_at_least","value":{"count":3,"since":"90 days ago"}}]}]}`
	segmentsPage := `{"object":"list","data":[` + segmentResponseFixture + `],"pagination":{"has_more":true,"next_cursor":"next","previous_cursor":null}}`

	assertSegmentsPage := func(t *testing.T, result any) {
		page := result.(*responses.PaginatedSegmentsResponse)
		require.Len(t, page.Data, 1)
		assert.Equal(t, segmentID, page.Data[0].ID)
		assert.Equal(t, common.SegmentStateBuilding, page.Data[0].State)
		assert.Equal(t, segmentTestDefinition(), page.Data[0].Definition)
		assert.True(t, page.Pagination.HasMore)
		require.NotNil(t, page.Pagination.NextCursor)
		assert.Equal(t, "next", *page.Pagination.NextCursor)
	}

	tests := []struct {
		name           string
		method         string
		path           string
		status         int
		response       string
		idempotencyKey string
		wantQuery      url.Values
		wantBody       string
		call           func(*APIClient) (any, *http.Response, error)
		assertResult   func(*testing.T, any)
	}{
		{
			name:      "get segments",
			method:    http.MethodGet,
			path:      segmentsPath,
			status:    http.StatusOK,
			response:  segmentsPage,
			wantQuery: url.Values{"limit": {"100"}},
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.GetSegments(context.Background(), accountID, requests.GetSegmentsParams{})
			},
			assertResult: assertSegmentsPage,
		},
		{
			name:      "get segments with a cursor",
			method:    http.MethodGet,
			path:      segmentsPath,
			status:    http.StatusOK,
			response:  segmentsPage,
			wantQuery: url.Values{"limit": {"25"}, "after": {"after-cursor"}},
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.GetSegments(context.Background(), accountID, requests.GetSegmentsParams{
					Limit: ahasend.Int32(25),
					After: ahasend.String("after-cursor"),
				})
			},
		},
		{
			name:           "create segment",
			method:         http.MethodPost,
			path:           segmentsPath,
			status:         http.StatusCreated,
			response:       segmentResponseFixture,
			idempotencyKey: "create-segment",
			wantBody:       `{"name":"Engaged customers","description":"Opened lately","definition":` + definitionJSON + `,"timezone":"Europe/Berlin"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.CreateSegment(context.Background(), accountID, requests.CreateSegmentRequest{
					Name:        "Engaged customers",
					Description: ahasend.String("Opened lately"),
					Definition:  segmentTestDefinition(),
					Timezone:    ahasend.String("Europe/Berlin"),
				}, WithIdempotencyKey("create-segment"))
			},
			assertResult: func(t *testing.T, result any) {
				segment := result.(*responses.Segment)
				assert.Equal(t, "segment", segment.Object)
				assert.Equal(t, common.SegmentStateBuilding, segment.State)
				assert.Nil(t, segment.LastEvaluatedAt)
			},
		},
		{
			name:           "preview segment",
			method:         http.MethodPost,
			path:           segmentsPath + "/preview",
			status:         http.StatusOK,
			response:       `{"object":"segment_preview","count":null,"contacts":[{"contact_id":"11111111-1111-4111-8111-111111111111","email":"user+tag@example.com"}]}`,
			idempotencyKey: "preview-segment",
			wantBody:       `{"definition":` + definitionJSON + `}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.PreviewSegment(context.Background(), accountID, requests.PreviewSegmentRequest{
					Definition: segmentTestDefinition(),
				}, WithIdempotencyKey("preview-segment"))
			},
			assertResult: func(t *testing.T, result any) {
				preview := result.(*responses.SegmentPreview)
				assert.Equal(t, "segment_preview", preview.Object)
				assert.Nil(t, preview.Count)
				require.Len(t, preview.Contacts, 1)
				assert.Equal(t, contactID, preview.Contacts[0].ContactID)
			},
		},
		{
			name:     "get segment",
			method:   http.MethodGet,
			path:     segmentPath,
			status:   http.StatusOK,
			response: segmentResponseFixture[:len(segmentResponseFixture)-1] + `,"marketable_count":null}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.GetSegment(context.Background(), accountID, segmentID)
			},
			assertResult: func(t *testing.T, result any) {
				segment := result.(*responses.SegmentDetail)
				assert.Equal(t, segmentID, segment.ID)
				assert.Nil(t, segment.MarketableCount)
			},
		},
		{
			name:     "update segment",
			method:   http.MethodPut,
			path:     segmentPath,
			status:   http.StatusOK,
			response: segmentResponseFixture,
			wantBody: `{"name":"Engaged customers","definition":` + definitionJSON + `}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				definition := segmentTestDefinition()
				return client.SegmentsAPI.UpdateSegment(context.Background(), accountID, segmentID, requests.UpdateSegmentRequest{
					Name:       ahasend.String("Engaged customers"),
					Definition: &definition,
				})
			},
		},
		{
			name:     "update segment naming nothing",
			method:   http.MethodPut,
			path:     segmentPath,
			status:   http.StatusOK,
			response: segmentResponseFixture,
			wantBody: `{}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.UpdateSegment(context.Background(), accountID, segmentID, requests.UpdateSegmentRequest{})
			},
		},
		{
			name:     "delete segment",
			method:   http.MethodDelete,
			path:     segmentPath,
			status:   http.StatusOK,
			response: `{"message":"segment deleted"}`,
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.DeleteSegment(context.Background(), accountID, segmentID)
			},
			assertResult: func(t *testing.T, result any) {
				assert.Equal(t, "segment deleted", result.(*common.SuccessResponse).Message)
			},
		},
		{
			name:      "get segment contacts",
			method:    http.MethodGet,
			path:      segmentPath + "/contacts",
			status:    http.StatusOK,
			response:  `{"object":"list","data":[` + segmentContactResponseFixture + `],"pagination":{"has_more":false,"next_cursor":null,"previous_cursor":"prev"}}`,
			wantQuery: url.Values{"limit": {"100"}, "before": {"before-cursor"}},
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.GetSegmentContacts(context.Background(), accountID, segmentID, requests.GetSegmentContactsParams{
					Before: ahasend.String("before-cursor"),
				})
			},
			assertResult: func(t *testing.T, result any) {
				page := result.(*responses.PaginatedSegmentContactsResponse)
				require.Len(t, page.Data, 1)
				assert.Equal(t, segmentID, page.Data[0].SegmentID)
				assert.Equal(t, contactID, page.Data[0].ContactID)
				assert.Equal(t, time.Date(2026, 10, 20, 11, 0, 0, 0, time.UTC), page.Data[0].EnteredAt.UTC())
				require.NotNil(t, page.Pagination.PreviousCursor)
				assert.Nil(t, page.Pagination.NextCursor)
			},
		},
		{
			name:      "get contact segments",
			method:    http.MethodGet,
			path:      "/v2/accounts/" + accountID.String() + "/contacts/" + escapedEmail + "/segments",
			status:    http.StatusOK,
			response:  segmentsPage,
			wantQuery: url.Values{"limit": {"10"}, "after": {"after-cursor"}},
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.GetContactSegments(context.Background(), accountID, rawEmail, requests.GetContactSegmentsParams{
					Limit: ahasend.Int32(10),
					After: ahasend.String("after-cursor"),
				})
			},
			assertResult: assertSegmentsPage,
		},
		{
			name:      "get contact segments by ID",
			method:    http.MethodGet,
			path:      "/v2/accounts/" + accountID.String() + "/contacts/" + contactID.String() + "/segments",
			status:    http.StatusOK,
			response:  segmentsPage,
			wantQuery: url.Values{"limit": {"100"}},
			call: func(client *APIClient) (any, *http.Response, error) {
				return client.SegmentsAPI.GetContactSegments(context.Background(), accountID, contactID.String(), requests.GetContactSegmentsParams{})
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
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.response))
			})
			defer cleanup()

			result, httpResponse, err := tt.call(client)

			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, httpResponse)
			require.NotNil(t, capturedRequest)
			assert.Equal(t, tt.status, httpResponse.StatusCode)
			assert.Equal(t, tt.method, capturedRequest.Method)
			assert.Equal(t, tt.path, capturedRequest.URL.EscapedPath())
			assert.Equal(t, tt.idempotencyKey, capturedRequest.Header.Get("Idempotency-Key"))
			if tt.wantQuery != nil {
				assert.Equal(t, tt.wantQuery, capturedRequest.URL.Query())
			} else {
				assert.Empty(t, capturedRequest.URL.RawQuery)
			}
			if tt.wantBody != "" {
				assert.JSONEq(t, tt.wantBody, string(capturedBody))
			} else {
				assert.Empty(t, capturedBody)
			}
			if tt.assertResult != nil {
				tt.assertResult(t, result)
			}
		})
	}

	t.Run("create segment without a key", func(t *testing.T) {
		var idempotencyKey string
		client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			idempotencyKey = r.Header.Get("Idempotency-Key")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(segmentResponseFixture))
		})
		defer cleanup()

		_, _, err := client.SegmentsAPI.CreateSegment(context.Background(), accountID, requests.CreateSegmentRequest{
			Name:       "Engaged customers",
			Definition: segmentTestDefinition(),
		})

		require.NoError(t, err)
		assert.NotEmpty(t, idempotencyKey)
	})
}

func TestCreateSegmentRetriesWithItsIdempotencyKey(t *testing.T) {
	var keys []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys = append(keys, r.Header.Get("Idempotency-Key"))
		w.Header().Set("Content-Type", "application/json")
		if len(keys) == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message":"try again"}`))
			return
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(segmentResponseFixture))
	}))
	defer server.Close()
	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	cfg := NewConfiguration()
	cfg.Host = serverURL.Host
	cfg.Scheme = serverURL.Scheme
	cfg.APIKey = "test-key"
	cfg.RetryConfig.Enabled = true
	cfg.RetryConfig.MaxRetries = 2
	cfg.RetryConfig.BaseDelay = time.Millisecond
	cfg.RetryConfig.MaxDelay = time.Millisecond
	client := NewAPIClientWithConfig(cfg)

	_, httpResponse, err := client.SegmentsAPI.CreateSegment(context.Background(), uuid.New(), requests.CreateSegmentRequest{
		Name:       "Engaged customers",
		Definition: segmentTestDefinition(),
	}, WithIdempotencyKey("create-segment"))

	require.NoError(t, err)
	require.NotNil(t, httpResponse)
	assert.Equal(t, http.StatusCreated, httpResponse.StatusCode)
	assert.Equal(t, []string{"create-segment", "create-segment"}, keys)
}

func TestSegmentsAPIClassifiesSegmentConflicts(t *testing.T) {
	createSegment := func(client *APIClient) error {
		_, _, err := client.SegmentsAPI.CreateSegment(context.Background(), uuid.New(), requests.CreateSegmentRequest{
			Name:       "Engaged customers",
			Definition: segmentTestDefinition(),
		}, WithIdempotencyKey("create-segment"))
		return err
	}

	tests := []struct {
		name           string
		headers        map[string]string
		message        string
		call           func(*APIClient) error
		wantType       ErrorType
		wantRetryAfter int
	}{
		{
			name:     "create refusal",
			message:  "An account can have at most 100 segments. Delete a segment to add another.",
			call:     createSegment,
			wantType: ErrorTypeConflict,
		},
		{
			name:           "create in progress",
			headers:        map[string]string{"Idempotent-Replayed": "false", "Retry-After": "5"},
			message:        "request with this idempotency key is in progress",
			call:           createSegment,
			wantType:       ErrorTypeIdempotencyConflict,
			wantRetryAfter: 5,
		},
		{
			name:    "update busy",
			message: "The segment is being updated. Try again in a minute.",
			call: func(client *APIClient) error {
				_, _, err := client.SegmentsAPI.UpdateSegment(context.Background(), uuid.New(), uuid.New(), requests.UpdateSegmentRequest{
					Name: ahasend.String("Engaged customers"),
				})
				return err
			},
			wantType: ErrorTypeConflict,
		},
		{
			name:    "delete busy",
			message: "The segment is being updated. Try again in a minute.",
			call: func(client *APIClient) error {
				_, _, err := client.SegmentsAPI.DeleteSegment(context.Background(), uuid.New(), uuid.New())
				return err
			},
			wantType: ErrorTypeConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
				for key, value := range tt.headers {
					w.Header().Set(key, value)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusConflict)
				_, _ = w.Write([]byte(`{"message":"` + tt.message + `"}`))
			})
			defer cleanup()

			err := tt.call(client)

			var apiErr *APIError
			require.ErrorAs(t, err, &apiErr)
			assert.Equal(t, http.StatusConflict, apiErr.StatusCode)
			assert.Equal(t, tt.wantType, apiErr.Type)
			assert.Equal(t, tt.wantRetryAfter, apiErr.RetryAfter)
			assert.Equal(t, tt.message, apiErr.Message)
		})
	}
}
