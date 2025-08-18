package responses

import (
	"github.com/AhaSend/ahasend-go/models/common"
)

// Paginated response type aliases using the generic PaginatedResponse[T]
type PaginatedMessagesResponse = common.PaginatedResponse[Message]
type PaginatedDomainsResponse = common.PaginatedResponse[Domain]
type PaginatedWebhooksResponse = common.PaginatedResponse[Webhook]
type PaginatedRoutesResponse = common.PaginatedResponse[Route]
type PaginatedSMTPCredentialsResponse = common.PaginatedResponse[SMTPCredential]
type PaginatedSuppressionsResponse = common.PaginatedResponse[Suppression]
type PaginatedAPIKeysResponse = common.PaginatedResponse[APIKey]
