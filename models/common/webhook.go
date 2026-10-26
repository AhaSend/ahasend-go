package common

// Webhook origins: who made a change that a contact or suppression event
// reports. A webhook's Origins lists the origins whose events it gets; nil
// means every origin. Message events and domain.dns_error have no origin, and
// a webhook always gets them.
//
//   - WebhookOriginRecipient: the contact or recipient, through an unsubscribe
//     link, a one-click unsubscribe, or a spam complaint.
//   - WebhookOriginAPI: you, through the AhaSend API.
//   - WebhookOriginDashboard: you, through the AhaSend dashboard.
//   - WebhookOriginSystem: AhaSend, for example after a bounce, an expired
//     message or an email validation, or when AhaSend support makes the
//     change.
const (
	WebhookOriginRecipient = "recipient"
	WebhookOriginAPI       = "api"
	WebhookOriginDashboard = "dashboard"
	WebhookOriginSystem    = "system"
)
