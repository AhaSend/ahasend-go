package api

import (
	"context"
	"net/http"

	"github.com/AhaSend/ahasend-go/models/common"
)

// UtilityAPIService UtilityAPI service
type UtilityAPIService service

/*
Ping Health check endpoint

Health check endpoint that returns a simple pong response

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param opts ...RequestOption - optional request options (timeout, retry, headers, etc.)
	@return SuccessResponse, *http.Response, error
*/
func (a *UtilityAPIService) Ping(
	ctx context.Context,
	opts ...RequestOption,
) (*common.SuccessResponse, *http.Response, error) {
	var result common.SuccessResponse

	config := RequestConfig{
		Method:       http.MethodGet,
		PathTemplate: "/v2/ping",
		PathParams:   nil, // No path parameters for ping
		QueryParams:  nil, // No query parameters for ping
		Body:         nil, // No request body for ping
		Result:       &result,
	}

	// Apply options
	for _, opt := range opts {
		opt(&config)
	}

	resp, err := a.client.Execute(ctx, config)
	return &result, resp, err
}
