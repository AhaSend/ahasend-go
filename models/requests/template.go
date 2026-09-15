package requests

// GetTemplatesParams represents the pagination cursors accepted by the
// transactional templates list endpoint.
type GetTemplatesParams struct {
	Limit  *int32
	After  *string
	Before *string
}
