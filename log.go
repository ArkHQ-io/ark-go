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

// LogService contains methods and other services that help with interacting with
// the ark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r LogService) {
	r = LogService{}
	r.Options = opts
	return
}

// Retrieve detailed information about a specific API request log, including the
// full request and response bodies.
//
// **Body decryption:** Request and response bodies are stored encrypted and
// automatically decrypted when retrieved. Bodies larger than 25KB are truncated at
// storage time with a `... [truncated]` marker.
//
// **Use cases:**
//
// - Debug a specific failed request
// - Review the exact payload sent/received
// - Share request details with support
//
// **Related endpoints:**
//
// - `GET /logs` - List logs with filters
func (r *LogService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *LogGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return
	}
	path := fmt.Sprintf("logs/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a paginated list of API request logs for debugging and monitoring.
// Results are ordered by timestamp, newest first.
//
// **Use cases:**
//
// - Debug integration issues by reviewing recent requests
// - Monitor error rates and response times
// - Audit API usage patterns
//
// **Filters:**
//
// - `status` - Filter by success or error category
// - `statusCode` - Filter by exact HTTP status code
// - `endpoint` - Filter by endpoint name (e.g., `emails.send`)
// - `credentialId` - Filter by API key
// - `startDate`/`endDate` - Filter by date range
//
// **Note:** Request and response bodies are only included when retrieving a single
// log entry with `GET /logs/{requestId}`.
//
// **Related endpoints:**
//
// - `GET /logs/{requestId}` - Get full log details with request/response bodies
func (r *LogService) List(ctx context.Context, query LogListParams, opts ...option.RequestOption) (res *pagination.PageNumberPagination[LogEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "logs"
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

// Retrieve a paginated list of API request logs for debugging and monitoring.
// Results are ordered by timestamp, newest first.
//
// **Use cases:**
//
// - Debug integration issues by reviewing recent requests
// - Monitor error rates and response times
// - Audit API usage patterns
//
// **Filters:**
//
// - `status` - Filter by success or error category
// - `statusCode` - Filter by exact HTTP status code
// - `endpoint` - Filter by endpoint name (e.g., `emails.send`)
// - `credentialId` - Filter by API key
// - `startDate`/`endDate` - Filter by date range
//
// **Note:** Request and response bodies are only included when retrieving a single
// log entry with `GET /logs/{requestId}`.
//
// **Related endpoints:**
//
// - `GET /logs/{requestId}` - Get full log details with request/response bodies
func (r *LogService) ListAutoPaging(ctx context.Context, query LogListParams, opts ...option.RequestOption) *pagination.PageNumberPaginationAutoPager[LogEntry] {
	return pagination.NewPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// API request log entry (list view)
type LogEntry struct {
	// Request context information
	Context LogEntryContext `json:"context" api:"required"`
	// API credential information
	Credential LogEntryCredential `json:"credential" api:"required"`
	// Request duration in milliseconds
	DurationMs int64 `json:"durationMs" api:"required"`
	// Semantic endpoint name
	Endpoint string `json:"endpoint" api:"required"`
	// HTTP method
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method LogEntryMethod `json:"method" api:"required"`
	// Request path
	Path string `json:"path" api:"required"`
	// Rate limit state at time of request
	RateLimit LogEntryRateLimit `json:"rateLimit" api:"required"`
	// Unique request identifier
	RequestID string `json:"requestId" api:"required"`
	// HTTP response status code
	StatusCode int64 `json:"statusCode" api:"required"`
	// When the request was made (ISO 8601)
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Email-specific data (for email endpoints)
	Email LogEntryEmail `json:"email" api:"nullable"`
	// Error details (null if request succeeded)
	Error LogEntryError `json:"error" api:"nullable"`
	// SDK information (null if not using an SDK)
	SDK LogEntrySDK `json:"sdk" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Credential  respjson.Field
		DurationMs  respjson.Field
		Endpoint    respjson.Field
		Method      respjson.Field
		Path        respjson.Field
		RateLimit   respjson.Field
		RequestID   respjson.Field
		StatusCode  respjson.Field
		Timestamp   respjson.Field
		Email       respjson.Field
		Error       respjson.Field
		SDK         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntry) RawJSON() string { return r.JSON.raw }
func (r *LogEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request context information
type LogEntryContext struct {
	// Idempotency key if provided
	IdempotencyKey string `json:"idempotencyKey" api:"nullable"`
	// Client IP address
	IPAddress string `json:"ipAddress" api:"nullable"`
	// Query parameters
	QueryParams map[string]any `json:"queryParams" api:"nullable"`
	// User-Agent header
	UserAgent string `json:"userAgent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IdempotencyKey respjson.Field
		IPAddress      respjson.Field
		QueryParams    respjson.Field
		UserAgent      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntryContext) RawJSON() string { return r.JSON.raw }
func (r *LogEntryContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// API credential information
type LogEntryCredential struct {
	// Credential ID
	ID string `json:"id" api:"required"`
	// API key prefix (first 8 characters)
	KeyPrefix string `json:"keyPrefix" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		KeyPrefix   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntryCredential) RawJSON() string { return r.JSON.raw }
func (r *LogEntryCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP method
type LogEntryMethod string

const (
	LogEntryMethodGet    LogEntryMethod = "GET"
	LogEntryMethodPost   LogEntryMethod = "POST"
	LogEntryMethodPut    LogEntryMethod = "PUT"
	LogEntryMethodPatch  LogEntryMethod = "PATCH"
	LogEntryMethodDelete LogEntryMethod = "DELETE"
)

// Rate limit state at time of request
type LogEntryRateLimit struct {
	// Rate limit ceiling
	Limit int64 `json:"limit" api:"nullable"`
	// Whether the request was rate limited
	Limited bool `json:"limited"`
	// Remaining requests in window
	Remaining int64 `json:"remaining" api:"nullable"`
	// Unix timestamp when limit resets
	Reset int64 `json:"reset" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Limited     respjson.Field
		Remaining   respjson.Field
		Reset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntryRateLimit) RawJSON() string { return r.JSON.raw }
func (r *LogEntryRateLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email-specific data (for email endpoints)
type LogEntryEmail struct {
	// Email message identifier (token)
	ID string `json:"id"`
	// Number of recipients
	RecipientCount int64 `json:"recipientCount" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		RecipientCount respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntryEmail) RawJSON() string { return r.JSON.raw }
func (r *LogEntryEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details (null if request succeeded)
type LogEntryError struct {
	// Error code
	Code string `json:"code"`
	// Error message
	Message string `json:"message" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntryError) RawJSON() string { return r.JSON.raw }
func (r *LogEntryError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SDK information (null if not using an SDK)
type LogEntrySDK struct {
	// SDK name
	Name string `json:"name"`
	// SDK version
	Version string `json:"version" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntrySDK) RawJSON() string { return r.JSON.raw }
func (r *LogEntrySDK) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full API request log entry with bodies
type LogEntryDetail struct {
	// Request body information
	Request LogEntryDetailRequest `json:"request"`
	// Response body information
	Response LogEntryDetailResponse `json:"response"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LogEntry
}

// Returns the unmodified JSON received from the API
func (r LogEntryDetail) RawJSON() string { return r.JSON.raw }
func (r *LogEntryDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body information
type LogEntryDetailRequest struct {
	// Decrypted request body (JSON or string). Bodies over 25KB are truncated.
	Body LogEntryDetailRequestBodyUnion `json:"body" api:"nullable"`
	// Original request body size in bytes
	BodySize int64 `json:"bodySize" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		BodySize    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntryDetailRequest) RawJSON() string { return r.JSON.raw }
func (r *LogEntryDetailRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogEntryDetailRequestBodyUnion contains all possible properties and values from
// [map[string]any], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfLogEntryDetailRequestBodyMapItem OfString]
type LogEntryDetailRequestBodyUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfLogEntryDetailRequestBodyMapItem any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfLogEntryDetailRequestBodyMapItem respjson.Field
		OfString                           respjson.Field
		raw                                string
	} `json:"-"`
}

func (u LogEntryDetailRequestBodyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogEntryDetailRequestBodyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LogEntryDetailRequestBodyUnion) RawJSON() string { return u.JSON.raw }

func (r *LogEntryDetailRequestBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response body information
type LogEntryDetailResponse struct {
	// Decrypted response body (JSON or string). Bodies over 25KB are truncated.
	Body LogEntryDetailResponseBodyUnion `json:"body" api:"nullable"`
	// Response body size in bytes
	BodySize int64 `json:"bodySize" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		BodySize    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogEntryDetailResponse) RawJSON() string { return r.JSON.raw }
func (r *LogEntryDetailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LogEntryDetailResponseBodyUnion contains all possible properties and values from
// [map[string]any], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfLogEntryDetailResponseBodyMapItem OfString]
type LogEntryDetailResponseBodyUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfLogEntryDetailResponseBodyMapItem any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfLogEntryDetailResponseBodyMapItem respjson.Field
		OfString                            respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u LogEntryDetailResponseBodyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LogEntryDetailResponseBodyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LogEntryDetailResponseBodyUnion) RawJSON() string { return u.JSON.raw }

func (r *LogEntryDetailResponseBodyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed API request log with request/response bodies
type LogGetResponse struct {
	// Full API request log entry with bodies
	Data    LogEntryDetail `json:"data" api:"required"`
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
func (r LogGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LogGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogListParams struct {
	// Filter by API credential ID
	CredentialID param.Opt[string] `query:"credentialId,omitzero" json:"-"`
	// Filter logs before this date (ISO 8601 format)
	EndDate param.Opt[time.Time] `query:"endDate,omitzero" format:"date-time" json:"-"`
	// Filter by endpoint name
	Endpoint param.Opt[string] `query:"endpoint,omitzero" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Results per page (max 100)
	PerPage param.Opt[int64] `query:"perPage,omitzero" json:"-"`
	// Filter by request ID (partial match)
	RequestID param.Opt[string] `query:"requestId,omitzero" json:"-"`
	// Filter logs after this date (ISO 8601 format)
	StartDate param.Opt[time.Time] `query:"startDate,omitzero" format:"date-time" json:"-"`
	// Filter by exact HTTP status code (100-599)
	StatusCode param.Opt[int64] `query:"statusCode,omitzero" json:"-"`
	// Filter by status category:
	//
	// - `success` - Status codes < 400
	// - `error` - Status codes >= 400
	//
	// Any of "success", "error".
	Status LogListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LogListParams]'s query parameters as `url.Values`.
func (r LogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status category:
//
// - `success` - Status codes < 400
// - `error` - Status codes >= 400
type LogListParamsStatus string

const (
	LogListParamsStatusSuccess LogListParamsStatus = "success"
	LogListParamsStatusError   LogListParamsStatus = "error"
)
