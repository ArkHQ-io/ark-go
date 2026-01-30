// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/ArkHQ-io/ark-go/internal/apijson"
	"github.com/ArkHQ-io/ark-go/internal/apiquery"
	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/option"
	"github.com/ArkHQ-io/ark-go/packages/pagination"
	"github.com/ArkHQ-io/ark-go/packages/param"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
	"github.com/ArkHQ-io/ark-go/shared"
)

// EmailService contains methods and other services that help with interacting with
// the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailService] method instead.
type EmailService struct {
	Options []option.RequestOption
}

// NewEmailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmailService(opts ...option.RequestOption) (r EmailService) {
	r = EmailService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of sent emails. Results are ordered by send time,
// newest first.
//
// Use filters to narrow down results by status, recipient, sender, or tag.
//
// **Related endpoints:**
//
// - `GET /emails/{id}` - Get full details of a specific email
// - `POST /emails` - Send a new email
func (r *EmailService) List(ctx context.Context, query EmailListParams, opts ...option.RequestOption) (res *pagination.PageNumberPagination[EmailListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "emails"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a paginated list of sent emails. Results are ordered by send time,
// newest first.
//
// Use filters to narrow down results by status, recipient, sender, or tag.
//
// **Related endpoints:**
//
// - `GET /emails/{id}` - Get full details of a specific email
// - `POST /emails` - Send a new email
func (r *EmailService) ListAutoPaging(ctx context.Context, query EmailListParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[EmailListResponse] {
	return pagination.NewPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// Send a single email message. The email is accepted for immediate delivery and
// typically delivered within seconds.
//
// **Example use case:** Send a password reset email to a user.
//
// **Required fields:** `from`, `to`, `subject`, and either `html` or `text`
//
// **Idempotency:** Supports `Idempotency-Key` header for safe retries.
//
// **Related endpoints:**
//
// - `GET /emails/{id}` - Track delivery status
// - `GET /emails/{id}/deliveries` - View delivery attempts
// - `POST /emails/{id}/retry` - Retry failed delivery
func (r *EmailService) Send(ctx context.Context, params EmailSendParams, opts ...option.RequestOption) (res *EmailSendResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%s", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "emails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Send up to 100 emails in a single request. Useful for sending personalized
// emails to multiple recipients efficiently.
//
// Each email in the batch can have different content and recipients. Failed emails
// don't affect other emails in the batch.
//
// **Idempotency:** Supports `Idempotency-Key` header for safe retries.
func (r *EmailService) SendBatch(ctx context.Context, params EmailSendBatchParams, opts ...option.RequestOption) (res *EmailSendBatchResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%s", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "emails/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Send a pre-formatted RFC 2822 MIME message. Use this for advanced use cases or
// when migrating from systems that generate raw email content.
//
// **Important:** The `rawMessage` field must be base64-encoded. Your raw MIME
// message (with headers like From, To, Subject, Content-Type, followed by a blank
// line and the body) must be encoded to base64 before sending.
func (r *EmailService) SendRaw(ctx context.Context, body EmailSendRawParams, opts ...option.RequestOption) (res *EmailSendRawResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "emails/raw"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EmailListResponse struct {
	// Internal message ID
	ID    string `json:"id,required"`
	Token string `json:"token,required"`
	From  string `json:"from,required"`
	// Current delivery status:
	//
	// - `pending` - Email accepted, waiting to be processed
	// - `sent` - Email transmitted to recipient's mail server
	// - `softfail` - Temporary delivery failure, will retry
	// - `hardfail` - Permanent delivery failure
	// - `bounced` - Email bounced back
	// - `held` - Held for manual review
	//
	// Any of "pending", "sent", "softfail", "hardfail", "bounced", "held".
	Status       EmailListResponseStatus `json:"status,required"`
	Subject      string                  `json:"subject,required"`
	Timestamp    float64                 `json:"timestamp,required"`
	TimestampISO time.Time               `json:"timestampIso,required" format:"date-time"`
	To           string                  `json:"to,required" format:"email"`
	Tag          string                  `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Token        respjson.Field
		From         respjson.Field
		Status       respjson.Field
		Subject      respjson.Field
		Timestamp    respjson.Field
		TimestampISO respjson.Field
		To           respjson.Field
		Tag          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current delivery status:
//
// - `pending` - Email accepted, waiting to be processed
// - `sent` - Email transmitted to recipient's mail server
// - `softfail` - Temporary delivery failure, will retry
// - `hardfail` - Permanent delivery failure
// - `bounced` - Email bounced back
// - `held` - Held for manual review
type EmailListResponseStatus string

const (
	EmailListResponseStatusPending  EmailListResponseStatus = "pending"
	EmailListResponseStatusSent     EmailListResponseStatus = "sent"
	EmailListResponseStatusSoftfail EmailListResponseStatus = "softfail"
	EmailListResponseStatusHardfail EmailListResponseStatus = "hardfail"
	EmailListResponseStatusBounced  EmailListResponseStatus = "bounced"
	EmailListResponseStatusHeld     EmailListResponseStatus = "held"
)

type EmailSendResponse struct {
	Data    EmailSendResponseData `json:"data,required"`
	Meta    shared.APIMeta        `json:"meta,required"`
	Success bool                  `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendResponseData struct {
	// Unique message ID (format: msg*{id}*{token})
	ID string `json:"id,required"`
	// Current delivery status
	//
	// Any of "pending", "sent".
	Status string `json:"status,required"`
	// List of recipient addresses
	To []string `json:"to,required" format:"email"`
	// SMTP Message-ID header value
	MessageID string `json:"messageId"`
	// Whether this email was sent in sandbox mode. Only present (and true) for sandbox
	// emails sent from @arkhq.io addresses.
	Sandbox bool `json:"sandbox"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		To          respjson.Field
		MessageID   respjson.Field
		Sandbox     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailSendResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendBatchResponse struct {
	Data    EmailSendBatchResponseData `json:"data,required"`
	Meta    shared.APIMeta             `json:"meta,required"`
	Success bool                       `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailSendBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendBatchResponseData struct {
	// Successfully accepted emails
	Accepted int64 `json:"accepted,required"`
	// Failed emails
	Failed int64 `json:"failed,required"`
	// Map of recipient email to message info
	Messages map[string]EmailSendBatchResponseDataMessage `json:"messages,required"`
	// Total emails in the batch
	Total int64 `json:"total,required"`
	// Whether this batch was sent in sandbox mode. Only present (and true) for sandbox
	// emails sent from @arkhq.io addresses.
	Sandbox bool `json:"sandbox"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Failed      respjson.Field
		Messages    respjson.Field
		Total       respjson.Field
		Sandbox     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendBatchResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailSendBatchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendBatchResponseDataMessage struct {
	// Message ID
	ID    string `json:"id,required"`
	Token string `json:"token,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Token       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendBatchResponseDataMessage) RawJSON() string { return r.JSON.raw }
func (r *EmailSendBatchResponseDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendRawResponse struct {
	Data    EmailSendRawResponseData `json:"data,required"`
	Meta    shared.APIMeta           `json:"meta,required"`
	Success bool                     `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendRawResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailSendRawResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendRawResponseData struct {
	// Unique message ID (format: msg*{id}*{token})
	ID string `json:"id,required"`
	// Current delivery status
	//
	// Any of "pending", "sent".
	Status string `json:"status,required"`
	// List of recipient addresses
	To []string `json:"to,required" format:"email"`
	// SMTP Message-ID header value
	MessageID string `json:"messageId"`
	// Whether this email was sent in sandbox mode. Only present (and true) for sandbox
	// emails sent from @arkhq.io addresses.
	Sandbox bool `json:"sandbox"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		To          respjson.Field
		MessageID   respjson.Field
		Sandbox     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSendRawResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailSendRawResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailListParams struct {
	// Return emails sent after this timestamp (Unix seconds or ISO 8601)
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Return emails sent before this timestamp
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter by sender email address
	From param.Opt[string] `query:"from,omitzero" format:"email" json:"-"`
	// Page number (starts at 1)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Results per page (max 100)
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	// Filter by tag
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Filter by recipient email address
	To param.Opt[string] `query:"to,omitzero" format:"email" json:"-"`
	// Filter by delivery status:
	//
	// - `pending` - Email accepted, waiting to be processed
	// - `sent` - Email transmitted to recipient's mail server
	// - `softfail` - Temporary delivery failure, will retry
	// - `hardfail` - Permanent delivery failure
	// - `bounced` - Email bounced back
	// - `held` - Held for manual review
	//
	// Any of "pending", "sent", "softfail", "hardfail", "bounced", "held".
	Status EmailListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailListParams]'s query parameters as `url.Values`.
func (r EmailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by delivery status:
//
// - `pending` - Email accepted, waiting to be processed
// - `sent` - Email transmitted to recipient's mail server
// - `softfail` - Temporary delivery failure, will retry
// - `hardfail` - Permanent delivery failure
// - `bounced` - Email bounced back
// - `held` - Held for manual review
type EmailListParamsStatus string

const (
	EmailListParamsStatusPending  EmailListParamsStatus = "pending"
	EmailListParamsStatusSent     EmailListParamsStatus = "sent"
	EmailListParamsStatusSoftfail EmailListParamsStatus = "softfail"
	EmailListParamsStatusHardfail EmailListParamsStatus = "hardfail"
	EmailListParamsStatusBounced  EmailListParamsStatus = "bounced"
	EmailListParamsStatusHeld     EmailListParamsStatus = "held"
)

type EmailSendParams struct {
	// Sender email address. Must be from a verified domain OR use sandbox mode.
	//
	// **Supported formats:**
	//
	// - Email only: `hello@yourdomain.com`
	// - With display name: `Acme <hello@yourdomain.com>`
	// - With quoted name: `"Acme Support" <support@yourdomain.com>`
	//
	// The domain portion must match a verified sending domain in your account.
	//
	// **Sandbox mode:** Use `sandbox@arkhq.io` to send test emails without domain
	// verification. Sandbox emails can only be sent to organization members and are
	// limited to 10 per day.
	From string `json:"from,required"`
	// Email subject line
	Subject string `json:"subject,required"`
	// Recipient email addresses (max 50)
	To []string `json:"to,omitzero,required" format:"email"`
	// HTML body content (accepts null). Maximum 5MB (5,242,880 characters). Combined
	// with attachments, the total message must not exceed 14MB.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Reply-to address (accepts null)
	ReplyTo param.Opt[string] `json:"replyTo,omitzero" format:"email"`
	// Tag for categorization and filtering (accepts null)
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Plain text body (accepts null, auto-generated from HTML if not provided).
	// Maximum 5MB (5,242,880 characters).
	Text           param.Opt[string] `json:"text,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// File attachments (accepts null)
	Attachments []EmailSendParamsAttachment `json:"attachments,omitzero"`
	// BCC recipients (accepts null)
	Bcc []string `json:"bcc,omitzero" format:"email"`
	// CC recipients (accepts null)
	Cc []string `json:"cc,omitzero" format:"email"`
	// Custom email headers (accepts null)
	Headers map[string]string `json:"headers,omitzero"`
	// Custom key-value pairs attached to an email for webhook correlation.
	//
	// When you send an email with metadata, these key-value pairs are:
	//
	// - **Stored** with the message
	// - **Returned** in all webhook event payloads (MessageSent, MessageBounced, etc.)
	// - **Never visible** to email recipients
	//
	// This is useful for correlating webhook events with your internal systems (e.g.,
	// user IDs, order IDs, campaign identifiers).
	//
	// **Validation Rules:**
	//
	//   - Maximum 10 keys per email
	//   - Keys: 1-40 characters, must start with a letter, only alphanumeric and
	//     underscores (`^[a-zA-Z][a-zA-Z0-9_]*$`)
	//   - Values: 1-500 characters, no control characters (newlines, tabs, etc.)
	//   - Total size: 4KB maximum (JSON-encoded)
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r EmailSendParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, ContentType, Filename are required.
type EmailSendParamsAttachment struct {
	// Base64-encoded file content
	Content string `json:"content,required"`
	// MIME type
	ContentType string `json:"contentType,required"`
	// Attachment filename
	Filename string `json:"filename,required"`
	paramObj
}

func (r EmailSendParamsAttachment) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendParamsAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendParamsAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendBatchParams struct {
	Emails []EmailSendBatchParamsEmail `json:"emails,omitzero,required"`
	// Sender email for all messages
	From           string            `json:"from,required"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r EmailSendBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Subject, To are required.
type EmailSendBatchParamsEmail struct {
	Subject string            `json:"subject,required"`
	To      []string          `json:"to,omitzero,required" format:"email"`
	HTML    param.Opt[string] `json:"html,omitzero"`
	// Tag for categorization and filtering
	Tag  param.Opt[string] `json:"tag,omitzero"`
	Text param.Opt[string] `json:"text,omitzero"`
	// Custom key-value pairs attached to an email for webhook correlation.
	//
	// When you send an email with metadata, these key-value pairs are:
	//
	// - **Stored** with the message
	// - **Returned** in all webhook event payloads (MessageSent, MessageBounced, etc.)
	// - **Never visible** to email recipients
	//
	// This is useful for correlating webhook events with your internal systems (e.g.,
	// user IDs, order IDs, campaign identifiers).
	//
	// **Validation Rules:**
	//
	//   - Maximum 10 keys per email
	//   - Keys: 1-40 characters, must start with a letter, only alphanumeric and
	//     underscores (`^[a-zA-Z][a-zA-Z0-9_]*$`)
	//   - Values: 1-500 characters, no control characters (newlines, tabs, etc.)
	//   - Total size: 4KB maximum (JSON-encoded)
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r EmailSendBatchParamsEmail) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendBatchParamsEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendBatchParamsEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendRawParams struct {
	// Sender email address. Must be from a verified domain.
	//
	// **Supported formats:**
	//
	// - Email only: `hello@yourdomain.com`
	// - With display name: `Acme <hello@yourdomain.com>`
	// - With quoted name: `"Acme Support" <support@yourdomain.com>`
	//
	// The domain portion must match a verified sending domain in your account.
	From string `json:"from,required"`
	// Base64-encoded RFC 2822 MIME message.
	//
	// **You must base64-encode your raw email before sending.** The raw email should
	// include headers (From, To, Subject, Content-Type, etc.) followed by a blank line
	// and the message body.
	RawMessage string `json:"rawMessage,required"`
	// Recipient email addresses
	To []string `json:"to,omitzero,required" format:"email"`
	// Whether this is a bounce message (accepts null)
	Bounce param.Opt[bool] `json:"bounce,omitzero"`
	paramObj
}

func (r EmailSendRawParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendRawParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendRawParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
