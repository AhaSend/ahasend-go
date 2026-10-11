package responses

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const campaignTemplateFixture = `{
	"object":"campaign_template",
	"id":"44444444-4444-4444-8444-444444444444",
	"created_at":"2026-10-01T10:00:00Z",
	"updated_at":"2026-10-01T11:00:00Z",
	"name":"Monthly newsletter",
	"editor":"advanced",
	"content":{"mjml":"<mjml></mjml>","html":"<html></html>","text":"News","text_is_custom":false}
}`

func TestCampaignTemplateDecodesAPIFixture(t *testing.T) {
	var template CampaignTemplate
	require.NoError(t, json.Unmarshal([]byte(campaignTemplateFixture), &template))

	assert.Equal(t, "campaign_template", template.Object)
	assert.Equal(t, uuid.MustParse("44444444-4444-4444-8444-444444444444"), template.ID)
	assert.Equal(t, time.Date(2026, 10, 1, 10, 0, 0, 0, time.UTC), template.CreatedAt.UTC())
	assert.Equal(t, "Monthly newsletter", template.Name)
	assert.Equal(t, common.TemplateEditorAdvanced, template.Editor)
	require.NotNil(t, template.Content)
	assert.Equal(t, TemplateContent{MJML: "<mjml></mjml>", HTML: "<html></html>", Text: "News"}, *template.Content)

	encoded, err := json.Marshal(template)
	require.NoError(t, err)
	assert.JSONEq(t, campaignTemplateFixture, string(encoded))
}

func TestCampaignTemplateDecodesNullContent(t *testing.T) {
	var template CampaignTemplate
	require.NoError(t, json.Unmarshal([]byte(`{"object":"campaign_template","name":"Empty","editor":"simple","content":null}`), &template))

	assert.Nil(t, template.Content)
}

// campaignFixture is a draft with every nullable field set.
const campaignFixture = `{
	"object":"campaign",
	"id":"55555555-5555-4555-8555-555555555555",
	"created_at":"2026-10-01T10:00:00Z",
	"updated_at":"2026-10-01T11:00:00Z",
	"name":"October newsletter",
	"subject":"Our news",
	"preheader":"Three things we shipped",
	"from":{"email":"news@example.com","name":""},
	"reply_to":"support@example.com",
	"status":"draft",
	"status_reason":null,
	"send_at":null,
	"send_resumes_at":null,
	"finished_at":null,
	"list_ids":["66666666-6666-4666-8666-666666666666"],
	"editor":"html",
	"content":{"html":"<p>News</p>","text":"News","text_is_custom":true},
	"ready":true,
	"issues":[]
}`

func TestCampaignDecodesAPIFixture(t *testing.T) {
	var campaign Campaign
	require.NoError(t, json.Unmarshal([]byte(campaignFixture), &campaign))

	assert.Equal(t, "campaign", campaign.Object)
	assert.Equal(t, uuid.MustParse("55555555-5555-4555-8555-555555555555"), campaign.ID)
	assert.Equal(t, "October newsletter", campaign.Name)
	require.NotNil(t, campaign.Subject)
	assert.Equal(t, "Our news", *campaign.Subject)
	require.NotNil(t, campaign.Preheader)
	assert.Equal(t, "Three things we shipped", *campaign.Preheader)
	require.NotNil(t, campaign.From)
	assert.Equal(t, "news@example.com", campaign.From.Email)
	require.NotNil(t, campaign.From.Name)
	assert.Empty(t, *campaign.From.Name)
	require.NotNil(t, campaign.ReplyTo)
	assert.Equal(t, "support@example.com", *campaign.ReplyTo)
	assert.Equal(t, common.CampaignStatusDraft, campaign.Status)
	assert.Nil(t, campaign.StatusReason)
	assert.Nil(t, campaign.SendAt)
	assert.Equal(t, []uuid.UUID{uuid.MustParse("66666666-6666-4666-8666-666666666666")}, campaign.ListIDs)
	require.NotNil(t, campaign.Editor)
	assert.Equal(t, common.TemplateEditorHTML, *campaign.Editor)
	require.NotNil(t, campaign.Content)
	assert.True(t, campaign.Content.TextIsCustom)
	require.NotNil(t, campaign.Ready)
	assert.True(t, *campaign.Ready)
	require.NotNil(t, campaign.Issues)
	assert.Empty(t, campaign.Issues)
}

