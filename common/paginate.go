package common

import (
	"context"

	"github.com/gin-gonic/gin"
)

type PaginatedParams struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}

func GetPaginationParams(c *gin.Context) PaginatedParams {
	return GetPaginationParamsWithDefault(c, 1, 10)
}

func GetPaginationParamsWithDefault(
	c *gin.Context, defaultPage, defaultPerPage int,
) PaginatedParams {
	var page PaginatedParams
	if c.ShouldBind(&page) != nil {
		page.Page = defaultPage
		page.PerPage = defaultPerPage
	}
	if page.Page < 1 {
		page.Page = defaultPage
	}
	if page.PerPage < 1 {
		page.PerPage = defaultPerPage
	}
	return page
}

type PaginatedList[T any] struct {
	Total        int    `json:"total" bson:"total" xml:"total" yaml:"total"`
	PerPage      int    `json:"per_page" bson:"per_page" xml:"per_page" yaml:"per_page"`
	CurrentPage  int    `json:"current_page" bson:"current_page" xml:"current_page" yaml:"current_page"`
	LastPage     int    `json:"last_page" bson:"last_page" xml:"last_page" yaml:"last_page"`
	FirstPageUrl string `json:"first_page_url" bson:"first_page_url" xml:"first_page_url" yaml:"first_page_url"`
	LastPageUrl  string `json:"last_page_url" bson:"last_page_url" xml:"last_page_url" yaml:"last_page_url"`
	NextPageUrl  string `json:"next_page_url" bson:"next_page_url" xml:"next_page_url" yaml:"next_page_url"`
	PrevPageUrl  string `json:"prev_page_url" bson:"prev_page_url" xml:"prev_page_url" yaml:"prev_page_url"`
	Path         string `json:"path" bson:"path" xml:"path" yaml:"path"`
	From         int    `json:"from" bson:"from" xml:"from" yaml:"from"`
	To           int    `json:"to" bson:"to" xml:"to" yaml:"to"`
	Data         []T    `json:"data" bson:"data" xml:"data" yaml:"data"`
}

type PQ[V any, Q any] interface {
	Offset(int) *Q
	Limit(int) *Q
	All(context.Context) ([]*V, error)
}

func GetPage[V any, Q any, T PQ[V, Q]](
	ctx context.Context, q *T, page PaginatedParams,
) (
	[]*T, error,
) {
	return q.Offset((page.Page - 1) * page.PerPage).Limit(page.PerPage).All(ctx)
}
