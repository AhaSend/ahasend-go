# Changelog

All notable changes to this package are documented here. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and
this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

### Fixed

- The doc comments of `CreateCampaign`, `UpdateCampaign` and `DuplicateCampaign`, and the bundled
  `openapi/openapi.yaml`, say that the 409 for a segment of the audience that is being rebuilt is
  not stored for the idempotency key, so a retry with the same key runs again; retry it after a
  minute. They said the refusal was stored and a retry needed a new key.

## [0.6.0] — 2026-11-09

### BREAKING

- The `ID` of a send result (`responses.CreateSingleMessageResponse.ID`, from `CreateMessage`,
  `CreateTemplateMessage` and `CreateConversationMessage`) is the message's ID as a bare UUID: the
  ID that `GetMessageByAPIID` and `CancelMessage` take, and that `GetMessage` takes parsed with
  `uuid.Parse`. It was the Message-ID header, in the form `<uuid@domain>`, which is now in
  `MessageIDHeader`. `ID` stays a `*string`. The API makes this change, so an older SDK gets the
  bare UUID in `ID` too. Code that passed a send result's `ID` to the `MessageIDHeader` filter of
  `GetMessages` (`requests.GetMessagesParams`) must pass its `MessageIDHeader` instead.

### Added

- `responses.CreateSingleMessageResponse` gains `MessageIDHeader`: the Message-ID header of the
  recipient's email, with its angle brackets, or nil when the message was not sent.

### Changed

- The `Data.ID` of a message webhook event (`webhooks.MessageEventData`,
  `webhooks.MessageClickedEventData`) and of a `message.routing` event (`webhooks.RouteEventData`)
  is the message's ID, a hyphenated UUID that `GetMessageByAPIID` takes. For a message sent through
  the API, it is the `ID` that the send returned. For a campaign message, it only identifies that
  message's events, because the API does not return campaign messages. `ID` stays a `string`.
  Events sent before the server change carry the older ID, so a message whose events span the
  change has events with both IDs.
- The bundled `openapi/openapi.yaml` and `openapi/webhooks.yaml` follow the current API. Beside
  the ID changes above, only their descriptions and Node.js samples change.

## [0.5.0] — 2026-10-26

### BREAKING

- `TemplatesAPI.GetTemplateDraft` and `responses.TemplateDraft` are removed, because the API no
  longer has the draft read. `GetTemplate` and the template writes return the template's working
  copy: the draft while one exists (`HasDraft`), and the published copy otherwise. Code that
  builds a send from `GetTemplate` should call `GetPublishedTemplate` instead.
- `responses.CampaignRecipient.ListName` is a `*string`: it is nil for a recipient sent to from a
  segment. Code that reads it as a `string` needs a nil check.
- `responses.CampaignTemplate` gains a slice field, `MarkupWarnings`, so it can no longer be
  compared with `==`.

### Added

- `SegmentsAPI` manages segments, saved rules over the account's contacts that AhaSend keeps up to
  date: `GetSegments`, `CreateSegment`, `PreviewSegment`, `GetSegment`, `UpdateSegment` and
  `DeleteSegment`. `GetSegmentContacts` lists a segment's members, and `GetContactSegments` lists
  the segments one contact is a member of. A definition is a `common.SegmentDefinition` of
  `common.SegmentRule` entries, each a `SegmentCondition` or a `SegmentGroup`, and a condition's
  value is a `common.SegmentValue`: a string, number, boolean or `SegmentCampaignCount`. A rule or
  value with none or more than one of its fields set fails to marshal. The constants
  `SegmentMatch*`, `SegmentState*`, `SegmentField*` and `SegmentOp*` name the match modes, states,
  fields and operators. `GetSegment` returns a `responses.SegmentDetail`, which adds
  `MarketableCount`, the members a campaign to the segment alone would reach. The segment
  operations need `segments:read`, `segments:write` or `segments:delete`, and `PreviewSegment`,
  `GetSegmentContacts` and `GetContactSegments` also need `contacts:read`. See "Segments" in the
  README.
