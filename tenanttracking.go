// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ark

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/ArkHQ-io/ark-go/internal/apijson"
	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/option"
	"github.com/ArkHQ-io/ark-go/packages/param"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
	"github.com/ArkHQ-io/ark-go/shared"
)

// TenantTrackingService contains methods and other services that help with
// interacting with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantTrackingService] method instead.
type TenantTrackingService struct {
	Options []option.RequestOption
}

// NewTenantTrackingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTenantTrackingService(opts ...option.RequestOption) (r TenantTrackingService) {
	r = TenantTrackingService{}
	r.Options = opts
	return
}

// Create a new track domain for open/click tracking for a tenant.
//
// After creation, you must configure a CNAME record pointing to the provided DNS
// value before tracking will work.
func (r *TenantTrackingService) New(ctx context.Context, tenantID string, body TenantTrackingNewParams, opts ...option.RequestOption) (res *TenantTrackingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/tracking", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details of a specific track domain including DNS configuration.
func (r *TenantTrackingService) Get(ctx context.Context, trackingID string, query TenantTrackingGetParams, opts ...option.RequestOption) (res *TenantTrackingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/tracking/%s", query.TenantID, trackingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update track domain settings.
//
// Use this to:
//
// - Enable/disable click tracking
// - Enable/disable open tracking
// - Enable/disable SSL
// - Set excluded click domains
func (r *TenantTrackingService) Update(ctx context.Context, trackingID string, params TenantTrackingUpdateParams, opts ...option.RequestOption) (res *TenantTrackingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/tracking/%s", params.TenantID, trackingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all track domains configured for a tenant. Track domains enable open and
// click tracking for emails.
func (r *TenantTrackingService) List(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *TenantTrackingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/tracking", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a track domain. This will disable tracking for any emails using this
// domain.
func (r *TenantTrackingService) Delete(ctx context.Context, trackingID string, body TenantTrackingDeleteParams, opts ...option.RequestOption) (res *TenantTrackingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/tracking/%s", body.TenantID, trackingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Check DNS configuration for the track domain.
//
// The track domain requires a CNAME record to be configured before open and click
// tracking will work. Use this endpoint to verify the DNS is correctly set up.
func (r *TenantTrackingService) Verify(ctx context.Context, trackingID string, body TenantTrackingVerifyParams, opts ...option.RequestOption) (res *TenantTrackingVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/tracking/%s/verify", body.TenantID, trackingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type TrackDomain struct {
	// Track domain ID
	ID string `json:"id,required"`
	// When the track domain was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Whether the tracking CNAME record is correctly configured. Must be true to use
	// tracking features.
	DNSOk bool `json:"dnsOk,required"`
	// ID of the parent sending domain
	DomainID string `json:"domainId,required"`
	// Full domain name
	FullName string `json:"fullName,required"`
	// Subdomain name
	Name string `json:"name,required"`
	// Whether SSL is enabled for tracking URLs
	SslEnabled bool `json:"sslEnabled,required"`
	// Whether click tracking is enabled
	TrackClicks bool `json:"trackClicks,required"`
	// Whether open tracking is enabled
	TrackOpens bool `json:"trackOpens,required"`
	// When DNS was last checked
	DNSCheckedAt time.Time `json:"dnsCheckedAt,nullable" format:"date-time"`
	// DNS error message if verification failed
	DNSError string `json:"dnsError,nullable"`
	// Required DNS record configuration
	DNSRecord TrackDomainDNSRecord `json:"dnsRecord,nullable"`
	// Current DNS verification status
	//
	// Any of "ok", "missing", "invalid".
	DNSStatus TrackDomainDNSStatus `json:"dnsStatus,nullable"`
	// Domains excluded from click tracking
	ExcludedClickDomains string `json:"excludedClickDomains,nullable"`
	// When the track domain was last updated
	UpdatedAt time.Time `json:"updatedAt,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		DNSOk                respjson.Field
		DomainID             respjson.Field
		FullName             respjson.Field
		Name                 respjson.Field
		SslEnabled           respjson.Field
		TrackClicks          respjson.Field
		TrackOpens           respjson.Field
		DNSCheckedAt         respjson.Field
		DNSError             respjson.Field
		DNSRecord            respjson.Field
		DNSStatus            respjson.Field
		ExcludedClickDomains respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackDomain) RawJSON() string { return r.JSON.raw }
func (r *TrackDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required DNS record configuration
type TrackDomainDNSRecord struct {
	// DNS record name
	Name string `json:"name"`
	// DNS record type
	Type string `json:"type"`
	// DNS record value (target)
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackDomainDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *TrackDomainDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current DNS verification status
type TrackDomainDNSStatus string

const (
	TrackDomainDNSStatusOk      TrackDomainDNSStatus = "ok"
	TrackDomainDNSStatusMissing TrackDomainDNSStatus = "missing"
	TrackDomainDNSStatusInvalid TrackDomainDNSStatus = "invalid"
)

type TenantTrackingNewResponse struct {
	Data    TrackDomain    `json:"data,required"`
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
func (r TenantTrackingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingGetResponse struct {
	Data    TrackDomain    `json:"data,required"`
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
func (r TenantTrackingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingUpdateResponse struct {
	Data    TrackDomain    `json:"data,required"`
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
func (r TenantTrackingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingListResponse struct {
	Data    TenantTrackingListResponseData `json:"data,required"`
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
func (r TenantTrackingListResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingListResponseData struct {
	TrackDomains []TrackDomain `json:"trackDomains,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TrackDomains respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantTrackingListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingDeleteResponse struct {
	Data    TenantTrackingDeleteResponseData `json:"data,required"`
	Meta    shared.APIMeta                   `json:"meta,required"`
	Success bool                             `json:"success,required"`
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
func (r TenantTrackingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingDeleteResponseData struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantTrackingDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingVerifyResponse struct {
	Data    TenantTrackingVerifyResponseData `json:"data,required"`
	Meta    shared.APIMeta                   `json:"meta,required"`
	Success bool                             `json:"success,required"`
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
func (r TenantTrackingVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingVerifyResponseData struct {
	// Track domain ID
	ID string `json:"id,required"`
	// Whether DNS is correctly configured
	DNSOk bool `json:"dnsOk,required"`
	// Current DNS verification status
	//
	// Any of "ok", "missing", "invalid".
	DNSStatus string `json:"dnsStatus,required"`
	// Full domain name
	FullName string `json:"fullName,required"`
	// When DNS was last checked
	DNSCheckedAt time.Time `json:"dnsCheckedAt,nullable" format:"date-time"`
	// DNS error message if verification failed
	DNSError string `json:"dnsError,nullable"`
	// Required DNS record configuration
	DNSRecord TenantTrackingVerifyResponseDataDNSRecord `json:"dnsRecord,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DNSOk        respjson.Field
		DNSStatus    respjson.Field
		FullName     respjson.Field
		DNSCheckedAt respjson.Field
		DNSError     respjson.Field
		DNSRecord    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantTrackingVerifyResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingVerifyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required DNS record configuration
type TenantTrackingVerifyResponseDataDNSRecord struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantTrackingVerifyResponseDataDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *TenantTrackingVerifyResponseDataDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingNewParams struct {
	// ID of the sending domain to attach this track domain to
	DomainID int64 `json:"domainId,required"`
	// Subdomain name (e.g., 'track' for track.yourdomain.com)
	Name string `json:"name,required"`
	// Enable SSL for tracking URLs (accepts null, defaults to true)
	SslEnabled param.Opt[bool] `json:"sslEnabled,omitzero"`
	// Enable click tracking (accepts null, defaults to true)
	TrackClicks param.Opt[bool] `json:"trackClicks,omitzero"`
	// Enable open tracking (tracking pixel, accepts null, defaults to true)
	TrackOpens param.Opt[bool] `json:"trackOpens,omitzero"`
	paramObj
}

func (r TenantTrackingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantTrackingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantTrackingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingGetParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	paramObj
}

type TenantTrackingUpdateParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	// Comma-separated list of domains to exclude from click tracking (accepts null)
	ExcludedClickDomains param.Opt[string] `json:"excludedClickDomains,omitzero"`
	// Enable or disable SSL for tracking URLs (accepts null)
	SslEnabled param.Opt[bool] `json:"sslEnabled,omitzero"`
	// Enable or disable click tracking (accepts null)
	TrackClicks param.Opt[bool] `json:"trackClicks,omitzero"`
	// Enable or disable open tracking (accepts null)
	TrackOpens param.Opt[bool] `json:"trackOpens,omitzero"`
	paramObj
}

func (r TenantTrackingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantTrackingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantTrackingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantTrackingDeleteParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	paramObj
}

type TenantTrackingVerifyParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	paramObj
}
