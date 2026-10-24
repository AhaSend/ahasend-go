package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// CampaignsAPIService CampaignsAPI service
type CampaignsAPIService service

/*
ListCampaigns List Campaigns

# Returns the account's marketing campaigns in newest-first order

Each campaign is listed without its audience, content and issues;
GetCampaign returns them.

Query Parameters:
- `status`: Only campaigns in this status
- `limit`: Maximum number of campaigns to return (1-100, default: 100)
- `after`: Cursor for the next page
- `before`: Cursor for the previous page

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params ListCampaignsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedCampaignsResponse, *http.Response, error
*/
func (a *CampaignsAPIService) ListCampaigns(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.ListCampaignsParams,
	opts ...RequestOption,
) (*responses.PaginatedCampaignsResponse, *http.Response, error) {
	var result responses.PaginatedCampaignsResponse

	queryParams := url.Values{}
	if params.Status != nil {
		queryParams.Set("status", *params.Status)
	}
	if params.Limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *params.Limit))
	} else {
		queryParams.Set("limit", "100")
	}
	if params.After != nil {
		queryParams.Set("after", *params.After)
	}
	if params.Before != nil {
		queryParams.Set("before", *params.Before)
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/campaigns",
		PathParams:   map[string]string{"account_id": accountId.String()},
		QueryParams:  queryParams,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetCampaign Get Campaign

# Returns one of the account's campaigns with its audience, its content and, for a draft, its issues

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Campaign, *http.Response, error
*/
func (a *CampaignsAPIService) GetCampaign(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	opts ...RequestOption,
) (*responses.Campaign, *http.Response, error) {
	var result responses.Campaign

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}",
		PathParams:   campaignPathParams(accountId, campaignId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
CreateCampaign Create Campaign

# Creates a draft marketing campaign

Every field but Name can be set later with UpdateCampaign; the campaign
returns its issues until it is ready to send. The design comes from Content,
from an Editor alone, or from a campaign template named by TemplateID, which
the campaign copies. The audience is ListIDs, SegmentIDs and the exclusions;
it needs at least one list or segment to send to once it is set.

Returns 400 for a list, segment or campaign template that does not exist, too
many lists or segments, or an audience with no list or segment to send to.
Returns 409 while a segment of the audience is being rebuilt. The API stores
that refusal for the idempotency key and replays it for 24 hours, so retry it
with a new key. Returns 503 when the AhaML in Content.Markup cannot be read
now; retry after Retry-After.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateCampaignRequest - the campaign to create
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Campaign, *http.Response, error
*/
func (a *CampaignsAPIService) CreateCampaign(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateCampaignRequest,
	opts ...RequestOption,
) (*responses.Campaign, *http.Response, error) {
	var result responses.Campaign

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/campaigns",
		PathParams:   map[string]string{"account_id": accountId.String()},
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
UpdateCampaign Update Campaign

# Changes a draft campaign

A campaign that is not a draft cannot be changed, and returns 409. An update
that changes the audience also returns 409 while a segment it names is being
rebuilt; the API stores that refusal for the idempotency key, so retry it with
a new key. Returns 400 for a list or segment that does not exist, too many
lists or segments, or an audience with no list or segment to send to. The
change waits for any other change of the same campaign, from the dashboard or the
API. Between this request and a dashboard editor open on the same campaign,
the later design save wins. It returns 503 when the AhaML in Content.Markup
cannot be read now; retry after Retry-After.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param request UpdateCampaignRequest - the fields to change
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Campaign, *http.Response, error
*/
func (a *CampaignsAPIService) UpdateCampaign(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	request requests.UpdateCampaignRequest,
	opts ...RequestOption,
) (*responses.Campaign, *http.Response, error) {
	var result responses.Campaign

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}",
		PathParams:   campaignPathParams(accountId, campaignId),
		Body:         request,
		Result:       &result,
		idempotent:   true,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DeleteCampaign Delete Campaign

# Deletes a draft campaign with its design and its audience choices

A campaign that is not a draft cannot be deleted, nor one that a segment
refers to in a campaign condition; both return 409, and the second names the
segments. A second delete returns 404.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *CampaignsAPIService) DeleteCampaign(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}",
		PathParams:   campaignPathParams(accountId, campaignId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DuplicateCampaign Duplicate Campaign

# Makes a new draft from a campaign in any status

The copy gets the campaign's design, its subject, preview text and sender, and
its lists and segments that still exist, included and excluded, the included
lists in the same order. When none of its included lists and segments still
exists, the copy has no lists or segments. The copy's design is its own.

Returns 409 while a segment of the audience is being rebuilt. The API stores
that refusal for the idempotency key and replays it for 24 hours, so retry it
with a new key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param request DuplicateCampaignRequest - the copy's name
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Campaign, *http.Response, error
*/
func (a *CampaignsAPIService) DuplicateCampaign(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	request requests.DuplicateCampaignRequest,
	opts ...RequestOption,
) (*responses.Campaign, *http.Response, error) {
	var result responses.Campaign

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}/duplicate",
		PathParams:   campaignPathParams(accountId, campaignId),
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
SendCampaign Send Campaign

# Sends a draft campaign now, or at SendAt

A campaign with any issue is refused with 400 and the issues; AsCampaignNotReady
returns them. The issue text_unchecked clears by itself: call again in a
moment. The send then checks what would stop the campaign at its send
time: its sender domain must still be verified and not paused, the account must
be able to send, and its lists must have subscribed contacts. For a send at a
later time, a block that a payment or the next billing period can clear is a
warning instead of a refusal.

The campaign moves to validating and goes through the same list validation and
review as a send from the dashboard. Read its status with GetCampaign. To change
the send time, unschedule the campaign and send it again.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param request SendCampaignRequest - when to send
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return CampaignSend, *http.Response, error
*/
func (a *CampaignsAPIService) SendCampaign(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	request requests.SendCampaignRequest,
	opts ...RequestOption,
) (*responses.CampaignSend, *http.Response, error) {
	var result responses.CampaignSend

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}/send",
		PathParams:   campaignPathParams(accountId, campaignId),
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
UnscheduleCampaign Unschedule Campaign

# Returns a validating or scheduled campaign to draft

It cancels the list validation that runs, or the scheduled send. A campaign
that has queued mail, or that is in any other status, cannot be unscheduled,
and returns 409.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Campaign, *http.Response, error
*/
func (a *CampaignsAPIService) UnscheduleCampaign(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	opts ...RequestOption,
) (*responses.Campaign, *http.Response, error) {
	var result responses.Campaign

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}/unschedule",
		PathParams:   campaignPathParams(accountId, campaignId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
TestCampaign Test Campaign

# Sends a test of a draft campaign to up to 5 addresses

Each address gets the design with placeholders for the contact's fields. A
campaign with any issue but unsubscribe_link_missing is refused with 400 and
the issues; AsCampaignNotReady returns them. The issue text_unchecked clears by
itself: call again in a moment. An account can send 50 test
emails per day (UTC), counted per address, from the API and the dashboard
together. Past that, the request returns 429; its message currently reads
"daily test send limit reached". The limit clears only when the UTC day
changes, and the client retries the 429 as it retries any 429.

Each test is checked with the template engine before it is sent. A test the
check refuses is answered 400 with a message, and nothing is sent: content
that cannot be rendered with a recipient's values, a call to random,
randrange, lipsum or now, namespace, a work, output, memory or time limit, or
an address the check cannot read. When the check cannot run, the API answers
503 with a Retry-After, and the client retries it with the same key.

With an idempotency key, a 2xx answer is replayed, so a retry with the same
key sends nothing again. A 4xx answer is not stored, so a retry with the same
key is answered afresh. A 502 means nothing was sent: retry it with the same
key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param request TestCampaignRequest - the addresses to send to
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return CampaignTest, *http.Response, error
*/
func (a *CampaignsAPIService) TestCampaign(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	request requests.TestCampaignRequest,
	opts ...RequestOption,
) (*responses.CampaignTest, *http.Response, error) {
	var result responses.CampaignTest

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}/test",
		PathParams:   campaignPathParams(accountId, campaignId),
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetCampaignStats Get Campaign Stats

# Returns a campaign's live counters, its stored report and the rates from both

The stored report is recounted from time to time. A request that finds it
missing or old asks for a recount and returns what is stored, so read again
later for the new numbers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return CampaignStats, *http.Response, error
*/
func (a *CampaignsAPIService) GetCampaignStats(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	opts ...RequestOption,
) (*responses.CampaignStats, *http.Response, error) {
	var result responses.CampaignStats

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}/stats",
		PathParams:   campaignPathParams(accountId, campaignId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
ListCampaignRecipients List Campaign Recipients

# Returns a campaign's recipients in one status or with one bounce classification, by contact ID, the highest first

Query Parameters:
- `status`: A key of the stats report's ByStatus, such as `Delivered`
- `bounce_classification`: A key of the stats report's ByBounceClassification, except `Unknown`
- `limit`: Maximum number of recipients to return (1-100, default: 100)
- `after`: Cursor for the next page
- `before`: Cursor for the previous page

Set exactly one of Status and BounceClassification.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param campaignId Campaign ID
	@param params ListCampaignRecipientsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedCampaignRecipientsResponse, *http.Response, error
*/
func (a *CampaignsAPIService) ListCampaignRecipients(
	ctx context.Context,
	accountId uuid.UUID,
	campaignId uuid.UUID,
	params requests.ListCampaignRecipientsParams,
	opts ...RequestOption,
) (*responses.PaginatedCampaignRecipientsResponse, *http.Response, error) {
	var result responses.PaginatedCampaignRecipientsResponse

	queryParams := url.Values{}
	if params.Status != nil {
		queryParams.Set("status", *params.Status)
	}
	if params.BounceClassification != nil {
		queryParams.Set("bounce_classification", *params.BounceClassification)
	}
	if params.Limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *params.Limit))
	} else {
		queryParams.Set("limit", "100")
	}
	if params.After != nil {
		queryParams.Set("after", *params.After)
	}
	if params.Before != nil {
		queryParams.Set("before", *params.Before)
	}

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/campaigns/{campaign_id}/recipients",
		PathParams:   campaignPathParams(accountId, campaignId),
		QueryParams:  queryParams,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

// AsCampaignNotReady reports whether err is the 400 that SendCampaign or
// TestCampaign returns for a campaign with issues, and returns its body with
// the issues that block the request. Every other refusal returns false.
func AsCampaignNotReady(err error) (*responses.CampaignNotReady, bool) {
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusBadRequest {
		return nil, false
	}
	var notReady responses.CampaignNotReady
	if json.Unmarshal(apiErr.Raw, &notReady) != nil || len(notReady.Issues) == 0 {
		return nil, false
	}
	return &notReady, true
}

func campaignPathParams(accountId, campaignId uuid.UUID) map[string]string {
	return map[string]string{
		"account_id":  accountId.String(),
		"campaign_id": campaignId.String(),
	}
}