- `Campaign`, `CreateCampaignRequest` and `UpdateCampaignRequest` gain `SegmentIDs`,
  `ExcludeListIDs` and `ExcludeSegmentIDs`. A campaign uses at most 50 lists and 10 segments,
  exclusions included, and `ListIDs` can be empty when `SegmentIDs` names a segment. On
  `UpdateCampaign`, each of `ListIDs`, `SegmentIDs`, `ExcludeListIDs` and `ExcludeSegmentIDs` is
  left as it is when nil and replaced when not nil; a non-nil empty slice empties it. The result
  needs at least one list or segment to send to.
- `responses.CampaignRecipient` gains `SegmentID` and `SegmentName`, which name the segment a
  recipient was sent to from, and are nil for a recipient sent to from a list.
- `responses.DeliveryEvent` gains `ReportMessageID`: on the delivery event of an out-of-band
  bounce, the ID of the bounce report message, which `GetMessage` reads.
- `responses.Webhook`, `requests.CreateWebhookRequest`, `requests.UpdateWebhookRequest` and
  `api.GetWebhooksParams` gain the settings of the new webhook events: `OnUnsubscribed`,
  `OnComplained`, `OnSuppressionDeleted`, `OnContactCreated`, `OnContactUpdated`,
  `OnContactDeleted`, `OnContactAddedToList`, `OnContactRemovedFromList`, `OnContactSubscribed`,
  `OnContactUnsubscribed` and `OnContactComplained`. Each is off unless set, so an existing webhook
  gets nothing new. The contact events need the contact read scopes on the API key. Code that
  builds `CreateWebhookRequest` with an unkeyed struct literal must name the fields.
- `responses.Webhook`, `requests.CreateWebhookRequest` and `requests.UpdateWebhookRequest` gain
  `Origins`: whose changes the contact and suppression events report, as
  `common.WebhookOriginRecipient`, `WebhookOriginAPI`, `WebhookOriginDashboard` or
  `WebhookOriginSystem`. On `Webhook`, nil means every origin. In a request, an empty `Origins` is
  left out: on create it means every origin, and on update it keeps the stored value.
- `WebhookVerifier.Parse` and `ParseRequest` return typed events for 11 more webhook events, which
  they answered with `ErrUnknownEventType` before: `MessageUnsubscribedEvent`,
  `MessageComplainedEvent`, `SuppressionDeletedEvent`, `ContactCreatedEvent`,
  `ContactUpdatedEvent`, `ContactDeletedEvent`, `ContactAddedToListEvent`,
  `ContactRemovedFromListEvent`, `ContactSubscribedEvent`, `ContactUnsubscribedEvent` and
  `ContactComplainedEvent`. The contact event data is `ContactEventData`,
  `ContactUpdatedEventData` (with `ChangedFields`) and `ContactListEventData` (with `List` and
  `ListContact`), built from `Contact`, `WebhookList` and `WebhookListContact`. `List` and
  `ListContact` are nil on an event about the contact's unsubscribe from all lists, and on
  `contact.complained` for a campaign sent to a segment. `IsContactEvent` matches the contact
  events, and `IsMessageEvent`, `IsSuppressionEvent` and `GetMessageEventData` cover the new
  message and suppression events.
- `webhooks.MessageEventData` and `MessageClickedEventData` gain `CampaignID`, `ListID` and
  `SegmentID` for a campaign message; each is nil when absent or null. `MessageEventData` also
  gains `Unsubscribe` (`WebhookUnsubscribe`, with an `UnsubscribeMethod`) on `message.unsubscribed`
  and `Complaint` (`WebhookComplaint`, with a `ComplaintReporter`) on `message.complained`.
- `webhooks.SuppressionEventData` and the contact event data carry `Origin` (`WebhookOrigin`): who
  made the change.
- `webhooks.WebhookOrigin`, `ComplaintReporter` and `ContactChangedField` are string types with
  constants for the values known today (`Origin*`, `ComplaintReporter*`, `ChangedField*`), so a
  value this release predates decodes without error. Keep a default branch for it.
