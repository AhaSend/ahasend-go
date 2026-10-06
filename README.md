# AhaSend Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/AhaSend/ahasend-go.svg)](https://pkg.go.dev/github.com/AhaSend/ahasend-go)
[![License: MIT](https://img.shields.io/github/license/ahasend/ahasend-go)](https://opensource.org/licenses/MIT)

The Go SDK for the [AhaSend](https://ahasend.com) API. Use it to send email and to manage templates, contacts, lists, domains, webhooks, routes, suppressions and sub accounts.

It also:

- Keeps under the API's rate limits.
- Retries failed requests, waiting longer between each try.
- Adds an idempotency key to every create request, so a retried request is not applied twice.
- Checks and reads webhook requests.

## Install

```bash
go get github.com/AhaSend/ahasend-go
```

Requires Go 1.18 or later.

## Send an email

```go
package main

import (
    "context"
    "log"
    "os"

    "github.com/AhaSend/ahasend-go"
    "github.com/AhaSend/ahasend-go/api"
    "github.com/AhaSend/ahasend-go/models/common"
    "github.com/AhaSend/ahasend-go/models/requests"
    "github.com/google/uuid"
)

func main() {
    accountID, err := uuid.Parse(os.Getenv("AHASEND_ACCOUNT_ID"))
    if err != nil {
        log.Fatalf("invalid AHASEND_ACCOUNT_ID: %v", err)
    }

    client := api.NewAPIClient(api.WithAPIKey(os.Getenv("AHASEND_API_KEY")))

    message := requests.CreateMessageRequest{
        From:        common.SenderAddress{Email: "sender@yourdomain.com"},
        Recipients:  []common.Recipient{{Email: "recipient@example.com"}},
        Subject:     "Hello from AhaSend",
        HtmlContent: ahasend.String("<h1>Welcome</h1>"),
        TextContent: ahasend.String("Welcome"),
    }

    response, _, err := client.MessagesAPI.CreateMessage(context.Background(), accountID, message)
    if err != nil {
        log.Fatal(err)
    }

    if len(response.Data) > 0 && response.Data[0].ID != nil {
        log.Printf("sent, message ID %s", *response.Data[0].ID)
    }
}
```

## Send from a template

A transactional template made in the dashboard holds the subject, the preview text and both bodies. A send names the template and gives values for its variables:

```go
// Read the template to see which variables a send must give.
template, _, err := client.TemplatesAPI.GetTemplate(ctx, accountID, templateID)
if err != nil {
    log.Fatal(err)
}
for _, variable := range template.Variables {
    log.Printf("%s (required: %t)", variable.Name, variable.Required)
}

message := requests.CreateMessageRequest{
    From: common.SenderAddress{Email: "sender@yourdomain.com"},
    Recipients: []common.Recipient{
        {
            Email:         "recipient@example.com",
            Substitutions: map[string]interface{}{"first_name": "Pat"},
        },
    },
    TemplateID: &templateID,
}

response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
```

A template can also hold a default sender and reply-to address: `template.From` and `template.ReplyTo`, each nil when the template has none. A send from a template with a sender can leave `From` out:

```go
message := requests.CreateMessageRequest{
    // No From: the send uses the template's sender.
    Recipients: []common.Recipient{
        {
            Email:         "recipient@example.com",
            Substitutions: map[string]interface{}{"first_name": "Pat"},
        },
    },
    TemplateID: &templateID,
}
```

- `TemplateID` cannot be used with `TextContent`, `HtmlContent` or `AmpContent`.
- Leave `Subject` empty to use the template's subject, or set it to replace it.
- Leave `From.Email` empty to use the template's sender; `From.Name` is then ignored too. A `From` with an `Email` replaces the template's sender.
- The template's reply-to applies to every send from it, with or without `From`, unless the request sets `ReplyTo` or a `reply-to` entry in `Headers`; either one replaces it.
- The template's sender is checked as a sender in the request is: its domain must be in your account, have valid DNS records and not be paused.
- Variable values come from the request's `Substitutions` and from each recipient's. When both give the same variable, the recipient's value is used.
- AhaSend fills in `email`, `view_browser_url` and `unsubscribe_url` itself.
- The send fails with:
  - 404 when the template does not exist.
  - 400 when neither the request nor the template has a subject, or the template has no saved design.
  - 400 when neither the request nor the template has a sender.
  - 400 `the template's default sender or reply-to is not valid, edit it on the template page` when a stored value the send uses no longer passes the address checks.
  - 403 `this api key is not authorized to send messages` when the request has no sender and the API key cannot send from any domain.
  - 400 for the whole request when any recipient is missing a required variable.
- `GetTemplate` and `GetTemplates` need the `templates:read` scope. Sending from a template needs only the normal send scope.

`client.TemplatesAPI.GetTemplates(ctx, accountID, requests.GetTemplatesParams{})` lists the account's templates, newest first. For the next page, pass the response's `Pagination.NextCursor` as `After`. For the previous page, pass `Pagination.PreviousCursor` as `Before`.

## API keys

Every request needs an API key. Get one from the [AhaSend dashboard](https://dashboard.ahasend.com). There are three ways to give it to the SDK.

From the environment:

```bash
export AHASEND_API_KEY="aha-sk-your-64-character-key"
```

```go
client := api.NewAPIClientFromEnv()
```

When you create the client:

```go
client := api.NewAPIClient(api.WithAPIKey(apiKey))
```

For one request only:

```go
ctx := context.WithValue(context.Background(), api.ContextAccessToken, "aha-sk-your-64-character-key")
response, _, err := client.MessagesAPI.CreateMessage(ctx, accountID, message)
```

### Sub account scopes

A parent or partner key that manages sub accounts needs one or more of these scopes:

| Scope | Allows |
|---|---|
| `sub-accounts:read` | List and read sub accounts |
| `sub-accounts:write` | Create and update sub accounts |
| `sub-accounts:delete` | Delete sub accounts |
| `sub-accounts:suspend` | Pause and resume sub accounts (`SuspendSubAccount`, `UnsuspendSubAccount`), and unpause their domains |
| `sub-accounts:usage` | Read each sub account's usage and cost |
| `sub-account-api-keys:read` | List and read sub accounts' API keys |
| `sub-account-api-keys:write` | Create and update sub accounts' API keys |
| `sub-account-api-keys:delete` | Delete sub accounts' API keys |

### IP allow list

An API key can carry `IPAllowList`, the source IPs that can authenticate with the key:

- Each entry is a CIDR block, such as `203.0.113.0/24`, or a bare IPv4 or IPv6 address (stored as a `/32` or `/128`). Entries are canonicalized (host bits are masked) and de-duplicated.
- The API refuses the allow-all prefixes `0.0.0.0/0` and `::/0`. At most 100 entries are allowed after de-duplication; the API answers 400 to a longer list.
- An empty list, the default, allows every source IP.
- When the list is not empty, the API answers 403 on every v2 endpoint to a request from an IP outside the list, whatever the key's scopes.
- On `CreateAPIKey` and `CreateSubAccountAPIKey`, leave `IPAllowList` nil or empty to allow every source IP.
- On `UpdateAPIKey` and `UpdateSubAccountAPIKey`:
  - A nil `IPAllowList` (omitted from the request) leaves the list as it is.
  - `&[]string{}` clears the list, so the key works from every IP.
  - A pointer to a non-empty slice replaces the list.
- A key that updates its own list so that it no longer covers the caller's IP gets 409. A parent key that updates a sub account key has no such check.

## Contacts

- Scopes: `contacts:read` to list and read, `contacts:write` to create, update and batch upsert, `contacts:delete` to delete.
- `UpdateContact` changes only the fields you set:
  - A nil field is left as it is.
  - A pointer to `""` clears a text field.
  - `Unsubscribed` set to `false` subscribes the contact again.
- `BatchUpsertContacts` returns 200 even when some items fail. Check `Failed` and each `Data[i].Outcome`.
- Find a contact by ID or by email. For an email that contains `/`, use the ID.

## Lists

- Scopes:
  - `lists:read` to read lists and their members.
  - `lists:write` to create and update lists, and to add or remove members.
  - `lists:delete` to delete lists.
  - `IncludeContacts` on `GetListContacts` also needs `contacts:read`.
- `ContactCount` is the number of members a campaign to the list would reach, not the number of all members.
- `BatchAddListContacts` returns 200 even when some items fail. Check `Failed` and each `Data[i].Outcome`. A contact already on the list is reported as `already_member` and left as it is.
- To stop mail to one member, use `UpsertListContact` with the status `unsubscribed`. This keeps the record of their choice. `DeleteListContact` removes the member and that record with it.
- A member whose status is `complained` cannot be changed or removed. Both calls return 409.

## Domains

- Each domain has a sending type, `common.DomainSendingTypeTransactional` or `common.DomainSendingTypeMarketing`. The sending type affects deliverability: marketing email from a transactional domain can get your account paused.
  - On `CreateDomain`, a nil `SendingType` creates a transactional domain.
  - On `UpdateDomain`, a nil `SendingType` leaves it as it is. A change applies to new messages within five minutes.
  - The API refuses an empty `SendingType`.
- To list only one sending type, use `GetDomainsWithParams`:

  ```go
  response, _, err := client.DomainsAPI.GetDomainsWithParams(ctx, accountID, requests.GetDomainsParams{
      SendingType: ahasend.String(common.DomainSendingTypeMarketing),
  })
  ```

- AhaSend can pause sending from one domain. Then `Paused` is true, `PausedAt` tells when, and `PauseReason` tells why. The only reason today is `responses.DomainPauseReasonBounceRate` (too many recent emails from the domain bounced), but AhaSend can add others: keep a default branch. While a domain is paused, the API refuses new email from it with 403 (except sandbox messages), and you cannot delete or rename it. Your other domains continue to send.
- A parent account can lift the pause on a sub account's domain with `client.SubAccountsAPI.UnpauseSubAccountDomain(ctx, accountID, subAccountID, "example.com")`. It needs the `sub-accounts:suspend` scope. The call is idempotent: on a domain that is not paused, it changes nothing and returns the domain. The change can take some minutes to apply to new email.
- `DKIMSelector` on `CreateDomainRequest` and `UpdateDomainRequest` sets a per-domain DKIM selector (Platform Partner accounts only). On create, nil or an empty string uses the default selector. On update, nil leaves it as it is, and a pointer to `""` clears the override. `Domain.DKIMSelector` reports the override, not always the selector used for signing.

## Services

| Service | Use it to | Main methods |
|---|---|---|
| `MessagesAPI` | Send and manage email | `CreateMessage`, `GetMessage`, `CancelMessage` |
| `TemplatesAPI` | Read transactional templates | `GetTemplates`, `GetTemplate` |
| `ContactsAPI` | Manage contacts | `GetContacts`, `GetContact`, `CreateContact`, `UpdateContact`, `DeleteContact`, `BatchUpsertContacts` |
| `ListsAPI` | Manage lists and their members | `GetLists`, `CreateList`, `GetList`, `UpdateList`, `DeleteList`, `GetListContacts`, `BatchAddListContacts`, `UpsertListContact`, `DeleteListContact`, `GetContactLists` |
| `DomainsAPI` | Add and check sending domains | `CreateDomain`, `CheckDomainDNS`, `GetDomain`, `GetDomainsWithParams` |
| `WebhooksAPI` | Manage webhooks | `CreateWebhook`, `UpdateWebhook`, `GetWebhooks` |
| `StatisticsAPI` | Read sending statistics | `GetDeliverabilityStatistics`, `GetBounceStatistics` |
| `SuppressionsAPI` | Manage addresses that must not be emailed | `CreateSuppression`, `DeleteSuppression`, `GetSuppressions` |
| `RoutesAPI` | Handle incoming email | `CreateRoute`, `UpdateRoute` |
| `AccountsAPI` | Manage the account and its members | `GetAccount`, `AddAccountMember` |
| `APIKeysAPI` | Manage API keys | `CreateAPIKey`, `UpdateAPIKey` |
| `SubAccountsAPI` | Manage sub accounts, their API keys and the pause of their domains | `ListSubAccounts`, `CreateSubAccount`, `CreateSubAccountAPIKey`, `GetSubAccountsUsage`, `UnpauseSubAccountDomain` |

## Examples

The [examples](./examples/) folder has a runnable program for each common task:

- [send_email.go](./examples/send_email.go): send an email
- [send_template.go](./examples/send_template.go): send from a transactional template
- [send_with_attachments.go](./examples/send_with_attachments.go): send with attachments
- [batch_send.go](./examples/batch_send.go): send to many recipients
- [scheduled_send.go](./examples/scheduled_send.go): send later
- [webhook_processing.go](./examples/webhook_processing.go): receive webhook requests
- [webhook_management.go](./examples/webhook_management.go): create and manage webhooks
- [domain_management.go](./examples/domain_management.go): add and check a domain
- [statistics.go](./examples/statistics.go): read statistics
- [error_handling.go](./examples/error_handling.go): handle errors
- [rate_limiting.go](./examples/rate_limiting.go): set rate limits
- [idempotency.go](./examples/idempotency.go): avoid sending twice
- [list_management.go](./examples/list_management.go): create lists and manage their members
- [sub_account_management.go](./examples/sub_account_management.go): manage sub accounts, usage and their API keys

To run one:

```bash
export AHASEND_API_KEY="your-api-key"
export AHASEND_ACCOUNT_ID="your-account-id"
go run examples/send_email.go
```

## Webhooks

The `webhooks` package checks each request's signature (it follows the Standard Webhooks spec) and reads the event:

```go
package main

import (
    "errors"
    "log"
    "net/http"

    "github.com/AhaSend/ahasend-go/webhooks"
)

func main() {
    verifier, err := webhooks.NewWebhookVerifier("your-webhook-secret")
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
        event, err := verifier.ParseRequest(r)
        switch {
        case errors.Is(err, webhooks.ErrUnknownEventType):
            // Signed, but this SDK version does not know the event type.
            log.Printf("ignoring webhook: %v", err)
            w.WriteHeader(http.StatusOK)
            return
        case errors.Is(err, webhooks.ErrMissingHeaders),
            errors.Is(err, webhooks.ErrInvalidSignature),
            errors.Is(err, webhooks.ErrExpiredTimestamp),
            errors.Is(err, webhooks.ErrInvalidTimestamp):
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        case err != nil:
            http.Error(w, "invalid webhook", http.StatusBadRequest)
            return
        }

        switch e := event.(type) {
        case *webhooks.MessageDeliveredEvent:
            log.Printf("delivered to %s", e.Data.Recipient)
        case *webhooks.MessageBouncedEvent:
            // DeliveryAttempt is nil when no delivery attempt was recorded.
            // Log the codes, not Response or Description: those are free text
            // that often contains the recipient and the message content.
            if a := e.Data.DeliveryAttempt; a != nil {
                log.Printf("bounced with SMTP code %d", a.SMTPCode)
                if a.Classification != nil {
                    log.Printf("  classification: %s", *a.Classification)
                }
            } else {
                log.Printf("bounced, no delivery attempt recorded")
            }
        }

        w.WriteHeader(http.StatusOK)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

New values can appear in `DeliveryAttempt.Classification` at any time. Handle the `webhooks.Classification*` values you care about, and keep a `default` case for the rest. Do not reject a request because of a value you do not know: when more than 100 attempts in a row fail, retries included, the webhook or route is automatically disabled.

For the same reason, answer 2xx when `ParseRequest` returns `webhooks.ErrUnknownEventType`. The SDK checks the signature first, so that error means a signed event that this SDK version does not know yet. The SDK wraps its errors: compare them with `errors.Is`, not `==`.

For a campaign message, `from` in message events includes the sender's name, as `Name <address>`. In route events, `to` and `reply_to` can carry display names and several addresses, and `spam_score` (a `*float64`) can be below 0 or above 10.

Events:
- `message.reception`, `message.delivered`, `message.transient_error`, `message.failed`, `message.bounced`, `message.suppressed`, `message.opened`, `message.clicked`
- `suppression.created`
- `domain.dns_error`
- `message.routing` (the SDK also reads the legacy name `route.message` as this event)

## Settings

Rate limits:

```go
client := api.NewAPIClient(api.WithAPIKey(apiKey))

client.SetSendMessageRateLimit(500, 1000) // 500 requests a second, bursts up to 1000
client.SetStatisticsRateLimit(10, 20)     // 10 requests a second, bursts up to 20
```

Retries:

```go
client := api.NewAPIClient(
    api.WithAPIKey(apiKey),
    api.WithRetryConfig(api.RetryConfig{
        Enabled:         true,
        MaxRetries:      3,
        BackoffStrategy: api.BackoffExponential,
        BaseDelay:       time.Second,
        MaxDelay:        30 * time.Second,
    }),
)
```

## Development

Run `make help` to see every command. The common ones:

- `make setup`: install the tools
- `make dev-test`: format, lint and run the tests
- `make full-test`: all tests with coverage
- `make test-unit`: unit tests
- `make test-integration`: tests against a mock server (needs Prism)
- `make test-coverage`: coverage report

### Code samples

Each operation's Go sample in the API reference is a program at `codesamples/<operationId>/main.go`. `go build ./...` compiles them, so a sample always matches the SDK. After you change one, run `make code-samples` to write it into `openapi/openapi.yaml`. CI runs `make check-code-samples` to check the file is up to date.

## Related

- [AhaSend CLI](https://github.com/AhaSend/ahasend-cli): a command-line tool built on this SDK

## Help

- [API documentation](https://ahasend.com/docs)
- [Go package documentation](https://pkg.go.dev/github.com/AhaSend/ahasend-go)
- [Email support](mailto:support@ahasend.com)
- [Report an issue](https://github.com/AhaSend/ahasend-go/issues)

## License

MIT. See [LICENSE](LICENSE).
