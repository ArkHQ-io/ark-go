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

// TenantDomainService contains methods and other services that help with
// interacting with the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantDomainService] method instead.
type TenantDomainService struct {
	Options []option.RequestOption
}

// NewTenantDomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTenantDomainService(opts ...option.RequestOption) (r TenantDomainService) {
	r = TenantDomainService{}
	r.Options = opts
	return
}

// Add a new sending domain to a tenant. Returns DNS records that must be
// configured before the domain can be verified.
//
// Each tenant gets their own isolated mail server for domain isolation.
//
// **Required DNS records:**
//
// - **SPF** - TXT record for sender authentication
// - **DKIM** - TXT record for email signing
// - **Return Path** - CNAME for bounce handling
//
// After adding DNS records, call
// `POST /tenants/{tenantId}/domains/{domainId}/verify` to verify.
func (r *TenantDomainService) New(ctx context.Context, tenantID string, body TenantDomainNewParams, opts ...option.RequestOption) (res *TenantDomainNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/domains", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get detailed information about a domain including DNS record status.
func (r *TenantDomainService) Get(ctx context.Context, domainID string, query TenantDomainGetParams, opts ...option.RequestOption) (res *TenantDomainGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/domains/%s", query.TenantID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all sending domains for a specific tenant with their verification status.
func (r *TenantDomainService) List(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *TenantDomainListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/domains", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a sending domain from a tenant. You will no longer be able to send emails
// from this domain.
//
// **Warning:** This action cannot be undone.
func (r *TenantDomainService) Delete(ctx context.Context, domainID string, body TenantDomainDeleteParams, opts ...option.RequestOption) (res *TenantDomainDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/domains/%s", body.TenantID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Check if DNS records are correctly configured and verify the domain. Returns the
// current status of each required DNS record.
//
// Call this after you've added the DNS records shown when creating the domain.
func (r *TenantDomainService) Verify(ctx context.Context, domainID string, body TenantDomainVerifyParams, opts ...option.RequestOption) (res *TenantDomainVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return
	}
	path := fmt.Sprintf("tenants/%s/domains/%s/verify", body.TenantID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// A DNS record that needs to be configured in your domain's DNS settings.
//
// The `name` field contains the relative hostname to enter in your DNS provider
// (which auto-appends the zone). The `fullName` field contains the complete
// fully-qualified domain name (FQDN) for reference.
//
// **Example for subdomain `mail.example.com`:**
//
// - `name`: `"mail"` (what you enter in DNS provider)
// - `fullName`: `"mail.example.com"` (the complete hostname)
//
// **Example for root domain `example.com`:**
//
// - `name`: `"@"` (DNS shorthand for apex/root)
// - `fullName`: `"example.com"`
type DNSRecord struct {
	// The complete fully-qualified domain name (FQDN). Use this as a reference to
	// verify the record is configured correctly.
	FullName string `json:"fullName,required"`
	// The relative hostname to enter in your DNS provider. Most DNS providers
	// auto-append the zone name, so you only need to enter this relative part.
	//
	// - `"@"` means the apex/root of the zone (for root domains)
	// - `"mail"` for a subdomain like `mail.example.com`
	// - `"ark-xyz._domainkey.mail"` for DKIM on a subdomain
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
		FullName    respjson.Field
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

type TenantDomainNewResponse struct {
	Data    TenantDomainNewResponseData `json:"data,required"`
	Meta    shared.APIMeta              `json:"meta,required"`
	Success bool                        `json:"success,required"`
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
func (r TenantDomainNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainNewResponseData struct {
	// Unique domain identifier
	ID int64 `json:"id,required"`
	// Timestamp when the domain was added
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// DNS records that must be added to your domain's DNS settings. Null if records
	// are not yet generated.
	//
	// **Important:** The `name` field contains the relative hostname that you should
	// enter in your DNS provider. Most DNS providers auto-append the zone name, so you
	// only need to enter the relative part.
	//
	// For subdomains like `mail.example.com`, the zone is `example.com`, so:
	//
	// - SPF `name` would be `mail` (not `@`)
	// - DKIM `name` would be `ark-xyz._domainkey.mail`
	// - Return Path `name` would be `psrp.mail`
	DNSRecords TenantDomainNewResponseDataDNSRecords `json:"dnsRecords,required"`
	// The domain name used for sending emails
	Name string `json:"name,required"`
	// UUID of the domain
	Uuid string `json:"uuid,required" format:"uuid"`
	// Whether all DNS records (SPF, DKIM, Return Path) are correctly configured.
	// Domain must be verified before sending emails.
	Verified bool `json:"verified,required"`
	// ID of the tenant this domain belongs to
	TenantID string `json:"tenant_id"`
	// Name of the tenant this domain belongs to
	TenantName string `json:"tenant_name"`
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
		TenantID    respjson.Field
		TenantName  respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS records that must be added to your domain's DNS settings. Null if records
// are not yet generated.
//
// **Important:** The `name` field contains the relative hostname that you should
// enter in your DNS provider. Most DNS providers auto-append the zone name, so you
// only need to enter the relative part.
//
// For subdomains like `mail.example.com`, the zone is `example.com`, so:
//
// - SPF `name` would be `mail` (not `@`)
// - DKIM `name` would be `ark-xyz._domainkey.mail`
// - Return Path `name` would be `psrp.mail`
type TenantDomainNewResponseDataDNSRecords struct {
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	Dkim DNSRecord `json:"dkim,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	ReturnPath DNSRecord `json:"returnPath,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	Spf DNSRecord `json:"spf,nullable"`
	// The DNS zone (registrable domain) where records should be added. This is the
	// root domain that your DNS provider manages. For `mail.example.com`, the zone is
	// `example.com`. For `example.co.uk`, the zone is `example.co.uk`.
	Zone string `json:"zone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dkim        respjson.Field
		ReturnPath  respjson.Field
		Spf         respjson.Field
		Zone        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainNewResponseDataDNSRecords) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainNewResponseDataDNSRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainGetResponse struct {
	Data    TenantDomainGetResponseData `json:"data,required"`
	Meta    shared.APIMeta              `json:"meta,required"`
	Success bool                        `json:"success,required"`
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
func (r TenantDomainGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainGetResponseData struct {
	// Unique domain identifier
	ID int64 `json:"id,required"`
	// Timestamp when the domain was added
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// DNS records that must be added to your domain's DNS settings. Null if records
	// are not yet generated.
	//
	// **Important:** The `name` field contains the relative hostname that you should
	// enter in your DNS provider. Most DNS providers auto-append the zone name, so you
	// only need to enter the relative part.
	//
	// For subdomains like `mail.example.com`, the zone is `example.com`, so:
	//
	// - SPF `name` would be `mail` (not `@`)
	// - DKIM `name` would be `ark-xyz._domainkey.mail`
	// - Return Path `name` would be `psrp.mail`
	DNSRecords TenantDomainGetResponseDataDNSRecords `json:"dnsRecords,required"`
	// The domain name used for sending emails
	Name string `json:"name,required"`
	// UUID of the domain
	Uuid string `json:"uuid,required" format:"uuid"`
	// Whether all DNS records (SPF, DKIM, Return Path) are correctly configured.
	// Domain must be verified before sending emails.
	Verified bool `json:"verified,required"`
	// ID of the tenant this domain belongs to
	TenantID string `json:"tenant_id"`
	// Name of the tenant this domain belongs to
	TenantName string `json:"tenant_name"`
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
		TenantID    respjson.Field
		TenantName  respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS records that must be added to your domain's DNS settings. Null if records
// are not yet generated.
//
// **Important:** The `name` field contains the relative hostname that you should
// enter in your DNS provider. Most DNS providers auto-append the zone name, so you
// only need to enter the relative part.
//
// For subdomains like `mail.example.com`, the zone is `example.com`, so:
//
// - SPF `name` would be `mail` (not `@`)
// - DKIM `name` would be `ark-xyz._domainkey.mail`
// - Return Path `name` would be `psrp.mail`
type TenantDomainGetResponseDataDNSRecords struct {
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	Dkim DNSRecord `json:"dkim,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	ReturnPath DNSRecord `json:"returnPath,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	Spf DNSRecord `json:"spf,nullable"`
	// The DNS zone (registrable domain) where records should be added. This is the
	// root domain that your DNS provider manages. For `mail.example.com`, the zone is
	// `example.com`. For `example.co.uk`, the zone is `example.co.uk`.
	Zone string `json:"zone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dkim        respjson.Field
		ReturnPath  respjson.Field
		Spf         respjson.Field
		Zone        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainGetResponseDataDNSRecords) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainGetResponseDataDNSRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainListResponse struct {
	Data    TenantDomainListResponseData `json:"data,required"`
	Meta    shared.APIMeta               `json:"meta,required"`
	Success bool                         `json:"success,required"`
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
func (r TenantDomainListResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainListResponseData struct {
	Domains []TenantDomainListResponseDataDomain `json:"domains,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainListResponseDataDomain struct {
	// Unique domain identifier
	ID int64 `json:"id,required"`
	// The domain name used for sending emails
	Name string `json:"name,required"`
	// Whether all DNS records (SPF, DKIM, Return Path) are correctly configured.
	// Domain must be verified before sending emails.
	Verified bool `json:"verified,required"`
	// ID of the tenant this domain belongs to (included when filtering by tenant_id)
	TenantID string `json:"tenant_id"`
	// Name of the tenant this domain belongs to (included when filtering by tenant_id)
	TenantName string `json:"tenant_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Verified    respjson.Field
		TenantID    respjson.Field
		TenantName  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainListResponseDataDomain) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainListResponseDataDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainDeleteResponse struct {
	Data    TenantDomainDeleteResponseData `json:"data,required"`
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
func (r TenantDomainDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainDeleteResponseData struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainVerifyResponse struct {
	Data    TenantDomainVerifyResponseData `json:"data,required"`
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
func (r TenantDomainVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainVerifyResponseData struct {
	// Unique domain identifier
	ID int64 `json:"id,required"`
	// Timestamp when the domain was added
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// DNS records that must be added to your domain's DNS settings. Null if records
	// are not yet generated.
	//
	// **Important:** The `name` field contains the relative hostname that you should
	// enter in your DNS provider. Most DNS providers auto-append the zone name, so you
	// only need to enter the relative part.
	//
	// For subdomains like `mail.example.com`, the zone is `example.com`, so:
	//
	// - SPF `name` would be `mail` (not `@`)
	// - DKIM `name` would be `ark-xyz._domainkey.mail`
	// - Return Path `name` would be `psrp.mail`
	DNSRecords TenantDomainVerifyResponseDataDNSRecords `json:"dnsRecords,required"`
	// The domain name used for sending emails
	Name string `json:"name,required"`
	// UUID of the domain
	Uuid string `json:"uuid,required" format:"uuid"`
	// Whether all DNS records (SPF, DKIM, Return Path) are correctly configured.
	// Domain must be verified before sending emails.
	Verified bool `json:"verified,required"`
	// ID of the tenant this domain belongs to
	TenantID string `json:"tenant_id"`
	// Name of the tenant this domain belongs to
	TenantName string `json:"tenant_name"`
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
		TenantID    respjson.Field
		TenantName  respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainVerifyResponseData) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainVerifyResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DNS records that must be added to your domain's DNS settings. Null if records
// are not yet generated.
//
// **Important:** The `name` field contains the relative hostname that you should
// enter in your DNS provider. Most DNS providers auto-append the zone name, so you
// only need to enter the relative part.
//
// For subdomains like `mail.example.com`, the zone is `example.com`, so:
//
// - SPF `name` would be `mail` (not `@`)
// - DKIM `name` would be `ark-xyz._domainkey.mail`
// - Return Path `name` would be `psrp.mail`
type TenantDomainVerifyResponseDataDNSRecords struct {
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	Dkim DNSRecord `json:"dkim,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	ReturnPath DNSRecord `json:"returnPath,nullable"`
	// A DNS record that needs to be configured in your domain's DNS settings.
	//
	// The `name` field contains the relative hostname to enter in your DNS provider
	// (which auto-appends the zone). The `fullName` field contains the complete
	// fully-qualified domain name (FQDN) for reference.
	//
	// **Example for subdomain `mail.example.com`:**
	//
	// - `name`: `"mail"` (what you enter in DNS provider)
	// - `fullName`: `"mail.example.com"` (the complete hostname)
	//
	// **Example for root domain `example.com`:**
	//
	// - `name`: `"@"` (DNS shorthand for apex/root)
	// - `fullName`: `"example.com"`
	Spf DNSRecord `json:"spf,nullable"`
	// The DNS zone (registrable domain) where records should be added. This is the
	// root domain that your DNS provider manages. For `mail.example.com`, the zone is
	// `example.com`. For `example.co.uk`, the zone is `example.co.uk`.
	Zone string `json:"zone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dkim        respjson.Field
		ReturnPath  respjson.Field
		Spf         respjson.Field
		Zone        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantDomainVerifyResponseDataDNSRecords) RawJSON() string { return r.JSON.raw }
func (r *TenantDomainVerifyResponseDataDNSRecords) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainNewParams struct {
	// Domain name (e.g., "mail.example.com")
	Name string `json:"name,required"`
	paramObj
}

func (r TenantDomainNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TenantDomainNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantDomainNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantDomainGetParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	paramObj
}

type TenantDomainDeleteParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	paramObj
}

type TenantDomainVerifyParams struct {
	TenantID string `path:"tenantId,required" json:"-"`
	paramObj
}
