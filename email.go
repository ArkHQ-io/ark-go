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
// The `data` field should contain the base64-encoded raw email.
func (r *EmailService) SendRaw(ctx context.Context, body EmailSendRawParams, opts ...option.RequestOption) (res *EmailSendRawResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "emails/raw"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EmailGetResponse struct {
	Data    EmailGetResponseData `json:"data,required"`
	Meta    shared.APIMeta       `json:"meta,required"`
	Success bool                 `json:"success,required"`
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
	// - `softfail` - Temporary delivery failure, will retry
	// - `hardfail` - Permanent delivery failure
	// - `bounced` - Email bounced back
	// - `held` - Held for manual review
	//
	// Any of "pending", "sent", "softfail", "hardfail", "bounced", "held".
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
	Deliveries []EmailGetResponseDataDelivery `json:"deliveries"`
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

type EmailGetResponseDataDelivery struct {
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
func (r EmailGetResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *EmailGetResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type EmailGetDeliveriesResponse struct {
	Data    EmailGetDeliveriesResponseData `json:"data,required"`
	Meta    shared.APIMeta                 `json:"meta,required"`
	Success bool                           `json:"success,required"`
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
	Deliveries []EmailGetDeliveriesResponseDataDelivery `json:"deliveries,required"`
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

type EmailGetDeliveriesResponseDataDelivery struct {
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
func (r EmailGetDeliveriesResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *EmailGetDeliveriesResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRetryResponse struct {
	Data    EmailRetryResponseData `json:"data,required"`
	Meta    shared.APIMeta         `json:"meta,required"`
	Success bool                   `json:"success,required"`
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
func (r EmailRetryResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailRetryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRetryResponseData struct {
	Message string `json:"message,required"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Failed      respjson.Field
		Messages    respjson.Field
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
func (r EmailSendRawResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailSendRawResponseData) UnmarshalJSON(data []byte) error {
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
	// Base64-encoded RFC 2822 message
	Data string `json:"data,required"`
	// Envelope sender address
	MailFrom string `json:"mailFrom,required" format:"email"`
	// Envelope recipient addresses
	RcptTo []string `json:"rcptTo,omitzero,required" format:"email"`
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