func TestCampaignRoundTripsWithoutLoss(t *testing.T) {
	var decoded Campaign
	require.NoError(t, json.Unmarshal([]byte(campaignFixture), &decoded))

	encoded, err := json.Marshal(decoded)
	require.NoError(t, err)
	assert.JSONEq(t, campaignFixture, string(encoded))
}

func TestCampaignDecodesASentCampaign(t *testing.T) {
	payload := []byte(`{
		"object":"campaign",
		"id":"55555555-5555-4555-8555-555555555555",
		"created_at":"2026-10-01T10:00:00Z",
		"updated_at":"2026-10-01T11:00:00Z",
		"name":"October newsletter",
		"subject":null,
		"preheader":null,
		"from":null,
		"reply_to":null,
		"status":"canceled",
		"status_reason":"The account was suspended.",
		"send_at":"2026-10-02T09:00:00Z",
		"send_resumes_at":"2026-10-03T00:00:00Z",
		"finished_at":"2026-10-02T09:30:00Z",
		"list_ids":[],
		"editor":null,
		"content":null,
		"ready":null,
		"issues":null
	}`)

	var campaign Campaign
	require.NoError(t, json.Unmarshal(payload, &campaign))

	assert.Nil(t, campaign.Subject)
	assert.Nil(t, campaign.Preheader)
	assert.Nil(t, campaign.From)
	assert.Nil(t, campaign.ReplyTo)
	assert.Equal(t, common.CampaignStatusCanceled, campaign.Status)
	require.NotNil(t, campaign.StatusReason)
	assert.Equal(t, "The account was suspended.", *campaign.StatusReason)
	require.NotNil(t, campaign.SendAt)
	assert.Equal(t, time.Date(2026, 10, 2, 9, 0, 0, 0, time.UTC), campaign.SendAt.UTC())
	require.NotNil(t, campaign.SendResumesAt)
	require.NotNil(t, campaign.FinishedAt)
	assert.Equal(t, time.Date(2026, 10, 2, 9, 30, 0, 0, time.UTC), campaign.FinishedAt.UTC())
	assert.Nil(t, campaign.Editor)
	assert.Nil(t, campaign.Content)
	assert.Nil(t, campaign.Ready)
	assert.Nil(t, campaign.Issues)

	encoded, err := json.Marshal(campaign)
	require.NoError(t, err)
	assert.JSONEq(t, string(payload), string(encoded))
}

func TestCampaignSendDecodesWarnings(t *testing.T) {
	payload := []byte(`{"object":"campaign","id":"55555555-5555-4555-8555-555555555555","status":"validating","ready":null,"issues":null,"warnings":["The view-in-browser page cannot show this design."]}`)

	var send CampaignSend
	require.NoError(t, json.Unmarshal(payload, &send))

	assert.Equal(t, uuid.MustParse("55555555-5555-4555-8555-555555555555"), send.ID)
	assert.Equal(t, common.CampaignStatusValidating, send.Status)
	assert.Equal(t, []string{"The view-in-browser page cannot show this design."}, send.Warnings)
}

func TestCampaignNotReadyDecodesIssues(t *testing.T) {
	var notReady CampaignNotReady
	require.NoError(t, json.Unmarshal([]byte(`{"message":"campaign is not ready to send","issues":[{"code":"design_missing","message":"Add a design before you send the campaign."}]}`), &notReady))

	assert.Equal(t, "campaign is not ready to send", notReady.Message)
	assert.Equal(t, []CampaignIssue{{Code: CampaignIssueDesignMissing, Message: "Add a design before you send the campaign."}}, notReady.Issues)
}

func TestCampaignTestDecodesAnUnknownResult(t *testing.T) {
	var test CampaignTest
	require.NoError(t, json.Unmarshal([]byte(`{"object":"campaign_test","sent_to":[],"failed":[],"result_unknown":true}`), &test))

	assert.Equal(t, "campaign_test", test.Object)
	assert.Empty(t, test.SentTo)
	assert.Empty(t, test.Failed)
	assert.True(t, test.ResultUnknown)
}