- Simple designs can be read and written as AhaSend Markup (AhaML). `responses.TemplateContent`, the
  content of a template, campaign template or campaign, gains `Markup`, the design of a `simple`
  design as AhaML, and `MarkupError` (`responses.MarkupError`, codes `MarkupError*`), which says why
  `Markup` is nil for a simple design. `requests.TemplateContentInput` gains `Markup`, which writes
  a simple design; on a create, and on a campaign with no design yet, it means
  `common.TemplateEditorSimple`. Writes that carry `Markup` count against a per-account limit,
  beside the API's own rate limit.
- `responses.Template`, `CampaignTemplate` and `Campaign` gain `MarkupWarnings`
  (`[]common.MarkupProblem`), set only in the answer to a write that carried `Content.Markup`.
  `common.MarkupWarningJinjaInText` is the only warning code.
- `api.APIError` gains `Errors` (`[]common.MarkupProblem`): the problems of the AhaML a write
  refused, up to the first 100. It is nil for every other error.
- `TemplatesAPI.GetPublishedTemplate` returns a template's published copy, the one a send uses,
  with the variables a send must supply. Read it to build a send.

### Changed

- The campaign, list and template operations document the API's new refusals: `CreateCampaign`,
  `UpdateCampaign` and `DuplicateCampaign` answer 409 while a segment of the audience is being
  rebuilt, and `DeleteCampaign` and `DeleteList` answer 409 while a segment refers to the campaign
  or list. A write that carries `Content.Markup` answers 503 when the AhaML cannot be read now;
  retry after `Retry-After`.
- The bundled `openapi/openapi.yaml` and `openapi/webhooks.yaml` follow the current API.

### Fixed

- The webhook receiver in `examples/webhook_management.go` verifies every delivery with
  `WebhookVerifier`. It checked an `X-AhaSend-Signature` header, which AhaSend does not send, so it
  verified nothing.

## [0.4.0] — 2026-10-19

### Added

- `responses.CampaignIssueTextUnchecked` (`text_unchecked`): the API could not check the
  campaign's subject, preview text and design right now. It clears by itself, so a `SendCampaign`
  or `TestCampaign` refused with it can be called again in a moment; under load it can come and go
  between reads. Code that switches over every campaign issue code needs a case for it.
- `responses.TemplateVariable` gains `Kind` (`TemplateVariableKindValue`,
  `TemplateVariableKindCondition` or `TemplateVariableKindList`), which says what shape of value a
  send gives for the variable, and `ItemFields`, set only on a list variable, which names the
  fields each element of its slice needs. `Kind` is a plain string, so a kind this release does
  not know still decodes. A version published before kinds were recorded shows every variable as
  `value`, with no `ItemFields`. The README and the `getTemplate` sample show how to fill a list
  variable.

### Changed

- `CreateMessage`, `CreateTemplateMessage`, `CreateConversationMessage` and `TestCampaign`
  document the API's check of every message before it is sent. A message the check refuses, such
  as one whose content cannot be rendered with a recipient's values, is answered 400, and nothing
  of the request is sent; for a send, the message reads `<email>: the message could not be
  rendered: <reason>`. A check that cannot run answers 503 with `Retry-After`, which the SDK
  retries with the same idempotency key.
- The bundled `openapi/openapi.yaml` follows the current API, which also changes descriptions:
  the rate-limit `Retry-After`, sender and recipient names, `bcc`, the `to` header, and
  substitutions for list variables.

## [0.3.2] — 2026-10-14

### Fixed

- The rate limiter keeps the configured rate when callers wait at the same time. Callers that
  found the token bucket empty all woke at the same moment and went ahead together, so concurrent
  requests ran far above the limit (about 1,000 requests a second from 10 workers against a 100/s
  limit), which the API refuses with 429. Callers that find no token now wait in a first-come,
  first-served queue and leave one after another at the configured rate. A caller whose context
  ends leaves the queue. A configuration change applies to the callers already waiting, and
  disabling a bucket, or a rate of zero, releases them.
- The client-wide rate limit switch works: `Configuration.EnableRateLimit = false`,
  `WithRateLimit(false)` and `SetGlobalRateLimit(false)` had no effect on requests. A
  `Configuration` built by hand, not through `NewConfiguration` or `NewAPIClient`, that leaves
  `EnableRateLimit` false now has no client-side rate limiting, as the field says. Clients created
  the usual way start with it enabled.

