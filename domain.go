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

// DomainService contains methods and other services that help with interacting
// with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r DomainService) {
	r = DomainService{}
	r.Options = opts
	return
}

// Add a new domain for sending emails. Returns DNS records that must be configured
// before the domain can be verified.
//
// **Required DNS records:**
//
// - **SPF** - TXT record for sender authentication
// - **DKIM** - TXT record for email signing
// - **Return Path** - CNAME for bounce handling
//
// After adding DNS records, call `POST /domains/{id}/verify` to verify.
func (r *DomainService) New(ctx context.Context, body DomainNewParams, opts ...option.RequestOption) (res *DomainNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get detailed information about a domain including DNS record status
func (r *DomainService) Get(ctx context.Context, domainID string, opts ...option.RequestOption) (res *DomainGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return
	}
	path := fmt.Sprintf("domains/%s", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all sending domains with their verification status
func (r *DomainService) List(ctx context.Context, opts ...option.RequestOption) (res *DomainListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a sending domain. You will no longer be able to send emails from this
// domain.
//
// **Warning:** This action cannot be undone.
func (r *DomainService) Delete(ctx context.Context, domainID string, opts ...option.RequestOption) (res *DomainDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return
	}
	path := fmt.Sprintf("domains/%s", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Check if DNS records are correctly configured and verify the domain. Returns the
// current status of each required DNS record.
//
// Call this after you've added the DNS records shown when creating the domain.
func (r *DomainService) Verify(ctx context.Context, domainID string, opts ...option.RequestOption) (res *DomainVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return
	}
	path := fmt.Sprintf("domains/%s/verify", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// A DNS record that needs to be configured in your domain's DNS settings
type DNSRecord struct {
	// The hostname where the record should be created (relative to your domain)
	Name string `json:"name,required"`
	// The DNS record type to create
	//
	// Any of "TXT", "CNAME", "MX".
	Type DNSRecordType `json:"type,required"`
	// The value to set for the DNS record
	Value string `json:"value,required"`
	// Current verification status of this DNS record:
	//
	// - `OK` - Record is correctly configured and verified
	// - `Missing` - Record was not found in your DNS
	// - `Invalid` - Record exists but has an incorrect value
	// - `null` - Record has not been checked yet
	//
	// Any of "OK", "Missing", "Invalid".
	Status DNSRecordStatus `json:"status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DNSRecord) RawJSON() string { return r.JSON.raw }
func (r *DNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The DNS record type to create
type DNSRecordType string

const (
	DNSRecordTypeTxt   DNSRecordType = "TXT"
	DNSRecordTypeCname DNSRecordType = "CNAME"
	DNSRecordTypeMx    DNSRecordType = "MX"
)

// Current verification status of this DNS record:
//
// - `OK` - Record is correctly configured and verified
// - `Missing` - Record was not found in your DNS
// - `Invalid` - Record exists but has an incorrect value
// - `null` - Record has not been checked yet
type DNSRecordStatus string

const (
	DNSRecordStatusOk      DNSRecordStatus = "OK"
	DNSRecordStatusMissing DNSRecordStatus = "Missing"
	DNSRecordStatusInvalid DNSRecordStatus = "Invalid"
)

type DomainNewResponse struct {
	Data    DomainNewResponseData `json:"data,required"`
	Meta    shared.APIMeta        `json:"meta,required"`
	Success bool                  `json:"success,required"`
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
func (r DomainNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainNewResponseData struct {
	// Unique domain identifier
	ID int64 `json:"id,required"`
	// Timestamp when the domain was added
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// DNS records that must be added to your domain's DNS settings. Null if records
	// are not yet generated.
	DNSRecords DomainNewResponseDataDNSRecords `json:"dnsRecords,required"`
	// The domain name used for sending emails
	Name string `json:"name,required"`
	// UUID of the domain
	Uuid string `json:"uuid,required" format:"uuid"`
	// Whether all DNS records (SPF, DKIM, Return Path) are correctly configured.
	// Domain must be verified before sending emails.
	Verified bool `json:"verified,required"`
	// Timestamp when the domain ownership was verified, or null if not yet verified
	VerifiedAt time.Time `json:"verifiedAt,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DNSRecords  respjson.Field
		Name        respjson.Field
		Uuid        respjson.Field
		Verified    respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *DomainNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS records that must be added to your domain's DNS settings. Null if records
// are not yet generated.
type DomainNewResponseDataDNSRecords struct {
	// A DNS record that needs to be configured in your domain's DNS settings
	Dkim DNSRecord `json:"dkim,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings
	ReturnPath DNSRecord `json:"returnPath,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings
	Spf DNSRecord `json:"spf,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dkim        respjson.Field
		ReturnPath  respjson.Field
		Spf         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainNewResponseDataDNSRecords) RawJSON() string { return r.JSON.raw }
func (r *DomainNewResponseDataDNSRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainGetResponse struct {
	Data    DomainGetResponseData `json:"data,required"`
	Meta    shared.APIMeta        `json:"meta,required"`
	Success bool                  `json:"success,required"`
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
func (r DomainGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainGetResponseData struct {
	// Unique domain identifier
	ID int64 `json:"id,required"`
	// Timestamp when the domain was added
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// DNS records that must be added to your domain's DNS settings. Null if records
	// are not yet generated.
	DNSRecords DomainGetResponseDataDNSRecords `json:"dnsRecords,required"`
	// The domain name used for sending emails
	Name string `json:"name,required"`
	// UUID of the domain
	Uuid string `json:"uuid,required" format:"uuid"`
	// Whether all DNS records (SPF, DKIM, Return Path) are correctly configured.
	// Domain must be verified before sending emails.
	Verified bool `json:"verified,required"`
	// Timestamp when the domain ownership was verified, or null if not yet verified
	VerifiedAt time.Time `json:"verifiedAt,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DNSRecords  respjson.Field
		Name        respjson.Field
		Uuid        respjson.Field
		Verified    respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *DomainGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS records that must be added to your domain's DNS settings. Null if records
// are not yet generated.
type DomainGetResponseDataDNSRecords struct {
	// A DNS record that needs to be configured in your domain's DNS settings
	Dkim DNSRecord `json:"dkim,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings
	ReturnPath DNSRecord `json:"returnPath,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings
	Spf DNSRecord `json:"spf,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dkim        respjson.Field
		ReturnPath  respjson.Field
		Spf         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainGetResponseDataDNSRecords) RawJSON() string { return r.JSON.raw }
func (r *DomainGetResponseDataDNSRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainListResponse struct {
	Data    DomainListResponseData `json:"data,required"`
	Meta    shared.APIMeta         `json:"meta,required"`
	Success bool                   `json:"success,required"`
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
func (r DomainListResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainListResponseData struct {
	Domains []DomainListResponseDataDomain `json:"domains,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainListResponseData) RawJSON() string { return r.JSON.raw }
func (r *DomainListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainListResponseDataDomain struct {
	// Unique domain identifier
	ID int64 `json:"id,required"`
	// The domain name used for sending emails
	Name string `json:"name,required"`
	// Whether all DNS records (SPF, DKIM, Return Path) are correctly configured.
	// Domain must be verified before sending emails.
	Verified bool `json:"verified,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Verified    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainListResponseDataDomain) RawJSON() string { return r.JSON.raw }
func (r *DomainListResponseDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainDeleteResponse struct {
	Data    DomainDeleteResponseData `json:"data,required"`
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
func (r DomainDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainDeleteResponseData struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *DomainDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainVerifyResponse struct {
	Data    DomainVerifyResponseData `json:"data,required"`
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
func (r DomainVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainVerifyResponseData struct {
	// Unique domain identifier
	ID int64 `json:"id,required"`
	// Timestamp when the domain was added
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// DNS records that must be added to your domain's DNS settings. Null if records
	// are not yet generated.
	DNSRecords DomainVerifyResponseDataDNSRecords `json:"dnsRecords,required"`
	// The domain name used for sending emails
	Name string `json:"name,required"`
	// UUID of the domain
	Uuid string `json:"uuid,required" format:"uuid"`
	// Whether all DNS records (SPF, DKIM, Return Path) are correctly configured.
	// Domain must be verified before sending emails.
	Verified bool `json:"verified,required"`
	// Timestamp when the domain ownership was verified, or null if not yet verified
	VerifiedAt time.Time `json:"verifiedAt,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DNSRecords  respjson.Field
		Name        respjson.Field
		Uuid        respjson.Field
		Verified    respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainVerifyResponseData) RawJSON() string { return r.JSON.raw }
func (r *DomainVerifyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS records that must be added to your domain's DNS settings. Null if records
// are not yet generated.
type DomainVerifyResponseDataDNSRecords struct {
	// A DNS record that needs to be configured in your domain's DNS settings
	Dkim DNSRecord `json:"dkim,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings
	ReturnPath DNSRecord `json:"returnPath,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings
	Spf DNSRecord `json:"spf,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dkim        respjson.Field
		ReturnPath  respjson.Field
		Spf         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainVerifyResponseDataDNSRecords) RawJSON() string { return r.JSON.raw }
func (r *DomainVerifyResponseDataDNSRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainNewParams struct {
	// Domain name (e.g., "mail.example.com")
	Name string `json:"name,required"`
	paramObj
}

func (r DomainNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DomainNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomainNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
