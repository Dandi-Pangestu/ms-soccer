package repository

import "ms-soccer/service/shared/domains"

type BaseRepository interface {
	FindByID(entity interface{}, id string) (interface{}, error)
	Fetch(entities interface{}) (interface{}, error)
	FetchWithFilter(entities interface{}, params map[string]interface{}) (interface{}, error)
	FetchWithPage(entities interface{}, page int, limit int) *domains.Paginator
	FetchWithPageAndFilter(entities interface{}, page int, limit int, params map[string]interface{}) *domains.Paginator
	SaveOrUpdate(entity interface{}, id *string) (interface{}, error)
	Delete(entity interface{}) error
	GetDatabaseInstance() interface{}
}
