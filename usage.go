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

// UsageService contains methods and other services that help with interacting with
// the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r UsageService) {
	r = UsageService{}
	r.Options = opts
	return
}

// > **Deprecated:** Use `GET /limits` instead for rate limits and send limits.
// > This endpoint will be removed in a future version.
//
// Returns current usage and limit information for your account.
//
// This endpoint is designed for:
//
// - **AI agents/MCP servers:** Check constraints before planning batch operations
// - **Monitoring dashboards:** Display current usage status
// - **Rate limit awareness:** Know remaining capacity before making requests
//
// **Response includes:**
//
// - `rateLimit` - API request rate limit (requests per second)
// - `sendLimit` - Email sending limit (emails per hour)
// - `billing` - Credit balance and auto-recharge configuration
//
// **Notes:**
//
// - This request counts against your rate limit
// - `sendLimit` may be null if Postal is temporarily unavailable
// - `billing` is null if billing is not configured
// - Send limit resets at the top of each hour
//
// **Migration:**
//
// - For rate limits and send limits, use `GET /limits`
// - For per-tenant usage analytics, use `GET /tenants/{tenantId}/usage`
// - For bulk tenant usage, use `GET /usage/by-tenant`
//
// Deprecated: deprecated
func (r *UsageService) Get(ctx context.Context, opts ...option.RequestOption) (res *UsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Export usage data for all tenants in a format suitable for billing systems.
//
// **Use cases:**
//
// - Import into billing systems (Stripe, Chargebee, etc.)
// - Generate invoices
// - Archive usage data
//
// **Export formats:**
//
// - `csv` - Comma-separated values (default)
// - `jsonl` - JSON Lines (one JSON object per line)
// - `json` - JSON array
//
// **Response headers:**
//
// - `X-Total-Tenants` - Total number of tenants in export
// - `X-Total-Sent` - Total emails sent across all tenants
// - `Content-Disposition` - Suggested filename for download
//
// This endpoint returns up to 10,000 tenants per request. For organizations with
// more tenants, use the `/usage/by-tenant` endpoint with pagination.
func (r *UsageService) Export(ctx context.Context, query UsageExportParams, opts ...option.RequestOption) (res *[]UsageExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "usage/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns email usage statistics for all tenants in your organization.
//
// **Use cases:**
//
// - Generate monthly billing reports
// - Build admin dashboards showing all customer usage
// - Identify high-volume or problematic tenants
//
// **Sorting options:**
//
// - `sent`, `-sent` - Sort by emails sent (ascending/descending)
// - `delivered`, `-delivered` - Sort by emails delivered
// - `bounce_rate`, `-bounce_rate` - Sort by bounce rate
// - `name`, `-name` - Sort alphabetically by tenant name
//
// **Filtering:**
//
// - `status` - Filter by tenant status (active, suspended, archived)
// - `min_sent` - Only include tenants with at least N emails sent
//
// Results are paginated. Use `limit` and `offset` for pagination.
func (r *UsageService) ListByTenant(ctx context.Context, query UsageListByTenantParams, opts ...option.RequestOption) (res *pagination.OffsetPagination[BulkTenantUsageTenant], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "usage/by-tenant"
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

// Returns email usage statistics for all tenants in your organization.
//
// **Use cases:**
//
// - Generate monthly billing reports
// - Build admin dashboards showing all customer usage
// - Identify high-volume or problematic tenants
//
// **Sorting options:**
//
// - `sent`, `-sent` - Sort by emails sent (ascending/descending)
// - `delivered`, `-delivered` - Sort by emails delivered
// - `bounce_rate`, `-bounce_rate` - Sort by bounce rate
// - `name`, `-name` - Sort alphabetically by tenant name
//
// **Filtering:**
//
// - `status` - Filter by tenant status (active, suspended, archived)
// - `min_sent` - Only include tenants with at least N emails sent
//
// Results are paginated. Use `limit` and `offset` for pagination.
func (r *UsageService) ListByTenantAutoPaging(ctx context.Context, query UsageListByTenantParams, opts ...option.RequestOption) *pagination.OffsetPaginationAutoPager[BulkTenantUsageTenant] {
	return pagination.NewOffsetPaginationAutoPager(r.ListByTenant(ctx, query, opts...))
}

// Returns time-bucketed email statistics for a specific tenant.
//
// **Use cases:**
//
// - Build usage charts and graphs
// - Identify sending patterns
// - Detect anomalies in delivery rates
//
// **Granularity options:**
//
// - `hour` - Hourly buckets (best for last 7 days)
// - `day` - Daily buckets (best for last 30-90 days)
// - `week` - Weekly buckets (best for last 6 months)
// - `month` - Monthly buckets (best for year-over-year)
//
// The response includes a data point for each time bucket with all email metrics.
func (r *UsageService) GetTenantTimeseries(ctx context.Context, tenantID string, query UsageGetTenantTimeseriesParams, opts ...option.RequestOption) (res *UsageGetTenantTimeseriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/usage/timeseries", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns email sending statistics for a specific tenant over a time period.
//
// **Use cases:**
//
// - Display usage dashboard to your customers
// - Calculate per-tenant billing
// - Monitor tenant health and delivery rates
//
// **Period formats:**
//
//   - Shortcuts: `today`, `yesterday`, `this_week`, `last_week`, `this_month`,
//     `last_month`, `last_7_days`, `last_30_days`, `last_90_days`
//   - Month: `2024-01` (full month)
//   - Date range: `2024-01-01..2024-01-31`
//   - Single day: `2024-01-15`
//
// **Response includes:**
//
// - `emails` - Counts for sent, delivered, soft_failed, hard_failed, bounced, held
// - `rates` - Delivery rate and bounce rate as decimals (0.95 = 95%)
func (r *UsageService) GetTenantUsage(ctx context.Context, tenantID string, query UsageGetTenantUsageParams, opts ...option.RequestOption) (res *UsageGetTenantUsageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/usage", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Bulk tenant usage data with pagination
type BulkTenantUsage struct {
	// Pagination information for usage queries
	Pagination BulkTenantUsagePagination `json:"pagination,required"`
	// Time period for usage data
	Period UsagePeriod `json:"period,required"`
	// Aggregate summary across all tenants
	Summary BulkTenantUsageSummary `json:"summary,required"`
	// Array of tenant usage records
	Tenants []BulkTenantUsageTenant `json:"tenants,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Period      respjson.Field
		Summary     respjson.Field
		Tenants     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkTenantUsage) RawJSON() string { return r.JSON.raw }
func (r *BulkTenantUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information for usage queries
type BulkTenantUsagePagination struct {
	// Whether more pages are available
	HasMore bool `json:"has_more,required"`
	// Maximum items per page
	Limit int64 `json:"limit,required"`
	// Number of items skipped
	Offset int64 `json:"offset,required"`
	// Total number of tenants matching the query
	Total int64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkTenantUsagePagination) RawJSON() string { return r.JSON.raw }
func (r *BulkTenantUsagePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregate summary across all tenants
type BulkTenantUsageSummary struct {
	// Total emails delivered across all tenants
	TotalDelivered int64 `json:"total_delivered,required"`
	// Total emails sent across all tenants
	TotalSent int64 `json:"total_sent,required"`
	// Total number of tenants in the query
	TotalTenants int64 `json:"total_tenants,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalDelivered respjson.Field
		TotalSent      respjson.Field
		TotalTenants   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkTenantUsageSummary) RawJSON() string { return r.JSON.raw }
func (r *BulkTenantUsageSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage record for a single tenant in bulk response
type BulkTenantUsageTenant struct {
	// Email delivery counts
	Emails EmailCounts `json:"emails,required"`
	// Email delivery rates (as decimals, e.g., 0.95 = 95%)
	Rates EmailRates `json:"rates,required"`
	// Current tenant status
	//
	// Any of "active", "suspended", "archived".
	Status string `json:"status,required"`
	// Unique tenant identifier
	TenantID string `json:"tenant_id,required"`
	// Tenant display name
	TenantName string `json:"tenant_name,required"`
	// Your external ID for this tenant
	ExternalID string `json:"external_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Emails      respjson.Field
		Rates       respjson.Field
		Status      respjson.Field
		TenantID    respjson.Field
		TenantName  respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkTenantUsageTenant) RawJSON() string { return r.JSON.raw }
func (r *BulkTenantUsageTenant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email delivery counts
type EmailCounts struct {
	// Emails that bounced
	Bounced int64 `json:"bounced,required"`
	// Emails successfully delivered
	Delivered int64 `json:"delivered,required"`
	// Emails that hard-failed (permanent failures)
	HardFailed int64 `json:"hard_failed,required"`
	// Emails currently held for review
	Held int64 `json:"held,required"`
	// Total emails sent
	Sent int64 `json:"sent,required"`
	// Emails that soft-failed (temporary failures, may be retried)
	SoftFailed int64 `json:"soft_failed,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced     respjson.Field
		Delivered   respjson.Field
		HardFailed  respjson.Field
		Held        respjson.Field
		Sent        respjson.Field
		SoftFailed  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailCounts) RawJSON() string { return r.JSON.raw }
func (r *EmailCounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email delivery rates (as decimals, e.g., 0.95 = 95%)
type EmailRates struct {
	// Percentage of sent emails that bounced (0-1)
	BounceRate float64 `json:"bounce_rate,required"`
	// Percentage of sent emails that were delivered (0-1)
	DeliveryRate float64 `json:"delivery_rate,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BounceRate   respjson.Field
		DeliveryRate respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailRates) RawJSON() string { return r.JSON.raw }
func (r *EmailRates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tenant usage statistics
type TenantUsage struct {
	// Email delivery counts
	Emails EmailCounts `json:"emails,required"`
	// Time period for usage data
	Period UsagePeriod `json:"period,required"`
	// Email delivery rates (as decimals, e.g., 0.95 = 95%)
	Rates EmailRates `json:"rates,required"`
	// Unique tenant identifier
	TenantID string `json:"tenant_id,required"`
	// Tenant display name
	TenantName string `json:"tenant_name,required"`
	// Your external ID for this tenant (from metadata)
	ExternalID string `json:"external_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Emails      respjson.Field
		Period      respjson.Field
		Rates       respjson.Field
		TenantID    respjson.Field
		TenantName  respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantUsage) RawJSON() string { return r.JSON.raw }
func (r *TenantUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Timeseries usage statistics
type TenantUsageTimeseries struct {
	// Array of time-bucketed data points
	Data []TenantUsageTimeseriesData `json:"data,required"`
	// Time bucket granularity
	//
	// Any of "hour", "day", "week", "month".
	Granularity TenantUsageTimeseriesGranularity `json:"granularity,required"`
	// Time period for usage data
	Period UsagePeriod `json:"period,required"`
	// Unique tenant identifier
	TenantID string `json:"tenant_id,required"`
	// Tenant display name
	TenantName string `json:"tenant_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Granularity respjson.Field
		Period      respjson.Field
		TenantID    respjson.Field
		TenantName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantUsageTimeseries) RawJSON() string { return r.JSON.raw }
func (r *TenantUsageTimeseries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single timeseries data point
type TenantUsageTimeseriesData struct {
	// Bounces in this bucket
	Bounced int64 `json:"bounced,required"`
	// Emails delivered in this bucket
	Delivered int64 `json:"delivered,required"`
	// Hard failures in this bucket
	HardFailed int64 `json:"hard_failed,required"`
	// Emails held in this bucket
	Held int64 `json:"held,required"`
	// Emails sent in this bucket
	Sent int64 `json:"sent,required"`
	// Soft failures in this bucket
	SoftFailed int64 `json:"soft_failed,required"`
	// Start of time bucket
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced     respjson.Field
		Delivered   respjson.Field
		HardFailed  respjson.Field
		Held        respjson.Field
		Sent        respjson.Field
		SoftFailed  respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantUsageTimeseriesData) RawJSON() string { return r.JSON.raw }
func (r *TenantUsageTimeseriesData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time bucket granularity
type TenantUsageTimeseriesGranularity string

const (
	TenantUsageTimeseriesGranularityHour  TenantUsageTimeseriesGranularity = "hour"
	TenantUsageTimeseriesGranularityDay   TenantUsageTimeseriesGranularity = "day"
	TenantUsageTimeseriesGranularityWeek  TenantUsageTimeseriesGranularity = "week"
	TenantUsageTimeseriesGranularityMonth TenantUsageTimeseriesGranularity = "month"
)

// Time period for usage data
type UsagePeriod struct {
	// Period end (inclusive)
	End time.Time `json:"end,required" format:"date-time"`
	// Period start (inclusive)
	Start time.Time `json:"start,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsagePeriod) RawJSON() string { return r.JSON.raw }
func (r *UsagePeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account usage and limits response
type UsageGetResponse struct {
	// Current usage and limit information
	Data    LimitsData     `json:"data,required"`
	Meta    shared.APIMeta `json:"meta,required"`
	Success bool           `json:"success,required"`
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
func (r UsageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single row in usage export (JSON format)
type UsageExportResponse struct {
	// Bounce rate (0-1)
	BounceRate float64 `json:"bounce_rate,required"`
	// Emails that bounced
	Bounced int64 `json:"bounced,required"`
	// Emails successfully delivered
	Delivered int64 `json:"delivered,required"`
	// Delivery rate (0-1)
	DeliveryRate float64 `json:"delivery_rate,required"`
	// Emails that hard-failed
	HardFailed int64 `json:"hard_failed,required"`
	// Emails currently held
	Held int64 `json:"held,required"`
	// Total emails sent
	Sent int64 `json:"sent,required"`
	// Emails that soft-failed
	SoftFailed int64 `json:"soft_failed,required"`
	// Current tenant status
	//
	// Any of "active", "suspended", "archived".
	Status UsageExportResponseStatus `json:"status,required"`
	// Unique tenant identifier
	TenantID string `json:"tenant_id,required"`
	// Tenant display name
	TenantName string `json:"tenant_name,required"`
	// Your external ID for this tenant
	ExternalID string `json:"external_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BounceRate   respjson.Field
		Bounced      respjson.Field
		Delivered    respjson.Field
		DeliveryRate respjson.Field
		HardFailed   respjson.Field
		Held         respjson.Field
		Sent         respjson.Field
		SoftFailed   respjson.Field
		Status       respjson.Field
		TenantID     respjson.Field
		TenantName   respjson.Field
		ExternalID   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageExportResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageExportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current tenant status
type UsageExportResponseStatus string

const (
	UsageExportResponseStatusActive    UsageExportResponseStatus = "active"
	UsageExportResponseStatusSuspended UsageExportResponseStatus = "suspended"
	UsageExportResponseStatusArchived  UsageExportResponseStatus = "archived"
)

// Timeseries usage data for a tenant
type UsageGetTenantTimeseriesResponse struct {
	// Timeseries usage statistics
	Data    TenantUsageTimeseries `json:"data,required"`
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
func (r UsageGetTenantTimeseriesResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageGetTenantTimeseriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage statistics for a single tenant
type UsageGetTenantUsageResponse struct {
	// Tenant usage statistics
	Data    TenantUsage    `json:"data,required"`
	Meta    shared.APIMeta `json:"meta,required"`
	Success bool           `json:"success,required"`
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
func (r UsageGetTenantUsageResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageGetTenantUsageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageExportParams struct {
	// Only include tenants with at least this many emails sent
	MinSent param.Opt[int64] `query:"min_sent,omitzero" json:"-"`
	// Time period for export. Defaults to current month.
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Timezone for period calculations (IANA format). Defaults to UTC.
	Timezone param.Opt[string] `query:"timezone,omitzero" json:"-"`
	// Export format
	//
	// Any of "csv", "jsonl", "json".
	Format UsageExportParamsFormat `query:"format,omitzero" json:"-"`
	// Filter by tenant status
	//
	// Any of "active", "suspended", "archived".
	Status UsageExportParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageExportParams]'s query parameters as `url.Values`.
func (r UsageExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Export format
type UsageExportParamsFormat string

const (
	UsageExportParamsFormatCsv   UsageExportParamsFormat = "csv"
	UsageExportParamsFormatJSONL UsageExportParamsFormat = "jsonl"
	UsageExportParamsFormatJson  UsageExportParamsFormat = "json"
)

// Filter by tenant status
type UsageExportParamsStatus string

const (
	UsageExportParamsStatusActive    UsageExportParamsStatus = "active"
	UsageExportParamsStatusSuspended UsageExportParamsStatus = "suspended"
	UsageExportParamsStatusArchived  UsageExportParamsStatus = "archived"
)

type UsageListByTenantParams struct {
	// Maximum number of tenants to return (1-100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only include tenants with at least this many emails sent
	MinSent param.Opt[int64] `query:"min_sent,omitzero" json:"-"`
	// Number of tenants to skip for pagination
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Time period for usage data. Defaults to current month.
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Timezone for period calculations (IANA format). Defaults to UTC.
	Timezone param.Opt[string] `query:"timezone,omitzero" json:"-"`
	// Sort order for results. Prefix with `-` for descending order.
	//
	// Any of "sent", "-sent", "delivered", "-delivered", "bounce_rate",
	// "-bounce_rate", "name", "-name".
	Sort UsageListByTenantParamsSort `query:"sort,omitzero" json:"-"`
	// Filter by tenant status
	//
	// Any of "active", "suspended", "archived".
	Status UsageListByTenantParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageListByTenantParams]'s query parameters as
// `url.Values`.
func (r UsageListByTenantParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order for results. Prefix with `-` for descending order.
type UsageListByTenantParamsSort string

const (
	UsageListByTenantParamsSortSent            UsageListByTenantParamsSort = "sent"
	UsageListByTenantParamsSortMinusSent       UsageListByTenantParamsSort = "-sent"
	UsageListByTenantParamsSortDelivered       UsageListByTenantParamsSort = "delivered"
	UsageListByTenantParamsSortMinusDelivered  UsageListByTenantParamsSort = "-delivered"
	UsageListByTenantParamsSortBounceRate      UsageListByTenantParamsSort = "bounce_rate"
	UsageListByTenantParamsSortMinusBounceRate UsageListByTenantParamsSort = "-bounce_rate"
	UsageListByTenantParamsSortName            UsageListByTenantParamsSort = "name"
	UsageListByTenantParamsSortMinusName       UsageListByTenantParamsSort = "-name"
)

// Filter by tenant status
type UsageListByTenantParamsStatus string

const (
	UsageListByTenantParamsStatusActive    UsageListByTenantParamsStatus = "active"
	UsageListByTenantParamsStatusSuspended UsageListByTenantParamsStatus = "suspended"
	UsageListByTenantParamsStatusArchived  UsageListByTenantParamsStatus = "archived"
)

type UsageGetTenantTimeseriesParams struct {
	// Time period for timeseries data. Defaults to current month.
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Timezone for period calculations (IANA format). Defaults to UTC.
	Timezone param.Opt[string] `query:"timezone,omitzero" json:"-"`
	// Time bucket size for data points
	//
	// Any of "hour", "day", "week", "month".
	Granularity UsageGetTenantTimeseriesParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageGetTenantTimeseriesParams]'s query parameters as
// `url.Values`.
func (r UsageGetTenantTimeseriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time bucket size for data points
type UsageGetTenantTimeseriesParamsGranularity string

const (
	UsageGetTenantTimeseriesParamsGranularityHour  UsageGetTenantTimeseriesParamsGranularity = "hour"
	UsageGetTenantTimeseriesParamsGranularityDay   UsageGetTenantTimeseriesParamsGranularity = "day"
	UsageGetTenantTimeseriesParamsGranularityWeek  UsageGetTenantTimeseriesParamsGranularity = "week"
	UsageGetTenantTimeseriesParamsGranularityMonth UsageGetTenantTimeseriesParamsGranularity = "month"
)

type UsageGetTenantUsageParams struct {
	// Time period for usage data. Defaults to current month.
	//
	// **Formats:**
	//
	//   - Shortcuts: `today`, `yesterday`, `this_week`, `last_week`, `this_month`,
	//     `last_month`, `last_7_days`, `last_30_days`, `last_90_days`
	//   - Month: `2024-01`
	//   - Range: `2024-01-01..2024-01-31`
	//   - Day: `2024-01-15`
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Timezone for period calculations (IANA format). Defaults to UTC.
	Timezone param.Opt[string] `query:"timezone,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageGetTenantUsageParams]'s query parameters as
// `url.Values`.
func (r UsageGetTenantUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
