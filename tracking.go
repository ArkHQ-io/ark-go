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

// TrackingService contains methods and other services that help with interacting
// with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrackingService] method instead.
type TrackingService struct {
	Options []option.RequestOption
}

// NewTrackingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTrackingService(opts ...option.RequestOption) (r TrackingService) {
	r = TrackingService{}
	r.Options = opts
	return
}

// Create a new track domain for open/click tracking.
//
// After creation, you must configure a CNAME record pointing to the provided DNS
// value before tracking will work.
func (r *TrackingService) New(ctx context.Context, body TrackingNewParams, opts ...option.RequestOption) (res *TrackingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tracking"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details of a specific track domain including DNS configuration
func (r *TrackingService) Get(ctx context.Context, trackingID string, opts ...option.RequestOption) (res *TrackingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("tracking/%s", trackingID)
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
func (r *TrackingService) Update(ctx context.Context, trackingID string, body TrackingUpdateParams, opts ...option.RequestOption) (res *TrackingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("tracking/%s", trackingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all track domains configured for your server. Track domains enable open and
// click tracking for your emails.
func (r *TrackingService) List(ctx context.Context, opts ...option.RequestOption) (res *TrackingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tracking"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a track domain. This will disable tracking for any emails using this
// domain.
func (r *TrackingService) Delete(ctx context.Context, trackingID string, opts ...option.RequestOption) (res *TrackingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("tracking/%s", trackingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Check DNS configuration for the track domain.
//
// The track domain requires a CNAME record to be configured before open and click
// tracking will work. Use this endpoint to verify the DNS is correctly set up.
func (r *TrackingService) Verify(ctx context.Context, trackingID string, opts ...option.RequestOption) (res *TrackingVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if trackingID == "" {
		err = errors.New("missing required trackingId parameter")
		return
	}
	path := fmt.Sprintf("tracking/%s/verify", trackingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type TrackDomain struct {
	// Track domain ID
	ID string `json:"id,required"`
	// When the track domain was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Whether DNS is correctly configured
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

type TrackingNewResponse struct {
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
func (r TrackingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TrackingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingGetResponse struct {
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
func (r TrackingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TrackingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingUpdateResponse struct {
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
func (r TrackingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TrackingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingListResponse struct {
	Data    TrackingListResponseData `json:"data,required"`
	Meta    shared.APIMeta           `json:"meta,required"`
	Success bool                     `json:"success,required"`
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
func (r TrackingListResponse) RawJSON() string { return r.JSON.raw }
func (r *TrackingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingListResponseData struct {
	TrackDomains []TrackDomain `json:"trackDomains,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TrackDomains respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackingListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TrackingListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingDeleteResponse struct {
	Data    TrackingDeleteResponseData `json:"data,required"`
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
func (r TrackingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TrackingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingDeleteResponseData struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackingDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *TrackingDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingVerifyResponse struct {
	Data    TrackingVerifyResponseData `json:"data,required"`
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
func (r TrackingVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *TrackingVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingVerifyResponseData struct {
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
	DNSRecord TrackingVerifyResponseDataDNSRecord `json:"dnsRecord,nullable"`
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
func (r TrackingVerifyResponseData) RawJSON() string { return r.JSON.raw }
func (r *TrackingVerifyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required DNS record configuration
type TrackingVerifyResponseDataDNSRecord struct {
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
func (r TrackingVerifyResponseDataDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *TrackingVerifyResponseDataDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingNewParams struct {
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

func (r TrackingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TrackingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrackingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackingUpdateParams struct {
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

func (r TrackingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TrackingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrackingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
