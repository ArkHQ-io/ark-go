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

// SuppressionService contains methods and other services that help with
// interacting with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuppressionService] method instead.
type SuppressionService struct {
	Options []option.RequestOption
}

// NewSuppressionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuppressionService(opts ...option.RequestOption) (r SuppressionService) {
	r = SuppressionService{}
	r.Options = opts
	return
}

// Add an email address to the suppression list. The address will not receive any
// emails until removed.
func (r *SuppressionService) New(ctx context.Context, body SuppressionNewParams, opts ...option.RequestOption) (res *SuppressionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "suppressions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check if a specific email address is on the suppression list
func (r *SuppressionService) Get(ctx context.Context, email string, opts ...option.RequestOption) (res *SuppressionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if email == "" {
		err = errors.New("missing required email parameter")
		return
	}
	path := fmt.Sprintf("suppressions/%s", email)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all email addresses on the suppression list. These addresses will not
// receive any emails.
func (r *SuppressionService) List(ctx context.Context, query SuppressionListParams, opts ...option.RequestOption) (res *pagination.PageNumberPagination[SuppressionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "suppressions"
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

// Get all email addresses on the suppression list. These addresses will not
// receive any emails.
func (r *SuppressionService) ListAutoPaging(ctx context.Context, query SuppressionListParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[SuppressionListResponse] {
	return pagination.NewPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// Remove an email address from the suppression list. The address will be able to
// receive emails again.
func (r *SuppressionService) Delete(ctx context.Context, email string, opts ...option.RequestOption) (res *SuppressionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if email == "" {
		err = errors.New("missing required email parameter")
		return
	}
	path := fmt.Sprintf("suppressions/%s", email)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add up to 1000 email addresses to the suppression list at once
func (r *SuppressionService) BulkNew(ctx context.Context, body SuppressionBulkNewParams, opts ...option.RequestOption) (res *SuppressionBulkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "suppressions/bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SuppressionNewResponse struct {
	Data    SuppressionNewResponseData `json:"data,required"`
	Meta    shared.APIMeta             `json:"meta,required"`
	Success bool                       `json:"success,required"`
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
func (r SuppressionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SuppressionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionNewResponseData struct {
	// Suppression ID
	ID        string    `json:"id,required"`
	Address   string    `json:"address,required" format:"email"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Reason for suppression
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		CreatedAt   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuppressionNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *SuppressionNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionGetResponse struct {
	Data    SuppressionGetResponseData `json:"data,required"`
	Meta    shared.APIMeta             `json:"meta,required"`
	Success bool                       `json:"success,required"`
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
func (r SuppressionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SuppressionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionGetResponseData struct {
	// The email address that was checked
	Address string `json:"address,required" format:"email"`
	// Whether the address is currently suppressed
	Suppressed bool `json:"suppressed,required"`
	// When the suppression was created (if suppressed)
	CreatedAt time.Time `json:"createdAt,nullable" format:"date-time"`
	// Reason for suppression (if suppressed)
	Reason string `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Suppressed  respjson.Field
		CreatedAt   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuppressionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *SuppressionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionListResponse struct {
	// Suppression ID
	ID        string    `json:"id,required"`
	Address   string    `json:"address,required" format:"email"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	Reason    string    `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		CreatedAt   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuppressionListResponse) RawJSON() string { return r.JSON.raw }
func (r *SuppressionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionDeleteResponse struct {
	Data    SuppressionDeleteResponseData `json:"data,required"`
	Meta    shared.APIMeta                `json:"meta,required"`
	Success bool                          `json:"success,required"`
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
func (r SuppressionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SuppressionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionDeleteResponseData struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuppressionDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *SuppressionDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionBulkNewResponse struct {
	Data    SuppressionBulkNewResponseData `json:"data,required"`
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
func (r SuppressionBulkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SuppressionBulkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionBulkNewResponseData struct {
	// Newly suppressed addresses
	Added int64 `json:"added,required"`
	// Invalid addresses skipped
	Failed int64 `json:"failed,required"`
	// Total addresses in request
	TotalRequested int64 `json:"totalRequested,required"`
	// Already suppressed addresses (updated reason)
	Updated int64 `json:"updated,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added          respjson.Field
		Failed         respjson.Field
		TotalRequested respjson.Field
		Updated        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuppressionBulkNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *SuppressionBulkNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionNewParams struct {
	// Email address to suppress
	Address string `json:"address,required" format:"email"`
	// Reason for suppression
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r SuppressionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SuppressionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SuppressionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionListParams struct {
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SuppressionListParams]'s query parameters as `url.Values`.
func (r SuppressionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SuppressionBulkNewParams struct {
	Suppressions []SuppressionBulkNewParamsSuppression `json:"suppressions,omitzero,required"`
	paramObj
}

func (r SuppressionBulkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SuppressionBulkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SuppressionBulkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Address is required.
type SuppressionBulkNewParamsSuppression struct {
	Address string            `json:"address,required" format:"email"`
	Reason  param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r SuppressionBulkNewParamsSuppression) MarshalJSON() (data []byte, err error) {
	type shadow SuppressionBulkNewParamsSuppression
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SuppressionBulkNewParamsSuppression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
