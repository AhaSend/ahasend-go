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
	// From is the default sender, used by a send that names no sender. It is
	// nil when the template has none, and then every send must name one.
	From *common.SenderAddress `json:"from,omitempty"`
	// ReplyTo is the default reply-to address, used by a send that sets no
	// reply-to. It is "" when the template has none.
	ReplyTo string `json:"reply_to"`
}

// PaginatedTemplatesResponse represents a page of transactional templates.
type PaginatedTemplatesResponse struct {
	Object     string                `json:"object"`
	Data       []Template            `json:"data"`
	Pagination common.PaginationInfo `json:"pagination"`
}
