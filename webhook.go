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
	"github.com/ArkHQ-io/ark-go/packages/param"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
	"github.com/ArkHQ-io/ark-go/shared"
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
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get webhook details
func (r *WebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookGetResponse, err error) {
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
func (r *WebhookService) Update(ctx context.Context, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
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
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a paginated list of delivery attempts for a specific webhook.
//
// Use this to:
//
// - Monitor webhook health and delivery success rate
// - Debug failed deliveries
// - Find specific events to replay
//
// **Filtering:**
//
// - Filter by success/failure to find problematic deliveries
// - Filter by event type to find specific events
// - Filter by time range for debugging recent issues
//
// **Retry behavior:** Failed deliveries are automatically retried with exponential
// backoff over ~3 days. Check `willRetry` to see if more attempts are scheduled.
func (r *WebhookService) ListDeliveries(ctx context.Context, webhookID string, query WebhookListDeliveriesParams, opts ...option.RequestOption) (res *WebhookListDeliveriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/deliveries", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Re-send a webhook delivery to your endpoint.
//
// **Use cases:**
//
// - Recover from transient failures after fixing your endpoint
// - Test endpoint changes with real historical data
// - Retry deliveries that failed due to downtime
//
// **How it works:**
//
// 1. Fetches the original payload from the delivery
// 2. Generates a new timestamp and signature
// 3. Sends to your webhook URL immediately
// 4. Returns the result (does not queue for retry if it fails)
//
// **Note:** The webhook must be enabled to replay deliveries.
func (r *WebhookService) ReplayDelivery(ctx context.Context, deliveryID string, body WebhookReplayDeliveryParams, opts ...option.RequestOption) (res *WebhookReplayDeliveryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.WebhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	if deliveryID == "" {
		err = errors.New("missing required deliveryId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/deliveries/%s/replay", body.WebhookID, deliveryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get detailed information about a specific webhook delivery attempt.
//
// Returns:
//
// - The complete request payload that was sent
// - Request headers including the signature
// - Response status code and body from your endpoint
// - Timing information
//
// Use this to debug why a delivery failed or verify what data was sent.
func (r *WebhookService) GetDelivery(ctx context.Context, deliveryID string, query WebhookGetDeliveryParams, opts ...option.RequestOption) (res *WebhookGetDeliveryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WebhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	if deliveryID == "" {
		err = errors.New("missing required deliveryId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/deliveries/%s", query.WebhookID, deliveryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

type WebhookNewResponse struct {
	Data    WebhookNewResponseData `json:"data,required"`
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
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewResponseData struct {
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
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError",
	// "SendLimitApproaching", "SendLimitExceeded".
	Events []string `json:"events,required"`
	// Webhook name for identification
	Name string `json:"name,required"`
	// Webhook endpoint URL
	URL  string `json:"url,required" format:"uri"`
	Uuid string `json:"uuid,required" format:"uuid"`
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
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetResponse struct {
	Data    WebhookGetResponseData `json:"data,required"`
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
func (r WebhookGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetResponseData struct {
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
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError",
	// "SendLimitApproaching", "SendLimitExceeded".
	Events []string `json:"events,required"`
	// Webhook name for identification
	Name string `json:"name,required"`
	// Webhook endpoint URL
	URL  string `json:"url,required" format:"uri"`
	Uuid string `json:"uuid,required" format:"uuid"`
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
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponse struct {
	Data    WebhookUpdateResponseData `json:"data,required"`
	Meta    shared.APIMeta            `json:"meta,required"`
	Success bool                      `json:"success,required"`
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
func (r WebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateResponseData struct {
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
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError",
	// "SendLimitApproaching", "SendLimitExceeded".
	Events []string `json:"events,required"`
	// Webhook name for identification
	Name string `json:"name,required"`
	// Webhook endpoint URL
	URL  string `json:"url,required" format:"uri"`
	Uuid string `json:"uuid,required" format:"uuid"`
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
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponse struct {
	Data    WebhookListResponseData `json:"data,required"`
	Meta    shared.APIMeta          `json:"meta,required"`
	Success bool                    `json:"success,required"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Enabled     respjson.Field
		Events      respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponseDataWebhook) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseDataWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeleteResponse struct {
	Data    WebhookDeleteResponseData `json:"data,required"`
	Meta    shared.APIMeta            `json:"meta,required"`
	Success bool                      `json:"success,required"`
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
func (r WebhookDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeleteResponseData struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of webhook delivery attempts
type WebhookListDeliveriesResponse struct {
	Data []WebhookListDeliveriesResponseData `json:"data,required"`
	Meta shared.APIMeta                      `json:"meta,required"`
	// Current page number
	Page int64 `json:"page,required"`
	// Items per page
	PerPage int64 `json:"perPage,required"`
	// Total number of deliveries matching the filter
	Total int64 `json:"total,required"`
	// Total number of pages
	TotalPages int64 `json:"totalPages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		Page        respjson.Field
		PerPage     respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListDeliveriesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListDeliveriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a webhook delivery attempt
type WebhookListDeliveriesResponseData struct {
	// Unique delivery ID (UUID)
	ID string `json:"id,required"`
	// Attempt number (1 for first attempt, increments with retries)
	Attempt int64 `json:"attempt,required"`
	// Event type that triggered this delivery
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError",
	// "SendLimitApproaching", "SendLimitExceeded".
	Event string `json:"event,required"`
	// HTTP status code returned by the endpoint (null if connection failed)
	StatusCode int64 `json:"statusCode,required"`
	// Whether the delivery was successful (2xx response)
	Success bool `json:"success,required"`
	// When this delivery attempt occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// URL the webhook was delivered to
	URL string `json:"url,required" format:"uri"`
	// ID of the webhook this delivery belongs to
	WebhookID string `json:"webhookId,required"`
	// Whether this delivery will be retried (true if failed and retries remaining)
	WillRetry bool `json:"willRetry,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attempt     respjson.Field
		Event       respjson.Field
		StatusCode  respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		URL         respjson.Field
		WebhookID   respjson.Field
		WillRetry   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListDeliveriesResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookListDeliveriesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of replaying a webhook delivery
type WebhookReplayDeliveryResponse struct {
	Data    WebhookReplayDeliveryResponseData `json:"data,required"`
	Meta    shared.APIMeta                    `json:"meta,required"`
	Success bool                              `json:"success,required"`
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
func (r WebhookReplayDeliveryResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookReplayDeliveryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookReplayDeliveryResponseData struct {
	// Request duration in milliseconds
	Duration int64 `json:"duration,required"`
	// ID of the new delivery created by the replay
	NewDeliveryID string `json:"newDeliveryId,required"`
	// ID of the original delivery that was replayed
	OriginalDeliveryID string `json:"originalDeliveryId,required"`
	// HTTP status code from your endpoint
	StatusCode int64 `json:"statusCode,required"`
	// Whether the replay was successful (2xx response from endpoint)
	Success bool `json:"success,required"`
	// When the replay was executed
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration           respjson.Field
		NewDeliveryID      respjson.Field
		OriginalDeliveryID respjson.Field
		StatusCode         respjson.Field
		Success            respjson.Field
		Timestamp          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookReplayDeliveryResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookReplayDeliveryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed information about a webhook delivery attempt
type WebhookGetDeliveryResponse struct {
	// Full details of a webhook delivery including request and response
	Data    WebhookGetDeliveryResponseData `json:"data,required"`
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
func (r WebhookGetDeliveryResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetDeliveryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full details of a webhook delivery including request and response
type WebhookGetDeliveryResponseData struct {
	// Unique delivery ID (UUID)
	ID string `json:"id,required"`
	// Attempt number for this delivery
	Attempt int64 `json:"attempt,required"`
	// Event type that triggered this delivery
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError",
	// "SendLimitApproaching", "SendLimitExceeded".
	Event string `json:"event,required"`
	// The request that was sent to your endpoint
	Request WebhookGetDeliveryResponseDataRequest `json:"request,required"`
	// The response received from your endpoint
	Response WebhookGetDeliveryResponseDataResponse `json:"response,required"`
	// HTTP status code returned by the endpoint
	StatusCode int64 `json:"statusCode,required"`
	// Whether the delivery was successful (2xx response)
	Success bool `json:"success,required"`
	// When this delivery attempt occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// URL the webhook was delivered to
	URL string `json:"url,required" format:"uri"`
	// ID of the webhook this delivery belongs to
	WebhookID string `json:"webhookId,required"`
	// Name of the webhook for easy identification
	WebhookName string `json:"webhookName,required"`
	// Whether this delivery will be retried
	WillRetry bool `json:"willRetry,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attempt     respjson.Field
		Event       respjson.Field
		Request     respjson.Field
		Response    respjson.Field
		StatusCode  respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		URL         respjson.Field
		WebhookID   respjson.Field
		WebhookName respjson.Field
		WillRetry   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetDeliveryResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetDeliveryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The request that was sent to your endpoint
type WebhookGetDeliveryResponseDataRequest struct {
	// HTTP headers that were sent with the request
	Headers map[string]string `json:"headers,required"`
	// The complete webhook payload that was sent
	Payload map[string]any `json:"payload,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetDeliveryResponseDataRequest) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetDeliveryResponseDataRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response received from your endpoint
type WebhookGetDeliveryResponseDataResponse struct {
	// HTTP status code from your endpoint
	StatusCode int64 `json:"statusCode,required"`
	// Response body from your endpoint (may be truncated)
	Body string `json:"body,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StatusCode  respjson.Field
		Body        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetDeliveryResponseDataResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetDeliveryResponseDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTestResponse struct {
	Data    WebhookTestResponseData `json:"data,required"`
	Meta    shared.APIMeta          `json:"meta,required"`
	Success bool                    `json:"success,required"`
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
	// Webhook name for identification
	Name string `json:"name,required"`
	// HTTPS endpoint URL
	URL string `json:"url,required" format:"uri"`
	// Subscribe to all events (ignores events array, accepts null)
	AllEvents param.Opt[bool] `json:"allEvents,omitzero"`
	// Whether the webhook is enabled (accepts null)
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Events to subscribe to (accepts null):
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
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError",
	// "SendLimitApproaching", "SendLimitExceeded".
	Events []string `json:"events,omitzero"`
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

type WebhookListDeliveriesParams struct {
	// Only deliveries after this Unix timestamp
	After param.Opt[int64] `query:"after,omitzero" json:"-"`
	// Only deliveries before this Unix timestamp
	Before param.Opt[int64] `query:"before,omitzero" json:"-"`
	// Page number (default 1)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Items per page (default 30, max 100)
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	// Filter by delivery success (true = 2xx response, false = non-2xx or error)
	Success param.Opt[bool] `query:"success,omitzero" json:"-"`
	// Filter by event type
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError",
	// "SendLimitApproaching", "SendLimitExceeded".
	Event WebhookListDeliveriesParamsEvent `query:"event,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListDeliveriesParams]'s query parameters as
// `url.Values`.
func (r WebhookListDeliveriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by event type
type WebhookListDeliveriesParamsEvent string

const (
	WebhookListDeliveriesParamsEventMessageSent           WebhookListDeliveriesParamsEvent = "MessageSent"
	WebhookListDeliveriesParamsEventMessageDelayed        WebhookListDeliveriesParamsEvent = "MessageDelayed"
	WebhookListDeliveriesParamsEventMessageDeliveryFailed WebhookListDeliveriesParamsEvent = "MessageDeliveryFailed"
	WebhookListDeliveriesParamsEventMessageHeld           WebhookListDeliveriesParamsEvent = "MessageHeld"
	WebhookListDeliveriesParamsEventMessageBounced        WebhookListDeliveriesParamsEvent = "MessageBounced"
	WebhookListDeliveriesParamsEventMessageLinkClicked    WebhookListDeliveriesParamsEvent = "MessageLinkClicked"
	WebhookListDeliveriesParamsEventMessageLoaded         WebhookListDeliveriesParamsEvent = "MessageLoaded"
	WebhookListDeliveriesParamsEventDomainDNSError        WebhookListDeliveriesParamsEvent = "DomainDNSError"
	WebhookListDeliveriesParamsEventSendLimitApproaching  WebhookListDeliveriesParamsEvent = "SendLimitApproaching"
	WebhookListDeliveriesParamsEventSendLimitExceeded     WebhookListDeliveriesParamsEvent = "SendLimitExceeded"
)

type WebhookReplayDeliveryParams struct {
	WebhookID string `path:"webhookId,required" json:"-"`
	paramObj
}

type WebhookGetDeliveryParams struct {
	WebhookID string `path:"webhookId,required" json:"-"`
	paramObj
}

type WebhookTestParams struct {
	// Event type to simulate
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError",
	// "SendLimitApproaching", "SendLimitExceeded".
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
	WebhookTestParamsEventSendLimitApproaching  WebhookTestParamsEvent = "SendLimitApproaching"
	WebhookTestParamsEventSendLimitExceeded     WebhookTestParamsEvent = "SendLimitExceeded"
)
