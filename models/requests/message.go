package requests

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
	"github.com/google/uuid"
)

// CreateMessageRequest represents a request to create and send an email message.
type CreateMessageRequest struct {
	From          common.SenderAddress    `json:"from"`
	Recipients    []common.Recipient      `json:"recipients"`
	Subject       string                  `json:"subject"`
	ReplyTo       *common.SenderAddress   `json:"reply_to,omitempty"`
	TextContent   *string                 `json:"text_content,omitempty"`
	HtmlContent   *string                 `json:"html_content,omitempty"`
	AmpContent    *string                 `json:"amp_content,omitempty"`
	Attachments   []common.Attachment     `json:"attachments,omitempty"`
	Headers       map[string]string       `json:"headers,omitempty"`
	Substitutions map[string]interface{}  `json:"substitutions,omitempty"`
	Tags          []string                `json:"tags,omitempty"`
	Sandbox       *bool                   `json:"sandbox,omitempty"`
	SandboxResult *string                 `json:"sandbox_result,omitempty"`
	Tracking      *common.Tracking        `json:"tracking,omitempty"`
	Retention     *common.Retention       `json:"retention,omitempty"`
	Schedule      *common.MessageSchedule `json:"schedule,omitempty"`
}

// CreateTemplateMessageRequest represents a request to send a stored
// transactional template to one or more recipients. The template supplies
// the body, so the request has no content fields and no request-level
// substitutions: each recipient's Substitutions hold its values for the
// template's variables.
type CreateTemplateMessageRequest struct {
	TemplateID uuid.UUID `json:"template_id"`
	// From is the sender. Leave it nil to send from the template's default
	// sender; the API rejects the send when the template has none. A From
	// given here is used even when the template has a default sender.
	From       *common.SenderAddress `json:"from,omitempty"`
	Recipients []common.Recipient    `json:"recipients"`
	// ReplyTo replaces the template's reply-to. The template's applies when
	// the request has neither ReplyTo nor a reply-to header.
	ReplyTo *common.SenderAddress `json:"reply_to,omitempty"`
	// Subject replaces the template's subject. Leave it empty to use the
	// template's.
	Subject       string                  `json:"subject,omitempty"`
	Attachments   []common.Attachment     `json:"attachments,omitempty"`
	Headers       map[string]string       `json:"headers,omitempty"`
	Tags          []string                `json:"tags,omitempty"`
	Sandbox       *bool                   `json:"sandbox,omitempty"`
	SandboxResult *string                 `json:"sandbox_result,omitempty"`
	Tracking      *common.Tracking        `json:"tracking,omitempty"`
	Retention     *common.Retention       `json:"retention,omitempty"`
	Schedule      *common.MessageSchedule `json:"schedule,omitempty"`
}

// CreateConversationMessageRequest represents a request to create and send a conversational email message.
type CreateConversationMessageRequest struct {
	From          common.SenderAddress    `json:"from"`
	To            []common.SenderAddress  `json:"to"`
	CC            []common.SenderAddress  `json:"cc,omitempty"`
	BCC           []common.SenderAddress  `json:"bcc,omitempty"`
	Subject       string                  `json:"subject"`
	ReplyTo       *common.SenderAddress   `json:"reply_to,omitempty"`
	TextContent   *string                 `json:"text_content,omitempty"`
	HtmlContent   *string                 `json:"html_content,omitempty"`
	AmpContent    *string                 `json:"amp_content,omitempty"`
	Attachments   []common.Attachment     `json:"attachments,omitempty"`
	Headers       map[string]string       `json:"headers,omitempty"`
	Tags          []string                `json:"tags,omitempty"`
	Sandbox       *bool                   `json:"sandbox,omitempty"`
	SandboxResult *string                 `json:"sandbox_result,omitempty"`
	Tracking      *common.Tracking        `json:"tracking,omitempty"`
	Retention     *common.Retention       `json:"retention,omitempty"`
	Schedule      *common.MessageSchedule `json:"schedule,omitempty"`
}

type GetMessagesParams struct {
	Status          *string
	Tags            []string
	Sender          *string
	Recipient       *string
	Subject         *string
	MessageIDHeader *string
	FromTime        *time.Time
	ToTime          *time.Time
	common.PaginationParams
}
