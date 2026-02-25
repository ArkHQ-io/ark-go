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

// TenantUsageService contains methods and other services that help with
// interacting with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantUsageService] method instead.
type TenantUsageService struct {
	Options []option.RequestOption
}

// NewTenantUsageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTenantUsageService(opts ...option.RequestOption) (r TenantUsageService) {
	r = TenantUsageService{}
	r.Options = opts
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
func (r *TenantUsageService) Get(ctx context.Context, tenantID string, query TenantUsageGetParams, opts ...option.RequestOption) (res *TenantUsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/usage", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
func (r *TenantUsageService) GetTimeseries(ctx context.Context, tenantID string, query TenantUsageGetTimeseriesParams, opts ...option.RequestOption) (res *TenantUsageGetTimeseriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/usage/timeseries", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Tenant usage statistics
type TenantUsage struct {
	// Email delivery counts
	Emails EmailCounts `json:"emails" api:"required"`
	// Time period for usage data
	Period UsagePeriod `json:"period" api:"required"`
	// Email delivery rates (as decimals, e.g., 0.95 = 95%)
	Rates EmailRates `json:"rates" api:"required"`
	// Unique tenant identifier
	TenantID string `json:"tenant_id" api:"required"`
	// Tenant display name
	TenantName string `json:"tenant_name" api:"required"`
	// Your external ID for this tenant (from metadata)
	ExternalID string `json:"external_id" api:"nullable"`
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
	Data []TenantUsageTimeseriesData `json:"data" api:"required"`
	// Time bucket granularity
	//
	// Any of "hour", "day", "week", "month".
	Granularity TenantUsageTimeseriesGranularity `json:"granularity" api:"required"`
	// Time period for usage data
	Period UsagePeriod `json:"period" api:"required"`
	// Unique tenant identifier
	TenantID string `json:"tenant_id" api:"required"`
	// Tenant display name
	TenantName string `json:"tenant_name" api:"required"`
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
	Bounced int64 `json:"bounced" api:"required"`
	// Emails delivered in this bucket
	Delivered int64 `json:"delivered" api:"required"`
	// Hard failures in this bucket
	HardFailed int64 `json:"hard_failed" api:"required"`
	// Emails held in this bucket
	Held int64 `json:"held" api:"required"`
	// Emails sent in this bucket
	Sent int64 `json:"sent" api:"required"`
	// Soft failures in this bucket
	SoftFailed int64 `json:"soft_failed" api:"required"`
	// Start of time bucket
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
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

// Usage statistics for a single tenant
type TenantUsageGetResponse struct {
	// Tenant usage statistics
	Data    TenantUsage    `json:"data" api:"required"`
	Meta    shared.APIMeta `json:"meta" api:"required"`
	Success bool           `json:"success" api:"required"`
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
func (r TenantUsageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantUsageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Timeseries usage data for a tenant
type TenantUsageGetTimeseriesResponse struct {
	// Timeseries usage statistics
	Data    TenantUsageTimeseries `json:"data" api:"required"`
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
func (r TenantUsageGetTimeseriesResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantUsageGetTimeseriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantUsageGetParams struct {
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

// URLQuery serializes [TenantUsageGetParams]'s query parameters as `url.Values`.
func (r TenantUsageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantUsageGetTimeseriesParams struct {
	// Time period for timeseries data. Defaults to current month.
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Timezone for period calculations (IANA format). Defaults to UTC.
	Timezone param.Opt[string] `query:"timezone,omitzero" json:"-"`
	// Time bucket size for data points
	//
	// Any of "hour", "day", "week", "month".
	Granularity TenantUsageGetTimeseriesParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantUsageGetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r TenantUsageGetTimeseriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time bucket size for data points
type TenantUsageGetTimeseriesParamsGranularity string

const (
	TenantUsageGetTimeseriesParamsGranularityHour  TenantUsageGetTimeseriesParamsGranularity = "hour"
	TenantUsageGetTimeseriesParamsGranularityDay   TenantUsageGetTimeseriesParamsGranularity = "day"
	TenantUsageGetTimeseriesParamsGranularityWeek  TenantUsageGetTimeseriesParamsGranularity = "week"
	TenantUsageGetTimeseriesParamsGranularityMonth TenantUsageGetTimeseriesParamsGranularity = "month"
)
