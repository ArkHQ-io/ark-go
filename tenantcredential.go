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

// TenantCredentialService contains methods and other services that help with
// interacting with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantCredentialService] method instead.
type TenantCredentialService struct {
	Options []option.RequestOption
}

// NewTenantCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTenantCredentialService(opts ...option.RequestOption) (r TenantCredentialService) {
	r = TenantCredentialService{}
	r.Options = opts
	return
}

// Create a new SMTP or API credential for a tenant. The credential can be used to
// send emails via Ark on behalf of the tenant.
//
// **Important:** The credential key is only returned once at creation time. Store
// it securely - you cannot retrieve it again.
//
// **Credential Types:**
//
// - `smtp` - For SMTP-based email sending. Returns both `key` and `smtpUsername`.
// - `api` - For API-based email sending. Returns only `key`.
func (r *TenantCredentialService) New(ctx context.Context, tenantID string, body TenantCredentialNewParams, opts ...option.RequestOption) (res *TenantCredentialNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/credentials", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details of a specific credential.
//
// **Revealing the key:** By default, the credential key is not returned. Pass
// `reveal=true` to include the key in the response. Use this sparingly and only
// when you need to retrieve the key (e.g., for configuration).
func (r *TenantCredentialService) Get(ctx context.Context, credentialID int64, params TenantCredentialGetParams, opts ...option.RequestOption) (res *TenantCredentialGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/credentials/%v", params.TenantID, credentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update a credential's name or hold status.
//
// **Hold Status:**
//
//   - When `hold: true`, the credential is disabled and cannot be used to send
//     emails.
//   - When `hold: false`, the credential is active and can send emails.
//   - Use this to temporarily disable a credential without deleting it.
func (r *TenantCredentialService) Update(ctx context.Context, credentialID int64, params TenantCredentialUpdateParams, opts ...option.RequestOption) (res *TenantCredentialUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/credentials/%v", params.TenantID, credentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all SMTP and API credentials for a tenant. Credentials are used to send
// emails via Ark on behalf of the tenant.
//
// **Security:** Credential keys are not returned in the list response. Use the
// retrieve endpoint with `reveal=true` to get the key.
func (r *TenantCredentialService) List(ctx context.Context, tenantID string, query TenantCredentialListParams, opts ...option.RequestOption) (res *pagination.PageNumberPagination[TenantCredentialListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/credentials", tenantID)
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

// List all SMTP and API credentials for a tenant. Credentials are used to send
// emails via Ark on behalf of the tenant.
//
// **Security:** Credential keys are not returned in the list response. Use the
// retrieve endpoint with `reveal=true` to get the key.
func (r *TenantCredentialService) ListAutoPaging(ctx context.Context, tenantID string, query TenantCredentialListParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[TenantCredentialListResponse] {
	return pagination.NewPageNumberPaginationAutoPager(r.List(ctx, tenantID, query, opts...))
}

// Permanently delete (revoke) a credential. The credential can no longer be used
// to send emails.
//
// **Warning:** This action is irreversible. If you want to temporarily disable a
// credential, use the update endpoint to set `hold: true` instead.
func (r *TenantCredentialService) Delete(ctx context.Context, credentialID int64, body TenantCredentialDeleteParams, opts ...option.RequestOption) (res *TenantCredentialDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/credentials/%v", body.TenantID, credentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type TenantCredentialNewResponse struct {
	Data    TenantCredentialNewResponseData `json:"data,required"`
	Meta    shared.APIMeta                  `json:"meta,required"`
	Success bool                            `json:"success,required"`
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
func (r TenantCredentialNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialNewResponseData struct {
	// Unique identifier for the credential
	ID int64 `json:"id,required"`
	// When the credential was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Whether the credential is on hold (disabled). When `true`, the credential cannot
	// be used to send emails.
	Hold bool `json:"hold,required"`
	// The credential key (secret). **Store this securely** - it will not be shown
	// again unless you use the reveal parameter.
	Key string `json:"key,required"`
	// When the credential was last used to send an email
	LastUsedAt time.Time `json:"lastUsedAt,required" format:"date-time"`
	// Name of the credential
	Name string `json:"name,required"`
	// Type of credential:
	//
	// - `smtp` - For SMTP-based email sending
	// - `api` - For API-based email sending
	//
	// Any of "smtp", "api".
	Type string `json:"type,required"`
	// When the credential was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// SMTP username for authentication. Only included for SMTP credentials. Format:
	// `{tenantId}/{key}`
	SmtpUsername string `json:"smtpUsername"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Hold         respjson.Field
		Key          respjson.Field
		LastUsedAt   respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		SmtpUsername respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantCredentialNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialGetResponse struct {
	Data    TenantCredentialGetResponseData `json:"data,required"`
	Meta    shared.APIMeta                  `json:"meta,required"`
	Success bool                            `json:"success,required"`
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
func (r TenantCredentialGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialGetResponseData struct {
	// Unique identifier for the credential
	ID int64 `json:"id,required"`
	// When the credential was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Whether the credential is on hold (disabled). When `true`, the credential cannot
	// be used to send emails.
	Hold bool `json:"hold,required"`
	// When the credential was last used to send an email
	LastUsedAt time.Time `json:"lastUsedAt,required" format:"date-time"`
	// Name of the credential
	Name string `json:"name,required"`
	// Type of credential:
	//
	// - `smtp` - For SMTP-based email sending
	// - `api` - For API-based email sending
	//
	// Any of "smtp", "api".
	Type string `json:"type,required"`
	// When the credential was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The credential key (secret). Only included when:
	//
	// - Creating a new credential (always returned)
	// - Retrieving with `reveal=true`
	Key string `json:"key"`
	// SMTP username for authentication. Only included for SMTP credentials when the
	// key is revealed.
	SmtpUsername string `json:"smtpUsername"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Hold         respjson.Field
		LastUsedAt   respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		Key          respjson.Field
		SmtpUsername respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantCredentialGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialUpdateResponse struct {
	Data    TenantCredentialUpdateResponseData `json:"data,required"`
	Meta    shared.APIMeta                     `json:"meta,required"`
	Success bool                               `json:"success,required"`
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
func (r TenantCredentialUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialUpdateResponseData struct {
	// Unique identifier for the credential
	ID int64 `json:"id,required"`
	// When the credential was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Whether the credential is on hold (disabled). When `true`, the credential cannot
	// be used to send emails.
	Hold bool `json:"hold,required"`
	// When the credential was last used to send an email
	LastUsedAt time.Time `json:"lastUsedAt,required" format:"date-time"`
	// Name of the credential
	Name string `json:"name,required"`
	// Type of credential:
	//
	// - `smtp` - For SMTP-based email sending
	// - `api` - For API-based email sending
	//
	// Any of "smtp", "api".
	Type string `json:"type,required"`
	// When the credential was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The credential key (secret). Only included when:
	//
	// - Creating a new credential (always returned)
	// - Retrieving with `reveal=true`
	Key string `json:"key"`
	// SMTP username for authentication. Only included for SMTP credentials when the
	// key is revealed.
	SmtpUsername string `json:"smtpUsername"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Hold         respjson.Field
		LastUsedAt   respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		Key          respjson.Field
		SmtpUsername respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantCredentialUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialListResponse struct {
	// Unique identifier for the credential
	ID int64 `json:"id,required"`
	// When the credential was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Whether the credential is on hold (disabled). When `true`, the credential cannot
	// be used to send emails.
	Hold bool `json:"hold,required"`
	// When the credential was last used to send an email
	LastUsedAt time.Time `json:"lastUsedAt,required" format:"date-time"`
	// Name of the credential
	Name string `json:"name,required"`
	// Type of credential:
	//
	// - `smtp` - For SMTP-based email sending
	// - `api` - For API-based email sending
	//
	// Any of "smtp", "api".
	Type TenantCredentialListResponseType `json:"type,required"`
	// When the credential was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Hold        respjson.Field
		LastUsedAt  respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantCredentialListResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of credential:
//
// - `smtp` - For SMTP-based email sending
// - `api` - For API-based email sending
type TenantCredentialListResponseType string

const (
	TenantCredentialListResponseTypeSmtp TenantCredentialListResponseType = "smtp"
	TenantCredentialListResponseTypeAPI  TenantCredentialListResponseType = "api"
)

type TenantCredentialDeleteResponse struct {
	Data    TenantCredentialDeleteResponseData `json:"data,required"`
	Meta    shared.APIMeta                     `json:"meta,required"`
	Success bool                               `json:"success,required"`
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
func (r TenantCredentialDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialDeleteResponseData struct {
	Deleted bool `json:"deleted,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantCredentialDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantCredentialDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialNewParams struct {
	// Name for the credential. Can only contain letters, numbers, hyphens, and
	// underscores. Max 50 characters.
	Name string `json:"name,required"`
	// Type of credential:
	//
	// - `smtp` - For SMTP-based email sending
	// - `api` - For API-based email sending
	//
	// Any of "smtp", "api".
	Type TenantCredentialNewParamsType `json:"type,omitzero,required"`
	paramObj
}

func (r TenantCredentialNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantCredentialNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantCredentialNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of credential:
//
// - `smtp` - For SMTP-based email sending
// - `api` - For API-based email sending
type TenantCredentialNewParamsType string

const (
	TenantCredentialNewParamsTypeSmtp TenantCredentialNewParamsType = "smtp"
	TenantCredentialNewParamsTypeAPI  TenantCredentialNewParamsType = "api"
)

type TenantCredentialGetParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	// Set to `true` to include the credential key in the response
	Reveal param.Opt[bool] `query:"reveal,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantCredentialGetParams]'s query parameters as
// `url.Values`.
func (r TenantCredentialGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantCredentialUpdateParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	// Set to `true` to disable the credential (put on hold). Set to `false` to enable
	// the credential (release from hold).
	Hold param.Opt[bool] `json:"hold,omitzero"`
	// New name for the credential
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r TenantCredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantCredentialUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantCredentialUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantCredentialListParams struct {
	// Page number (1-indexed)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page (max 100)
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	// Filter by credential type
	//
	// Any of "smtp", "api".
	Type TenantCredentialListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantCredentialListParams]'s query parameters as
// `url.Values`.
func (r TenantCredentialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by credential type
type TenantCredentialListParamsType string

const (
	TenantCredentialListParamsTypeSmtp TenantCredentialListParamsType = "smtp"
	TenantCredentialListParamsTypeAPI  TenantCredentialListParamsType = "api"
)

type TenantCredentialDeleteParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	paramObj
}
