// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ArkHQ-io/ark-go/internal/apijson"
	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/packages/param"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PageNumberPaginationData struct {
	Pagination PageNumberPaginationDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageNumberPaginationData) RawJSON() string { return r.JSON.raw }
func (r *PageNumberPaginationData) UnmarshalJSON(data []byte) error {
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
	Data PageNumberPaginationData `json:"data"`
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
	if len(r.Data) == 0 {
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
