package api

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// emailStatisticsFixture is a response without a breakdown: two daily rows
// with a null sender_domain, and a null breakdown.
const emailStatisticsFixture = `{
	"object":"email_statistics",
	"from_time":"2026-10-12T00:00:00Z",
	"to_time":"2026-10-14T00:00:00Z",
	"group_by":"day",
	"breakdown":null,
	"totals":{"reception_count":30,"delivered_count":27,"bounced_count":1,"suppressed_count":2,"complained_count":1,"deferred_count":4,"opened_count":12,"clicked_count":3},
	"data":[
		{"from_timestamp":"2026-10-12T00:00:00Z","to_timestamp":"2026-10-13T00:00:00Z","sender_domain":null,"reception_count":20,"delivered_count":18,"bounced_count":1,"suppressed_count":1,"complained_count":1,"deferred_count":3,"opened_count":9,"clicked_count":2},
		{"from_timestamp":"2026-10-13T00:00:00Z","to_timestamp":"2026-10-14T00:00:00Z","sender_domain":null,"reception_count":10,"delivered_count":9,"bounced_count":0,"suppressed_count":1,"complained_count":0,"deferred_count":1,"opened_count":3,"clicked_count":1}
	],
	"pagination":{"has_more":false}
}`

// emailStatisticsBreakdownPageFixture is the first page of a breakdown by
// domain, with a next page.
const emailStatisticsBreakdownPageFixture = `{
	"object":"email_statistics",
	"from_time":"2026-08-01T00:00:00Z",
	"to_time":"2026-11-01T00:00:00Z",
	"group_by":"month",
	"breakdown":"domain",
	"totals":{"reception_count":184200,"delivered_count":181050,"bounced_count":1900,"suppressed_count":1250,"complained_count":12,"deferred_count":2310,"opened_count":60210,"clicked_count":8400},
	"data":[
		{"from_timestamp":"2026-08-01T00:00:00Z","to_timestamp":"2026-09-01T00:00:00Z","sender_domain":"example.com","reception_count":61000,"delivered_count":60100,"bounced_count":610,"suppressed_count":290,"complained_count":4,"deferred_count":800,"opened_count":20100,"clicked_count":2800},
		{"from_timestamp":"2026-09-01T00:00:00Z","to_timestamp":"2026-10-01T00:00:00Z","sender_domain":"example.com","reception_count":0,"delivered_count":0,"bounced_count":0,"suppressed_count":0,"complained_count":0,"deferred_count":0,"opened_count":0,"clicked_count":0}
	],
	"pagination":{"has_more":true,"next_cursor":"ZXhhbXBsZS5jb20="}
}`

func TestStatisticsAPIGetEmailStatisticsSendsItsParameters(t *testing.T) {
	accountID := uuid.MustParse("aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa")

	var lastRequest *http.Request
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(emailStatisticsBreakdownPageFixture))
	})
	defer cleanup()

	t.Run("every parameter reaches the query", func(t *testing.T) {
		fromTime := time.Date(2026, 8, 1, 2, 0, 0, 0, time.FixedZone("CEST", 2*60*60))
		toTime := time.Date(2026, 11, 1, 0, 0, 0, 0, time.UTC)

		_, _, err := client.StatisticsAPI.GetEmailStatistics(context.Background(), accountID, requests.GetEmailStatisticsParams{
			FromTime:     &fromTime,
			ToTime:       &toTime,
			GroupBy:      ahasend.String("month"),
			SenderDomain: ahasend.String("example.com,example.org"),
			Breakdown:    ahasend.String("domain"),
			Limit:        ahasend.Int32(10),
			After:        ahasend.String("ZXhhbXBsZS5jb20="),
		})

		require.NoError(t, err)
		require.NotNil(t, lastRequest)
		assert.Equal(t, http.MethodGet, lastRequest.Method)
		assert.Equal(t, "/v2/accounts/"+accountID.String()+"/statistics/emails", lastRequest.URL.EscapedPath())
		assert.Equal(t, url.Values{
			"from_time":     {"2026-08-01T02:00:00+02:00"},
			"to_time":       {"2026-11-01T00:00:00Z"},
			"group_by":      {"month"},
			"sender_domain": {"example.com,example.org"},
			"breakdown":     {"domain"},
			"limit":         {"10"},
			"after":         {"ZXhhbXBsZS5jb20="},
		}, lastRequest.URL.Query())
	})

	t.Run("unset parameters are not sent", func(t *testing.T) {
		_, _, err := client.StatisticsAPI.GetEmailStatistics(context.Background(), accountID, requests.GetEmailStatisticsParams{})

		require.NoError(t, err)
		require.NotNil(t, lastRequest)
		assert.Equal(t, "/v2/accounts/"+accountID.String()+"/statistics/emails", lastRequest.URL.EscapedPath())
		assert.Empty(t, lastRequest.URL.RawQuery)
	})
}

