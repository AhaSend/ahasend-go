package requests

// AddMemberRequest represents a request to add a member to an account.
type AddMemberRequest struct {
	Email string  `json:"email"`
	Role  string  `json:"role"`
	Name  *string `json:"name,omitempty"`
}

// UpdateAccountRequest represents a request to update account settings.
type UpdateAccountRequest struct {
	Name                     *string `json:"name,omitempty"`
	Website                  *string `json:"website,omitempty"`
	About                    *string `json:"about,omitempty"`
	TrackOpens               *bool   `json:"track_opens,omitempty"`
	TrackClicks              *bool   `json:"track_clicks,omitempty"`
	RejectBadRecipients      *bool   `json:"reject_bad_recipients,omitempty"`
	RejectMistypedRecipients *bool   `json:"reject_mistyped_recipients,omitempty"`
	MessageMetadataRetention *int32  `json:"message_metadata_retention,omitempty"`
	MessageDataRetention     *int32  `json:"message_data_retention,omitempty"`
}
