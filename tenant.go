// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"context"
	"encoding/json"
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
//
// TenantService contains methods and other services that help with interacting
// with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	Options     []option.RequestOption
	Credentials TenantCredentialService
	// Manage sending domains.
	//
	// Before you can send emails, you need to:
	//
	// 1. Add a domain
	// 2. Configure DNS records (SPF, DKIM, Return Path)
	// 3. Verify the domain
	//
	// **Quick Reference:**
	//
	// - `POST /domains` - Add a new domain
	// - `GET /domains` - List all domains
	// - `POST /domains/{id}/verify` - Check DNS and verify domain
	// - `DELETE /domains/{id}` - Remove a domain
	Domains TenantDomainService
	// Manage the suppression list.
	//
	// Suppressed email addresses will not receive any emails. Addresses are
	// automatically suppressed when they hard bounce or file spam complaints.
	//
	// **Quick Reference:**
	//
	// - `GET /suppressions` - List suppressed addresses
	// - `POST /suppressions` - Add to suppression list
	// - `DELETE /suppressions/{email}` - Remove from suppression list
	// - `GET /suppressions/{email}` - Check if address is suppressed
	Suppressions TenantSuppressionService
	// Configure webhook endpoints for real-time notifications.
	//
	// Webhooks notify your application when email events occur:
	//
	// - Email delivered, bounced, or failed
	// - Email opened or link clicked
	// - Spam complaint received
	//
	// **Quick Reference:**
	//
	// - `POST /webhooks` - Create a webhook endpoint
	// - `GET /webhooks` - List all webhooks
	// - `POST /webhooks/{id}/test` - Test a webhook with sample data
	// - `PATCH /webhooks/{id}` - Update webhook configuration
	// - `DELETE /webhooks/{id}` - Remove a webhook
	// - `GET /webhooks/{id}/deliveries` - List delivery attempts
	// - `GET /webhooks/{id}/deliveries/{deliveryId}` - Get delivery details
	// - `POST /webhooks/{id}/deliveries/{deliveryId}/replay` - Replay a delivery
	//
	// ## Webhook Signatures
	//
	// All webhooks are cryptographically signed using RSA-SHA256 for security. Each
	// webhook request includes:
	//
	// | Header                | Description                                             |
	// | --------------------- | ------------------------------------------------------- |
	// | `X-Ark-Signature`     | Base64-encoded RSA-SHA256 signature of the request body |
	// | `X-Ark-Signature-KID` | Key ID identifying which public key was used            |
	//
	// Verify signatures by fetching the public key from:
	//
	// ```
	// GET https://mail.arkhq.io/.well-known/jwks.json
	// ```
	//
	// ```javascript
	// const crypto = require("crypto");
	//
	//	async function verifyWebhook(payload, signatureBase64, publicKey) {
	//	  const signature = Buffer.from(signatureBase64, "base64");
	//	  const verifier = crypto.createVerify("RSA-SHA256");
	//	  verifier.update(payload);
	//	  return verifier.verify(publicKey, signature);
	//	}
	//
	// // In your webhook handler:
	// const isValid = await verifyWebhook(
	//
	//	rawBody,
	//	req.headers["x-ark-signature"],
	//	cachedPublicKey
	//
	// );
	// ```
	//
	// **Important:** Always verify signatures before processing webhook data. See the
	// [Webhook Integration Guide](/guides/webhook-integration) for complete examples.
	Webhooks TenantWebhookService
	// Manage track domains for open and click tracking.
	//
	// Track domains enable you to track when recipients:
	//
	// - Open your emails (tracking pixel)
	// - Click links in your emails
	//
	// **Setup Process:**
	//
	// 1. Create a track domain with `POST /tracking`
	// 2. Add the CNAME record to your DNS
	// 3. Verify DNS with `POST /tracking/{id}/verify`
	// 4. Track domain is ready when `dnsOk` is true
	//
	// **Quick Reference:**
	//
	// - `POST /tracking` - Create a new track domain
	// - `GET /tracking` - List all track domains
	// - `GET /tracking/{id}` - Get track domain details
	// - `POST /tracking/{id}/verify` - Verify DNS configuration
	// - `PATCH /tracking/{id}` - Enable/disable tracking features
	// - `DELETE /tracking/{id}` - Remove a track domain
	Tracking TenantTrackingService
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
	Usage TenantUsageService
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r TenantService) {
	r = TenantService{}
	r.Options = opts
	r.Credentials = NewTenantCredentialService(opts...)
	r.Domains = NewTenantDomainService(opts...)
	r.Suppressions = NewTenantSuppressionService(opts...)
	r.Webhooks = NewTenantWebhookService(opts...)
	r.Tracking = NewTenantTrackingService(opts...)
	r.Usage = NewTenantUsageService(opts...)
	return
}

