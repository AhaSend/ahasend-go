# AhaSend Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/AhaSend/ahasend-go.svg)](https://pkg.go.dev/github.com/AhaSend/ahasend-go)
[![License: MIT](https://img.shields.io/github/license/ahasend/ahasend-go)](https://opensource.org/licenses/MIT)

The Go SDK for the [AhaSend](https://ahasend.com) API. Use it to send email and to manage templates, contacts, lists, segments, campaigns, domains, webhooks, routes, suppressions and sub accounts.

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

Every message is checked before it is sent, by `CreateMessage`, `CreateConversationMessage` and `CreateTemplateMessage` alike:

- A message whose content cannot be rendered with a recipient's values, calls `random`, `randrange`, `lipsum` or `now`, uses `namespace`, passes a work, output, memory or time limit, or has a recipient address the check cannot read is refused with 400, and nothing of the request is sent. The message reads `<email>: the message could not be rendered: <reason>`, or has no `<email>: ` when no recipient had started rendering.
- When the check cannot run, the API answers 503 with a `Retry-After`. Nothing was sent. The SDK retries it with the same idempotency key, as [Settings](#settings) describes for retries.

## Send from a template

A transactional template, made in the dashboard or through the API, holds the subject, the preview text and both bodies. `CreateTemplateMessage` names the template, and each recipient gives values for its variables:

```go
// Read the published copy, which sends use, to see which variables a send must give.
template, _, err := client.TemplatesAPI.GetPublishedTemplate(ctx, accountID, templateID)
if err != nil {
    log.Fatal(err)
}
for _, variable := range template.Variables {
    log.Printf("%s (required: %t, kind: %s)", variable.Name, variable.Required, variable.Kind)
}

message := requests.CreateTemplateMessageRequest{
    TemplateID: templateID,
    From:       &common.SenderAddress{Email: "sender@yourdomain.com"},
    Recipients: []common.Recipient{
        {
            Email:         "recipient@example.com",
            Substitutions: map[string]interface{}{"first_name": "Pat"},
        },
    },
}

response, _, err := client.MessagesAPI.CreateTemplateMessage(ctx, accountID, message)
```

Each variable's `Kind` says what shape of value to send:

- `responses.TemplateVariableKindValue`: text or a number that appears in the email, such as a first name.
- `responses.TemplateVariableKindCondition`: used only to decide whether a part of the email is shown. Send a value that is not empty, such as `true`, to make it set. Leave it out, or send `""`, `false`, `0`, `nil`, an empty slice or an empty map, to make it not set. The strings `"false"` and `"0"` make it set.
- `responses.TemplateVariableKindList`: a slice. A section of the email is repeated once for each element. `ItemFields` names the fields each element should have.

For example, if `items` has `ItemFields` `["name", "price"]`, send:

```go
Substitutions: map[string]interface{}{
    "items": []map[string]interface{}{
        {"name": "Linen shirt", "price": "€49"},
        {"name": "Canvas tote", "price": "€19"},
    },
},
```

A field an element does not have is taken from the recipient's substitution of the same name. If `ItemFields` is `["value"]`, the section shows each element as it is, so you can send plain values such as `[]string{"new", "sale"}`. A version published before kinds were recorded shows every variable as `value`.

A template can also hold a default sender and reply-to address: `template.From` and `template.ReplyTo`, each nil when the template has none. A send from a template with a sender can leave `From` nil:

```go
message := requests.CreateTemplateMessageRequest{
    TemplateID: templateID,
    // No From: the send uses the template's sender.
    Recipients: []common.Recipient{
        {
            Email:         "recipient@example.com",
            Substitutions: map[string]interface{}{"first_name": "Pat"},
        },
    },
}
```

- The template supplies the body. The request has no content fields and no request-level `Substitutions`.
- Leave `Subject` empty to use the template's published subject, or set it to replace it.
- Leave `From` nil to use the template's sender. A `From` replaces the template's sender and must have an `Email`.
- The template's reply-to applies to every send from it, with or without `From`, unless the request sets `ReplyTo` or a `reply-to` entry in `Headers`; either one replaces it.
- The template's sender is checked as a sender in the request is: its domain must be in your account, have valid DNS records and not be paused.
- Variable values come only from each recipient's `Substitutions`. Every recipient must give a value for each required variable.
- AhaSend fills in `email` and `unsubscribe_url` itself.
- In the HTML body, each value is HTML-escaped. Write `{{ name|safe }}` in the design to insert a value as HTML. The subject and the text body get values as they are.
- The send fails with:
  - 404 when the template does not exist.
  - 400 when neither the request nor the template has a subject, or the template has no published content.
  - 400 when neither the request nor the template has a sender.
  - 400 `template's published design holds a {{, {% or {# that cannot be rendered, fix it and publish it` when the published design cannot be rendered.
  - 400 `the template's default sender or reply-to is not valid, edit it on the template page and publish it` when a published value the send uses no longer passes the address checks.
  - 403 `this api key is not authorized to send messages` when the request has no sender and the API key cannot send from any domain.
  - 400 for the whole request when any recipient is missing a required variable.
  - 400 `<email>: the message could not be rendered: <reason>` when the check before sending refuses a message, and 503 when the check cannot run, as for any send.
- Sending from a template needs only the normal send scope.

`client.TemplatesAPI.GetTemplates(ctx, accountID, requests.GetTemplatesParams{})` lists the account's templates, newest first. For the next page, pass the response's `Pagination.NextCursor` as `After`. For the previous page, pass `Pagination.PreviousCursor` as `Before`.

## Manage templates

A template has a published copy, which sends use, and can have a draft: changes that are not published yet. `TemplatesAPI` writes go to the draft, the same draft the dashboard edits. Set `Publish` to publish the draft in the same request, or call `PublishTemplate`:

```go
template, _, err := client.TemplatesAPI.CreateTemplate(ctx, accountID, requests.CreateTemplateRequest{
    Name:    "Password reset",
    Subject: ahasend.String("Reset your password, {{ first_name }}"),
    From:    &common.SenderAddress{Email: "hello@yourdomain.com"},
    Content: &requests.TemplateContentInput{
        HTML: ahasend.String(`<p>Hi {{ first_name }}, <a href="{{ reset_url }}">reset your password</a>.</p>`),
    },
    Publish: true,
})
```

- `Content.Markup` makes a `simple` template, `Content.MJML` an `advanced` template and `Content.HTML` an `html` template. A template with none of them must set `Editor`. The editor never changes after create.
- The design of a `simple` template is read and written as AhaSend Markup (AhaML), described in the [AhaML reference](https://ahasend.com/docs/send-api/ahaml). A read of one template or version, and the answer to a write, give it in `Content.Markup`; when AhaML cannot be made, `Content.Markup` is nil and `Content.MarkupError` says why. A write sets it with `Content.Markup`, and its HTML is made from the design. The answer to such a write lists warnings in `MarkupWarnings`, and refused AhaML gives an `*api.APIError` whose `Errors` list each problem with its line, column and path.
- MJML is compiled in strict mode. Images and stylesheets must use `https://` URLs, such as `https://example.com/logo.png`: the API cannot upload files. A template variable such as `{{ logo_url }}` is fine.
- Beside a new Markup, MJML or HTML, a nil `Content.Text` keeps a custom text, which must still render, and otherwise makes the text from the new HTML. Point `Content.Text` at `""` to make the text again from the new HTML.
- `UpdateTemplate` changes only the fields you set; the name changes at once. To clear a field, point it at the empty value: `ahasend.String("")` for `Subject` and `Preheader`, `&common.SenderAddress{}` for `From` (which also clears the reply-to) and `ReplyTo`. `Content.Text` pointing at `""` makes the text again from the HTML.
- A publish publishes the whole draft, including changes made in the dashboard.
- `GetTemplate` and every write return the working copy: the draft while one exists, and the published copy otherwise. So a template read, changed and sent back with `UpdateTemplate` builds on what was read. `GetPublishedTemplate` returns the published copy, which sends use, and `GetTemplates` lists the published copies. `Template.HasDraft` tells whether a draft exists, and `DiscardTemplateDraft` drops it.
- `GetTemplateVersions` lists the published versions, newest first, and `GetTemplateVersion` returns one with its content. `RestoreTemplateVersion` copies a version into the draft; set `Publish` to publish it too.
- The SDK sends an `Idempotency-Key` with `CreateTemplate`, `UpdateTemplate`, `PublishTemplate` and `RestoreTemplateVersion`, so a retried request is not applied twice.
- A write returns 503 when the template keeps changing while the request writes it, or when its AhaML cannot be read now. The SDK retries it a limited number of times, and a request with an idempotency key keeps the same key.
- Reads need the `templates:read` scope, writes `templates:write`, and `DeleteTemplate` needs `templates:delete`.

## Campaign templates

A campaign template is a design that a campaign copies when it is made from it. It is a different resource from a transactional template: it has no draft, no versions, no subject and no sender. `CampaignTemplatesAPI` writes change it at once:

```go
template, _, err := client.CampaignTemplatesAPI.CreateCampaignTemplate(ctx, accountID, requests.CreateCampaignTemplateRequest{
    Name: "Monthly newsletter",
    Content: &requests.TemplateContentInput{
        HTML: ahasend.String(`<p>Hi {{ first_name }}</p><p><a href="{{ unsubscribe_url }}">Unsubscribe</a></p>`),
    },
})
```

- `Content` follows the rules of transactional templates: `Content.Markup` makes a `simple` campaign template, `Content.MJML` an `advanced` one and `Content.HTML` an `html` one, a campaign template with none of them must set `Editor`, and the editor never changes. Without `Content`, the campaign template gets an empty design that the dashboard's editor opens.
- `UpdateCampaignTemplate` changes only the fields you set. Campaigns already made from the template keep their own copies.
- `ListCampaignTemplates` lists campaign templates without their content; `GetCampaignTemplate` returns it. A campaign can be made from a campaign template only when it has HTML.
- The SDK sends an `Idempotency-Key` with `CreateCampaignTemplate` and `UpdateCampaignTemplate`, so a retried request is not applied twice.
- A write that carries `Content.Markup` returns 503 when its AhaML cannot be read now. The SDK retries it a limited number of times with the same key; after that, retry after `RetryAfter` seconds.
- Campaign templates use the template scopes: `templates:read`, `templates:write` and `templates:delete`.

## Campaigns

A campaign sends a design to an audience: the lists in `ListIDs` and the segments in `SegmentIDs`, less the contacts of the lists in `ExcludeListIDs` and the members of the segments in `ExcludeSegmentIDs`. `CreateCampaign` makes a draft. Every field but `Name` can be set later, and the draft's `Issues` tell what is missing before it can be sent:

```go
campaign, _, err := client.CampaignsAPI.CreateCampaign(ctx, accountID, requests.CreateCampaignRequest{
    Name:              "October newsletter",
    Subject:           ahasend.String("Our news for October"),
    Preheader:         ahasend.String("Three things we shipped this month"),
    From:              &common.SenderAddress{Email: "news@yourdomain.com"},
    ListIDs:           []uuid.UUID{listID},
    SegmentIDs:        []uuid.UUID{segmentID},
    ExcludeListIDs:    []uuid.UUID{excludedListID},
    ExcludeSegmentIDs: []uuid.UUID{excludedSegmentID},
    TemplateID:        &campaignTemplateID,
})
if err != nil {
    log.Fatal(err)
}
for _, issue := range campaign.Issues {
    log.Printf("%s: %s", issue.Code, issue.Message)
}

_, _, err = client.CampaignsAPI.SendCampaign(ctx, accountID, campaign.ID, requests.SendCampaignRequest{})
if notReady, ok := api.AsCampaignNotReady(err); ok {
    for _, issue := range notReady.Issues {
        log.Printf("%s: %s", issue.Code, issue.Message)
    }
}
```

- The design comes from `Content`, from `Editor` alone (an empty design), or from a campaign template named by `TemplateID`. The copy is the campaign's own.
- An audience that is set needs at least one list or segment to send to, so `ListIDs` can be empty when `SegmentIDs` is set. A contact reached through a list and a segment gets one message. An exclusion wins: a contact with any membership of an excluded list, whatever its status, gets no message, and neither does a member of an excluded segment. A list or segment cannot be both included and excluded. See [Segments](#segments) for how segments behave.
- Only a draft (`common.CampaignStatusDraft`) can be changed or deleted. `UpdateCampaign` changes only the fields you set. To clear a field, point it at the empty value: `ahasend.String("")` for `Subject`, `Preheader` and `ReplyTo`, and `&common.SenderAddress{}` for `From`. The SDK sends each of them as `null`.
- A field that the API sends as `null` is nil: a campaign's `Subject`, `Preheader`, `From` and `ReplyTo` until they are set, and `Ready` and `Issues` unless the campaign is a draft. A draft that is ready has an empty `Issues`.
- `SendCampaign` sends now, or at `SendAt`. The campaign goes to `validating` and then through list validation and review, as a send from the dashboard does; read its `Status` and `StatusReason` with `GetCampaign`. The response's `Warnings` tell what the send accepted but you should know. To change the send time, call `UnscheduleCampaign` and send again.
- `SendCampaign` refuses a campaign with any issue with 400, and `TestCampaign` refuses a campaign with any issue but `unsubscribe_link_missing`. `api.AsCampaignNotReady(err)` returns the issues; every other refusal has only a message. The issue `text_unchecked` means the API could not check the text right now and clears by itself; the SDK does not retry a 400, so call again in a moment.
- A 409 means the campaign's status does not allow the request, a segment of the audience is being rebuilt, a segment refers to the draft you delete, or an earlier request with the same idempotency key is still running. The last case has the type `ErrorTypeIdempotencyConflict` and a `RetryAfter`; retry with the same key after it. A new call makes a new key, so to retry by hand, pass the key yourself with `api.WithIdempotencyKey` on both calls. `CreateCampaign`, `DuplicateCampaign` and an `UpdateCampaign` that changes the audience get the audience refusal while a segment of the audience is being rebuilt. The API stores that refusal for the key and replays it for 24 hours, so a retry needs a new key: call again in a minute without `api.WithIdempotencyKey`, or with a new key.
- `TestCampaign` sends a test to up to 5 addresses, with placeholders for the contact's fields. An account can send 50 test emails per day (UTC). Past that, the API answers 429; its message currently reads `daily test send limit reached`. The limit clears only when the UTC day changes, and the SDK retries the 429 as it retries any 429. A test is checked before it is sent as a message send is: a test the check refuses gets 400 with a message and nothing is sent, and a 503 when the check cannot run is retried with the same key. When `ResultUnknown` is true, the test may have gone out: check the inboxes before you send another.
- `DuplicateCampaign` makes a new draft from a campaign in any status.
- `GetCampaignStats` returns live counters and a stored report that is counted again from time to time. `Report` is nil until the first count, and each rate is nil while it has no base. `ListCampaignRecipients` lists the recipients in one `Status` or one `BounceClassification`, each a key of the report's `ByStatus` or `ByBounceClassification`.
- The SDK sends an `Idempotency-Key` with `CreateCampaign`, `UpdateCampaign`, `DuplicateCampaign`, `SendCampaign`, `UnscheduleCampaign` and `TestCampaign`, so a retried request is not applied twice.
- `Content` follows the rules of transactional templates: a `simple` design is written as AhaML in `Content.Markup`. `CreateCampaign` and `UpdateCampaign` return 503 when that AhaML cannot be read now. The SDK retries it a limited number of times with the same key; after that, retry after `RetryAfter` seconds.
- Scopes: `campaigns:read` to read campaigns, their stats and their recipients, `campaigns:write` to create, update and duplicate them, `campaigns:delete` to delete drafts, and `campaigns:send` to send, unschedule and test them.

## Segments

A segment is a saved set of rules over the account's contacts: their fields, their list memberships, and the campaigns they received, opened or clicked. AhaSend keeps its members up to date as that data changes. A definition holds conditions and groups of conditions. A condition on `attributes.<key>` needs a contact field with that key. `PreviewSegment` tries a definition without saving it, and `CreateSegment` saves it:

```go
definition := common.SegmentDefinition{
    SchemaVersion: common.SegmentSchemaVersion,
    Match:         common.SegmentMatchAll,
    Conditions: []common.SegmentRule{
        {Condition: &common.SegmentCondition{
            Field: common.SegmentFieldList,
            Op:    common.SegmentOpIn,
            Value: &common.SegmentValue{String: ahasend.String(listID.String())},
        }},
        {Group: &common.SegmentGroup{
            Match: common.SegmentMatchAny,
            Conditions: []common.SegmentCondition{
                {
                    Field: "attributes.country",
                    Op:    common.SegmentOpEqual,
                    Value: &common.SegmentValue{String: ahasend.String("DE")},
                },
                {
                    Field: common.SegmentFieldCampaignCount,
                    Op:    common.SegmentOpOpenedAtLeast,
                    Value: &common.SegmentValue{Count: &common.SegmentCampaignCount{Count: 3, Since: "90 days ago"}},
                },
            },
        }},
    },
}

preview, _, err := client.SegmentsAPI.PreviewSegment(ctx, accountID, requests.PreviewSegmentRequest{
    Definition: definition,
    Timezone:   ahasend.String("Europe/Berlin"),
})
if err != nil {
    log.Fatal(err)
}
if preview.Count != nil {
    log.Printf("%d contacts match", *preview.Count)
}

segment, _, err := client.SegmentsAPI.CreateSegment(ctx, accountID, requests.CreateSegmentRequest{
    Name:       "Engaged customers",
    Definition: definition,
    Timezone:   ahasend.String("Europe/Berlin"),
})
if err != nil {
    log.Fatal(err)
}

// Later, when the segment is ready:
detail, _, err := client.SegmentsAPI.GetSegment(ctx, accountID, segment.ID)
if err != nil {
    log.Fatal(err)
}
if detail.MarketableCount != nil {
    log.Printf("A campaign to %s reaches %d contacts", detail.Name, *detail.MarketableCount)
}

members, _, err := client.SegmentsAPI.GetSegmentContacts(ctx, accountID, segment.ID, requests.GetSegmentContactsParams{})
if err != nil {
    log.Fatal(err)
}
for _, member := range members.Data {
    log.Println(member.Email)
}
```

- A segment's `State` is `common.SegmentStateBuilding` while its members are found: after `CreateSegment`, and after an `UpdateSegment` that changes `Definition` or `Timezone`. A large account can take some minutes. A campaign that uses a segment cannot be sent until the segment is ready. While a segment is building, `GetSegmentContacts` may still return the members of its previous rules. A segment in `common.SegmentStateFailed` has a `FailureReason`; any `UpdateSegment`, an empty one included, builds it again.
- `MemberCount` and `GetSegmentContacts` include every contact that matches the rules, whatever its consent. `MarketableCount` counts the members a campaign to the segment alone would reach, It is nil unless the segment is ready, and also when they cannot be counted within 10 seconds.
- On `UpdateCampaign`, each of `ListIDs`, `SegmentIDs`, `ExcludeListIDs` and `ExcludeSegmentIDs` is left as it is when it is nil, and replaced when it is not. A non-nil empty slice empties it: `ExcludeSegmentIDs: []uuid.UUID{}` removes every segment exclusion. The result needs at least one list or segment to send to.
- `CampaignRecipient.ListName` is a `*string`. It is nil when the contact was sent to from a segment, and then `SegmentID` and `SegmentName` name the segment. A contact reached through a list and a segment is recorded under the list.
- Scopes: `segments:read` to list and read segments, preview a definition, and list a segment's members or a contact's segments, `segments:write` to create and update segments, and `segments:delete` to delete them. `PreviewSegment`, `GetSegmentContacts` and `GetContactSegments` also need `contacts:read`.
- `CreateSegment` returns 409 when the account already has 100 segments, or when a list or campaign that the definition names is being changed. Unlike most writes, that refusal is not stored for the idempotency key, so a retry with the same key runs again, for example after you delete a segment. A 409 of type `ErrorTypeIdempotencyConflict` means an earlier request with the same key is still running; retry with the same key after its `RetryAfter`. `UpdateSegment` returns 409 while the segment is being changed, and `DeleteSegment` while the segment is being changed or a campaign still needs it. These refusals change nothing; call again later, or after the campaign no longer needs the segment.

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
| `MessagesAPI` | Send and manage email | `CreateMessage`, `CreateTemplateMessage`, `GetMessage`, `CancelMessage` |
| `TemplatesAPI` | Manage transactional templates, their drafts and versions | `GetTemplates`, `GetTemplate`, `GetPublishedTemplate`, `CreateTemplate`, `UpdateTemplate`, `PublishTemplate`, `GetTemplateVersions`, `RestoreTemplateVersion` |
| `CampaignTemplatesAPI` | Manage the designs that campaigns copy | `ListCampaignTemplates`, `GetCampaignTemplate`, `CreateCampaignTemplate`, `UpdateCampaignTemplate`, `DeleteCampaignTemplate` |
| `CampaignsAPI` | Create, test and send marketing campaigns, and read their results | `ListCampaigns`, `CreateCampaign`, `UpdateCampaign`, `TestCampaign`, `SendCampaign`, `UnscheduleCampaign`, `GetCampaignStats`, `ListCampaignRecipients` |
| `ContactsAPI` | Manage contacts | `GetContacts`, `GetContact`, `CreateContact`, `UpdateContact`, `DeleteContact`, `BatchUpsertContacts` |
| `ListsAPI` | Manage lists and their members | `GetLists`, `CreateList`, `GetList`, `UpdateList`, `DeleteList`, `GetListContacts`, `BatchAddListContacts`, `UpsertListContact`, `DeleteListContact`, `GetContactLists` |
| `SegmentsAPI` | Manage segments and read their members | `GetSegments`, `CreateSegment`, `PreviewSegment`, `GetSegment`, `UpdateSegment`, `DeleteSegment`, `GetSegmentContacts`, `GetContactSegments` |
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

The SDK retries a network error, a 429 and a 5xx, and every other 4xx when `RetryClientErrors` is set. It waits the backoff delay, or, on a 429 or 503, the response's `Retry-After` when that is longer. A response whose `Retry-After` is longer than `MaxDelay` is not retried: you get the error with its `RetryAfter`. A TLS certificate the client does not trust, and a context that is cancelled or past its deadline, end the request without a retry. Cancelling the context also ends a wait between tries.

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
