package common

// Domain sending types. The sending type tells AhaSend which type of email a
// domain sends, and it affects deliverability: marketing email from a
// transactional domain can get the account paused.
const (
	DomainSendingTypeTransactional = "transactional"
	DomainSendingTypeMarketing     = "marketing"
)
