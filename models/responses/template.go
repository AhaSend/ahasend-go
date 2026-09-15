package responses

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// TemplateVariable represents one variable a template's design uses.
type TemplateVariable struct {
	Name string `json:"name"`
	// Required reports whether a send has to supply a value for this
	// variable. A variable the design wraps in a fallback is not required,
	// and neither are the names the send supplies itself.
	Required bool `json:"required"`
}

// Template represents an account's transactional template.
type Template struct {
	Object    string             `json:"object"`
	ID        uuid.UUID          `json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Name      string             `json:"name"`
	Subject   string             `json:"subject"`
	Preheader string             `json:"preheader"`
	Variables []TemplateVariable `json:"variables"`
}

// PaginatedTemplatesResponse represents a page of transactional templates.
type PaginatedTemplatesResponse struct {
	Object     string                `json:"object"`
	Data       []Template            `json:"data"`
	Pagination common.PaginationInfo `json:"pagination"`
}
