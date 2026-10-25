/*
AhaSend API v2

Testing SegmentsAPIService against the OpenAPI Prism mock.
*/

package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/internal/prismmock"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ahasend_SegmentsAPIService(t *testing.T) {
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
	segmentID := uuid.MustParse("33333333-3333-4333-8333-333333333333")
	rawEmail := "User+Tag@Example.COM"
	definition := common.SegmentDefinition{
		SchemaVersion: common.SegmentSchemaVersion,
		Match:         common.SegmentMatchAll,
		Conditions: []common.SegmentRule{
			{Condition: &common.SegmentCondition{
				Field: "attributes.country",
				Op:    common.SegmentOpEqual,
				Value: &common.SegmentValue{String: ahasend.String("DE")},
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

	// Prism answers each operation with its first documented 2xx example.
	validatePrismResponse := func(t *testing.T, wantStatus int, response any, httpResponse *http.Response, err error) {
		t.Helper()
		require.NoError(t, err)
		require.NotNil(t, httpResponse)
		assert.Equal(t, wantStatus, httpResponse.StatusCode)
		assert.NotNil(t, response)
	}

	t.Run("GetSegments", func(t *testing.T) {
		response, httpResponse, err := apiClient.SegmentsAPI.GetSegments(auth, accountID, requests.GetSegmentsParams{})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("CreateSegment", func(t *testing.T) {
		response, httpResponse, err := apiClient.SegmentsAPI.CreateSegment(auth, accountID, requests.CreateSegmentRequest{
			Name:       "Engaged customers",
			Definition: definition,
			Timezone:   ahasend.String("Europe/Berlin"),
		}, WithIdempotencyKey("prism-create-segment"))
		validatePrismResponse(t, http.StatusCreated, response, httpResponse, err)
	})

	t.Run("PreviewSegment", func(t *testing.T) {
		response, httpResponse, err := apiClient.SegmentsAPI.PreviewSegment(auth, accountID, requests.PreviewSegmentRequest{
			Definition: definition,
		})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("GetSegment", func(t *testing.T) {
		response, httpResponse, err := apiClient.SegmentsAPI.GetSegment(auth, accountID, segmentID)
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("UpdateSegment", func(t *testing.T) {
		response, httpResponse, err := apiClient.SegmentsAPI.UpdateSegment(auth, accountID, segmentID, requests.UpdateSegmentRequest{
			Name:       ahasend.String("Engaged customers"),
			Definition: &definition,
		})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("DeleteSegment", func(t *testing.T) {
		response, httpResponse, err := apiClient.SegmentsAPI.DeleteSegment(auth, accountID, segmentID)
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("GetSegmentContacts", func(t *testing.T) {
		response, httpResponse, err := apiClient.SegmentsAPI.GetSegmentContacts(auth, accountID, segmentID, requests.GetSegmentContactsParams{})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})

	t.Run("GetContactSegments", func(t *testing.T) {
		response, httpResponse, err := apiClient.SegmentsAPI.GetContactSegments(auth, accountID, rawEmail, requests.GetContactSegmentsParams{})
		validatePrismResponse(t, http.StatusOK, response, httpResponse, err)
	})
}
