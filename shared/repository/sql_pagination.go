package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"ms-soccer/service/shared/domains"
)

type sqlPagination struct {
	DB *gorm.DB
}

func NewSqlPagination(db *gorm.DB) Pagination {
	return &sqlPagination{
		DB: db,
	}
}

func (pg *sqlPagination) Paging(entities interface{}, p PagingParam) *domains.Paginator {
	db := pg.DB
	count := 0
	p.Page = GetPage(p.Page)
	p.Limit = GetLimit(p.Limit)
	offset := GetOffset(p.Page, p.Limit)
	countChan := make(chan int)
	var paginator domains.Paginator

	if p.ShowSql {
		db = db.Debug()
	}

	for key, value := range p.Params {
		db = db.Where(fmt.Sprintf("%s like '%v'", key, value))
	}

	for _, o := range p.OrderBy {
		db = db.Order(o)
	}

	go pg.countRecords(entities, countChan)

	db.Limit(p.Limit).Offset(offset).Find(entities)
	count = <-countChan

	paginator.TotalRecord = count
	paginator.TotalPage = count / p.Limit
	paginator.Data = entities
	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.Page = p.Page
	paginator.PrevPage = GetPrevPage(p.Page)
	paginator.NextPage = GetNextPage(p.Page, paginator.TotalPage)

	return &paginator
}

func (pg *sqlPagination) countRecords(entities interface{}, countChan chan int) {
	count := 0
	pg.DB.Model(entities).Count(&count)
	countChan <- count
}
