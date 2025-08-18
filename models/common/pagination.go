package common

// PaginationInfo represents pagination metadata for list responses.
type PaginationInfo struct {
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor,omitempty"`
}

// PaginatedResponse represents a generic paginated response.
type PaginatedResponse[T any] struct {
	Object     string         `json:"object"`
	Data       []T            `json:"data"`
	Pagination PaginationInfo `json:"pagination"`
}
