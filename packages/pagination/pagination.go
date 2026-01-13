// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ArkHQ-io/ark-go/internal/apijson"
	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/packages/param"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PageNumberPaginationData[T any] struct {
	Messages   []T                                `json:"messages"`
	Pagination PageNumberPaginationDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageNumberPaginationData[T]) RawJSON() string { return r.JSON.raw }
func (r *PageNumberPaginationData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageNumberPaginationDataMessage struct {
	// Internal message ID
	ID    string `json:"id,required"`
	Token string `json:"token,required"`
	From  string `json:"from,required"`
	// Current delivery status:
	//
	// - `pending` - Email accepted, waiting to be processed
	// - `sent` - Email transmitted to recipient's mail server
	// - `softfail` - Temporary delivery failure, will retry
	// - `hardfail` - Permanent delivery failure
	// - `bounced` - Email bounced back
	// - `held` - Held for manual review
	//
	// Any of "pending", "sent", "softfail", "hardfail", "bounced", "held".
	Status       string    `json:"status,required"`
	Subject      string    `json:"subject,required"`
	Timestamp    float64   `json:"timestamp,required"`
	TimestampISO time.Time `json:"timestampIso,required" format:"date-time"`
	To           string    `json:"to,required" format:"email"`
	Tag          string    `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Token        respjson.Field
		From         respjson.Field
		Status       respjson.Field
		Subject      respjson.Field
		Timestamp    respjson.Field
		TimestampISO respjson.Field
		To           respjson.Field
		Tag          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageNumberPaginationDataMessage) RawJSON() string { return r.JSON.raw }
func (r *PageNumberPaginationDataMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageNumberPaginationDataPagination struct {
	Page       int64 `json:"page"`
	TotalPages int64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageNumberPaginationDataPagination) RawJSON() string { return r.JSON.raw }
func (r *PageNumberPaginationDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageNumberPagination[T any] struct {
	Data PageNumberPaginationData[T] `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PageNumberPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *PageNumberPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumberPagination[T]) GetNextPage() (res *PageNumberPagination[T], err error) {
	if len(r.Data.Messages) == 0 {
		return nil, nil
	}
	currentPage := r.Data.Pagination.Page
	if currentPage >= r.Data.Pagination.TotalPages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PageNumberPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumberPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberPaginationAutoPager[T any] struct {
	page *PageNumberPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPageNumberPaginationAutoPager[T any](page *PageNumberPagination[T], err error) *PageNumberPaginationAutoPager[T] {
	return &PageNumberPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data.Messages) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data.Messages) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data.Messages) == 0 {
			return false
		}
	}
	r.cur = r.page.Data.Messages[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberPaginationAutoPager[T]) Index() int {
	return r.run
}

type SuppressionsPaginationData[T any] struct {
	Pagination   SuppressionsPaginationDataPagination `json:"pagination"`
	Suppressions []T                                  `json:"suppressions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination   respjson.Field
		Suppressions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuppressionsPaginationData[T]) RawJSON() string { return r.JSON.raw }
func (r *SuppressionsPaginationData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionsPaginationDataPagination struct {
	Page       int64 `json:"page"`
	TotalPages int64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuppressionsPaginationDataPagination) RawJSON() string { return r.JSON.raw }
func (r *SuppressionsPaginationDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionsPaginationDataSuppression struct {
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
func (r SuppressionsPaginationDataSuppression) RawJSON() string { return r.JSON.raw }
func (r *SuppressionsPaginationDataSuppression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionsPagination[T any] struct {
	Data SuppressionsPaginationData[T] `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r SuppressionsPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *SuppressionsPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SuppressionsPagination[T]) GetNextPage() (res *SuppressionsPagination[T], err error) {
	if len(r.Data.Suppressions) == 0 {
		return nil, nil
	}
	currentPage := r.Data.Pagination.Page
	if currentPage >= r.Data.Pagination.TotalPages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SuppressionsPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SuppressionsPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SuppressionsPaginationAutoPager[T any] struct {
	page *SuppressionsPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewSuppressionsPaginationAutoPager[T any](page *SuppressionsPagination[T], err error) *SuppressionsPaginationAutoPager[T] {
	return &SuppressionsPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SuppressionsPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data.Suppressions) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data.Suppressions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data.Suppressions) == 0 {
			return false
		}
	}
	r.cur = r.page.Data.Suppressions[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SuppressionsPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *SuppressionsPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *SuppressionsPaginationAutoPager[T]) Index() int {
	return r.run
}
