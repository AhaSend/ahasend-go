package common

// PaginationInfo represents pagination metadata for list responses.
type PaginationInfo struct {
	HasMore        bool    `json:"has_more"`
	NextCursor     *string `json:"next_cursor,omitempty"`
	PreviousCursor *string `json:"previous_cursor,omitempty"`
}

// PaginationParams represents pagination parameters for list requests.
type PaginationParams struct {
	Limit  *int32  `json:"limit,omitempty"`
	Cursor *string `json:"cursor,omitempty"` // Backward compatibility
	After  *string `json:"after,omitempty"`  // New bidirectional pagination
	Before *string `json:"before,omitempty"` // New bidirectional pagination
}

// PaginatedResponse represents a generic paginated response.
type PaginatedResponse[T any] struct {
	Object     string         `json:"object"`
	Data       []T            `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
}
