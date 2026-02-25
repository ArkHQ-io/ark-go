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

// PlatformWebhookService contains methods and other services that help with
// interacting with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlatformWebhookService] method instead.
type PlatformWebhookService struct {
	Options []option.RequestOption
}

// NewPlatformWebhookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlatformWebhookService(opts ...option.RequestOption) (r PlatformWebhookService) {
	r = PlatformWebhookService{}
	r.Options = opts
	return
}

// Create a platform webhook to receive email event notifications from all tenants.
//
// Platform webhooks receive events from **all tenants** in your organization. Each
// webhook payload includes a `tenant_id` field to identify which tenant the event
// belongs to.
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
//
// **Webhook payload includes:**
//
// - `event` - The event type
// - `tenant_id` - The tenant that sent the email
// - `timestamp` - Unix timestamp of the event
// - `payload` - Event-specific data (message details, status, etc.)
func (r *PlatformWebhookService) New(ctx context.Context, body PlatformWebhookNewParams, opts ...option.RequestOption) (res *PlatformWebhookNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "platform/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get detailed information about a specific platform webhook.
func (r *PlatformWebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *PlatformWebhookGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("platform/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a platform webhook's configuration.
//
// You can update:
//
// - `name` - Display name for the webhook
// - `url` - The endpoint URL (must be HTTPS)
// - `events` - Array of event types to receive (empty array = all events)
// - `enabled` - Enable or disable the webhook
func (r *PlatformWebhookService) Update(ctx context.Context, webhookID string, body PlatformWebhookUpdateParams, opts ...option.RequestOption) (res *PlatformWebhookUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("platform/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all platform webhook endpoints configured for your organization.
//
// Platform webhooks receive events from **all tenants** in your organization,
// unlike tenant webhooks which only receive events for a specific tenant. This is
// useful for centralized event processing and monitoring.
func (r *PlatformWebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *PlatformWebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "platform/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a platform webhook. This stops all event delivery to the webhook URL.
// This action cannot be undone.
func (r *PlatformWebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *PlatformWebhookDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("platform/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a paginated list of platform webhook delivery attempts.
//
// Filter by:
//
// - `webhookId` - Specific webhook
// - `tenantId` - Specific tenant
// - `event` - Specific event type
// - `success` - Successful (2xx) or failed deliveries
// - `before`/`after` - Time range (Unix timestamps)
//
// Deliveries are returned in reverse chronological order.
func (r *PlatformWebhookService) ListDeliveries(ctx context.Context, query PlatformWebhookListDeliveriesParams, opts ...option.RequestOption) (res *pagination.PageNumberPagination[PlatformWebhookListDeliveriesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "platform/webhooks/deliveries"
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

// Get a paginated list of platform webhook delivery attempts.
//
// Filter by:
//
// - `webhookId` - Specific webhook
// - `tenantId` - Specific tenant
// - `event` - Specific event type
// - `success` - Successful (2xx) or failed deliveries
// - `before`/`after` - Time range (Unix timestamps)
//
// Deliveries are returned in reverse chronological order.
func (r *PlatformWebhookService) ListDeliveriesAutoPaging(ctx context.Context, query PlatformWebhookListDeliveriesParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[PlatformWebhookListDeliveriesResponse] {
	return pagination.NewPageNumberPaginationAutoPager(r.ListDeliveries(ctx, query, opts...))
}

// Replay a previous platform webhook delivery.
//
// This re-sends the original payload with a new timestamp and delivery ID. Useful
// for recovering from temporary endpoint failures.
func (r *PlatformWebhookService) ReplayDelivery(ctx context.Context, deliveryID string, opts ...option.RequestOption) (res *PlatformWebhookReplayDeliveryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if deliveryID == "" {
		err = errors.New("missing required deliveryId parameter")
		return
	}
	path := fmt.Sprintf("platform/webhooks/deliveries/%s/replay", deliveryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get detailed information about a specific platform webhook delivery.
//
// Returns the complete request payload, headers, response, and timing info.
func (r *PlatformWebhookService) GetDelivery(ctx context.Context, deliveryID string, opts ...option.RequestOption) (res *PlatformWebhookGetDeliveryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if deliveryID == "" {
		err = errors.New("missing required deliveryId parameter")
		return
	}
	path := fmt.Sprintf("platform/webhooks/deliveries/%s", deliveryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Send a test payload to your platform webhook endpoint.
//
// Use this to:
//
// - Verify your webhook URL is accessible
// - Test your payload handling code
// - Ensure your server responds correctly
//
// The test payload is marked with `_test: true` so you can distinguish test events
// from real events.
func (r *PlatformWebhookService) Test(ctx context.Context, webhookID string, body PlatformWebhookTestParams, opts ...option.RequestOption) (res *PlatformWebhookTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return
	}
	path := fmt.Sprintf("platform/webhooks/%s/test", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PlatformWebhookNewResponse struct {
	Data    PlatformWebhookNewResponseData `json:"data" api:"required"`
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
func (r PlatformWebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookNewResponseData struct {
	// Platform webhook ID
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events (empty = all events)
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events" api:"required"`
	// Webhook name for identification
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Webhook endpoint URL
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Enabled     respjson.Field
		Events      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookGetResponse struct {
	Data    PlatformWebhookGetResponseData `json:"data" api:"required"`
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
func (r PlatformWebhookGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookGetResponseData struct {
	// Platform webhook ID
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events (empty = all events)
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events" api:"required"`
	// Webhook name for identification
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Webhook endpoint URL
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Enabled     respjson.Field
		Events      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookUpdateResponse struct {
	Data    PlatformWebhookUpdateResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                    `json:"meta" api:"required"`
	Success bool                              `json:"success" api:"required"`
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
func (r PlatformWebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookUpdateResponseData struct {
	// Platform webhook ID
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the webhook is active
	Enabled bool `json:"enabled" api:"required"`
	// Subscribed events (empty = all events)
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events" api:"required"`
	// Webhook name for identification
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Webhook endpoint URL
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Enabled     respjson.Field
		Events      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookListResponse struct {
	Data    []PlatformWebhookListResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                    `json:"meta" api:"required"`
	Success bool                              `json:"success" api:"required"`
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
func (r PlatformWebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookListResponseData struct {
	// Platform webhook ID
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Enabled   bool      `json:"enabled" api:"required"`
	Events    []string  `json:"events" api:"required"`
	Name      string    `json:"name" api:"required"`
	URL       string    `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Enabled     respjson.Field
		Events      respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookDeleteResponse struct {
	Data    PlatformWebhookDeleteResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                    `json:"meta" api:"required"`
	Success bool                              `json:"success" api:"required"`
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
func (r PlatformWebhookDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookDeleteResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Summary of a platform webhook delivery attempt
type PlatformWebhookListDeliveriesResponse struct {
	// Unique delivery ID
	ID string `json:"id" api:"required"`
	// Attempt number (1 for first attempt, higher for retries)
	Attempt int64 `json:"attempt" api:"required"`
	// Event type
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Event PlatformWebhookListDeliveriesResponseEvent `json:"event" api:"required"`
	// HTTP status code returned by your endpoint (null if connection failed)
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether delivery was successful (2xx response)
	Success bool `json:"success" api:"required"`
	// Tenant that triggered the event
	TenantID string `json:"tenantId" api:"required"`
	// When the delivery was attempted
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Endpoint URL the delivery was sent to
	URL string `json:"url" api:"required" format:"uri"`
	// Platform webhook ID
	WebhookID string `json:"webhookId" api:"required"`
	// Whether this delivery will be retried
	WillRetry bool `json:"willRetry" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attempt     respjson.Field
		Event       respjson.Field
		StatusCode  respjson.Field
		Success     respjson.Field
		TenantID    respjson.Field
		Timestamp   respjson.Field
		URL         respjson.Field
		WebhookID   respjson.Field
		WillRetry   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookListDeliveriesResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookListDeliveriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event type
type PlatformWebhookListDeliveriesResponseEvent string

const (
	PlatformWebhookListDeliveriesResponseEventMessageSent           PlatformWebhookListDeliveriesResponseEvent = "MessageSent"
	PlatformWebhookListDeliveriesResponseEventMessageDelayed        PlatformWebhookListDeliveriesResponseEvent = "MessageDelayed"
	PlatformWebhookListDeliveriesResponseEventMessageDeliveryFailed PlatformWebhookListDeliveriesResponseEvent = "MessageDeliveryFailed"
	PlatformWebhookListDeliveriesResponseEventMessageHeld           PlatformWebhookListDeliveriesResponseEvent = "MessageHeld"
	PlatformWebhookListDeliveriesResponseEventMessageBounced        PlatformWebhookListDeliveriesResponseEvent = "MessageBounced"
	PlatformWebhookListDeliveriesResponseEventMessageLinkClicked    PlatformWebhookListDeliveriesResponseEvent = "MessageLinkClicked"
	PlatformWebhookListDeliveriesResponseEventMessageLoaded         PlatformWebhookListDeliveriesResponseEvent = "MessageLoaded"
	PlatformWebhookListDeliveriesResponseEventDomainDNSError        PlatformWebhookListDeliveriesResponseEvent = "DomainDNSError"
)

// Result of replaying a platform webhook delivery
type PlatformWebhookReplayDeliveryResponse struct {
	Data    PlatformWebhookReplayDeliveryResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                            `json:"meta" api:"required"`
	Success bool                                      `json:"success" api:"required"`
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
func (r PlatformWebhookReplayDeliveryResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookReplayDeliveryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookReplayDeliveryResponseData struct {
	// Request duration in milliseconds
	Duration int64 `json:"duration" api:"required"`
	// ID of the new delivery created by the replay
	NewDeliveryID string `json:"newDeliveryId" api:"required"`
	// ID of the original delivery that was replayed
	OriginalDeliveryID string `json:"originalDeliveryId" api:"required"`
	// HTTP status code from your endpoint
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether the replay was successful
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
func (r PlatformWebhookReplayDeliveryResponseData) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookReplayDeliveryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookGetDeliveryResponse struct {
	Data    PlatformWebhookGetDeliveryResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                         `json:"meta" api:"required"`
	Success bool                                   `json:"success" api:"required"`
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
func (r PlatformWebhookGetDeliveryResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookGetDeliveryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookGetDeliveryResponseData struct {
	// Unique delivery ID
	ID string `json:"id" api:"required"`
	// Attempt number
	Attempt int64 `json:"attempt" api:"required"`
	// Event type
	Event string `json:"event" api:"required"`
	// Request details
	Request PlatformWebhookGetDeliveryResponseDataRequest `json:"request" api:"required"`
	// Response details
	Response PlatformWebhookGetDeliveryResponseDataResponse `json:"response" api:"required"`
	// HTTP status code from your endpoint
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether delivery was successful
	Success bool `json:"success" api:"required"`
	// Tenant that triggered the event
	TenantID string `json:"tenantId" api:"required"`
	// When delivery was attempted
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Endpoint URL
	URL string `json:"url" api:"required" format:"uri"`
	// Platform webhook ID
	WebhookID string `json:"webhookId" api:"required"`
	// Platform webhook name
	WebhookName string `json:"webhookName" api:"required"`
	// Whether this will be retried
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
		TenantID    respjson.Field
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
func (r PlatformWebhookGetDeliveryResponseData) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookGetDeliveryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request details
type PlatformWebhookGetDeliveryResponseDataRequest struct {
	// Request headers including signature
	Headers map[string]string `json:"headers"`
	// The complete webhook payload that was sent
	Payload map[string]any `json:"payload"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookGetDeliveryResponseDataRequest) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookGetDeliveryResponseDataRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response details
type PlatformWebhookGetDeliveryResponseDataResponse struct {
	// Response body (truncated if too large)
	Body string `json:"body" api:"nullable"`
	// Response time in milliseconds
	Duration int64 `json:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Duration    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookGetDeliveryResponseDataResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookGetDeliveryResponseDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookTestResponse struct {
	Data    PlatformWebhookTestResponseData `json:"data" api:"required"`
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
func (r PlatformWebhookTestResponse) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookTestResponseData struct {
	// Request duration in milliseconds
	DurationMs int64 `json:"durationMs" api:"required"`
	// HTTP status code from the webhook endpoint
	StatusCode int64 `json:"statusCode" api:"required"`
	// Whether the webhook endpoint responded with a 2xx status
	Success bool `json:"success" api:"required"`
	// Error message if the request failed
	Error string `json:"error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs  respjson.Field
		StatusCode  respjson.Field
		Success     respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformWebhookTestResponseData) RawJSON() string { return r.JSON.raw }
func (r *PlatformWebhookTestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookNewParams struct {
	// Display name for the webhook
	Name string `json:"name" api:"required"`
	// Webhook endpoint URL (must be HTTPS)
	URL string `json:"url" api:"required" format:"uri"`
	// Events to subscribe to. Empty array means all events.
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r PlatformWebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PlatformWebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlatformWebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookUpdateParams struct {
	// Enable or disable the webhook
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Display name for the webhook
	Name param.Opt[string] `json:"name,omitzero"`
	// Webhook endpoint URL (must be HTTPS)
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Events to subscribe to. Empty array means all events.
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r PlatformWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PlatformWebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlatformWebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformWebhookListDeliveriesParams struct {
	// Only deliveries after this Unix timestamp
	After param.Opt[int64] `query:"after,omitzero" json:"-"`
	// Only deliveries before this Unix timestamp
	Before param.Opt[int64] `query:"before,omitzero" json:"-"`
	// Page number (default 1)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Items per page (default 30, max 100)
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	// Filter by delivery success
	Success param.Opt[bool] `query:"success,omitzero" json:"-"`
	// Filter by tenant ID
	TenantID param.Opt[string] `query:"tenantId,omitzero" json:"-"`
	// Filter by platform webhook ID
	WebhookID param.Opt[string] `query:"webhookId,omitzero" json:"-"`
	// Filter by event type
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Event PlatformWebhookListDeliveriesParamsEvent `query:"event,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlatformWebhookListDeliveriesParams]'s query parameters as
// `url.Values`.
func (r PlatformWebhookListDeliveriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by event type
type PlatformWebhookListDeliveriesParamsEvent string

const (
	PlatformWebhookListDeliveriesParamsEventMessageSent           PlatformWebhookListDeliveriesParamsEvent = "MessageSent"
	PlatformWebhookListDeliveriesParamsEventMessageDelayed        PlatformWebhookListDeliveriesParamsEvent = "MessageDelayed"
	PlatformWebhookListDeliveriesParamsEventMessageDeliveryFailed PlatformWebhookListDeliveriesParamsEvent = "MessageDeliveryFailed"
	PlatformWebhookListDeliveriesParamsEventMessageHeld           PlatformWebhookListDeliveriesParamsEvent = "MessageHeld"
	PlatformWebhookListDeliveriesParamsEventMessageBounced        PlatformWebhookListDeliveriesParamsEvent = "MessageBounced"
	PlatformWebhookListDeliveriesParamsEventMessageLinkClicked    PlatformWebhookListDeliveriesParamsEvent = "MessageLinkClicked"
	PlatformWebhookListDeliveriesParamsEventMessageLoaded         PlatformWebhookListDeliveriesParamsEvent = "MessageLoaded"
	PlatformWebhookListDeliveriesParamsEventDomainDNSError        PlatformWebhookListDeliveriesParamsEvent = "DomainDNSError"
)

type PlatformWebhookTestParams struct {
	// Event type to simulate
	//
	// Any of "MessageSent", "MessageDelayed", "MessageDeliveryFailed", "MessageHeld",
	// "MessageBounced", "MessageLinkClicked", "MessageLoaded", "DomainDNSError".
	Event PlatformWebhookTestParamsEvent `json:"event,omitzero" api:"required"`
	paramObj
}

func (r PlatformWebhookTestParams) MarshalJSON() (data []byte, err error) {
	type shadow PlatformWebhookTestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlatformWebhookTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event type to simulate
type PlatformWebhookTestParamsEvent string

const (
	PlatformWebhookTestParamsEventMessageSent           PlatformWebhookTestParamsEvent = "MessageSent"
	PlatformWebhookTestParamsEventMessageDelayed        PlatformWebhookTestParamsEvent = "MessageDelayed"
	PlatformWebhookTestParamsEventMessageDeliveryFailed PlatformWebhookTestParamsEvent = "MessageDeliveryFailed"
	PlatformWebhookTestParamsEventMessageHeld           PlatformWebhookTestParamsEvent = "MessageHeld"
	PlatformWebhookTestParamsEventMessageBounced        PlatformWebhookTestParamsEvent = "MessageBounced"
	PlatformWebhookTestParamsEventMessageLinkClicked    PlatformWebhookTestParamsEvent = "MessageLinkClicked"
	PlatformWebhookTestParamsEventMessageLoaded         PlatformWebhookTestParamsEvent = "MessageLoaded"
	PlatformWebhookTestParamsEventDomainDNSError        PlatformWebhookTestParamsEvent = "DomainDNSError"
)