func TestCampaignStatsDecodesAReport(t *testing.T) {
	payload := []byte(`{
		"object":"campaign_stats",
		"campaign_id":"55555555-5555-4555-8555-555555555555",
		"recipients":100,
		"delivered":95,
		"opens":60,
		"clicks":12,
		"unsubscribes":1,
		"complaints":0,
		"report":{
			"computed_at":"2026-10-03T00:00:00Z",
			"recipients":100,
			"openers":40,
			"clickers":10,
			"bounces":5,
			"deferrals":2,
			"by_status":{"Delivered":95,"Bounced":5},
			"by_bounce_classification":{"InvalidRecipient":4,"Unknown":1}
		},
		"rates":{"open_rate":42.11,"click_rate":10.53,"bounce_rate":5,"deferral_rate":2},
		"recount_failed":true
	}`)

	var stats CampaignStats
	require.NoError(t, json.Unmarshal(payload, &stats))

	assert.Equal(t, "campaign_stats", stats.Object)
	assert.Equal(t, 95, stats.Delivered)
	require.NotNil(t, stats.Report)
	assert.Equal(t, time.Date(2026, 10, 3, 0, 0, 0, 0, time.UTC), stats.Report.ComputedAt.UTC())
	assert.Equal(t, 40, stats.Report.Openers)
	assert.Equal(t, map[string]int{"Delivered": 95, "Bounced": 5}, stats.Report.ByStatus)
	assert.Equal(t, map[string]int{"InvalidRecipient": 4, "Unknown": 1}, stats.Report.ByBounceClassification)
	require.NotNil(t, stats.Rates.OpenRate)
	assert.InDelta(t, 42.11, *stats.Rates.OpenRate, 0.001)
	require.NotNil(t, stats.Rates.BounceRate)
	assert.InDelta(t, 5.0, *stats.Rates.BounceRate, 0.001)
	assert.True(t, stats.RecountFailed)

	encoded, err := json.Marshal(stats)
	require.NoError(t, err)
	assert.JSONEq(t, string(payload), string(encoded))
}

func TestCampaignStatsDecodesNoReport(t *testing.T) {
	payload := []byte(`{
		"object":"campaign_stats",
		"campaign_id":"55555555-5555-4555-8555-555555555555",
		"recipients":0,
		"delivered":0,
		"opens":0,
		"clicks":0,
		"unsubscribes":0,
		"complaints":0,
		"report":null,
		"rates":{"open_rate":null,"click_rate":null,"bounce_rate":null,"deferral_rate":null},
		"recount_failed":false
	}`)

	var stats CampaignStats
	require.NoError(t, json.Unmarshal(payload, &stats))

	assert.Nil(t, stats.Report)
	assert.Equal(t, CampaignStatsRates{}, stats.Rates)

	encoded, err := json.Marshal(stats)
	require.NoError(t, err)
	assert.JSONEq(t, string(payload), string(encoded))
}

func TestPaginatedCampaignRecipientsResponseDecodesPage(t *testing.T) {
	payload := []byte(`{
		"object":"list",
		"data":[
			{"object":"campaign_recipient","contact_id":"77777777-7777-4777-8777-777777777777","email":"ann@example.com","list_name":"Customers","status":"Bounced","bounce_classification":"InvalidRecipient","bounce_reason":"bad mailbox","bounce_message":"550 5.1.1 user unknown","attempts":1,"created_at":"2026-10-02T09:00:00Z","updated_at":"2026-10-02T09:05:00Z"},
			{"object":"campaign_recipient","contact_id":"88888888-8888-4888-8888-888888888888","email":"bob@example.com","list_name":"Customers","status":"Delivered","bounce_classification":null,"bounce_reason":null,"bounce_message":null,"attempts":1,"created_at":"2026-10-02T09:00:00Z","updated_at":"2026-10-02T09:01:00Z"}
		],
		"pagination":{"has_more":true,"next_cursor":"next"}
	}`)

	var page PaginatedCampaignRecipientsResponse
	require.NoError(t, json.Unmarshal(payload, &page))

	require.Len(t, page.Data, 2)
	assert.Equal(t, uuid.MustParse("77777777-7777-4777-8777-777777777777"), page.Data[0].ContactID)
	require.NotNil(t, page.Data[0].BounceClassification)
	assert.Equal(t, "InvalidRecipient", *page.Data[0].BounceClassification)
	require.NotNil(t, page.Data[0].BounceMessage)
	assert.Equal(t, "550 5.1.1 user unknown", *page.Data[0].BounceMessage)
	assert.Nil(t, page.Data[1].BounceClassification)
	assert.Nil(t, page.Data[1].BounceReason)
	assert.Nil(t, page.Data[1].BounceMessage)
	assert.True(t, page.Pagination.HasMore)
	require.NotNil(t, page.Pagination.NextCursor)
	assert.Nil(t, page.Pagination.PreviousCursor)
}
