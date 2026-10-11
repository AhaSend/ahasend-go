package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/internal/prismmock"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_CampaignTemplatesAPIService(t *testing.T) {
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
	templateID := uuid.MustParse("44444444-4444-4444-8444-444444444444")

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

	t.Run("ListCampaignTemplates", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignTemplatesAPI.ListCampaignTemplates(auth, accountID, requests.ListCampaignTemplatesParams{})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("GetCampaignTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignTemplatesAPI.GetCampaignTemplate(auth, accountID, templateID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("CreateCampaignTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignTemplatesAPI.CreateCampaignTemplate(auth, accountID, requests.CreateCampaignTemplateRequest{
			Name:    "Monthly newsletter",
			Content: &requests.TemplateContentInput{HTML: ahasend.String("<p>News</p>")},
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("UpdateCampaignTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignTemplatesAPI.UpdateCampaignTemplate(auth, accountID, templateID, requests.UpdateCampaignTemplateRequest{
			Name: ahasend.String("Monthly newsletter, 2026"),
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("DeleteCampaignTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.CampaignTemplatesAPI.DeleteCampaignTemplate(auth, accountID, templateID)
		validatePrismResponse(t, response, httpResponse, err)
	})
}

func TestCampaignTemplatesAPIListCampaignTemplatesSerializesPagination(t *testing.T) {
	accountID := uuid.New()

	tests := []struct {
		name         string
		params       requests.ListCampaignTemplatesParams
		wantRawQuery string
	}{
		{
			name:         "defaults the limit when none is given",
			params:       requests.ListCampaignTemplatesParams{},
			wantRawQuery: "limit=100",
		},
		{
			name: "sends the caller's limit and cursors",
			params: requests.ListCampaignTemplatesParams{
				Limit:  ahasend.Int32(25),
				After:  ahasend.String("after-cursor"),
				Before: ahasend.String("before-cursor"),
			},
			wantRawQuery: "after=after-cursor&before=before-cursor&limit=25",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lastRequest *http.Request
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				lastRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				// List Campaign Templates leaves out each item's content.
				_, _ = w.Write([]byte(`{"object":"list","data":[{"object":"campaign_template","id":"44444444-4444-4444-8444-444444444444","created_at":"2026-10-01T10:00:00Z","updated_at":"2026-10-01T11:00:00Z","name":"Monthly newsletter","editor":"advanced"}],"pagination":{"has_more":true,"next_cursor":"next"}}`))
			})
			defer cleanup()

			page, httpResponse, err := client.CampaignTemplatesAPI.ListCampaignTemplates(context.Background(), accountID, tt.params)

			require.NoError(t, err)
			require.NotNil(t, httpResponse)
			require.NotNil(t, lastRequest)
			assert.Equal(t, http.MethodGet, lastRequest.Method)
			assert.Equal(t, "/v2/accounts/"+accountID.String()+"/campaign-templates", lastRequest.URL.Path)
			// The raw query is asserted rather than a parsed map, so a cursor
			// sent as an empty value cannot pass for one that was never sent.
			assert.Equal(t, tt.wantRawQuery, lastRequest.URL.RawQuery)

			require.NotNil(t, page)
			require.Len(t, page.Data, 1)
			assert.Equal(t, "Monthly newsletter", page.Data[0].Name)
			assert.Nil(t, page.Data[0].Content)
			assert.True(t, page.Pagination.HasMore)
			require.NotNil(t, page.Pagination.NextCursor)
			assert.Equal(t, "next", *page.Pagination.NextCursor)
			assert.Nil(t, page.Pagination.PreviousCursor)
		})
	}
}
