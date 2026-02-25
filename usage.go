// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"context"
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

// Returns aggregated email sending statistics for your entire organization. For
// per-tenant breakdown, use `GET /usage/tenants`.
//
// **Use cases:**
//
// - Platform dashboards showing org-wide metrics
// - Quick health check on overall sending
// - Monitoring total volume and delivery rates
//
// **Response includes:**
//
// - `emails` - Aggregated email counts across all tenants
// - `rates` - Overall delivery and bounce rates
// - `tenants` - Tenant count summary (total, active, with activity)
//
// **Related endpoints:**
//
// - `GET /usage/tenants` - Paginated usage per tenant
// - `GET /usage/export` - Export usage data for billing
// - `GET /tenants/{tenantId}/usage` - Single tenant usage details
// - `GET /limits` - Rate limits and send limits
func (r *UsageService) Get(ctx context.Context, query UsageGetParams, opts ...option.RequestOption) (res *OrgUsageSummary, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Export email usage data for all tenants in CSV or JSON Lines format. Designed
// for billing system integration, data warehousing, and analytics.
//
// **Jobs to be done:**
//
// - Import usage data into billing systems (Stripe, Chargebee, etc.)
// - Load into data warehouses (Snowflake, BigQuery, etc.)
// - Process in spreadsheets (Excel, Google Sheets)
// - Feed into BI tools (Looker, Metabase, etc.)
//
// **Export formats:**
//
// - `csv` - UTF-8 with BOM for Excel compatibility (default)
// - `jsonl` - JSON Lines (one JSON object per line, streamable)
//
// **CSV columns:** `tenant_id`, `tenant_name`, `external_id`, `status`, `sent`,
// `delivered`, `soft_failed`, `hard_failed`, `bounced`, `held`, `delivery_rate`,
// `bounce_rate`, `period_start`, `period_end`
//
// **Response headers:**
//
// - `Content-Disposition` - Filename for download
// - `Content-Type` - `text/csv` or `application/x-ndjson`
func (r *UsageService) Export(ctx context.Context, query UsageExportParams, opts ...option.RequestOption) (res *[]UsageExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "usage/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns email usage statistics for all tenants in your organization. Results are
// paginated with page-based navigation.
//
// **Jobs to be done:**
//
// - Generate monthly billing invoices per tenant
// - Build admin dashboards showing all customer usage
// - Identify high-volume or problematic tenants
// - Track usage against plan limits
//
// **Sorting options:**
//
// - `sent`, `-sent` - Sort by emails sent (ascending/descending)
// - `delivered`, `-delivered` - Sort by emails delivered
// - `bounce_rate`, `-bounce_rate` - Sort by bounce rate
// - `tenant_name`, `-tenant_name` - Sort alphabetically by tenant name
//
// **Filtering:**
//
// - `status` - Filter by tenant status (active, suspended, archived)
// - `minSent` - Only include tenants with at least N emails sent
//
// **Auto-pagination:** SDKs support iterating over all pages automatically.
func (r *UsageService) ListTenants(ctx context.Context, query UsageListTenantsParams, opts ...option.RequestOption) (res *pagination.PageNumberPagination[TenantUsageItem], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "usage/tenants"
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

// Returns email usage statistics for all tenants in your organization. Results are
// paginated with page-based navigation.
//
// **Jobs to be done:**
//
// - Generate monthly billing invoices per tenant
// - Build admin dashboards showing all customer usage
// - Identify high-volume or problematic tenants
// - Track usage against plan limits
//
// **Sorting options:**
//
// - `sent`, `-sent` - Sort by emails sent (ascending/descending)
// - `delivered`, `-delivered` - Sort by emails delivered
// - `bounce_rate`, `-bounce_rate` - Sort by bounce rate
// - `tenant_name`, `-tenant_name` - Sort alphabetically by tenant name
//
// **Filtering:**
//
// - `status` - Filter by tenant status (active, suspended, archived)
// - `minSent` - Only include tenants with at least N emails sent
//
// **Auto-pagination:** SDKs support iterating over all pages automatically.
func (r *UsageService) ListTenantsAutoPaging(ctx context.Context, query UsageListTenantsParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[TenantUsageItem] {
	return pagination.NewPageNumberPaginationAutoPager(r.ListTenants(ctx, query, opts...))
}

// Email delivery counts
type EmailCounts struct {
	// Emails that bounced
	Bounced int64 `json:"bounced" api:"required"`
	// Emails successfully delivered
	Delivered int64 `json:"delivered" api:"required"`
	// Emails that hard-failed (permanent failures)
	HardFailed int64 `json:"hard_failed" api:"required"`
	// Emails currently held for review
	Held int64 `json:"held" api:"required"`
	// Total emails sent
	Sent int64 `json:"sent" api:"required"`
	// Emails that soft-failed (temporary failures, may be retried)
	SoftFailed int64 `json:"soft_failed" api:"required"`
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
	BounceRate float64 `json:"bounce_rate" api:"required"`
	// Percentage of sent emails that were delivered (0-1)
	DeliveryRate float64 `json:"delivery_rate" api:"required"`
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

// Org-wide usage summary response
type OrgUsageSummary struct {
	Data    OrgUsageSummaryData `json:"data" api:"required"`
	Meta    shared.APIMeta      `json:"meta" api:"required"`
	Success bool                `json:"success" api:"required"`
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
func (r OrgUsageSummary) RawJSON() string { return r.JSON.raw }
func (r *OrgUsageSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrgUsageSummaryData struct {
	// Email delivery counts
	Emails EmailCounts `json:"emails" api:"required"`
	// Time period for usage data
	Period UsagePeriod `json:"period" api:"required"`
	// Email delivery rates (as decimals, e.g., 0.95 = 95%)
	Rates   EmailRates                 `json:"rates" api:"required"`
	Tenants OrgUsageSummaryDataTenants `json:"tenants" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Emails      respjson.Field
		Period      respjson.Field
		Rates       respjson.Field
		Tenants     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrgUsageSummaryData) RawJSON() string { return r.JSON.raw }
func (r *OrgUsageSummaryData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrgUsageSummaryDataTenants struct {
	// Number of active tenants
	Active int64 `json:"active" api:"required"`
	// Total number of tenants
	Total int64 `json:"total" api:"required"`
	// Number of tenants with sending activity
	WithActivity int64 `json:"withActivity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active       respjson.Field
		Total        respjson.Field
		WithActivity respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrgUsageSummaryDataTenants) RawJSON() string { return r.JSON.raw }
func (r *OrgUsageSummaryDataTenants) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage record for a single tenant (camelCase for SDK)
type TenantUsageItem struct {
	// Email delivery counts
	Emails EmailCounts `json:"emails" api:"required"`
	// Email delivery rates (as decimals, e.g., 0.95 = 95%)
	Rates EmailRates `json:"rates" api:"required"`
	// Current tenant status
	//
	// Any of "active", "suspended", "archived".
	Status TenantUsageItemStatus `json:"status" api:"required"`
	// Unique tenant identifier
	TenantID string `json:"tenantId" api:"required"`
	// Tenant display name
	TenantName string `json:"tenantName" api:"required"`
	// Your external ID for this tenant
	ExternalID string `json:"externalId" api:"nullable"`
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
func (r TenantUsageItem) RawJSON() string { return r.JSON.raw }
func (r *TenantUsageItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current tenant status
type TenantUsageItemStatus string

const (
	TenantUsageItemStatusActive    TenantUsageItemStatus = "active"
	TenantUsageItemStatusSuspended TenantUsageItemStatus = "suspended"
	TenantUsageItemStatusArchived  TenantUsageItemStatus = "archived"
)

// Time period for usage data
type UsagePeriod struct {
	// Period end (inclusive)
	End time.Time `json:"end" api:"required" format:"date-time"`
	// Period start (inclusive)
	Start time.Time `json:"start" api:"required" format:"date-time"`
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

// Single row in usage export (JSON format)
type UsageExportResponse struct {
	// Bounce rate (0-1)
	BounceRate float64 `json:"bounce_rate" api:"required"`
	// Emails that bounced
	Bounced int64 `json:"bounced" api:"required"`
	// Emails successfully delivered
	Delivered int64 `json:"delivered" api:"required"`
	// Delivery rate (0-1)
	DeliveryRate float64 `json:"delivery_rate" api:"required"`
	// Emails that hard-failed
	HardFailed int64 `json:"hard_failed" api:"required"`
	// Emails currently held
	Held int64 `json:"held" api:"required"`
	// Total emails sent
	Sent int64 `json:"sent" api:"required"`
	// Emails that soft-failed
	SoftFailed int64 `json:"soft_failed" api:"required"`
	// Current tenant status
	//
	// Any of "active", "suspended", "archived".
	Status UsageExportResponseStatus `json:"status" api:"required"`
	// Unique tenant identifier
	TenantID string `json:"tenant_id" api:"required"`
	// Tenant display name
	TenantName string `json:"tenant_name" api:"required"`
	// Your external ID for this tenant
	ExternalID string `json:"external_id" api:"nullable"`
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

type UsageGetParams struct {
	// Time period for usage data.
	//
	// **Shortcuts:** `today`, `yesterday`, `this_week`, `last_week`, `this_month`,
	// `last_month`, `last_7_days`, `last_30_days`, `last_90_days`
	//
	// **Month format:** `2024-01` (YYYY-MM)
	//
	// **Custom range:** `2024-01-01..2024-01-15`
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Timezone for period calculations (IANA format)
	Timezone param.Opt[string] `query:"timezone,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageGetParams]'s query parameters as `url.Values`.
func (r UsageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UsageExportParams struct {
	// Only include tenants with at least this many emails sent
	MinSent param.Opt[int64] `query:"minSent,omitzero" json:"-"`
	// Time period for export.
	//
	// **Shortcuts:** `this_month`, `last_month`, `last_30_days`, etc.
	//
	// **Month format:** `2024-01` (YYYY-MM)
	//
	// **Custom range:** `2024-01-01..2024-01-15`
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Timezone for period calculations (IANA format)
	Timezone param.Opt[string] `query:"timezone,omitzero" json:"-"`
	// Export format
	//
	// Any of "csv", "jsonl".
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
)

// Filter by tenant status
type UsageExportParamsStatus string

const (
	UsageExportParamsStatusActive    UsageExportParamsStatus = "active"
	UsageExportParamsStatusSuspended UsageExportParamsStatus = "suspended"
	UsageExportParamsStatusArchived  UsageExportParamsStatus = "archived"
)

type UsageListTenantsParams struct {
	// Only include tenants with at least this many emails sent
	MinSent param.Opt[int64] `query:"minSent,omitzero" json:"-"`
	// Page number (1-indexed)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Time period for usage data. Defaults to current month.
	//
	// **Shortcuts:** `today`, `yesterday`, `this_week`, `last_week`, `this_month`,
	// `last_month`, `last_7_days`, `last_30_days`, `last_90_days`
	//
	// **Month format:** `2024-01` (YYYY-MM)
	//
	// **Custom range:** `2024-01-01..2024-01-15`
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Results per page (max 100)
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	// Timezone for period calculations (IANA format). Defaults to UTC.
	Timezone param.Opt[string] `query:"timezone,omitzero" json:"-"`
	// Sort order for results. Prefix with `-` for descending order.
	//
	// Any of "sent", "-sent", "delivered", "-delivered", "bounce_rate",
	// "-bounce_rate", "delivery_rate", "-delivery_rate", "tenant_name",
	// "-tenant_name".
	Sort UsageListTenantsParamsSort `query:"sort,omitzero" json:"-"`
	// Filter by tenant status
	//
	// Any of "active", "suspended", "archived".
	Status UsageListTenantsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageListTenantsParams]'s query parameters as `url.Values`.
func (r UsageListTenantsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order for results. Prefix with `-` for descending order.
type UsageListTenantsParamsSort string

const (
	UsageListTenantsParamsSortSent              UsageListTenantsParamsSort = "sent"
	UsageListTenantsParamsSortMinusSent         UsageListTenantsParamsSort = "-sent"
	UsageListTenantsParamsSortDelivered         UsageListTenantsParamsSort = "delivered"
	UsageListTenantsParamsSortMinusDelivered    UsageListTenantsParamsSort = "-delivered"
	UsageListTenantsParamsSortBounceRate        UsageListTenantsParamsSort = "bounce_rate"
	UsageListTenantsParamsSortMinusBounceRate   UsageListTenantsParamsSort = "-bounce_rate"
	UsageListTenantsParamsSortDeliveryRate      UsageListTenantsParamsSort = "delivery_rate"
	UsageListTenantsParamsSortMinusDeliveryRate UsageListTenantsParamsSort = "-delivery_rate"
	UsageListTenantsParamsSortTenantName        UsageListTenantsParamsSort = "tenant_name"
	UsageListTenantsParamsSortMinusTenantName   UsageListTenantsParamsSort = "-tenant_name"
)

// Filter by tenant status
type UsageListTenantsParamsStatus string

const (
	UsageListTenantsParamsStatusActive    UsageListTenantsParamsStatus = "active"
	UsageListTenantsParamsStatusSuspended UsageListTenantsParamsStatus = "suspended"
	UsageListTenantsParamsStatusArchived  UsageListTenantsParamsStatus = "archived"
)
