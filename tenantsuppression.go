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

// TenantSuppressionService contains methods and other services that help with
// interacting with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantSuppressionService] method instead.
type TenantSuppressionService struct {
	Options []option.RequestOption
}

// NewTenantSuppressionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTenantSuppressionService(opts ...option.RequestOption) (r TenantSuppressionService) {
	r = TenantSuppressionService{}
	r.Options = opts
	return
}

// Add an email address to the tenant's suppression list. The address will not
// receive any emails from this tenant until removed.
func (r *TenantSuppressionService) New(ctx context.Context, tenantID string, body TenantSuppressionNewParams, opts ...option.RequestOption) (res *TenantSuppressionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/suppressions", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check if a specific email address is on the tenant's suppression list.
func (r *TenantSuppressionService) Get(ctx context.Context, email string, query TenantSuppressionGetParams, opts ...option.RequestOption) (res *TenantSuppressionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if email == "" {
		err = errors.New("missing required email parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/suppressions/%s", query.TenantID, email)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all email addresses on the tenant's suppression list. These addresses will
// not receive any emails from this tenant.
func (r *TenantSuppressionService) List(ctx context.Context, tenantID string, query TenantSuppressionListParams, opts ...option.RequestOption) (res *pagination.PageNumberPagination[TenantSuppressionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/suppressions", tenantID)
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

// Get all email addresses on the tenant's suppression list. These addresses will
// not receive any emails from this tenant.
func (r *TenantSuppressionService) ListAutoPaging(ctx context.Context, tenantID string, query TenantSuppressionListParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[TenantSuppressionListResponse] {
	return pagination.NewPageNumberPaginationAutoPager(r.List(ctx, tenantID, query, opts...))
}

// Remove an email address from the tenant's suppression list. The address will be
// able to receive emails from this tenant again.
func (r *TenantSuppressionService) Delete(ctx context.Context, email string, body TenantSuppressionDeleteParams, opts ...option.RequestOption) (res *TenantSuppressionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if email == "" {
		err = errors.New("missing required email parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/suppressions/%s", body.TenantID, email)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type TenantSuppressionNewResponse struct {
	Data    TenantSuppressionNewResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                   `json:"meta" api:"required"`
	Success bool                             `json:"success" api:"required"`
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
func (r TenantSuppressionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantSuppressionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantSuppressionNewResponseData struct {
	// Suppression ID
	ID        string    `json:"id" api:"required"`
	Address   string    `json:"address" api:"required" format:"email"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
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
func (r TenantSuppressionNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantSuppressionNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantSuppressionGetResponse struct {
	Data    TenantSuppressionGetResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                   `json:"meta" api:"required"`
	Success bool                             `json:"success" api:"required"`
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
func (r TenantSuppressionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantSuppressionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantSuppressionGetResponseData struct {
	// The email address that was checked
	Address string `json:"address" api:"required" format:"email"`
	// Whether the address is currently suppressed
	Suppressed bool `json:"suppressed" api:"required"`
	// When the suppression was created (if suppressed)
	CreatedAt time.Time `json:"createdAt" api:"nullable" format:"date-time"`
	// Reason for suppression (if suppressed)
	Reason string `json:"reason" api:"nullable"`
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
func (r TenantSuppressionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantSuppressionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantSuppressionListResponse struct {
	// Suppression ID
	ID        string    `json:"id" api:"required"`
	Address   string    `json:"address" api:"required" format:"email"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
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
func (r TenantSuppressionListResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantSuppressionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantSuppressionDeleteResponse struct {
	Data    TenantSuppressionDeleteResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta                      `json:"meta" api:"required"`
	Success bool                                `json:"success" api:"required"`
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
func (r TenantSuppressionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantSuppressionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantSuppressionDeleteResponseData struct {
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantSuppressionDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantSuppressionDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantSuppressionNewParams struct {
	// Email address to suppress
	Address string `json:"address" api:"required" format:"email"`
	// Reason for suppression (accepts null)
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r TenantSuppressionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantSuppressionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantSuppressionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantSuppressionGetParams struct {
	TenantID string `path:"tenantId" api:"required" json:"-"`
	paramObj
}

type TenantSuppressionListParams struct {
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantSuppressionListParams]'s query parameters as
// `url.Values`.
func (r TenantSuppressionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantSuppressionDeleteParams struct {
	TenantID string `path:"tenantId" api:"required" json:"-"`
	paramObj
}
