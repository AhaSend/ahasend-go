package api

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/AhaSend/ahasend-go"
	"github.com/AhaSend/ahasend-go/internal/prismmock"
	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const templateAPIResponseJSON = `{
	"object":"template",
	"id":"11111111-1111-4111-8111-111111111111",
	"created_at":"2026-09-10T10:00:00Z",
	"updated_at":"2026-09-10T11:00:00Z",
	"name":"Password reset",
	"subject":"Reset your password",
	"preheader":"It expires in an hour",
	"variables":[{"name":"first_name","required":true,"kind":"value"},{"name":"reset_url","required":false,"kind":"value"}],
	"from":null,
	"reply_to":null,
	"editor":"html",
	"has_draft":false
}`

func Test_ahasend_TemplatesAPIService(t *testing.T) {
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
	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	versionID := uuid.MustParse("22222222-2222-4222-8222-222222222222")

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

	t.Run("GetTemplates", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.GetTemplates(auth, accountID, requests.GetTemplatesParams{})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("GetTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.GetTemplate(auth, accountID, templateID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("CreateTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.CreateTemplate(auth, accountID, requests.CreateTemplateRequest{
			Name:    "Password reset",
			Content: &requests.TemplateContentInput{HTML: ahasend.String("<p>Hi</p>")},
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("UpdateTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.UpdateTemplate(auth, accountID, templateID, requests.UpdateTemplateRequest{
			Subject: ahasend.String("Reset your password"),
			Content: &requests.TemplateContentInput{Text: ahasend.String("")},
		})
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("DeleteTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.DeleteTemplate(auth, accountID, templateID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("GetPublishedTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.GetPublishedTemplate(auth, accountID, templateID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("DiscardTemplateDraft", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.DiscardTemplateDraft(auth, accountID, templateID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("PublishTemplate", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.PublishTemplate(auth, accountID, templateID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("GetTemplateVersions", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.GetTemplateVersions(auth, accountID, templateID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("GetTemplateVersion", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.GetTemplateVersion(auth, accountID, templateID, versionID)
		validatePrismResponse(t, response, httpResponse, err)
	})

	t.Run("RestoreTemplateVersion", func(t *testing.T) {
		response, httpResponse, err := apiClient.TemplatesAPI.RestoreTemplateVersion(auth, accountID, templateID, versionID, requests.RestoreTemplateVersionRequest{Publish: true})
		validatePrismResponse(t, response, httpResponse, err)
	})
}

func TestTemplatesAPIGetTemplatesSerializesPagination(t *testing.T) {
	accountID := uuid.New()

	tests := []struct {
		name          string
		params        requests.GetTemplatesParams
		wantRawQuery  string
		wantAfterKey  bool
		wantBeforeKey bool
	}{
		{
			name:         "defaults the limit when none is given",
			params:       requests.GetTemplatesParams{},
			wantRawQuery: "limit=100",
		},
		{
			name: "sends the caller's limit and cursors",
			params: requests.GetTemplatesParams{
				Limit:  ahasend.Int32(25),
				After:  ahasend.String("after-cursor"),
				Before: ahasend.String("before-cursor"),
			},
			wantRawQuery:  "after=after-cursor&before=before-cursor&limit=25",
			wantAfterKey:  true,
			wantBeforeKey: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lastRequest *http.Request
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				lastRequest = r
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"object":"list","data":[` + templateAPIResponseJSON + `],"pagination":{"has_more":true,"next_cursor":"next","previous_cursor":null}}`))
			})
			defer cleanup()

			page, httpResponse, err := client.TemplatesAPI.GetTemplates(context.Background(), accountID, tt.params)

			require.NoError(t, err)
			require.NotNil(t, httpResponse)
			require.NotNil(t, lastRequest)
			assert.Equal(t, http.MethodGet, lastRequest.Method)
			assert.Equal(t, "/v2/accounts/"+accountID.String()+"/templates", lastRequest.URL.Path)
			// The raw query is asserted rather than a parsed map, so a cursor
			// sent as an empty value cannot pass for one that was never sent.
			assert.Equal(t, tt.wantRawQuery, lastRequest.URL.RawQuery)
			query := lastRequest.URL.Query()
			_, hasAfter := query["after"]
			_, hasBefore := query["before"]
			assert.Equal(t, tt.wantAfterKey, hasAfter)
			assert.Equal(t, tt.wantBeforeKey, hasBefore)

			require.NotNil(t, page)
			require.Len(t, page.Data, 1)
			assert.Equal(t, "Password reset", page.Data[0].Name)
			assert.True(t, page.Pagination.HasMore)
			require.NotNil(t, page.Pagination.NextCursor)
			assert.Equal(t, "next", *page.Pagination.NextCursor)
		})
	}
}

func TestTemplatesAPIGetTemplateUsesUUIDPathIDAndNoQuery(t *testing.T) {
	accountID := uuid.New()
	templateID := uuid.MustParse("11111111-1111-4111-8111-111111111111")
	var lastRequest *http.Request

	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		lastRequest = r
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(templateAPIResponseJSON))
	})
	defer cleanup()

	template, httpResponse, err := client.TemplatesAPI.GetTemplate(context.Background(), accountID, templateID)

	require.NoError(t, err)
	require.NotNil(t, httpResponse)
	require.NotNil(t, lastRequest)
	assert.Equal(t, http.MethodGet, lastRequest.Method)
	assert.Equal(t, "/v2/accounts/"+accountID.String()+"/templates/"+templateID.String(), lastRequest.URL.Path)
	assert.Empty(t, lastRequest.URL.RawQuery)

	require.NotNil(t, template)
	assert.Equal(t, templateID, template.ID)
	assert.Equal(t, "Password reset", template.Name)
	require.Len(t, template.Variables, 2)
	assert.True(t, template.Variables[0].Required)
	assert.False(t, template.Variables[1].Required)
	assert.Equal(t, responses.TemplateVariableKindValue, template.Variables[0].Kind)
}

func TestTemplatesAPIErrorTypes(t *testing.T) {
	accountID := uuid.New()
	templateID := uuid.New()

	tests := []struct {
		statusCode int
		wantType   ErrorType
	}{
		{statusCode: http.StatusBadRequest, wantType: ErrorTypeValidation},
		{statusCode: http.StatusUnauthorized, wantType: ErrorTypeAuthentication},
		{statusCode: http.StatusForbidden, wantType: ErrorTypePermission},
		{statusCode: http.StatusNotFound, wantType: ErrorTypeNotFound},
		{statusCode: http.StatusInternalServerError, wantType: ErrorTypeServer},
	}

	for _, tt := range tests {
		t.Run(http.StatusText(tt.statusCode), func(t *testing.T) {
			client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(`{"message":"template request failed"}`))
			})
			defer cleanup()

			template, httpResponse, err := client.TemplatesAPI.GetTemplate(context.Background(), accountID, templateID)

			require.Error(t, err)
			assert.NotNil(t, template)
			require.NotNil(t, httpResponse)
			assert.Equal(t, tt.statusCode, httpResponse.StatusCode)

			var apiErr *APIError
			require.True(t, errors.As(err, &apiErr))
			assert.Equal(t, tt.wantType, apiErr.Type)
			assert.Equal(t, tt.statusCode, apiErr.StatusCode)
			assert.Equal(t, "template request failed", apiErr.Message)
		})
	}
}

func TestTemplatesAPIRefusedMarkupListsEveryProblem(t *testing.T) {
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{
			"message":"content.markup has 2 problems",
			"errors":[
				{"code":"unknown_attribute","message":"Paragraph has no attribute \"color\"","line":4,"column":14,"path":"Email > Section[2] > Paragraph[1]","hint":"wrap the text in <Color value=\"...\"> instead"},
				{"code":"too_large","message":"content.markup is larger than 1 MiB"}
			]
		}`))
	})
	defer cleanup()

	markup := `<Email version="1"><Paragraph color="red">Hi</Paragraph></Email>`
	_, httpResponse, err := client.TemplatesAPI.UpdateTemplate(context.Background(), uuid.New(), uuid.New(), requests.UpdateTemplateRequest{
		Content: &requests.TemplateContentInput{Markup: &markup},
	})

	require.Error(t, err)
	require.NotNil(t, httpResponse)
	var apiErr *APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, ErrorTypeValidation, apiErr.Type)
	assert.Equal(t, "content.markup has 2 problems", apiErr.Message)
	assert.Equal(t, []common.MarkupProblem{
		{
			Code:    "unknown_attribute",
			Message: `Paragraph has no attribute "color"`,
			Line:    4,
			Column:  14,
			Path:    "Email > Section[2] > Paragraph[1]",
			Hint:    `wrap the text in <Color value="..."> instead`,
		},
		{Code: "too_large", Message: "content.markup is larger than 1 MiB"},
	}, apiErr.Errors)
}

func TestTemplatesAPIOtherErrorsHaveNoProblems(t *testing.T) {
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"content.mjml is not allowed on a simple template"}`))
	})
	defer cleanup()

	_, _, err := client.TemplatesAPI.UpdateTemplate(context.Background(), uuid.New(), uuid.New(), requests.UpdateTemplateRequest{
		Content: &requests.TemplateContentInput{MJML: ahasend.String("<mjml></mjml>")},
	})

	var apiErr *APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, "content.mjml is not allowed on a simple template", apiErr.Message)
	assert.Nil(t, apiErr.Errors)
}

func TestTemplatesAPIErrorsOfAnotherShapeKeepTheMessage(t *testing.T) {
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"x","errors":["y"]}`))
	})
	defer cleanup()

	_, _, err := client.TemplatesAPI.GetTemplate(context.Background(), uuid.New(), uuid.New())

	var apiErr *APIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, "x", apiErr.Message)
	assert.Nil(t, apiErr.Errors)
}

func TestTemplatesAPIUpdateSendsMarkupAndDecodesWarnings(t *testing.T) {
	var body string
	client, cleanup := newContractTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		body = string(raw)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"object":"template","editor":"simple","has_draft":true,"content":null,"markup_warnings":[{"code":"jinja_in_text","message":"Text holds {{","line":1,"column":30,"path":"Email > Paragraph[1]"}]}`))
	})
	defer cleanup()

	markup := `<Email version="1"><Paragraph>{{ name }}</Paragraph></Email>`
	template, _, err := client.TemplatesAPI.UpdateTemplate(context.Background(), uuid.New(), uuid.New(), requests.UpdateTemplateRequest{
		Content: &requests.TemplateContentInput{Markup: &markup},
	})

	require.NoError(t, err)
	assert.JSONEq(t, `{"content":{"markup":"<Email version=\"1\"><Paragraph>{{ name }}</Paragraph></Email>"}}`, body)
	assert.Equal(t, []common.MarkupProblem{{Code: common.MarkupWarningJinjaInText, Message: "Text holds {{", Line: 1, Column: 30, Path: "Email > Paragraph[1]"}}, template.MarkupWarnings)
}
