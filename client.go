// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the ark API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	// Send and manage email messages.
	//
	// **Quick Reference:**
	//
	// - `POST /emails` - Send a single email
	// - `POST /emails/batch` - Send up to 100 emails
	// - `GET /emails/{emailId}` - Get email status and details
	// - `GET /emails` - List sent emails
	// - `POST /emails/{emailId}/retry` - Retry failed delivery
	Emails EmailService
	// Access API request logs for debugging and monitoring.
	//
	// Every API request is logged with details including:
	//
	// - Request method, path, and endpoint
	// - Response status code and duration
	// - Error details (code, message) for failed requests
	// - SDK information (name, version)
	// - Rate limit state at time of request
	// - Request and response bodies (for single log retrieval)
	//
	// **Retention:** Logs are retained for 90 days.
	//
	// **Body storage:** Request and response bodies are stored encrypted and truncated
	// at 25KB. Bodies are only returned when retrieving a single log entry.
	//
	// **Quick Reference:**
	//
	// - `GET /logs` - List API request logs with filters
	// - `GET /logs/{requestId}` - Get full details including request/response bodies
	Logs LogService
	// Per-tenant usage analytics and bulk reporting.
	//
	// Track email sending statistics for each tenant to power billing, dashboards, and
	// monitoring.
	//
	// **Single Tenant Usage:**
	//
	// - `GET /tenants/{id}/usage` - Get usage stats for a specific tenant
	// - `GET /tenants/{id}/usage/timeseries` - Get time-bucketed data for charts
	//
	// **Bulk Usage:**
	//
	// - `GET /usage/tenants` - Get usage for all tenants (paginated, sortable)
	// - `GET /usage/export` - Export usage data as CSV, JSONL, or JSON
	//
	// **Period Formats:**
	//
	//   - Shortcuts: `today`, `yesterday`, `this_month`, `last_month`, `last_7_days`,
	//     `last_30_days`
	//   - Month: `2024-01`
	//   - Date range: `2024-01-01..2024-01-15`
	Usage UsageService
	// Check account rate limits and send limits.
	//
	// The limits endpoint returns current status for operational limits:
	//
	// - **Rate limit:** API requests per second (currently 10/sec)
	// - **Send limit:** Emails per hour (default 100/hour for new accounts)
	// - **Billing:** Credit balance and auto-recharge configuration
	//
	// **AI Integration Note:** This endpoint is designed for AI agents and MCP servers
	// to understand account constraints before taking actions. Call this endpoint
	// first when planning batch operations to avoid hitting limits unexpectedly.
	//
	// **Quick Reference:**
	//
	// - `GET /limits` - Get current rate limits and send limits
	// - `GET /usage` - (Deprecated) Use `/limits` instead
	Limits LimitService
	// Manage tenants (your customers).
	//
	// Create a tenant for each of your customers to track their email sending
	// separately. Store the tenant `id` in your database and use `metadata` for any
	// custom data.
	//
	// **Quick Reference:**
	//
	// - `POST /tenants` - Create a new tenant
	// - `GET /tenants` - List all tenants (paginated)
	// - `GET /tenants/{id}` - Get tenant details
	// - `PATCH /tenants/{id}` - Update tenant name, metadata, or status
	// - `DELETE /tenants/{id}` - Delete a tenant
	Tenants  TenantService
	Platform PlatformService
}

// DefaultClientOptions read from the environment (ARK_API_KEY, ARK_BASE_URL). This
// should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("ARK_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("ARK_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (ARK_API_KEY, ARK_BASE_URL). The option passed in as arguments are
// applied after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Emails = NewEmailService(opts...)
	r.Logs = NewLogService(opts...)
	r.Usage = NewUsageService(opts...)
	r.Limits = NewLimitService(opts...)
	r.Tenants = NewTenantService(opts...)
	r.Platform = NewPlatformService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
