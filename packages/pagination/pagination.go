// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ArkHQ-io/ark-go/internal/apijson"
	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/option"
	"github.com/ArkHQ-io/ark-go/packages/param"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PageNumberPaginationMeta struct {
	RequestID string `json:"requestId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageNumberPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *PageNumberPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PageNumberPagination[T any] struct {
	Data       []T                      `json:"data"`
	Page       int64                    `json:"page"`
	PerPage    int64                    `json:"perPage"`
	Total      int64                    `json:"total"`
	TotalPages int64                    `json:"totalPages"`
	Meta       PageNumberPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Page        respjson.Field
		PerPage     respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		Meta        respjson.Field
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
	if len(r.Data) == 0 {
		return nil, nil
	}
	currentPage := r.Page
	if currentPage >= r.TotalPages {
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
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
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

type OffsetPaginationData[T any] struct {
	Pagination OffsetPaginationDataPagination `json:"pagination"`
	Tenants    []T                            `json:"tenants"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Tenants     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OffsetPaginationData[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPaginationData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OffsetPaginationDataPagination struct {
	HasMore bool  `json:"has_more"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
	Total   int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OffsetPaginationDataPagination) RawJSON() string { return r.JSON.raw }
func (r *OffsetPaginationDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OffsetPagination[T any] struct {
	Data OffsetPaginationData[T] `json:"data"`
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
func (r OffsetPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPagination[T]) GetNextPage() (res *OffsetPagination[T], err error) {
	if len(r.Data.Tenants) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	offset := r.Data.Pagination.Offset
	length := int64(len(r.Tenants))
	next := offset + length

	if next < r.Data.Pagination.Total && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
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

func (r *OffsetPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPaginationAutoPager[T any] struct {
	page *OffsetPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPaginationAutoPager[T any](page *OffsetPagination[T], err error) *OffsetPaginationAutoPager[T] {
	return &OffsetPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data.Tenants) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data.Tenants) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data.Tenants) == 0 {
			return false
		}
	}
	r.cur = r.page.Data.Tenants[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPaginationAutoPager[T]) Index() int {
	return r.run
}
