// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ArkHQ-io/ark-go/internal/apijson"
	"github.com/ArkHQ-io/ark-go/internal/requestconfig"
	"github.com/ArkHQ-io/ark-go/packages/param"
	"github.com/ArkHQ-io/ark-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PageNumber[T any] struct {
	Messages []T `json:"messages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Messages    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PageNumber[T]) RawJSON() string { return r.JSON.raw }
func (r *PageNumber[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PageNumber[T]) GetNextPage() (res *PageNumber[T], err error) {
	if len(r.Messages) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
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

func (r *PageNumber[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PageNumber[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PageNumberAutoPager[T any] struct {
	page *PageNumber[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPageNumberAutoPager[T any](page *PageNumber[T], err error) *PageNumberAutoPager[T] {
	return &PageNumberAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageNumberAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Messages) == 0 {
		return false
	}
	if r.idx >= len(r.page.Messages) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Messages) == 0 {
			return false
		}
	}
	r.cur = r.page.Messages[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageNumberAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageNumberAutoPager[T]) Err() error {
	return r.err
}

func (r *PageNumberAutoPager[T]) Index() int {
	return r.run
}

type SuppressionsPage[T any] struct {
	Suppressions []T `json:"suppressions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Suppressions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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
	if len(r.Suppressions) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
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
	if r.page == nil || len(r.page.Suppressions) == 0 {
		return false
	}
	if r.idx >= len(r.page.Suppressions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Suppressions) == 0 {
			return false
		}
	}
	r.cur = r.page.Suppressions[r.idx]
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
