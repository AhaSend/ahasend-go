package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/AhaSend/ahasend-go/models/requests"
	"github.com/AhaSend/ahasend-go/models/responses"
	"github.com/google/uuid"
)

// SegmentsAPIService SegmentsAPI service
type SegmentsAPIService service

/*
GetSegments List Segments

# Returns the account's segments in newest-first order

Query Parameters:
- `limit`: Maximum number of segments to return (1-100, default: 100)
- `after`: Cursor for the next page; cannot be combined with `before`
- `before`: Cursor for the previous page; cannot be combined with `after`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param params GetSegmentsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedSegmentsResponse, *http.Response, error
*/
func (a *SegmentsAPIService) GetSegments(
	ctx context.Context,
	accountId uuid.UUID,
	params requests.GetSegmentsParams,
	opts ...RequestOption,
) (*responses.PaginatedSegmentsResponse, *http.Response, error) {
	var result responses.PaginatedSegmentsResponse

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/segments",
		PathParams:   map[string]string{"account_id": accountId.String()},
		QueryParams:  segmentPageQuery(params.Limit, params.After, params.Before),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
CreateSegment Create Segment

# Creates a segment and starts finding its members

A new segment is building until its members are found; a large account can
take some minutes. An account has at most 100 segments.

Returns 409 when the account already has 100 segments, or a list or campaign
the definition names is being changed. A 409 with Idempotent-Replayed: false
means a request with the same key is still running; the error carries its
Retry-After. Unlike most operations, a 4xx answer
is not stored for the idempotency key, so a retry with the same key runs
again.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request CreateSegmentRequest - the segment to create
	@param opts ...RequestOption - optional request options, including WithIdempotencyKey
	@return Segment, *http.Response, error
*/
func (a *SegmentsAPIService) CreateSegment(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.CreateSegmentRequest,
	opts ...RequestOption,
) (*responses.Segment, *http.Response, error) {
	var result responses.Segment

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/segments",
		PathParams:   map[string]string{"account_id": accountId.String()},
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
PreviewSegment Preview Segment

# Returns how many contacts a definition matches now, and up to 10 of them, without saving anything

The count stops after 10 seconds and is nil then; the contacts are a separate
read, returned either way. It needs the contacts:read scope as well as
segments:read.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param request PreviewSegmentRequest - the definition to try
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SegmentPreview, *http.Response, error
*/
func (a *SegmentsAPIService) PreviewSegment(
	ctx context.Context,
	accountId uuid.UUID,
	request requests.PreviewSegmentRequest,
	opts ...RequestOption,
) (*responses.SegmentPreview, *http.Response, error) {
	var result responses.SegmentPreview

	config := RequestConfig{
		Method:       http.MethodPost,
		PathTemplate: "/v2/accounts/{account_id}/segments/preview",
		PathParams:   map[string]string{"account_id": accountId.String()},
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetSegment Get Segment

# Returns one segment, with the members a campaign to it alone would reach

MarketableCount is nil while the segment is not ready, and when the members
cannot be counted within 10 seconds.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param segmentId Segment ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SegmentDetail, *http.Response, error
*/
func (a *SegmentsAPIService) GetSegment(
	ctx context.Context,
	accountId uuid.UUID,
	segmentId uuid.UUID,
	opts ...RequestOption,
) (*responses.SegmentDetail, *http.Response, error) {
	var result responses.SegmentDetail

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/segments/{segment_id}",
		PathParams:   segmentPathParams(accountId, segmentId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
UpdateSegment Update Segment

# Changes the fields the request names and leaves the others as they are

A changed definition or time zone builds the segment's members again; the
segment is building until that ends. Any update of a failed segment, an
empty one included, builds it again the same way.

Returns 409 while the segment is being changed, such as by a member update
that holds it. Nothing was changed; retry the request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param segmentId Segment ID
	@param request UpdateSegmentRequest - the fields to change
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return Segment, *http.Response, error
*/
func (a *SegmentsAPIService) UpdateSegment(
	ctx context.Context,
	accountId uuid.UUID,
	segmentId uuid.UUID,
	request requests.UpdateSegmentRequest,
	opts ...RequestOption,
) (*responses.Segment, *http.Response, error) {
	var result responses.Segment

	config := RequestConfig{
		Method:       http.MethodPut,
		PathTemplate: "/v2/accounts/{account_id}/segments/{segment_id}",
		PathParams:   segmentPathParams(accountId, segmentId),
		Body:         request,
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
DeleteSegment Delete Segment

# Deletes one segment; its members are removed shortly after

Returns 409 while a campaign still needs the segment, as an included or an
excluded segment: campaigns in draft, validating, scheduled, sending, paused
or waiting_for_review block the delete. It also returns 409 while the
segment is being changed. A second delete returns 404.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param segmentId Segment ID
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *SegmentsAPIService) DeleteSegment(
	ctx context.Context,
	accountId uuid.UUID,
	segmentId uuid.UUID,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodDelete,
		PathTemplate: "/v2/accounts/{account_id}/segments/{segment_id}",
		PathParams:   segmentPathParams(accountId, segmentId),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetSegmentContacts List Segment Contacts

# Returns the segment's members by contact ID, from the lowest

A member matches the rules whatever its consent. While the segment is
building, its members may still follow its previous rules. It needs the
contacts:read scope as well as segments:read.

Query Parameters:
- `limit`: Maximum number of members to return (1-100, default: 100)
- `after`: Cursor for the next page; cannot be combined with `before`
- `before`: Cursor for the previous page; cannot be combined with `after`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param segmentId Segment ID
	@param params GetSegmentContactsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedSegmentContactsResponse, *http.Response, error
*/
func (a *SegmentsAPIService) GetSegmentContacts(
	ctx context.Context,
	accountId uuid.UUID,
	segmentId uuid.UUID,
	params requests.GetSegmentContactsParams,
	opts ...RequestOption,
) (*responses.PaginatedSegmentContactsResponse, *http.Response, error) {
	var result responses.PaginatedSegmentContactsResponse

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/segments/{segment_id}/contacts",
		PathParams:   segmentPathParams(accountId, segmentId),
		QueryParams:  segmentPageQuery(params.Limit, params.After, params.Before),
		Result:       &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

/*
GetContactSegments List a Contact's Segments

# Returns the segments one contact is a member of, in the order of GetSegments

A segment that is building may still list the contact by its previous rules.
It needs the contacts:read scope as well as segments:read.

Query Parameters:
- `limit`: Maximum number of segments to return (1-100, default: 100)
- `after`: Cursor for the next page; cannot be combined with `before`
- `before`: Cursor for the previous page; cannot be combined with `after`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountId Account ID
	@param idOrEmail Contact UUID or raw email address
	@param params GetContactSegmentsParams - query parameters
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return PaginatedSegmentsResponse, *http.Response, error
*/
func (a *SegmentsAPIService) GetContactSegments(
	ctx context.Context,
	accountId uuid.UUID,
	idOrEmail string,
	params requests.GetContactSegmentsParams,
	opts ...RequestOption,
) (*responses.PaginatedSegmentsResponse, *http.Response, error) {
	var result responses.PaginatedSegmentsResponse

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/accounts/{account_id}/contacts/{id_or_email}/segments",
		PathParams: map[string]string{
			"account_id":  accountId.String(),
			"id_or_email": idOrEmail,
		},
		QueryParams: segmentPageQuery(params.Limit, params.After, params.Before),
		Result:      &result,
	}
	applyRequestOptions(&config, opts)

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}

// segmentPathParams returns the path parameters of one segment.
func segmentPathParams(accountId, segmentId uuid.UUID) map[string]string {
	return map[string]string{
		"account_id": accountId.String(),
		"segment_id": segmentId.String(),
	}
}

// segmentPageQuery returns the query of a segment page, with limit=100 when
// limit is nil.
func segmentPageQuery(limit *int32, after, before *string) url.Values {
	queryParams := url.Values{}
	if limit != nil {
		queryParams.Set("limit", fmt.Sprintf("%d", *limit))
	} else {
		queryParams.Set("limit", "100")
	}
	if after != nil {
		queryParams.Set("after", *after)
	}
	if before != nil {
		queryParams.Set("before", *before)
	}
	return queryParams
}