## [0.3.1] — 2026-10-14

### Changed

- On a 429 or 503, a retry waits the longer of the backoff delay and the response's `Retry-After`,
  given as seconds or as an HTTP date. A response whose `Retry-After` is longer than
  `RetryConfig.MaxDelay` is not retried: the error is returned with `APIError.RetryAfter` set.
  `APIError.RetryAfter` is also set for a 503.
- A TLS certificate verification failure, and a request whose context was cancelled or passed its
  deadline, are not retried.

### Fixed

- Every retried response is closed, including the one held when the context ends during the wait
  between tries.

## [0.3.0] — 2026-10-12

### BREAKING

- `webhooks.RouteEventData.SpamScore` is a `*float64`, not a `*float32`. Code that stores it in a
  `float32` needs a conversion.

### Added

- `ContactsAPI` manages contacts: `GetContacts`, `CreateContact`, `GetContact`, `UpdateContact`,
  `DeleteContact`, and `BatchUpsertContacts`, which creates or updates up to 1,000 contacts. A
  contact is found by ID or by email. `UpdateContact` changes only the fields you set: a pointer to
  `""` clears a text field, and `Unsubscribed` set to `false` subscribes the contact again.
  `BatchUpsertContacts` answers 200 even when some items fail, so check `Failed` and each
  `Data[i].Outcome`. Scopes: `contacts:read`, `contacts:write` and `contacts:delete`.
- `ListsAPI` manages contact lists: `GetLists`, `CreateList`, `GetList`, `UpdateList` and
  `DeleteList`; a list's members, with `GetListContacts`, `UpsertListContact`, `DeleteListContact`
  and `BatchAddListContacts`; and `GetContactLists`, the lists one contact belongs to.
  `ContactCount` is the number of members a campaign to the list would reach. Scopes:
  `lists:read`, `lists:write` and `lists:delete`, and `IncludeContacts` on `GetListContacts` also
  needs `contacts:read`.
- `TemplatesAPI` manages transactional templates. `GetTemplates` and `GetTemplate` read them, each
  with the `Variables` a send must supply, its default sender and reply-to (`From` and `ReplyTo`,
  nil when it has none), its `Editor` (`common.TemplateEditor*`), `HasDraft` and, outside the
  list, its published `Content`. `CreateTemplate`, `UpdateTemplate` and `DeleteTemplate` write
  templates: a write changes the template's draft, and `Publish` publishes it in the same request.
  `GetTemplateDraft`, `DiscardTemplateDraft` and `PublishTemplate` act on the draft, and
  `GetTemplateVersions`, `GetTemplateVersion` and `RestoreTemplateVersion` read and restore
  published versions. `UpdateTemplate` sends an automatic `Idempotency-Key`, as the POST
  operations do. Scopes: `templates:read`, `templates:write` and `templates:delete`. See "Manage
  templates" in the README.
- `MessagesAPI.CreateTemplateMessage` sends a transactional template
  (`requests.CreateTemplateMessageRequest`): `TemplateID` and `Recipients` are required, and each
  recipient's `Substitutions` hold its values for the template's variables. A nil `From` sends from
  the template's default sender, and an empty `Subject` uses the template's subject.
  `responses.Message` gains `TemplateID`, the template a message was sent from, or nil. See "Send
  from a template" in the README.
- `CampaignTemplatesAPI` manages campaign templates, the designs a campaign copies:
  `ListCampaignTemplates`, `CreateCampaignTemplate`, `GetCampaignTemplate`,
  `UpdateCampaignTemplate` and `DeleteCampaignTemplate`. A write applies at once. They use the
  template scopes.
- `CampaignsAPI` manages marketing campaigns: `ListCampaigns`, `CreateCampaign`, `GetCampaign`,
  `UpdateCampaign`, `DeleteCampaign`, `DuplicateCampaign`, `SendCampaign`, `UnscheduleCampaign`,
  `TestCampaign`, `GetCampaignStats` and `ListCampaignRecipients`. A draft carries `Issues` and
  `Ready`; `api.AsCampaignNotReady(err)` returns the issues (`responses.CampaignIssue*`) of a
  `SendCampaign` or `TestCampaign` refused with 400. Statuses are `common.CampaignStatus*`, and only
  a draft can be changed or deleted. The SDK sends an `Idempotency-Key` with `CreateCampaign`,
  `UpdateCampaign`, `DuplicateCampaign`, `SendCampaign`, `UnscheduleCampaign` and `TestCampaign`.
  Scopes: `campaigns:read`, `campaigns:write`, `campaigns:delete` and `campaigns:send`. See
  "Campaigns" in the README.
