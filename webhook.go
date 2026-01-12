// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/ArkHQ-io/ark-go/internal/apijson"
	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/option"
	"github.com/ArkHQ-io/ark-go/packages/param"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Create a webhook endpoint to receive email event notifications.
//
// **Available events:**
//
// - `MessageSent` - Email accepted by recipient server
// - `MessageDeliveryFailed` - Delivery permanently failed
// - `MessageDelayed` - Delivery temporarily failed, will retry
// - `MessageBounced` - Email bounced
// - `MessageHeld` - Email held for review
// - `MessageLinkClicked` - Recipient clicked a link
// - `MessageLoaded` - Recipient opened the email
// - `DomainDNSError` - Domain DNS issue detected
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get webhook details
func (r *WebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a webhook
func (r *WebhookService) Update(ctx context.Context, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all configured webhook endpoints
func (r *WebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a webhook
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Send a test payload to your webhook endpoint and verify it receives the data
// correctly.
//
// Use this to:
//
// - Verify your webhook URL is accessible
// - Test your signature verification code
// - Ensure your server handles the payload format correctly
//
// **Test payload format:** The test payload is identical to real webhook payloads,
// containing sample data for the specified event type. Your webhook should respond
// with a 2xx status code.
func (r *WebhookService) Test(ctx context.Context, webhookID string, body WebhookTestParams, opts ...option.RequestOption) (res *WebhookTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/test", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type WebhookResponse struct {
	Data WebhookResponseData `json:"data,required"`
	Meta APIMeta             `json:"meta,required"`
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
func (r WebhookResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookResponseData struct {
	// Webhook ID
	ID string `json:"id,required"`
	// Whether subscribed to all events
	AllEvents bool      `json:"allEvents,required"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled,required"`
	// Subscribed events
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events,required"`
	// Webhook name for identification
	Name string `json:"name,required"`
	// Webhook endpoint URL
	URL  string `json:"url,required" format:"uri"`
	Uuid string `json:"uuid,required" format:"uuid"`
	// Whether the webhook payloads are signed (always true for new webhooks)
	Signed bool `json:"signed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AllEvents   respjson.Field
		CreatedAt   respjson.Field
		Enabled     respjson.Field
		Events      respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		Uuid        respjson.Field
		Signed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponse struct {
	Data WebhookListResponseData `json:"data,required"`
	Meta APIMeta                 `json:"meta,required"`
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
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponseData struct {
	Webhooks []WebhookListResponseDataWebhook `json:"webhooks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponseDataWebhook struct {
	// Webhook ID
	ID      string   `json:"id,required"`
	Enabled bool     `json:"enabled,required"`
	Events  []string `json:"events,required"`
	Name    string   `json:"name,required"`
	URL     string   `json:"url,required" format:"uri"`
	// Whether webhook payloads are signed
	Signed bool `json:"signed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Enabled     respjson.Field
		Events      respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		Signed      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseDataWebhook) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseDataWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestResponse struct {
	Data WebhookTestResponseData `json:"data,required"`
	Meta APIMeta                 `json:"meta,required"`
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
func (r WebhookTestResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestResponseData struct {
	// Request duration in milliseconds
	Duration int64 `json:"duration,required"`
	// Event type that was tested
	Event string `json:"event,required"`
	// HTTP status code from the webhook endpoint
	StatusCode int64 `json:"statusCode,required"`
	// Whether the webhook endpoint responded with a 2xx status
	Success bool `json:"success,required"`
	// Response body from the webhook endpoint (truncated if too long)
	Body string `json:"body,nullable"`
	// Error message if the request failed
	Error string `json:"error,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Event       respjson.Field
		StatusCode  respjson.Field
		Success     respjson.Field
		Body        respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTestResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookTestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// Events to subscribe to:
	//
	// - `MessageSent` - Email successfully delivered to recipient's server
	// - `MessageDelayed` - Temporary delivery failure, will retry
	// - `MessageDeliveryFailed` - Permanent delivery failure
	// - `MessageHeld` - Email held for manual review
	// - `MessageBounced` - Email bounced back
	// - `MessageLinkClicked` - Recipient clicked a tracked link
	// - `MessageLoaded` - Recipient opened the email (tracking pixel loaded)
	// - `DomainDNSError` - DNS configuration issue detected
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events,omitzero,required"`
	// Webhook name for identification
	Name string `json:"name,required"`
	// HTTPS endpoint URL
	URL string `json:"url,required" format:"uri"`
	// Subscribe to all events (ignores events array)
	AllEvents param.Opt[bool] `json:"allEvents,omitzero"`
	Enabled   param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	AllEvents param.Opt[bool]   `json:"allEvents,omitzero"`
	Enabled   param.Opt[bool]   `json:"enabled,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	URL       param.Opt[string] `json:"url,omitzero" format:"uri"`
	Events    []string          `json:"events,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestParams struct {
	// Event type to simulate
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Event WebhookTestParamsEvent `json:"event,omitzero,required"`
	paramObj
}

func (r WebhookTestParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event type to simulate
type WebhookTestParamsEvent string

const (
	WebhookTestParamsEventMessageSent           WebhookTestParamsEvent = "MessageSent"
	WebhookTestParamsEventMessageDelayed        WebhookTestParamsEvent = "MessageDelayed"
	WebhookTestParamsEventMessageDeliveryFailed WebhookTestParamsEvent = "MessageDeliveryFailed"
	WebhookTestParamsEventMessageHeld           WebhookTestParamsEvent = "MessageHeld"
	WebhookTestParamsEventMessageBounced        WebhookTestParamsEvent = "MessageBounced"
	WebhookTestParamsEventMessageLinkClicked    WebhookTestParamsEvent = "MessageLinkClicked"
	WebhookTestParamsEventMessageLoaded         WebhookTestParamsEvent = "MessageLoaded"
	WebhookTestParamsEventDomainDNSError        WebhookTestParamsEvent = "DomainDNSError"
)
