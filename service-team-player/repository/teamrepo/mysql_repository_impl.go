package teamrepo

import (
	"github.com/jinzhu/gorm"
	"ms-soccer/service/service-team-player/models"
	"ms-soccer/service/shared/domains"
	"ms-soccer/service/shared/repository"
)

type mysqlRepositoryImpl struct {
	BaseRepo repository.BaseRepository
}

func NewMysqlRepositoryImpl(baseRepo repository.BaseRepository) TeamRepository {
	return &mysqlRepositoryImpl{
		BaseRepo: baseRepo,
	}
}

func (r *mysqlRepositoryImpl) FindByID(id string) (*models.Team, error) {
	entity, err := r.BaseRepo.FindByID(&models.Team{}, id)
	if err != nil {
		return nil, err
	}

	return entity.(*models.Team), nil
}

func (r *mysqlRepositoryImpl) Fetch() (*[]models.Team, error) {
	entities, err := r.BaseRepo.Fetch(&[]models.Team{})
	if err != nil {
		return nil, err
	}

	return entities.(*[]models.Team), nil
}

func (r *mysqlRepositoryImpl) FetchWithFilter(params map[string]interface{}) (*[]models.Team, error) {
	entities, err := r.BaseRepo.FetchWithFilter(&[]models.Team{}, params)
	if err != nil {
		return nil, err
	}

	return entities.(*[]models.Team), nil
}

func (r *mysqlRepositoryImpl) FetchWithPage(page int, limit int) *domains.Paginator {
	return r.BaseRepo.FetchWithPage(&[]models.Team{}, page, limit)
}

func (r *mysqlRepositoryImpl) FetchWithPageAndFilter(page int, limit int,
	params map[string]interface{}) *domains.Paginator {

	return r.BaseRepo.FetchWithPageAndFilter(&[]models.Team{}, page, limit, params)
}

func (r *mysqlRepositoryImpl) SaveOrUpdate(entity *models.Team, id *string) (*models.Team, error) {
	ent, err := r.BaseRepo.SaveOrUpdate(entity, id)
	if err != nil {
		return nil, err
	}

	return ent.(*models.Team), nil
}

func (r *mysqlRepositoryImpl) Delete(entity *models.Team) error {
	return r.BaseRepo.Delete(entity)
}

func (r *mysqlRepositoryImpl) GetDatabaseInstance() *gorm.DB {
	return r.BaseRepo.GetDatabaseInstance().(*gorm.DB)
}
