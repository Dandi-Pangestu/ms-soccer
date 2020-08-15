package repository

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	"ms-soccer/service/shared/domains"
)

type sqlBaseRepository struct {
	DB         *gorm.DB
	Auth       *domains.Auth
	Pagination Pagination
}

func NewSqlBaseRepository(db *gorm.DB, auth *domains.Auth, pagination Pagination) BaseRepository {
	return &sqlBaseRepository{
		DB:         db,
		Auth:       auth,
		Pagination: pagination,
	}
}

func (r *sqlBaseRepository) FindByID(entity interface{}, id string) (interface{}, error) {
	err := r.DB.Where("id = ?", id).Find(entity).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	return entity, nil
}

func (r *sqlBaseRepository) Fetch(entities interface{}) (interface{}, error) {
	err := r.DB.Order("created_at desc").Find(entities).Error
	return entities, err
}

func (r *sqlBaseRepository) FetchWithFilter(entities interface{}, params map[string]interface{}) (interface{}, error) {
	db := r.DB

	for key, value := range params {
		db = db.Where(fmt.Sprintf("%s like '%v'", key, value))
	}

	err := db.Order("created_at desc").Find(entities).Error

	return entities, err
}

func (r *sqlBaseRepository) FetchWithPage(entities interface{}, page int, limit int) *domains.Paginator {
	return r.Pagination.Paging(entities, PagingParam{
		Limit:   limit,
		Page:    page,
		OrderBy: []string{"created_at desc"},
		ShowSql: true,
	})
}

func (r *sqlBaseRepository) FetchWithPageAndFilter(entities interface{}, page int, limit int,
	params map[string]interface{}) *domains.Paginator {

	return r.Pagination.Paging(entities, PagingParam{
		Limit:   limit,
		Page:    page,
		Params:  params,
		OrderBy: []string{"created_at desc"},
		ShowSql: true,
	})
}

func (r *sqlBaseRepository) SaveOrUpdate(entity interface{}, id *string) (interface{}, error) {
	var err error
	reflectEntity := reflect.ValueOf(entity)
	now := time.Now().UTC()

	if id == nil {
		createdAtMethod := reflectEntity.MethodByName("SetCreatedAt")
		createdAtMethod.Call([]reflect.Value{
			reflect.ValueOf(now),
		})

		createdByMethod := reflectEntity.MethodByName("SetCreatedBy")
		createdByMethod.Call([]reflect.Value{
			reflect.ValueOf(r.Auth.GetUserID()),
		})

		updatedAtMethod := reflectEntity.MethodByName("SetUpdatedAt")
		updatedAtMethod.Call([]reflect.Value{
			reflect.ValueOf(now),
		})

		err = r.DB.Create(entity).Error
	} else {
		updatedAtMethod := reflectEntity.MethodByName("SetUpdatedAt")
		updatedAtMethod.Call([]reflect.Value{
			reflect.ValueOf(now),
		})

		updatedAtByMethod := reflectEntity.MethodByName("SetUpdatedBy")
		updatedAtByMethod.Call([]reflect.Value{
			reflect.ValueOf(r.Auth.GetUserID()),
		})

		err = r.DB.Model(entity).Where("user = ?", id).UpdateColumn(entity).Error
	}

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *sqlBaseRepository) Delete(entity interface{}) error {
	return r.DB.Delete(entity).Error
}

func (r *sqlBaseRepository) GetDatabaseInstance() interface{} {
	return r.DB
}
