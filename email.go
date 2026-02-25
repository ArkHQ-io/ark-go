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
// - `GET /emails/{emailId}` - Get full details of a specific email
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
// - `GET /emails/{emailId}` - Get full details of a specific email
// - `POST /emails` - Send a new email
func (r *EmailService) ListAutoPaging(ctx context.Context, query EmailListParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[EmailListResponse] {
	return pagination.NewPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// Get the complete delivery history for an email, including SMTP response codes,
// timestamps, and current retry state.
//
// ## Response Fields
//
// ### Status
//
// The current status of the email:
//
// - `pending` - Awaiting first delivery attempt
// - `sent` - Successfully delivered to recipient server
// - `softfail` - Temporary failure, automatic retry scheduled
// - `hardfail` - Permanent failure, will not retry
// - `held` - Held for manual review
// - `bounced` - Bounced by recipient server
//
// ### Retry State
//
// When the email is in the delivery queue (`pending` or `softfail` status),
// `retryState` provides information about the retry schedule:
//
// - `attempt` - Current attempt number (0 = first attempt)
// - `maxAttempts` - Maximum attempts before hard-fail (typically 18)
// - `attemptsRemaining` - Attempts left before hard-fail
// - `nextRetryAt` - When the next retry is scheduled (Unix timestamp)
// - `processing` - Whether the email is currently being processed
// - `manual` - Whether this was triggered by a manual retry
//
// When the email has finished processing (`sent`, `hardfail`, `held`, `bounced`),
// `retryState` is `null`.
//
// ### Can Retry Manually
//
// Indicates whether you can call `POST /emails/{emailId}/retry` to manually retry
// the email. This is `true` when the raw message content is still available (not
// expired due to retention policy).
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
// - `GET /emails/{emailId}` - Track delivery status
// - `GET /emails/{emailId}/deliveries` - View delivery attempts
// - `POST /emails/{emailId}/retry` - Retry failed delivery
func (r *EmailService) Send(ctx context.Context, params EmailSendParams, opts ...option.RequestOption) (res *EmailSendResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
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
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
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

type EmailGetResponse struct {
	Data    EmailGetResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta       `json:"meta" api:"required"`
	Success bool                 `json:"success" api:"required"`
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
	// Unique message identifier (token)
	ID string `json:"id" api:"required"`
	// Sender address
	From string `json:"from" api:"required"`
	// Message direction
	//
	// Any of "outgoing", "incoming".
	Scope string `json:"scope" api:"required"`
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
	Status string `json:"status" api:"required"`
	// Email subject line
	Subject string `json:"subject" api:"required"`
	// The tenant ID this email belongs to
	TenantID string `json:"tenantId" api:"required"`
	// Unix timestamp when the email was sent
	Timestamp float64 `json:"timestamp" api:"required"`
	// ISO 8601 formatted timestamp
	TimestampISO time.Time `json:"timestampIso" api:"required" format:"date-time"`
	// Recipient address
	To string `json:"to" api:"required" format:"email"`
	// Opens and clicks tracking data (included if expand=activity)
	Activity EmailGetResponseDataActivity `json:"activity"`
	// File attachments (included if expand=attachments)
	Attachments []EmailGetResponseDataAttachment `json:"attachments"`
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
	// Complete raw MIME message, base64 encoded (included if expand=raw). Decode this
	// to get the original RFC 2822 formatted email.
	RawMessage string `json:"rawMessage"`
	// Whether the message was flagged as spam
	Spam bool `json:"spam"`
	// Spam score (if applicable)
	SpamScore float64 `json:"spamScore"`
	// Optional categorization tag
	Tag string `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		From         respjson.Field
		Scope        respjson.Field
		Status       respjson.Field
		Subject      respjson.Field
		TenantID     respjson.Field
		Timestamp    respjson.Field
		TimestampISO respjson.Field
		To           respjson.Field
		Activity     respjson.Field
		Attachments  respjson.Field
		Deliveries   respjson.Field
		Headers      respjson.Field
		HTMLBody     respjson.Field
		MessageID    respjson.Field
		PlainBody    respjson.Field
		RawMessage   respjson.Field
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

// Opens and clicks tracking data (included if expand=activity)
type EmailGetResponseDataActivity struct {
	// List of link click events
	Clicks []EmailGetResponseDataActivityClick `json:"clicks"`
	// List of email open events
	Opens []EmailGetResponseDataActivityOpen `json:"opens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clicks      respjson.Field
		Opens       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetResponseDataActivity) RawJSON() string { return r.JSON.raw }
func (r *EmailGetResponseDataActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetResponseDataActivityClick struct {
	// IP address of the clicker
	IPAddress string `json:"ipAddress"`
	// Unix timestamp of the click event
	Timestamp float64 `json:"timestamp"`
	// ISO 8601 timestamp of the click event
	TimestampISO time.Time `json:"timestampIso" format:"date-time"`
	// URL that was clicked
	URL string `json:"url" format:"uri"`
	// User agent of the email client
	UserAgent string `json:"userAgent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress    respjson.Field
		Timestamp    respjson.Field
		TimestampISO respjson.Field
		URL          respjson.Field
		UserAgent    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetResponseDataActivityClick) RawJSON() string { return r.JSON.raw }
func (r *EmailGetResponseDataActivityClick) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetResponseDataActivityOpen struct {
	// IP address of the opener
	IPAddress string `json:"ipAddress"`
	// Unix timestamp of the open event
	Timestamp float64 `json:"timestamp"`
	// ISO 8601 timestamp of the open event
	TimestampISO time.Time `json:"timestampIso" format:"date-time"`
	// User agent of the email client
	UserAgent string `json:"userAgent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress    respjson.Field
		Timestamp    respjson.Field
		TimestampISO respjson.Field
		UserAgent    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetResponseDataActivityOpen) RawJSON() string { return r.JSON.raw }
func (r *EmailGetResponseDataActivityOpen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An email attachment retrieved from a sent message
type EmailGetResponseDataAttachment struct {
	// MIME type of the attachment
	ContentType string `json:"contentType" api:"required"`
	// Base64 encoded attachment content. Decode this to get the raw file bytes.
	Data string `json:"data" api:"required"`
	// Original filename of the attachment
	Filename string `json:"filename" api:"required"`
	// SHA256 hash of the attachment content for verification
	Hash string `json:"hash" api:"required"`
	// Size of the attachment in bytes
	Size int64 `json:"size" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Data        respjson.Field
		Filename    respjson.Field
		Hash        respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetResponseDataAttachment) RawJSON() string { return r.JSON.raw }
func (r *EmailGetResponseDataAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetResponseDataDelivery struct {
	// Delivery attempt ID
	ID string `json:"id" api:"required"`
	// Delivery status (lowercase)
	Status string `json:"status" api:"required"`
	// Unix timestamp
	Timestamp float64 `json:"timestamp" api:"required"`
	// ISO 8601 timestamp
	TimestampISO time.Time `json:"timestampIso" api:"required" format:"date-time"`
	// Bounce classification category (present for failed deliveries). Helps understand
	// why delivery failed for analytics and automated handling.
	//
	// Any of "invalid_recipient", "mailbox_full", "message_too_large", "spam_block",
	// "policy_violation", "no_mailbox", "not_accepting_mail",
	// "temporarily_unavailable", "protocol_error", "tls_required", "connection_error",
	// "dns_error", "unclassified".
	Classification string `json:"classification" api:"nullable"`
	// Numeric bounce classification code for programmatic handling. Codes:
	// 10=invalid_recipient, 11=no_mailbox, 12=not_accepting_mail, 20=mailbox_full,
	// 21=message_too_large, 30=spam_block, 31=policy_violation, 32=tls_required,
	// 40=connection_error, 41=dns_error, 42=temporarily_unavailable,
	// 50=protocol_error, 99=unclassified
	ClassificationCode int64 `json:"classificationCode" api:"nullable"`
	// SMTP response code
	Code int64 `json:"code"`
	// Human-readable delivery summary. Format varies by status:
	//
	//   - **sent**: `Message for {recipient} accepted by {ip}:{port} ({hostname})`
	//   - **softfail/hardfail**:
	//     `{code} {classification}: Delivery to {recipient} failed at {ip}:{port} ({hostname})`
	Details string `json:"details"`
	// Raw SMTP response from the receiving mail server
	Output string `json:"output"`
	// Hostname of the remote mail server that processed the delivery. Present for all
	// delivery attempts (successful and failed).
	RemoteHost string `json:"remoteHost" api:"nullable"`
	// Whether TLS was used
	SentWithSsl bool `json:"sentWithSsl"`
	// RFC 3463 enhanced status code from SMTP response (e.g., "5.1.1", "4.2.2"). First
	// digit: 2=success, 4=temporary, 5=permanent. Second digit: category (1=address,
	// 2=mailbox, 7=security, etc.).
	SmtpEnhancedCode string `json:"smtpEnhancedCode" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Status             respjson.Field
		Timestamp          respjson.Field
		TimestampISO       respjson.Field
		Classification     respjson.Field
		ClassificationCode respjson.Field
		Code               respjson.Field
		Details            respjson.Field
		Output             respjson.Field
		RemoteHost         respjson.Field
		SentWithSsl        respjson.Field
		SmtpEnhancedCode   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *EmailGetResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailListResponse struct {
	// Unique message identifier (token)
	ID   string `json:"id" api:"required"`
	From string `json:"from" api:"required"`
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
	Status  EmailListResponseStatus `json:"status" api:"required"`
	Subject string                  `json:"subject" api:"required"`
	// The tenant ID this email belongs to
	TenantID     string    `json:"tenantId" api:"required"`
	Timestamp    float64   `json:"timestamp" api:"required"`
	TimestampISO time.Time `json:"timestampIso" api:"required" format:"date-time"`
	To           string    `json:"to" api:"required" format:"email"`
	Tag          string    `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		From         respjson.Field
		Status       respjson.Field
		Subject      respjson.Field
		TenantID     respjson.Field
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
	Data    EmailGetDeliveriesResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                 `json:"meta" api:"required"`
	Success bool                           `json:"success" api:"required"`
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
	// Message identifier (token)
	ID string `json:"id" api:"required"`
	// Whether the message can be manually retried via `POST /emails/{emailId}/retry`.
	// `true` when the raw message content is still available (not expired). Messages
	// older than the retention period cannot be retried.
	CanRetryManually bool `json:"canRetryManually" api:"required"`
	// Chronological list of delivery attempts for this message. Each attempt includes
	// SMTP response codes and timestamps.
	Deliveries []EmailGetDeliveriesResponseDataDelivery `json:"deliveries" api:"required"`
	// Information about the current retry state of a message that is queued for
	// delivery. Only present when the message is in the delivery queue.
	RetryState EmailGetDeliveriesResponseDataRetryState `json:"retryState" api:"required"`
	// Current message status (lowercase). Possible values:
	//
	// - `pending` - Initial state, awaiting first delivery attempt
	// - `sent` - Successfully delivered
	// - `softfail` - Temporary failure, will retry automatically
	// - `hardfail` - Permanent failure, will not retry
	// - `held` - Held for manual review (suppression list, etc.)
	// - `bounced` - Bounced by recipient server
	//
	// Any of "pending", "sent", "softfail", "hardfail", "held", "bounced".
	Status string `json:"status" api:"required"`
	// The tenant ID this email belongs to
	TenantID string `json:"tenantId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CanRetryManually respjson.Field
		Deliveries       respjson.Field
		RetryState       respjson.Field
		Status           respjson.Field
		TenantID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetDeliveriesResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailGetDeliveriesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailGetDeliveriesResponseDataDelivery struct {
	// Delivery attempt ID
	ID string `json:"id" api:"required"`
	// Delivery status (lowercase)
	Status string `json:"status" api:"required"`
	// Unix timestamp
	Timestamp float64 `json:"timestamp" api:"required"`
	// ISO 8601 timestamp
	TimestampISO time.Time `json:"timestampIso" api:"required" format:"date-time"`
	// Bounce classification category (present for failed deliveries). Helps understand
	// why delivery failed for analytics and automated handling.
	//
	// Any of "invalid_recipient", "mailbox_full", "message_too_large", "spam_block",
	// "policy_violation", "no_mailbox", "not_accepting_mail",
	// "temporarily_unavailable", "protocol_error", "tls_required", "connection_error",
	// "dns_error", "unclassified".
	Classification string `json:"classification" api:"nullable"`
	// Numeric bounce classification code for programmatic handling. Codes:
	// 10=invalid_recipient, 11=no_mailbox, 12=not_accepting_mail, 20=mailbox_full,
	// 21=message_too_large, 30=spam_block, 31=policy_violation, 32=tls_required,
	// 40=connection_error, 41=dns_error, 42=temporarily_unavailable,
	// 50=protocol_error, 99=unclassified
	ClassificationCode int64 `json:"classificationCode" api:"nullable"`
	// SMTP response code
	Code int64 `json:"code"`
	// Human-readable delivery summary. Format varies by status:
	//
	//   - **sent**: `Message for {recipient} accepted by {ip}:{port} ({hostname})`
	//   - **softfail/hardfail**:
	//     `{code} {classification}: Delivery to {recipient} failed at {ip}:{port} ({hostname})`
	Details string `json:"details"`
	// Raw SMTP response from the receiving mail server
	Output string `json:"output"`
	// Hostname of the remote mail server that processed the delivery. Present for all
	// delivery attempts (successful and failed).
	RemoteHost string `json:"remoteHost" api:"nullable"`
	// Whether TLS was used
	SentWithSsl bool `json:"sentWithSsl"`
	// RFC 3463 enhanced status code from SMTP response (e.g., "5.1.1", "4.2.2"). First
	// digit: 2=success, 4=temporary, 5=permanent. Second digit: category (1=address,
	// 2=mailbox, 7=security, etc.).
	SmtpEnhancedCode string `json:"smtpEnhancedCode" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Status             respjson.Field
		Timestamp          respjson.Field
		TimestampISO       respjson.Field
		Classification     respjson.Field
		ClassificationCode respjson.Field
		Code               respjson.Field
		Details            respjson.Field
		Output             respjson.Field
		RemoteHost         respjson.Field
		SentWithSsl        respjson.Field
		SmtpEnhancedCode   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetDeliveriesResponseDataDelivery) RawJSON() string { return r.JSON.raw }
func (r *EmailGetDeliveriesResponseDataDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the current retry state of a message that is queued for
// delivery. Only present when the message is in the delivery queue.
type EmailGetDeliveriesResponseDataRetryState struct {
	// Current attempt number (0-indexed). The first delivery attempt is 0, the first
	// retry is 1, and so on.
	Attempt int64 `json:"attempt" api:"required"`
	// Number of attempts remaining before the message is hard-failed. Calculated as
	// `maxAttempts - attempt`.
	AttemptsRemaining int64 `json:"attemptsRemaining" api:"required"`
	// Whether this queue entry was created by a manual retry request. Manual retries
	// bypass certain hold conditions like suppression lists.
	Manual bool `json:"manual" api:"required"`
	// Maximum number of delivery attempts before the message is hard-failed.
	// Configured at the server level.
	MaxAttempts int64 `json:"maxAttempts" api:"required"`
	// Whether the message is currently being processed by a delivery worker. When
	// `true`, the message is actively being sent.
	Processing bool `json:"processing" api:"required"`
	// Unix timestamp of when the next retry attempt is scheduled. `null` if the
	// message is ready for immediate processing or currently being processed.
	NextRetryAt float64 `json:"nextRetryAt" api:"nullable"`
	// ISO 8601 formatted timestamp of the next retry attempt. `null` if the message is
	// ready for immediate processing.
	NextRetryAtISO time.Time `json:"nextRetryAtIso" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt           respjson.Field
		AttemptsRemaining respjson.Field
		Manual            respjson.Field
		MaxAttempts       respjson.Field
		Processing        respjson.Field
		NextRetryAt       respjson.Field
		NextRetryAtISO    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailGetDeliveriesResponseDataRetryState) RawJSON() string { return r.JSON.raw }
func (r *EmailGetDeliveriesResponseDataRetryState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRetryResponse struct {
	Data    EmailRetryResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta         `json:"meta" api:"required"`
	Success bool                   `json:"success" api:"required"`
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
	// Email identifier (token)
	ID      string `json:"id" api:"required"`
	Message string `json:"message" api:"required"`
	// The tenant ID this email belongs to
	TenantID string `json:"tenantId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Message     respjson.Field
		TenantID    respjson.Field
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
	Data    EmailSendResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta        `json:"meta" api:"required"`
	Success bool                  `json:"success" api:"required"`
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
	// Unique message identifier (token)
	ID string `json:"id" api:"required"`
	// Current delivery status
	//
	// Any of "pending", "sent".
	Status string `json:"status" api:"required"`
	// The tenant ID this email was sent from
	TenantID string `json:"tenantId" api:"required"`
	// List of recipient addresses
	To []string `json:"to" api:"required" format:"email"`
	// SMTP Message-ID header value
	MessageID string `json:"messageId"`
	// Whether this email was sent in sandbox mode. Only present (and true) for sandbox
	// emails sent from @arkhq.io addresses.
	Sandbox bool `json:"sandbox"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		TenantID    respjson.Field
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
	Data    EmailSendBatchResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta             `json:"meta" api:"required"`
	Success bool                       `json:"success" api:"required"`
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
	Accepted int64 `json:"accepted" api:"required"`
	// Failed emails
	Failed int64 `json:"failed" api:"required"`
	// Map of recipient email to message info
	Messages map[string]EmailSendBatchResponseDataMessage `json:"messages" api:"required"`
	// The tenant ID this batch was sent from
	TenantID string `json:"tenantId" api:"required"`
	// Total emails in the batch
	Total int64 `json:"total" api:"required"`
	// Whether this batch was sent in sandbox mode. Only present (and true) for sandbox
	// emails sent from @arkhq.io addresses.
	Sandbox bool `json:"sandbox"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		Failed      respjson.Field
		Messages    respjson.Field
		TenantID    respjson.Field
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
	// Message identifier (token)
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
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
	Data    EmailSendRawResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta           `json:"meta" api:"required"`
	Success bool                     `json:"success" api:"required"`
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
	// Unique message identifier (token)
	ID string `json:"id" api:"required"`
	// Current delivery status
	//
	// Any of "pending", "sent".
	Status string `json:"status" api:"required"`
	// The tenant ID this email was sent from
	TenantID string `json:"tenantId" api:"required"`
	// List of recipient addresses
	To []string `json:"to" api:"required" format:"email"`
	// SMTP Message-ID header value
	MessageID string `json:"messageId"`
	// Whether this email was sent in sandbox mode. Only present (and true) for sandbox
	// emails sent from @arkhq.io addresses.
	Sandbox bool `json:"sandbox"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		TenantID    respjson.Field
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

type EmailGetParams struct {
	// Comma-separated list of fields to include:
	//
	// - `full` - Include all expanded fields in a single request
	// - `content` - HTML and plain text body
	// - `headers` - Email headers
	// - `deliveries` - Delivery attempt history
	// - `activity` - Opens and clicks tracking data
	// - `attachments` - File attachments with content (base64 encoded)
	// - `raw` - Complete raw MIME message (base64 encoded)
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
	From string `json:"from" api:"required"`
	// Email subject line
	Subject string `json:"subject" api:"required"`
	// Recipient email addresses (max 50)
	To []string `json:"to,omitzero" api:"required" format:"email"`
	// HTML body content (accepts null). Maximum 5MB (5,242,880 characters). Combined
	// with attachments, the total message must not exceed 14MB.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Reply-to address (accepts null)
	ReplyTo param.Opt[string] `json:"replyTo,omitzero" format:"email"`
	// Tag for categorization and filtering (accepts null)
	Tag param.Opt[string] `json:"tag,omitzero"`
	// The tenant ID to send this email from. Determines which tenant's configuration
	// (domains, webhooks, tracking) is used.
	//
	//   - If your API key is scoped to a specific tenant, this must match that tenant or
	//     be omitted.
	//   - If your API key is org-level, specify the tenant to send from.
	//   - If omitted, the organization's default tenant is used.
	TenantID param.Opt[string] `json:"tenantId,omitzero"`
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
	Content string `json:"content" api:"required"`
	// MIME type
	ContentType string `json:"contentType" api:"required"`
	// Attachment filename
	Filename string `json:"filename" api:"required"`
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
	Emails []EmailSendBatchParamsEmail `json:"emails,omitzero" api:"required"`
	// Sender email for all messages
	From string `json:"from" api:"required"`
	// The tenant ID to send this batch from. Determines which tenant's configuration
	// (domains, webhooks, tracking) is used.
	//
	//   - If your API key is scoped to a specific tenant, this must match that tenant or
	//     be omitted.
	//   - If your API key is org-level, specify the tenant to send from.
	//   - If omitted, the organization's default tenant is used.
	TenantID       param.Opt[string] `json:"tenantId,omitzero"`
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
	Subject string            `json:"subject" api:"required"`
	To      []string          `json:"to,omitzero" api:"required" format:"email"`
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
	From string `json:"from" api:"required"`
	// Base64-encoded RFC 2822 MIME message.
	//
	// **You must base64-encode your raw email before sending.** The raw email should
	// include headers (From, To, Subject, Content-Type, etc.) followed by a blank line
	// and the message body.
	RawMessage string `json:"rawMessage" api:"required"`
	// Recipient email addresses
	To []string `json:"to,omitzero" api:"required" format:"email"`
	// Whether this is a bounce message (accepts null)
	Bounce param.Opt[bool] `json:"bounce,omitzero"`
	// The tenant ID to send this email from. Determines which tenant's configuration
	// (domains, webhooks, tracking) is used.
	//
	//   - If your API key is scoped to a specific tenant, this must match that tenant or
	//     be omitted.
	//   - If your API key is org-level, specify the tenant to send from.
	//   - If omitted, the organization's default tenant is used.
	TenantID param.Opt[string] `json:"tenantId,omitzero"`
	paramObj
}

func (r EmailSendRawParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailSendRawParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailSendRawParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
