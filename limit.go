// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/ArkHQ-io/ark-go/internal/apijson"
	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/option"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
	"github.com/ArkHQ-io/ark-go/shared"
)

// LimitService contains methods and other services that help with interacting with
// the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLimitService] method instead.
type LimitService struct {
	Options []option.RequestOption
}

// NewLimitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLimitService(opts ...option.RequestOption) (r LimitService) {
	r = LimitService{}
	r.Options = opts
	return
}

// Returns current rate limit and send limit information for your account.
//
// This endpoint is the recommended way to check your account's operational limits.
// Use `/usage` endpoints for historical usage analytics.
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
// - `sendLimit` may be null if the service is temporarily unavailable
// - `billing` is null if billing is not configured
// - Send limit resets at the top of each hour
func (r *LimitService) Get(ctx context.Context, opts ...option.RequestOption) (res *LimitGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Current usage and limit information
type LimitsData struct {
	// Billing and credit information
	Billing LimitsDataBilling `json:"billing,required"`
	// API rate limit status
	RateLimit LimitsDataRateLimit `json:"rateLimit,required"`
	// Email send limit status (hourly cap)
	SendLimit LimitsDataSendLimit `json:"sendLimit,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing     respjson.Field
		RateLimit   respjson.Field
		SendLimit   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LimitsData) RawJSON() string { return r.JSON.raw }
func (r *LimitsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing and credit information
type LimitsDataBilling struct {
	// Auto-recharge configuration
	AutoRecharge LimitsDataBillingAutoRecharge `json:"autoRecharge,required"`
	// Current credit balance as formatted string (e.g., "25.50")
	CreditBalance string `json:"creditBalance,required"`
	// Current credit balance in cents for precise calculations
	CreditBalanceCents int64 `json:"creditBalanceCents,required"`
	// Whether a payment method is configured
	HasPaymentMethod bool `json:"hasPaymentMethod,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoRecharge       respjson.Field
		CreditBalance      respjson.Field
		CreditBalanceCents respjson.Field
		HasPaymentMethod   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LimitsDataBilling) RawJSON() string { return r.JSON.raw }
func (r *LimitsDataBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Auto-recharge configuration
type LimitsDataBillingAutoRecharge struct {
	// Amount to recharge when triggered
	Amount string `json:"amount,required"`
	// Whether auto-recharge is enabled
	Enabled bool `json:"enabled,required"`
	// Balance threshold that triggers recharge
	Threshold string `json:"threshold,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Enabled     respjson.Field
		Threshold   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LimitsDataBillingAutoRecharge) RawJSON() string { return r.JSON.raw }
func (r *LimitsDataBillingAutoRecharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API rate limit status
type LimitsDataRateLimit struct {
	// Maximum requests allowed per period
	Limit int64 `json:"limit,required"`
	// Time period for the limit
	//
	// Any of "second".
	Period string `json:"period,required"`
	// Requests remaining in current window
	Remaining int64 `json:"remaining,required"`
	// Unix timestamp when the limit resets
	Reset int64 `json:"reset,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Period      respjson.Field
		Remaining   respjson.Field
		Reset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LimitsDataRateLimit) RawJSON() string { return r.JSON.raw }
func (r *LimitsDataRateLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email send limit status (hourly cap)
type LimitsDataSendLimit struct {
	// Whether approaching the limit (>90%)
	Approaching bool `json:"approaching,required"`
	// Whether the limit has been exceeded
	Exceeded bool `json:"exceeded,required"`
	// Maximum emails allowed per hour (null = unlimited)
	Limit int64 `json:"limit,required"`
	// Time period for the limit
	//
	// Any of "hour".
	Period string `json:"period,required"`
	// Emails remaining in current period (null if unlimited)
	Remaining int64 `json:"remaining,required"`
	// ISO timestamp when the limit window resets (top of next hour)
	ResetsAt time.Time `json:"resetsAt,required" format:"date-time"`
	// Usage as a percentage (null if unlimited)
	UsagePercent float64 `json:"usagePercent,required"`
	// Emails sent in current period
	Used int64 `json:"used,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Approaching  respjson.Field
		Exceeded     respjson.Field
		Limit        respjson.Field
		Period       respjson.Field
		Remaining    respjson.Field
		ResetsAt     respjson.Field
		UsagePercent respjson.Field
		Used         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LimitsDataSendLimit) RawJSON() string { return r.JSON.raw }
func (r *LimitsDataSendLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account rate limits and send limits response
type LimitGetResponse struct {
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
func (r LimitGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LimitGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