- Domain sending types: `responses.Domain` gains `SendingType`, `Paused`, `PausedAt` and
  `PauseReason` (`responses.DomainPauseReasonBounceRate`; other values can appear).
  `CreateDomainRequest` and `UpdateDomainRequest` gain `SendingType`
  (`common.DomainSendingTypeTransactional` or `common.DomainSendingTypeMarketing`); a nil
  `SendingType` on create makes a transactional domain. `DomainsAPI.GetDomainsWithParams`
  (`requests.GetDomainsParams`) filters by sending type.
- `responses.Domain`, `CreateDomainRequest` and `UpdateDomainRequest` gain `DKIMSelector`, the
  domain's DKIM selector override, supported on Platform Partner accounts.
- `SubAccountsAPI.UnpauseSubAccountDomain` lifts the pause on a sub account's domain and returns
  the domain. It is idempotent and needs `sub-accounts:suspend`.
- `responses.Suppression` gains `Protected`, set when the recipient unsubscribed or reported spam.
  Deleting all suppressions keeps protected ones.
- `webhooks.ErrInvalidTimestamp`, returned when the `webhook-timestamp` header is not a Unix
  timestamp in seconds.
- `api.NetworkError` gains `Unwrap`, so `errors.Is` and `errors.As` reach the error it wraps.

### Changed

- `common.Recipient.Substitutions` and `requests.CreateMessageRequest.Substitutions` document that
  the API HTML-escapes substitution values in the HTML and AMP bodies; write `{{ name|safe }}` in
  the design to insert a value as HTML. The subject, the headers and the text body get values as
  they are.
- The README's webhook handler and `examples/webhook_processing.go` compare errors with
  `errors.Is`, and answer 2xx to `ErrUnknownEventType`, a signed event this SDK version does not
  know.
- The bundled `openapi/openapi.yaml` and `openapi/webhooks.yaml` follow the current API, and the
  doc comments with them: a campaign message event's `From` can carry the sender's name
  (`Name <address>`); a route event's `ReplyTo` is the Reply-To header as received, and its `To` is
  the address the route matched or, on a route that groups by Message-ID, the To header as
  received; `SpamScore` can be below 0 or above 10. `SuspendSubAccount` and `UnsuspendSubAccount`
  are described as pausing and resuming sending. `webhooks.yaml` no longer lists `route.message`,
  and the SDK still parses it as `message.routing`.

### Fixed

- `api.WithTimeout` is the deadline of the whole request, rate-limit wait and retries included,
  and replaces the HTTP client's own timeout for that request without changing the shared client.
  A rate-limit wait that ends with the context returns a `NetworkError` that wraps the context's
  error.
- `NewAPIClientFromEnv` honours the variables `ENV.md` documents. It loads `AHASEND_API_KEY` or
  `AHASEND_TOKEN`, so its requests no longer fail with "No API key provided"; an unset variable
  keeps a key set in code. `AHASEND_BASE_URL`, `AHASEND_HOST` and `AHASEND_SCHEME` reach the
  client. `AHASEND_TIMEOUT` and `AHASEND_CONNECT_TIMEOUT` set the HTTP client's request timeout and
  the transport's dial timeout; a caller-supplied client is copied, not modified.
- A path parameter, such as a contact's email, containing `//` or `..` is no longer refused. Each
  value is escaped into one path segment, and only an empty value, `.`, `..` or a value with a
  control character is refused.
- `GetSuppressions` sends the deprecated `FromDate` and `ToDate` when `FromTime` and `ToTime` are
  nil. It ignored them.

## [0.2.0] — 2026-08-20

### Added

