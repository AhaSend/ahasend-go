package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/internal/prismmock"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_CampaignsAPIService(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping API integration tests (SKIP_INTEGRATION_TESTS=true)")
	}
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	configuration := NewConfiguration()
	configuration.Host = prismmock.Addr()
	configuration.Scheme = "http"
	apiClient := NewAPIClientWithConfig(configuration)
	auth := context.WithValue(context.Background(), ContextAccessToken, "test-api-key")
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")
	campaignID := uuid.MustParse("55555555-5555-4555-8555-555555555555")
	listID := uuid.MustParse("66666666-6666-4666-8666-666666666666")

	validatePrismResponse := func(t *testing.T, response any, httpResponse *http.Response, err error) {
		t.Helper()
		if httpResponse == nil {
			assert.Error(t, err)
			return
		}
		assert.GreaterOrEqual(t, httpResponse.StatusCode, http.StatusOK)
		assert.Less(t, httpResponse.StatusCode, http.StatusInternalServerError)
		if httpResponse.StatusCode < http.StatusBadRequest {
			assert.NotNil(t, response)
		}
	}

	t.Run("ListCampaigns", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.ListCampaigns(auth, accountID, requests.ListCampaignsParams{
			Status: ahasend.String(common.CampaignStatusDraft),
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("GetCampaign", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.GetCampaign(auth, accountID, campaignID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("CreateCampaign", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.CreateCampaign(auth, accountID, requests.CreateCampaignRequest{
			Name:    "October newsletter",
			ListIDs: []uuid.UUID{listID},
			Content: &requests.TemplateContentInput{HTML: ahasend.String("<p>News</p>")},
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("UpdateCampaign", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.UpdateCampaign(auth, accountID, campaignID, requests.UpdateCampaignRequest{
			Subject: ahasend.String("Our news"),
			ReplyTo: ahasend.String(""),
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("DeleteCampaign", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.DeleteCampaign(auth, accountID, campaignID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("DuplicateCampaign", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.DuplicateCampaign(auth, accountID, campaignID, requests.DuplicateCampaignRequest{})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("SendCampaign", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.SendCampaign(auth, accountID, campaignID, requests.SendCampaignRequest{
			SendAt: ahasend.Time(time.Now().Add(time.Hour)),
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("UnscheduleCampaign", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.UnscheduleCampaign(auth, accountID, campaignID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("TestCampaign", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.TestCampaign(auth, accountID, campaignID, requests.TestCampaignRequest{
			To: []string{"ann@example.com"},
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("GetCampaignStats", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.GetCampaignStats(auth, accountID, campaignID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("ListCampaignRecipients", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignsAPI.ListCampaignRecipients(auth, accountID, campaignID, requests.ListCampaignRecipientsParams{
			Status: ahasend.String("Delivered"),
		})
		validatePrismResponse(t, response, httpResponse, err)
	})
}

func TestCampaignsAPIListCampaignsSerializesFilterAndPagination(t *testing.T) {
	accountID := uuid.New()

	tests := []struct {
		name         string
		params       requests.ListCampaignsParams
		wantRawQuery string
	}{
		{
			name:         "defaults the limit when none is given",
			params:       requests.ListCampaignsParams{},
			wantRawQuery: "limit=100",
		},
		{
			name: "sends the caller's status, limit and cursors",
			params: requests.ListCampaignsParams{
				Status: ahasend.String(common.CampaignStatusSent),
				Limit:  ahasend.Int32(25),
				After:  ahasend.String("after-cursor"),
				Before: ahasend.String("before-cursor"),
			},
			wantRawQuery: "after=after-cursor&before=before-cursor&limit=25&status=sent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lastRequest *http.Request
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				lastRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"object":"list","data":[{"object":"campaign","id":"55555555-5555-4555-8555-555555555555","created_at":"2026-10-01T10:00:00Z","updated_at":"2026-10-01T11:00:00Z","name":"October newsletter","subject":null,"preheader":null,"from":null,"reply_to":null,"status":"sent","status_reason":null,"send_at":"2026-10-02T09:00:00Z","send_resumes_at":null,"finished_at":"2026-10-02T09:30:00Z"}],"pagination":{"has_more":false}}`))
			})
			defer cleanup()

			page, httpResponse, err := client.CampaignsAPI.ListCampaigns(context.Background(), accountID, tt.params)

			require.NoError(t, err)
			require.NotNil(t, httpResponse)
			require.NotNil(t, lastRequest)
			assert.Equal(t, http.MethodGet, lastRequest.Method)
			assert.Equal(t, "/v2/accounts/"+accountID.String()+"/campaigns", lastRequest.URL.Path)
			assert.Equal(t, tt.wantRawQuery, lastRequest.URL.RawQuery)

			require.NotNil(t, page)
			require.Len(t, page.Data, 1)
			assert.Equal(t, common.CampaignStatusSent, page.Data[0].Status)
			assert.Nil(t, page.Data[0].Subject)
			require.NotNil(t, page.Data[0].FinishedAt)
			assert.False(t, page.Pagination.HasMore)
			assert.Nil(t, page.Pagination.NextCursor)
		})
	}
}

func TestCampaignsAPIListCampaignRecipientsSerializesFilterAndPagination(t *testing.T) {
	accountID := uuid.New()
	campaignID := uuid.New()

	tests := []struct {
		name         string
		params       requests.ListCampaignRecipientsParams
		wantRawQuery string
	}{
		{
			name:         "sends a status with the default limit",
			params:       requests.ListCampaignRecipientsParams{Status: ahasend.String("Administratively Bounced")},
			wantRawQuery: "limit=100&status=Administratively+Bounced",
		},
		{
			name: "sends a bounce classification, limit and cursors",
			params: requests.ListCampaignRecipientsParams{
				BounceClassification: ahasend.String("InvalidRecipient"),
				Limit:                ahasend.Int32(10),
				After:                ahasend.String("after-cursor"),
				Before:               ahasend.String("before-cursor"),
			},
			wantRawQuery: "after=after-cursor&before=before-cursor&bounce_classification=InvalidRecipient&limit=10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lastRequest *http.Request
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				lastRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				// An empty page past the last recipient has no previous cursor.
				_, _ = w.Write([]byte(`{"object":"list","data":[],"pagination":{"has_more":false}}`))
			})
			defer cleanup()

			page, httpResponse, err := client.CampaignsAPI.ListCampaignRecipients(context.Background(), accountID, campaignID, tt.params)

			require.NoError(t, err)
			require.NotNil(t, httpResponse)
			require.NotNil(t, lastRequest)
			assert.Equal(t, http.MethodGet, lastRequest.Method)
			assert.Equal(t, "/v2/accounts/"+accountID.String()+"/campaigns/"+campaignID.String()+"/recipients", lastRequest.URL.Path)
			assert.Equal(t, tt.wantRawQuery, lastRequest.URL.RawQuery)

			require.NotNil(t, page)
			assert.Empty(t, page.Data)
			assert.Nil(t, page.Pagination.PreviousCursor)
		})
	}
}

func TestCampaignsAPIErrorTypes(t *testing.T) {
	tests := []struct {
		statusCode int
		wantType   ErrorType
	}{
		{statusCode: http.StatusBadRequest, wantType: ErrorTypeValidation},
		{statusCode: http.StatusUnauthorized, wantType: ErrorTypeAuthentication},
		{statusCode: http.StatusForbidden, wantType: ErrorTypePermission},
		{statusCode: http.StatusNotFound, wantType: ErrorTypeNotFound},
		{statusCode: http.StatusConflict, wantType: ErrorTypeConflict},
		{statusCode: http.StatusInternalServerError, wantType: ErrorTypeServer},
	}

	for _, tt := range tests {
		t.Run(http.StatusText(tt.statusCode), func(t *testing.T) {
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(`{"message":"campaign request failed"}`))
			})
			defer cleanup()

			campaign, httpResponse, err := client.CampaignsAPI.GetCampaign(context.Background(), uuid.New(), uuid.New())

			require.Error(t, err)
			assert.NotNil(t, campaign)
			require.NotNil(t, httpResponse)
			assert.Equal(t, tt.statusCode, httpResponse.StatusCode)

			var apiErr *APIError
			require.True(t, errors.As(err, &apiErr))
			assert.Equal(t, tt.wantType, apiErr.Type)
			assert.Equal(t, tt.statusCode, apiErr.StatusCode)
			assert.Equal(t, "campaign request failed", apiErr.Message)
		})
	}
}

func TestAsCampaignNotReady(t *testing.T) {
	tests := []struct {
		name       string
		status     int
		body       string
		wantOK     bool
		wantIssues []responses.CampaignIssue
	}{
		{
			name:   "a campaign with issues",
			status: http.StatusBadRequest,
			body:   `{"message":"campaign is not ready to send","issues":[{"code":"subject_missing","message":"Add a subject before you send the campaign."},{"code":"unsubscribe_link_missing","message":"Add an unsubscribe link."}]}`,
			wantOK: true,
			wantIssues: []responses.CampaignIssue{
				{Code: responses.CampaignIssueSubjectMissing, Message: "Add a subject before you send the campaign."},
				{Code: responses.CampaignIssueUnsubscribeLinkMissing, Message: "Add an unsubscribe link."},
			},
		},
		{
			name:   "a refusal with only a message",
			status: http.StatusBadRequest,
			body:   `{"message":"send_at must be at least 1 minute ahead"}`,
		},
		{
			name:   "a refusal that is not a 400",
			status: http.StatusConflict,
			body:   `{"message":"campaign is not a draft","issues":[{"code":"subject_missing","message":"Add a subject."}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.body))
			})
			defer cleanup()

			_, _, err := client.CampaignsAPI.SendCampaign(context.Background(), uuid.New(), uuid.New(), requests.SendCampaignRequest{})
			require.Error(t, err)

			notReady, ok := AsCampaignNotReady(err)

			assert.Equal(t, tt.wantOK, ok)
			if tt.wantOK {
				require.NotNil(t, notReady)
				assert.Equal(t, "campaign is not ready to send", notReady.Message)
				assert.Equal(t, tt.wantIssues, notReady.Issues)
			} else {
				assert.Nil(t, notReady)
			}
		})
	}

	t.Run("an error that is not an API error", func(t *testing.T) {
		notReady, ok := AsCampaignNotReady(errors.New("network down"))

		assert.False(t, ok)
		assert.Nil(t, notReady)
	})
}
