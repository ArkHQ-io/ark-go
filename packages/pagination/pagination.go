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

type EmailsPageData[T any] struct {
	Messages   []T                      `json:"messages"`
	Pagination EmailsPageDataPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailsPageData[T]) RawJSON() string { return r.JSON.raw }
func (r *EmailsPageData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailsPageDataPagination struct {
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
func (r EmailsPageDataPagination) RawJSON() string { return r.JSON.raw }
func (r *EmailsPageDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailsPage[T any] struct {
	Data EmailsPageData[T] `json:"data"`
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
func (r EmailsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *EmailsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EmailsPage[T]) GetNextPage() (res *EmailsPage[T], err error) {
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

func (r *EmailsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EmailsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EmailsPageAutoPager[T any] struct {
	page *EmailsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewEmailsPageAutoPager[T any](page *EmailsPage[T], err error) *EmailsPageAutoPager[T] {
	return &EmailsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EmailsPageAutoPager[T]) Next() bool {
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

func (r *EmailsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *EmailsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *EmailsPageAutoPager[T]) Index() int {
	return r.run
}

type SuppressionsPageData[T any] struct {
	Pagination   SuppressionsPageDataPagination `json:"pagination"`
	Suppressions []T                            `json:"suppressions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination   respjson.Field
		Suppressions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SuppressionsPageData[T]) RawJSON() string { return r.JSON.raw }
func (r *SuppressionsPageData[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionsPageDataPagination struct {
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
func (r SuppressionsPageDataPagination) RawJSON() string { return r.JSON.raw }
func (r *SuppressionsPageDataPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SuppressionsPage[T any] struct {
	Data SuppressionsPageData[T] `json:"data"`
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
func (r SuppressionsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *SuppressionsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SuppressionsPage[T]) GetNextPage() (res *SuppressionsPage[T], err error) {
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

func (r *SuppressionsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SuppressionsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SuppressionsPageAutoPager[T any] struct {
	page *SuppressionsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewSuppressionsPageAutoPager[T any](page *SuppressionsPage[T], err error) *SuppressionsPageAutoPager[T] {
	return &SuppressionsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SuppressionsPageAutoPager[T]) Next() bool {
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

func (r *SuppressionsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SuppressionsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SuppressionsPageAutoPager[T]) Index() int {
	return r.run
}
