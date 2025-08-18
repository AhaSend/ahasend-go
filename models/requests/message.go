package requests

import (
	"time"

	"github.com/AhaSend/ahasend-go/models/common"
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

type GetMessagesParams struct {
	Status          *string
	Sender          *string
	Recipient       *string
	Subject         *string
	MessageIdHeader *string
	FromTime        *time.Time
	ToTime          *time.Time
	Limit           *int32
	Cursor          *string
}
