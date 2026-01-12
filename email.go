// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/ark-go/internal/apijson"
	"github.com/stainless-sdks/ark-go/internal/apiquery"
	"github.com/stainless-sdks/ark-go/internal/requestconfig"
	"github.com/stainless-sdks/ark-go/option"
	"github.com/stainless-sdks/ark-go/packages/param"
	"github.com/stainless-sdks/ark-go/packages/respjson"
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

// Retrieve detailed information about a specific email including delivery status,
// timestamps, and optionally the email content.
//
// Use the `expand` parameter to include additional data like the HTML/text body,
// headers, or delivery attempts.
func (r *EmailService) Get(ctx context.Context, emailID string, query EmailGetParams, opts ...option.RequestOption) (res *EmailGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return
	}
	path := fmt.Sprintf("emails/%s", emailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *EmailService) List(ctx context.Context, query EmailListParams, opts ...option.RequestOption) (res *EmailListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "emails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the history of delivery attempts for an email, including SMTP response codes
// and timestamps.
func (r *EmailService) GetDeliveries(ctx context.Context, emailID string, opts ...option.RequestOption) (res *EmailGetDeliveriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return
	}
	path := fmt.Sprintf("emails/%s/deliveries", emailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retry delivery of a failed or soft-bounced email. Creates a new delivery
// attempt.
//
// Only works for emails that have failed or are in a retryable state.
func (r *EmailService) Retry(ctx context.Context, emailID string, opts ...option.RequestOption) (res *EmailRetryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailID == "" {
		err = errors.New("missing required emailId parameter")
		return
	}
	path := fmt.Sprintf("emails/%s/retry", emailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Send a single email message. The email is queued for immediate delivery and
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
func (r *EmailService) Send(ctx context.Context, params EmailSendParams, opts ...option.RequestOption) (res *SendEmail, err error) {
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
// The `data` field should contain the base64-encoded raw email.
func (r *EmailService) SendRaw(ctx context.Context, body EmailSendRawParams, opts ...option.RequestOption) (res *SendEmail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "emails/raw"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Delivery struct {
	// Delivery attempt ID
	ID string `json:"id,required"`
	// Delivery status (lowercase)
	Status string `json:"status,required"`
	// Unix timestamp
	Timestamp float64 `json:"timestamp,required"`
	// ISO 8601 timestamp
	TimestampISO time.Time `json:"timestampIso,required" format:"date-time"`
	// SMTP response code
	Code int64 `json:"code"`
	// Status details
	Details string `json:"details"`
	// SMTP server response from the receiving mail server
	Output string `json:"output"`
	// Whether TLS was used
	SentWithSsl bool `json:"sentWithSsl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Status       respjson.Field
		Timestamp    respjson.Field
		TimestampISO respjson.Field
		Code         respjson.Field
		Details      respjson.Field
		Output       respjson.Field
		SentWithSsl  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Delivery) RawJSON() string { return r.JSON.raw }
func (r *Delivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Pagination struct {
	// Current page number (1-indexed)
	Page int64 `json:"page,required"`
	// Items per page
	PerPage int64 `json:"perPage,required"`
	// Total number of items
	Total int64 `json:"total,required"`
	// Total number of pages
	TotalPages int64 `json:"totalPages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		PerPage     respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pagination) RawJSON() string { return r.JSON.raw }
func (r *Pagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendEmail struct {
	Data SendEmailData `json:"data,required"`
	Meta APIMeta       `json:"meta,required"`
	// Any of true.
	Success bool `json:"success,required"`
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
func (r SendEmail) RawJSON() string { return r.JSON.raw }
func (r *SendEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendEmailData struct {
	// Unique message ID (format: msg*{id}*{token})
	ID string `json:"id,required"`
	// Current delivery status
	//
	// Any of "queued", "sent".
	Status string `json:"status,required"`
	// List of recipient addresses
	To []string `json:"to,required" format:"email"`
	// SMTP Message-ID header value
	MessageID string `json:"messageId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		To          respjson.Field
		MessageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SendEmailData) RawJSON() string { return r.JSON.raw }
func (r *SendEmailData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetResponse struct {
	Data EmailGetResponseData `json:"data,required"`
	Meta APIMeta              `json:"meta,required"`
	// Any of true.
	Success bool `json:"success,required"`
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
func (r EmailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetResponseData struct {
	// Internal message ID
	ID string `json:"id,required"`
	// Unique message token used to retrieve this email via API. Combined with id to
	// form the full message identifier: msg*{id}*{token} Use this token with GET
	// /emails/{emailId} where emailId = "msg*{id}*{token}"
	Token string `json:"token,required"`
	// Sender address
	From string `json:"from,required"`
	// Message direction
	//
	// Any of "outgoing", "incoming".
	Scope string `json:"scope,required"`
	// Current delivery status:
	//
	// - `pending` - Email accepted, waiting to be processed
	// - `sent` - Email transmitted to recipient's mail server
	// - `delivered` - Recipient's server confirmed receipt
	// - `bounced` - Permanently rejected (hard bounce)
	// - `failed` - Delivery failed after all retry attempts
	// - `delayed` - Temporary failure, will retry automatically
	// - `held` - Held for manual review
	//
	// Any of "pending", "sent", "delivered", "bounced", "failed", "delayed", "held".
	Status string `json:"status,required"`
	// Email subject line
	Subject string `json:"subject,required"`
	// Unix timestamp when the email was sent
	Timestamp float64 `json:"timestamp,required"`
	// ISO 8601 formatted timestamp
	TimestampISO time.Time `json:"timestampIso,required" format:"date-time"`
	// Recipient address
	To string `json:"to,required" format:"email"`
	// Delivery attempt history (included if expand=deliveries)
	Deliveries []Delivery `json:"deliveries"`
	// Email headers (included if expand=headers)
	Headers map[string]string `json:"headers"`
	// HTML body content (included if expand=content)
	HTMLBody string `json:"htmlBody"`
	// SMTP Message-ID header
	MessageID string `json:"messageId"`
	// Plain text body (included if expand=content)
	PlainBody string `json:"plainBody"`
	// Whether the message was flagged as spam
	Spam bool `json:"spam"`
	// Spam score (if applicable)
	SpamScore float64 `json:"spamScore"`
	// Optional categorization tag
	Tag string `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Token        respjson.Field
		From         respjson.Field
		Scope        respjson.Field
		Status       respjson.Field
		Subject      respjson.Field
		Timestamp    respjson.Field
		TimestampISO respjson.Field
		To           respjson.Field
		Deliveries   respjson.Field
		Headers      respjson.Field
		HTMLBody     respjson.Field
		MessageID    respjson.Field
		PlainBody    respjson.Field
		Spam         respjson.Field
		SpamScore    respjson.Field
		Tag          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailListResponse struct {
	Data EmailListResponseData `json:"data,required"`
	Meta APIMeta               `json:"meta,required"`
	// Any of true.
	Success bool `json:"success,required"`
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
func (r EmailListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailListResponseData struct {
	Messages   []EmailListResponseDataMessage `json:"messages,required"`
	Pagination Pagination                     `json:"pagination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailListResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailListResponseDataMessage struct {
	// Internal message ID
	ID    string `json:"id,required"`
	Token string `json:"token,required"`
	From  string `json:"from,required"`
	// Current delivery status:
	//
	// - `pending` - Email accepted, waiting to be processed
	// - `sent` - Email transmitted to recipient's mail server
	// - `delivered` - Recipient's server confirmed receipt
	// - `bounced` - Permanently rejected (hard bounce)
	// - `failed` - Delivery failed after all retry attempts
	// - `delayed` - Temporary failure, will retry automatically
	// - `held` - Held for manual review
	//
	// Any of "pending", "sent", "delivered", "bounced", "failed", "delayed", "held".
	Status       string    `json:"status,required"`
	Subject      string    `json:"subject,required"`
	Timestamp    float64   `json:"timestamp,required"`
	TimestampISO time.Time `json:"timestampIso,required" format:"date-time"`
	To           string    `json:"to,required" format:"email"`
	Tag          string    `json:"tag"`
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
func (r EmailListResponseDataMessage) RawJSON() string { return r.JSON.raw }
func (r *EmailListResponseDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetDeliveriesResponse struct {
	Data EmailGetDeliveriesResponseData `json:"data,required"`
	Meta APIMeta                        `json:"meta,required"`
	// Any of true.
	Success bool `json:"success,required"`
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
func (r EmailGetDeliveriesResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailGetDeliveriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetDeliveriesResponseData struct {
	Deliveries []Delivery `json:"deliveries,required"`
	// Internal message ID
	MessageID string `json:"messageId,required"`
	// Message token
	MessageToken string `json:"messageToken,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deliveries   respjson.Field
		MessageID    respjson.Field
		MessageToken respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetDeliveriesResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailGetDeliveriesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRetryResponse struct {
	Data    EmailRetryResponseData `json:"data"`
	Success bool                   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailRetryResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailRetryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRetryResponseData struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailRetryResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailRetryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailSendBatchResponse struct {
	Data EmailSendBatchResponseData `json:"data,required"`
	Meta APIMeta                    `json:"meta,required"`
	// Any of true.
	Success bool `json:"success,required"`
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
	// Failed emails
	Failed int64 `json:"failed,required"`
	// Map of recipient email to message info
	Messages map[string]EmailSendBatchResponseDataMessage `json:"messages,required"`
	// Successfully queued emails
	Queued int64 `json:"queued,required"`
	// Total emails in the batch
	Total int64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Messages    respjson.Field
		Queued      respjson.Field
		Total       respjson.Field
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

type EmailGetParams struct {
	// Comma-separated list of fields to include:
	//
	// - `content` - HTML and plain text body
	// - `headers` - Email headers
	// - `deliveries` - Delivery attempt history
	// - `activity` - Opens and clicks
	Expand param.Opt[string] `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailGetParams]'s query parameters as `url.Values`.
func (r EmailGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
	// - `queued` - Email accepted and waiting to be sent
	// - `sent` - Email transmitted to recipient's mail server
	// - `delivered` - Recipient's server confirmed receipt
	// - `bounced` - Permanently rejected (hard bounce)
	// - `failed` - Delivery failed after all retry attempts
	// - `delayed` - Temporary failure, will retry
	// - `held` - Held for manual review
	//
	// Any of "queued", "sent", "delivered", "bounced", "failed", "delayed", "held".
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
// - `queued` - Email accepted and waiting to be sent
// - `sent` - Email transmitted to recipient's mail server
// - `delivered` - Recipient's server confirmed receipt
// - `bounced` - Permanently rejected (hard bounce)
// - `failed` - Delivery failed after all retry attempts
// - `delayed` - Temporary failure, will retry
// - `held` - Held for manual review
type EmailListParamsStatus string

const (
	EmailListParamsStatusQueued    EmailListParamsStatus = "queued"
	EmailListParamsStatusSent      EmailListParamsStatus = "sent"
	EmailListParamsStatusDelivered EmailListParamsStatus = "delivered"
	EmailListParamsStatusBounced   EmailListParamsStatus = "bounced"
	EmailListParamsStatusFailed    EmailListParamsStatus = "failed"
	EmailListParamsStatusDelayed   EmailListParamsStatus = "delayed"
	EmailListParamsStatusHeld      EmailListParamsStatus = "held"
)

type EmailSendParams struct {
	// Sender email. Can include name: "Name <email@domain.com>" Must be from a
	// verified domain.
	From string `json:"from,required"`
	// Email subject line
	Subject string `json:"subject,required"`
	// Recipient email addresses (max 50)
	To []string `json:"to,omitzero,required" format:"email"`
	// HTML body content. Maximum 5MB (5,242,880 characters). Combined with
	// attachments, the total message must not exceed 14MB.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Reply-to address
	ReplyTo param.Opt[string] `json:"replyTo,omitzero" format:"email"`
	// Tag for categorization and filtering
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Plain text body (auto-generated from HTML if not provided). Maximum 5MB
	// (5,242,880 characters).
	Text           param.Opt[string] `json:"text,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// File attachments
	Attachments []EmailSendParamsAttachment `json:"attachments,omitzero"`
	// BCC recipients
	Bcc []string `json:"bcc,omitzero" format:"email"`
	// CC recipients
	Cc []string `json:"cc,omitzero" format:"email"`
	// Custom email headers
	Headers map[string]string `json:"headers,omitzero"`
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
	Tag     param.Opt[string] `json:"tag,omitzero"`
	Text    param.Opt[string] `json:"text,omitzero"`
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
	// Base64-encoded RFC 2822 message
	Data string `json:"data,required"`
	// Envelope sender address
	MailFrom string `json:"mailFrom,required" format:"email"`
	// Envelope recipient addresses
	RcptTo []string `json:"rcptTo,omitzero,required" format:"email"`
	paramObj
}

func (r EmailSendRawParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendRawParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendRawParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
