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
func (r *UsageService) Get(ctx context.Context, opts ...option.RequestOption) (res *UsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Account usage and limits response
type UsageGetResponse struct {
	// Current usage and limit information
	Data    UsageGetResponseData `json:"data,required"`
	Meta    shared.APIMeta       `json:"meta,required"`
	Success bool                 `json:"success,required"`
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

// Current usage and limit information
type UsageGetResponseData struct {
	// Billing and credit information
	Billing UsageGetResponseDataBilling `json:"billing,required"`
	// API rate limit status
	RateLimit UsageGetResponseDataRateLimit `json:"rateLimit,required"`
	// Email send limit status (hourly cap)
	SendLimit UsageGetResponseDataSendLimit `json:"sendLimit,required"`
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
func (r UsageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing and credit information
type UsageGetResponseDataBilling struct {
	// Auto-recharge configuration
	AutoRecharge UsageGetResponseDataBillingAutoRecharge `json:"autoRecharge,required"`
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
func (r UsageGetResponseDataBilling) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponseDataBilling) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Auto-recharge configuration
type UsageGetResponseDataBillingAutoRecharge struct {
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
func (r UsageGetResponseDataBillingAutoRecharge) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponseDataBillingAutoRecharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API rate limit status
type UsageGetResponseDataRateLimit struct {
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
func (r UsageGetResponseDataRateLimit) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponseDataRateLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email send limit status (hourly cap)
type UsageGetResponseDataSendLimit struct {
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
func (r UsageGetResponseDataSendLimit) RawJSON() string { return r.JSON.raw }
func (r *UsageGetResponseDataSendLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
