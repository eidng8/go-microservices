package common

import (
	"github.com/gin-gonic/gin"
)

type PaginatedParams struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}

type PaginatedList[T any] struct {
	Total        int    `json:"total"`
	PerPage      int    `json:"per_page"`
	CurrentPage  int    `json:"current_page"`
	LastPage     int    `json:"last_page"`
	FirstPageUrl string `json:"first_page_url"`
	LastPageUrl  string `json:"last_page_url"`
	NextPageUrl  string `json:"next_page_url"`
	PrevPageUrl  string `json:"prev_page_url"`
	Path         string `json:"path"`
	From         int    `json:"from"`
	To           int    `json:"to"`
	Data         []T    `json:"data"`
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
