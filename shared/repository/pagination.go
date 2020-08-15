package repository

import "ms-soccer/service/shared/domains"

type PagingParam struct {
	Limit   int
	Page    int
	Params  map[string]interface{}
	OrderBy []string
	ShowSql bool
}

type Pagination interface {
	Paging(entities interface{}, p PagingParam) *domains.Paginator
	countRecords(entities interface{}, countChan chan int)
}

func GetPage(page int) int {
	if page <= 0 {
		return 1
	}

	return page
}

func GetLimit(limit int) int {
	if limit <= 0 {
		return 1
	}

	return limit
}

func GetOffset(page int, limit int) int {
	if page > 1 {
		return (page - 1) * limit
	}

	return 0
}

func GetPrevPage(page int) int {
	if page > 1 {
		return page - 1
	}

	return page
}

func GetNextPage(page int, totalPage int) int {
	if page == totalPage {
		return page
	}

	return page + 1
}