func TestStatisticsAPIGetEmailStatisticsDecodesTheResponse(t *testing.T) {
	get := func(t *testing.T, body string) *responses.EmailStatisticsResponse {
		t.Helper()
		client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(body))
		})
		defer cleanup()

		resp, httpResp, err := client.StatisticsAPI.GetEmailStatistics(context.Background(), uuid.New(), requests.GetEmailStatisticsParams{})
		require.NoError(t, err)
		require.NotNil(t, httpResp)
		require.NotNil(t, resp)
		return resp
	}

	t.Run("a null breakdown and null sender domains are nil", func(t *testing.T) {
		resp := get(t, emailStatisticsFixture)

		assert.Equal(t, "email_statistics", resp.Object)
		assert.Equal(t, time.Date(2026, 10, 12, 0, 0, 0, 0, time.UTC), resp.FromTime)
		assert.Equal(t, time.Date(2026, 10, 14, 0, 0, 0, 0, time.UTC), resp.ToTime)
		assert.Equal(t, "day", resp.GroupBy)
		assert.Nil(t, resp.Breakdown)
		assert.Equal(t, responses.EmailStatisticsTotals{
			ReceptionCount: 30, DeliveredCount: 27, BouncedCount: 1, SuppressedCount: 2,
			ComplainedCount: 1, DeferredCount: 4, OpenedCount: 12, ClickedCount: 3,
		}, resp.Totals)

		require.Len(t, resp.Data, 2)
		assert.Nil(t, resp.Data[0].SenderDomain)
		assert.Nil(t, resp.Data[1].SenderDomain)
		assert.Equal(t, time.Date(2026, 10, 12, 0, 0, 0, 0, time.UTC), resp.Data[0].FromTimestamp)
		assert.Equal(t, time.Date(2026, 10, 13, 0, 0, 0, 0, time.UTC), resp.Data[0].ToTimestamp)
		assert.Equal(t, responses.EmailStatisticsTotals{
			ReceptionCount: 20, DeliveredCount: 18, BouncedCount: 1, SuppressedCount: 1,
			ComplainedCount: 1, DeferredCount: 3, OpenedCount: 9, ClickedCount: 2,
		}, resp.Data[0].EmailStatisticsTotals)

		assert.False(t, resp.Pagination.HasMore)
		assert.Nil(t, resp.Pagination.NextCursor)
		assert.Nil(t, resp.Pagination.PreviousCursor)
	})

	t.Run("a breakdown page carries its domains and the next cursor", func(t *testing.T) {
		resp := get(t, emailStatisticsBreakdownPageFixture)

		assert.Equal(t, "month", resp.GroupBy)
		require.NotNil(t, resp.Breakdown)
		assert.Equal(t, "domain", *resp.Breakdown)
		assert.Equal(t, responses.EmailStatisticsTotals{
			ReceptionCount: 184200, DeliveredCount: 181050, BouncedCount: 1900, SuppressedCount: 1250,
			ComplainedCount: 12, DeferredCount: 2310, OpenedCount: 60210, ClickedCount: 8400,
		}, resp.Totals)

		require.Len(t, resp.Data, 2)
		for _, row := range resp.Data {
			require.NotNil(t, row.SenderDomain)
			assert.Equal(t, "example.com", *row.SenderDomain)
		}
		assert.Equal(t, time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC), resp.Data[0].ToTimestamp)
		assert.Equal(t, int64(61000), resp.Data[0].ReceptionCount)
		assert.Equal(t, int64(2800), resp.Data[0].ClickedCount)

		assert.True(t, resp.Pagination.HasMore)
		require.NotNil(t, resp.Pagination.NextCursor)
		assert.Equal(t, "ZXhhbXBsZS5jb20=", *resp.Pagination.NextCursor)
		assert.Nil(t, resp.Pagination.PreviousCursor)
	})

	t.Run("a refused request is an APIError with the API's message", func(t *testing.T) {
		client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"message":"limit is accepted only with breakdown=domain"}`))
		})
		defer cleanup()

		_, httpResp, err := client.StatisticsAPI.GetEmailStatistics(context.Background(), uuid.New(), requests.GetEmailStatisticsParams{
			Limit: ahasend.Int32(10),
		})

		require.Error(t, err)
		require.NotNil(t, httpResp)
		var apiErr *APIError
		require.ErrorAs(t, err, &apiErr)
		assert.Equal(t, ErrorTypeValidation, apiErr.Type)
		assert.Equal(t, "limit is accepted only with breakdown=domain", apiErr.Message)
	})
}

func TestOpenAPIGetEmailStatisticsOperation(t *testing.T) {
	contents, err := os.ReadFile("../openapi/openapi.yaml")
	require.NoError(t, err)

	var spec struct {
		Paths map[string]map[string]struct {
			OperationID string                      `yaml:"operationId"`
			CodeSamples []struct{ Lang string }     `yaml:"x-code-samples"`
			Security    []map[string][]string       `yaml:"security"`
			Responses   map[string]responseContract `yaml:"responses"`
			Parameters  []struct {
				Name string `yaml:"name"`
				In   string `yaml:"in"`
			} `yaml:"parameters"`
		} `yaml:"paths"`
	}
	require.NoError(t, yaml.Unmarshal(contents, &spec))

	path, ok := spec.Paths["/v2/accounts/{account_id}/statistics/emails"]
	require.True(t, ok)
	operation, ok := path["get"]
	require.True(t, ok)
	assert.Equal(t, "getEmailStatistics", operation.OperationID)

	var goSamples int
	for _, sample := range operation.CodeSamples {
		if sample.Lang == "go" {
			goSamples++
		}
	}
	assert.Equal(t, 1, goSamples)

	assert.Equal(t, []map[string][]string{
		{"BearerAuth": {"statistics:read:all"}},
		{"BearerAuth": {"statistics:read:{domain}"}},
	}, operation.Security)

	statuses := make([]string, 0, len(operation.Responses))
	for status := range operation.Responses {
		statuses = append(statuses, status)
	}
	sort.Strings(statuses)
	assert.Equal(t, []string{"200", "400", "401", "403", "429", "500"}, statuses)

	// GetEmailStatisticsParams sends exactly these query parameters.
	var query []string
	for _, parameter := range operation.Parameters {
		if parameter.In == "query" {
			query = append(query, parameter.Name)
		}
	}
	sort.Strings(query)
	assert.Equal(t, []string{"after", "breakdown", "from_time", "group_by", "limit", "sender_domain", "to_time"}, query)
}

func TestOpenAPIEmailStatisticsSchemasMatchModels(t *testing.T) {
	tests := []struct {
		schemas []string
		model   any
	}{
		{schemas: []string{"EmailStatisticsTotals"}, model: responses.EmailStatisticsTotals{}},
		// EmailStatistics is an allOf of the period fields and a $ref to
		// EmailStatisticsTotals.
		{schemas: []string{"EmailStatistics", "EmailStatisticsTotals"}, model: responses.EmailStatistics{}},
		{schemas: []string{"EmailStatisticsResponse"}, model: responses.EmailStatisticsResponse{}},
	}

	for _, tt := range tests {
		t.Run(tt.schemas[0], func(t *testing.T) {
			assert.Equal(t, openAPISchemaProperties(t, tt.schemas...), modelJSONFields(tt.model))
		})
	}
}
