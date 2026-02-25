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

// TenantWebhookService contains methods and other services that help with
// interacting with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantWebhookService] method instead.
type TenantWebhookService struct {
	Options []option.RequestOption
}

// NewTenantWebhookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTenantWebhookService(opts ...option.RequestOption) (r TenantWebhookService) {
	r = TenantWebhookService{}
	r.Options = opts
	return
}

// Create a webhook endpoint to receive email event notifications for a tenant.
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
func (r *TenantWebhookService) New(ctx context.Context, tenantID string, body TenantWebhookNewParams, opts ...option.RequestOption) (res *TenantWebhookNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get webhook details
func (r *TenantWebhookService) Get(ctx context.Context, webhookID string, query TenantWebhookGetParams, opts ...option.RequestOption) (res *TenantWebhookGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks/%s", query.TenantID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a webhook
func (r *TenantWebhookService) Update(ctx context.Context, webhookID string, params TenantWebhookUpdateParams, opts ...option.RequestOption) (res *TenantWebhookUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks/%s", params.TenantID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get all configured webhook endpoints for a tenant.
func (r *TenantWebhookService) List(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *TenantWebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a webhook
func (r *TenantWebhookService) Delete(ctx context.Context, webhookID string, body TenantWebhookDeleteParams, opts ...option.RequestOption) (res *TenantWebhookDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks/%s", body.TenantID, webhookID)
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
func (r *TenantWebhookService) ListDeliveries(ctx context.Context, webhookID string, params TenantWebhookListDeliveriesParams, opts ...option.RequestOption) (res *TenantWebhookListDeliveriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks/%s/deliveries", params.TenantID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
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
func (r *TenantWebhookService) ReplayDelivery(ctx context.Context, deliveryID string, body TenantWebhookReplayDeliveryParams, opts ...option.RequestOption) (res *TenantWebhookReplayDeliveryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if body.WebhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	if deliveryID == "" {
		err = errors.New("missing required deliveryId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks/%s/deliveries/%s/replay", body.TenantID, body.WebhookID, deliveryID)
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
func (r *TenantWebhookService) GetDelivery(ctx context.Context, deliveryID string, query TenantWebhookGetDeliveryParams, opts ...option.RequestOption) (res *TenantWebhookGetDeliveryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if query.WebhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	if deliveryID == "" {
		err = errors.New("missing required deliveryId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks/%s/deliveries/%s", query.TenantID, query.WebhookID, deliveryID)
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
func (r *TenantWebhookService) Test(ctx context.Context, webhookID string, params TenantWebhookTestParams, opts ...option.RequestOption) (res *TenantWebhookTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/webhooks/%s/test", params.TenantID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type TenantWebhookNewResponse struct {
	Data    TenantWebhookNewResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta               `json:"meta" api:"required"`
	Success bool                         `json:"success" api:"required"`
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
func (r TenantWebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookNewResponseData struct {
	// Webhook ID
	ID string `json:"id" api:"required"`
	// Whether subscribed to all events
	AllEvents bool      `json:"allEvents" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events" api:"required"`
	// Webhook name for identification
	Name string `json:"name" api:"required"`
	// Webhook endpoint URL
	URL  string `json:"url" api:"required" format:"uri"`
	Uuid string `json:"uuid" api:"required" format:"uuid"`
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
func (r TenantWebhookNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookGetResponse struct {
	Data    TenantWebhookGetResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta               `json:"meta" api:"required"`
	Success bool                         `json:"success" api:"required"`
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
func (r TenantWebhookGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookGetResponseData struct {
	// Webhook ID
	ID string `json:"id" api:"required"`
	// Whether subscribed to all events
	AllEvents bool      `json:"allEvents" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events" api:"required"`
	// Webhook name for identification
	Name string `json:"name" api:"required"`
	// Webhook endpoint URL
	URL  string `json:"url" api:"required" format:"uri"`
	Uuid string `json:"uuid" api:"required" format:"uuid"`
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
func (r TenantWebhookGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookUpdateResponse struct {
	Data    TenantWebhookUpdateResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                  `json:"meta" api:"required"`
	Success bool                            `json:"success" api:"required"`
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
func (r TenantWebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookUpdateResponseData struct {
	// Webhook ID
	ID string `json:"id" api:"required"`
	// Whether subscribed to all events
	AllEvents bool      `json:"allEvents" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events" api:"required"`
	// Webhook name for identification
	Name string `json:"name" api:"required"`
	// Webhook endpoint URL
	URL  string `json:"url" api:"required" format:"uri"`
	Uuid string `json:"uuid" api:"required" format:"uuid"`
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
func (r TenantWebhookUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookListResponse struct {
	Data    TenantWebhookListResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                `json:"meta" api:"required"`
	Success bool                          `json:"success" api:"required"`
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
func (r TenantWebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookListResponseData struct {
	Webhooks []TenantWebhookListResponseDataWebhook `json:"webhooks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantWebhookListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookListResponseDataWebhook struct {
	// Webhook ID
	ID      string   `json:"id" api:"required"`
	Enabled bool     `json:"enabled" api:"required"`
	Events  []string `json:"events" api:"required"`
	Name    string   `json:"name" api:"required"`
	URL     string   `json:"url" api:"required" format:"uri"`
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
func (r TenantWebhookListResponseDataWebhook) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookListResponseDataWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookDeleteResponse struct {
	Data    TenantWebhookDeleteResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                  `json:"meta" api:"required"`
	Success bool                            `json:"success" api:"required"`
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
func (r TenantWebhookDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookDeleteResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantWebhookDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of webhook delivery attempts
type TenantWebhookListDeliveriesResponse struct {
	Data []TenantWebhookListDeliveriesResponseData `json:"data" api:"required"`
	Meta shared.APIMeta                            `json:"meta" api:"required"`
	// Current page number
	Page int64 `json:"page" api:"required"`
	// Items per page
	PerPage int64 `json:"perPage" api:"required"`
	// Total number of deliveries matching the filter
	Total int64 `json:"total" api:"required"`
	// Total number of pages
	TotalPages int64 `json:"totalPages" api:"required"`
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
func (r TenantWebhookListDeliveriesResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookListDeliveriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a webhook delivery attempt
type TenantWebhookListDeliveriesResponseData struct {
	// Unique delivery ID (UUID)
	ID string `json:"id" api:"required"`
	// Attempt number (1 for first attempt, increments with retries)
	Attempt int64 `json:"attempt" api:"required"`
	// Event type that triggered this delivery
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Event string `json:"event" api:"required"`
	// HTTP status code returned by the endpoint (null if connection failed)
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether the delivery was successful (2xx response)
	Success bool `json:"success" api:"required"`
	// When this delivery attempt occurred
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// URL the webhook was delivered to
	URL string `json:"url" api:"required" format:"uri"`
	// ID of the webhook this delivery belongs to
	WebhookID string `json:"webhookId" api:"required"`
	// Whether this delivery will be retried (true if failed and retries remaining)
	WillRetry bool `json:"willRetry" api:"required"`
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
func (r TenantWebhookListDeliveriesResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookListDeliveriesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of replaying a webhook delivery
type TenantWebhookReplayDeliveryResponse struct {
	Data    TenantWebhookReplayDeliveryResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                          `json:"meta" api:"required"`
	Success bool                                    `json:"success" api:"required"`
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
func (r TenantWebhookReplayDeliveryResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookReplayDeliveryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookReplayDeliveryResponseData struct {
	// Request duration in milliseconds
	Duration int64 `json:"duration" api:"required"`
	// ID of the new delivery created by the replay
	NewDeliveryID string `json:"newDeliveryId" api:"required"`
	// ID of the original delivery that was replayed
	OriginalDeliveryID string `json:"originalDeliveryId" api:"required"`
	// HTTP status code from your endpoint
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether the replay was successful (2xx response from endpoint)
	Success bool `json:"success" api:"required"`
	// When the replay was executed
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
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
func (r TenantWebhookReplayDeliveryResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookReplayDeliveryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed information about a webhook delivery attempt
type TenantWebhookGetDeliveryResponse struct {
	// Full details of a webhook delivery including request and response
	Data    TenantWebhookGetDeliveryResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                       `json:"meta" api:"required"`
	Success bool                                 `json:"success" api:"required"`
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
func (r TenantWebhookGetDeliveryResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookGetDeliveryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full details of a webhook delivery including request and response
type TenantWebhookGetDeliveryResponseData struct {
	// Unique delivery ID (UUID)
	ID string `json:"id" api:"required"`
	// Attempt number for this delivery
	Attempt int64 `json:"attempt" api:"required"`
	// Event type that triggered this delivery
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Event string `json:"event" api:"required"`
	// The request that was sent to your endpoint
	Request TenantWebhookGetDeliveryResponseDataRequest `json:"request" api:"required"`
	// The response received from your endpoint
	Response TenantWebhookGetDeliveryResponseDataResponse `json:"response" api:"required"`
	// HTTP status code returned by the endpoint
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether the delivery was successful (2xx response)
	Success bool `json:"success" api:"required"`
	// When this delivery attempt occurred
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// URL the webhook was delivered to
	URL string `json:"url" api:"required" format:"uri"`
	// ID of the webhook this delivery belongs to
	WebhookID string `json:"webhookId" api:"required"`
	// Name of the webhook for easy identification
	WebhookName string `json:"webhookName" api:"required"`
	// Whether this delivery will be retried
	WillRetry bool `json:"willRetry" api:"required"`
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
func (r TenantWebhookGetDeliveryResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookGetDeliveryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The request that was sent to your endpoint
type TenantWebhookGetDeliveryResponseDataRequest struct {
	// HTTP headers that were sent with the request
	Headers map[string]string `json:"headers" api:"required"`
	// The complete webhook payload that was sent
	Payload map[string]any `json:"payload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantWebhookGetDeliveryResponseDataRequest) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookGetDeliveryResponseDataRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response received from your endpoint
type TenantWebhookGetDeliveryResponseDataResponse struct {
	// HTTP status code from your endpoint
	StatusCode int64 `json:"statusCode" api:"required"`
	// Response body from your endpoint (may be truncated)
	Body string `json:"body" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StatusCode  respjson.Field
		Body        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantWebhookGetDeliveryResponseDataResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookGetDeliveryResponseDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookTestResponse struct {
	Data    TenantWebhookTestResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                `json:"meta" api:"required"`
	Success bool                          `json:"success" api:"required"`
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
func (r TenantWebhookTestResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookTestResponseData struct {
	// Request duration in milliseconds
	Duration int64 `json:"duration" api:"required"`
	// Event type that was tested
	Event string `json:"event" api:"required"`
	// HTTP status code from the webhook endpoint
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether the webhook endpoint responded with a 2xx status
	Success bool `json:"success" api:"required"`
	// Response body from the webhook endpoint (truncated if too long)
	Body string `json:"body" api:"nullable"`
	// Error message if the request failed
	Error string `json:"error" api:"nullable"`
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
func (r TenantWebhookTestResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantWebhookTestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookNewParams struct {
	// Webhook name for identification
	Name string `json:"name" api:"required"`
	// HTTPS endpoint URL
	URL string `json:"url" api:"required" format:"uri"`
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
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r TenantWebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantWebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantWebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookGetParams struct {
	TenantID string `path:"tenantId" api:"required" json:"-"`
	paramObj
}

type TenantWebhookUpdateParams struct {
	TenantID  string            `path:"tenantId" api:"required" json:"-"`
	AllEvents param.Opt[bool]   `json:"allEvents,omitzero"`
	Enabled   param.Opt[bool]   `json:"enabled,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	URL       param.Opt[string] `json:"url,omitzero" format:"uri"`
	Events    []string          `json:"events,omitzero"`
	paramObj
}

func (r TenantWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantWebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantWebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantWebhookDeleteParams struct {
	TenantID string `path:"tenantId" api:"required" json:"-"`
	paramObj
}

type TenantWebhookListDeliveriesParams struct {
	TenantID string `path:"tenantId" api:"required" json:"-"`
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
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Event TenantWebhookListDeliveriesParamsEvent `query:"event,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantWebhookListDeliveriesParams]'s query parameters as
// `url.Values`.
func (r TenantWebhookListDeliveriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by event type
type TenantWebhookListDeliveriesParamsEvent string

const (
	TenantWebhookListDeliveriesParamsEventMessageSent           TenantWebhookListDeliveriesParamsEvent = "MessageSent"
	TenantWebhookListDeliveriesParamsEventMessageDelayed        TenantWebhookListDeliveriesParamsEvent = "MessageDelayed"
	TenantWebhookListDeliveriesParamsEventMessageDeliveryFailed TenantWebhookListDeliveriesParamsEvent = "MessageDeliveryFailed"
	TenantWebhookListDeliveriesParamsEventMessageHeld           TenantWebhookListDeliveriesParamsEvent = "MessageHeld"
	TenantWebhookListDeliveriesParamsEventMessageBounced        TenantWebhookListDeliveriesParamsEvent = "MessageBounced"
	TenantWebhookListDeliveriesParamsEventMessageLinkClicked    TenantWebhookListDeliveriesParamsEvent = "MessageLinkClicked"
	TenantWebhookListDeliveriesParamsEventMessageLoaded         TenantWebhookListDeliveriesParamsEvent = "MessageLoaded"
	TenantWebhookListDeliveriesParamsEventDomainDNSError        TenantWebhookListDeliveriesParamsEvent = "DomainDNSError"
)

type TenantWebhookReplayDeliveryParams struct {
	TenantID  string `path:"tenantId" api:"required" json:"-"`
	WebhookID string `path:"webhookId" api:"required" json:"-"`
	paramObj
}

type TenantWebhookGetDeliveryParams struct {
	TenantID  string `path:"tenantId" api:"required" json:"-"`
	WebhookID string `path:"webhookId" api:"required" json:"-"`
	paramObj
}

type TenantWebhookTestParams struct {
	TenantID string `path:"tenantId" api:"required" json:"-"`
	// Event type to simulate
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Event TenantWebhookTestParamsEvent `json:"event,omitzero" api:"required"`
	paramObj
}

func (r TenantWebhookTestParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantWebhookTestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantWebhookTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event type to simulate
type TenantWebhookTestParamsEvent string

const (
	TenantWebhookTestParamsEventMessageSent           TenantWebhookTestParamsEvent = "MessageSent"
	TenantWebhookTestParamsEventMessageDelayed        TenantWebhookTestParamsEvent = "MessageDelayed"
	TenantWebhookTestParamsEventMessageDeliveryFailed TenantWebhookTestParamsEvent = "MessageDeliveryFailed"
	TenantWebhookTestParamsEventMessageHeld           TenantWebhookTestParamsEvent = "MessageHeld"
	TenantWebhookTestParamsEventMessageBounced        TenantWebhookTestParamsEvent = "MessageBounced"
	TenantWebhookTestParamsEventMessageLinkClicked    TenantWebhookTestParamsEvent = "MessageLinkClicked"
	TenantWebhookTestParamsEventMessageLoaded         TenantWebhookTestParamsEvent = "MessageLoaded"
	TenantWebhookTestParamsEventDomainDNSError        TenantWebhookTestParamsEvent = "DomainDNSError"
)
