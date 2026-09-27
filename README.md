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

- `TemplateID` cannot be used with `TextContent`, `HtmlContent` or `AmpContent`.
- Leave `Subject` empty to use the template's subject, or set it to replace it.
- Variable values come from the request's `Substitutions` and from each recipient's. When both give the same variable, the recipient's value is used.
- AhaSend fills in `email`, `view_browser_url` and `unsubscribe_url` itself.
- The send fails with:
  - 404 when the template does not exist.
  - 400 when neither the request nor the template has a subject, or the template has no saved design.
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
| `sub-accounts:suspend` | Suspend and unsuspend sub accounts |
| `sub-accounts:usage` | Read each sub account's usage and cost |
| `sub-account-api-keys:read` | List and read sub accounts' API keys |
| `sub-account-api-keys:write` | Create and update sub accounts' API keys |
| `sub-account-api-keys:delete` | Delete sub accounts' API keys |

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

## Services

| Service | Use it to | Main methods |
|---|---|---|
| `MessagesAPI` | Send and manage email | `CreateMessage`, `GetMessage`, `CancelMessage` |
| `TemplatesAPI` | Read transactional templates | `GetTemplates`, `GetTemplate` |
| `ContactsAPI` | Manage contacts | `GetContacts`, `GetContact`, `CreateContact`, `UpdateContact`, `DeleteContact`, `BatchUpsertContacts` |
| `ListsAPI` | Manage lists and their members | `GetLists`, `CreateList`, `GetList`, `UpdateList`, `DeleteList`, `GetListContacts`, `BatchAddListContacts`, `UpsertListContact`, `DeleteListContact`, `GetContactLists` |
| `DomainsAPI` | Add and check sending domains | `CreateDomain`, `CheckDomainDNS`, `GetDomain` |
| `WebhooksAPI` | Manage webhooks | `CreateWebhook`, `UpdateWebhook`, `GetWebhooks` |
| `StatisticsAPI` | Read sending statistics | `GetDeliverabilityStatistics`, `GetBounceStatistics` |
| `SuppressionsAPI` | Manage addresses that must not be emailed | `CreateSuppression`, `DeleteSuppression`, `GetSuppressions` |
| `RoutesAPI` | Handle incoming email | `CreateRoute`, `UpdateRoute` |
| `AccountsAPI` | Manage the account and its members | `GetAccount`, `AddAccountMember` |
| `APIKeysAPI` | Manage API keys | `CreateAPIKey`, `UpdateAPIKey` |
| `SubAccountsAPI` | Manage sub accounts and their API keys | `ListSubAccounts`, `CreateSubAccount`, `CreateSubAccountAPIKey`, `GetSubAccountsUsage` |

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
        if err != nil {
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

New values can appear in `DeliveryAttempt.Classification` at any time. Handle the `webhooks.Classification*` values you care about, and keep a `default` case for the rest. Do not reject a request because of a value you do not know: a webhook that fails 100 times in a row is turned off.

Events:
- `message.reception`, `message.delivered`, `message.transient_error`, `message.failed`, `message.bounced`, `message.suppressed`, `message.opened`, `message.clicked`
- `suppression.created`
- `domain.dns_error`
- `message.routing`

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

The Go samples in the API reference come from this repository. Each operation has one program at `codesamples/<operationId>/main.go`. `go build ./...` compiles them, so a sample always matches the SDK.

- `make sync-spec` downloads `openapi.yaml` from the `master` branch of the API repository and writes the samples into it. The API repository owns everything in that file except the Go samples. While a change has not reached `master`, use `make sync-spec REF=devel`. It uses the `gh` command, which needs read access to the private `AhaSend/AhaSend` repository.
- `make code-samples` writes the samples into `openapi/openapi.yaml` after you change one.
- `make check-code-samples` fails if `openapi/openapi.yaml` is out of date. CI runs it.

The API repository's `scripts/sync-code-samples` copies the Go samples from `openapi/openapi.yaml` on this repository's `main` branch into its own copy.

## Related

- [AhaSend CLI](https://github.com/AhaSend/ahasend-cli): a command-line tool built on this SDK

## Help

- [API documentation](https://ahasend.com/docs)
- [Go package documentation](https://pkg.go.dev/github.com/AhaSend/ahasend-go)
- [Email support](mailto:support@ahasend.com)
- [Report an issue](https://github.com/AhaSend/ahasend-go/issues)

## License

MIT. See [LICENSE](LICENSE).
