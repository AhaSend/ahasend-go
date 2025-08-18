package responses

import (
	"time"

	"github.com/google/uuid"
)

// SMTPCredential represents an AhaSend SMTP credential
type SMTPCredential struct {
	Object    string    `json:"object"`
	ID        uint64    `json:"id"`
	AccountID uuid.UUID `json:"account_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Sandbox   bool      `json:"sandbox"`
	Scope     string    `json:"scope"`
	Domains   []string  `json:"scoped_domains"`
}