- `webhooks.MessageEventData` gains `DeliveryAttempt`, the diagnostics of the delivery attempt an
  event reports on: `SMTPCode`, and `Classification`, `EnhancedStatusCode`, `Response`,
  `Description` and `Command` when they were recorded. It is optional and absent more often than
  present; nil covers both a missing field and an explicit null, so check it before reading.
  `MessageClickedEventData` has no `DeliveryAttempt`, and on `MessageFailedEvent`, which reports
  retry exhaustion, it is always nil.
- `webhooks.DeliveryAttemptClassification`, a string type with a constant for each of the 15
  buckets AhaSend emits today (`ClassificationInvalidRecipient` through
  `ClassificationUncategorized`). The set is open, so a value this release predates decodes
  without error: keep a default branch for it, and never reject a delivery because of it.
- `Response` and `Description` are free text that often holds the recipient address and parts of
  the message, so do not log them. `SMTPCode`, `EnhancedStatusCode`, `Classification` and
  `Command` are safe to log.

### Fixed

- The README's webhook handler compiles: it logged a `Reason` field that `MessageEventData` does
  not have. The statistics example uses the real bounce classifications, not `hard`, `soft` and
  `admin`.

## [0.1.0] — 2026-08-06

Initial release.

### Added

- Clients for the AhaSend API v2 on `api.APIClient`: `MessagesAPI` (`CreateMessage`,
  `CreateConversationMessage`, `GetMessages`, `GetMessage`, `GetMessageByAPIID`, `CancelMessage`),
  `DomainsAPI` (with `CheckDomainDNS`), `WebhooksAPI`, `RoutesAPI`, `SuppressionsAPI`,
  `StatisticsAPI` (deliverability, bounce and delivery-time statistics), `AccountsAPI` (with
  members), `APIKeysAPI`, `SMTPCredentialsAPI`, `SubAccountsAPI` (sub account lifecycle, suspend
  and unsuspend, usage, and the sub accounts' API keys) and `UtilityAPI.Ping`. An API key carries
  `IPAllowList`, the source IPs allowed to use it.
- Client setup with `NewAPIClient` and options such as `WithAPIKey`, `WithRetryConfig`,
  `WithRateLimit` and `WithHTTPClient`, or with `NewAPIClientWithConfig` and
  `NewAPIClientFromEnv`. Per-request options: `WithTimeout`, `WithRetry`, `WithHeaders`,
  `WithIdempotencyKey`, `WithRequestAPIKey` and `WithoutRateLimit`.
- Client-side rate limiting per endpoint category (`GeneralAPI`, `StatisticsAPI`,
  `SendMessageAPI`), adjustable with `SetGeneralRateLimit`, `SetStatisticsRateLimit`,
  `SetSendMessageRateLimit` and `SetGlobalRateLimit`.
- Retries with exponential, linear or constant backoff (`RetryConfig`, `BackoffExponential`,
  `BackoffLinear`, `BackoffConstant`).
- An automatic `Idempotency-Key` on every POST, reused across retries, and
  `IdempotencyKeyBuilder` for related operations. See `IDEMPOTENCY.md`.
- Typed errors: `APIError` with an `ErrorType`. A 409 for an idempotent request that is still in
  flight is `ErrorTypeIdempotencyConflict`, with `IsRetryable()` true and the wait in `RetryAfter`;
  the SDK does not retry it on its own. Every other 409 is `ErrorTypeConflict`.
- The `webhooks` package: `WebhookVerifier` checks Standard Webhooks HMAC-SHA256 signatures
  (`Verify`, `VerifyRequest`) and parses events (`Parse`, `ParseRequest`). A timestamp outside the
  tolerance (`DefaultTolerance`, 5 minutes, or `SetTolerance`) in either direction fails with
  `ErrExpiredTimestamp`. Typed events cover `message.reception`, `message.delivered`,
  `message.transient_error`, `message.failed`, `message.bounced`, `message.suppressed`,
  `message.opened`, `message.clicked`, `suppression.created`, `domain.dns_error` and
  `message.routing`, which also accepts the legacy name `route.message`. Route attachments carry
  `ContentID` and `Disposition`.
- Pointer helpers `ahasend.String`, `Bool`, `Int`, `Int32`, `Int64` and `Time`.
- Go 1.18 or later.