// Create a new tenant.
//
// Returns the created tenant with a unique `id`. Store this ID in your database to
// reference this tenant later.
func (r *TenantService) New(ctx context.Context, body TenantNewParams, opts ...option.RequestOption) (res *TenantNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tenants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a tenant by ID.
func (r *TenantService) Get(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *TenantGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a tenant's name, metadata, or status. At least one field is required.
//
// Metadata is replaced entirely—include all keys you want to keep.
func (r *TenantService) Update(ctx context.Context, tenantID string, body TenantUpdateParams, opts ...option.RequestOption) (res *TenantUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all tenants with pagination. Filter by `status` if needed.
func (r *TenantService) List(ctx context.Context, query TenantListParams, opts ...option.RequestOption) (res *pagination.PageNumberPagination[Tenant], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "tenants"
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

// List all tenants with pagination. Filter by `status` if needed.
func (r *TenantService) ListAutoPaging(ctx context.Context, query TenantListParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[Tenant] {
	return pagination.NewPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete a tenant. This cannot be undone.
func (r *TenantService) Delete(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *TenantDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Tenant struct {
	// Unique identifier for the tenant
	ID string `json:"id" api:"required"`
	// When the tenant was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Custom key-value pairs for storing additional data
	Metadata map[string]TenantMetadataUnion `json:"metadata" api:"required"`
	// Display name for the tenant
	Name string `json:"name" api:"required"`
	// Current status of the tenant:
	//
	// - `active` - Normal operation
	// - `suspended` - Temporarily disabled
	// - `archived` - Soft-deleted
	//
	// Any of "active", "suspended", "archived".
	Status TenantStatus `json:"status" api:"required"`
	// When the tenant was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Metadata    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tenant) RawJSON() string { return r.JSON.raw }
func (r *Tenant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TenantMetadataUnion contains all possible properties and values from [string],
// [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type TenantMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u TenantMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TenantMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TenantMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TenantMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *TenantMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the tenant:
//
// - `active` - Normal operation
// - `suspended` - Temporarily disabled
// - `archived` - Soft-deleted
type TenantStatus string

const (
	TenantStatusActive    TenantStatus = "active"
	TenantStatusSuspended TenantStatus = "suspended"
	TenantStatusArchived  TenantStatus = "archived"
)

type TenantNewResponse struct {
	Data    Tenant         `json:"data" api:"required"`
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
func (r TenantNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantGetResponse struct {
	Data    Tenant         `json:"data" api:"required"`
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
func (r TenantGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantUpdateResponse struct {
	Data    Tenant         `json:"data" api:"required"`
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
func (r TenantUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDeleteResponse struct {
	Data    TenantDeleteResponseData `json:"data" api:"required"`
	Meta    shared.APIMeta           `json:"meta" api:"required"`
	Success bool                     `json:"success" api:"required"`
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
func (r TenantDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDeleteResponseData struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantNewParams struct {
	// Display name for the tenant (e.g., your customer's company name)
	Name string `json:"name" api:"required"`
	// Custom key-value pairs. Useful for storing references to your internal systems.
	//
	// **Limits:**
	//
	// - Max 50 keys
	// - Key names max 40 characters
	// - String values max 500 characters
	// - Total size max 8KB
	Metadata map[string]TenantNewParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r TenantNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TenantNewParamsMetadataUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u TenantNewParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *TenantNewParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TenantNewParamsMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type TenantUpdateParams struct {
	// Display name for the tenant
	Name param.Opt[string] `json:"name,omitzero"`
	// Custom key-value pairs. Useful for storing references to your internal systems.
	//
	// **Limits:**
	//
	// - Max 50 keys
	// - Key names max 40 characters
	// - String values max 500 characters
	// - Total size max 8KB
	Metadata map[string]TenantUpdateParamsMetadataUnion `json:"metadata,omitzero"`
	// Tenant status
	//
	// Any of "active", "suspended", "archived".
	Status TenantUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r TenantUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TenantUpdateParamsMetadataUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u TenantUpdateParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *TenantUpdateParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TenantUpdateParamsMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Tenant status
type TenantUpdateParamsStatus string

const (
	TenantUpdateParamsStatusActive    TenantUpdateParamsStatus = "active"
	TenantUpdateParamsStatusSuspended TenantUpdateParamsStatus = "suspended"
	TenantUpdateParamsStatusArchived  TenantUpdateParamsStatus = "archived"
)

type TenantListParams struct {
	// Page number (1-indexed)
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page (max 100)
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	// Filter by tenant status
	//
	// Any of "active", "suspended", "archived".
	Status TenantListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantListParams]'s query parameters as `url.Values`.
func (r TenantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by tenant status
type TenantListParamsStatus string

const (
	TenantListParamsStatusActive    TenantListParamsStatus = "active"
	TenantListParamsStatusSuspended TenantListParamsStatus = "suspended"
	TenantListParamsStatusArchived  TenantListParamsStatus = "archived"
)
