package playerrepo

import (
	"ms-soccer/service/service-team-player/models"
	"ms-soccer/service/shared/domains"
	"ms-soccer/service/shared/repository"

	"github.com/jinzhu/gorm"
)

type mysqlRepositoryImpl struct {
	BaseRepo repository.BaseRepository
}

func NewMysqlRepositoryImpl(baseRepo repository.BaseRepository) PlayerRepository {
	return &mysqlRepositoryImpl{
		BaseRepo: baseRepo,
	}
}

func (r *mysqlRepositoryImpl) FindByID(id string) (*models.Player, error) {
	entity, err := r.BaseRepo.FindByID(&models.Player{}, id)
	if err != nil {
		return nil, err
	}

	return entity.(*models.Player), nil
}

func (r *mysqlRepositoryImpl) Fetch() (*[]models.Player, error) {
	entities, err := r.BaseRepo.Fetch(&[]models.Player{})
	if err != nil {
		return nil, err
	}

	return entities.(*[]models.Player), nil
}

func (r *mysqlRepositoryImpl) FetchWithFilter(params map[string]interface{}) (*[]models.Player, error) {
	entities, err := r.BaseRepo.FetchWithFilter(&[]models.Player{}, params)
	if err != nil {
		return nil, err
	}

	return entities.(*[]models.Player), nil
}

func (r *mysqlRepositoryImpl) FetchWithPage(page int, limit int) *domains.Paginator {
	return r.BaseRepo.FetchWithPage(&[]models.Player{}, page, limit)
}

func (r *mysqlRepositoryImpl) FetchWithPageAndFilter(page int, limit int,
	params map[string]interface{}) *domains.Paginator {

	return r.BaseRepo.FetchWithPageAndFilter(&[]models.Player{}, page, limit, params)
}

func (r *mysqlRepositoryImpl) SaveOrUpdate(entity *models.Player, id *string) (*models.Player, error) {
	ent, err := r.BaseRepo.SaveOrUpdate(entity, id)
	if err != nil {
		return nil, err
	}

	return ent.(*models.Player), nil
}

func (r *mysqlRepositoryImpl) Delete(entity *models.Player) error {
	return r.BaseRepo.Delete(entity)
}

func (r *mysqlRepositoryImpl) GetDatabaseInstance() *gorm.DB {
	return r.BaseRepo.GetDatabaseInstance().(*gorm.DB)
}

func (r *mysqlRepositoryImpl) FindByTeamIDAndNumber(teamID string, number int) (*models.Player, error) {
	db := r.GetDatabaseInstance()
	var entity models.Player

	err := db.Where("team_id = ?", teamID).Where("number = ?", number).Find(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *mysqlRepositoryImpl) FetchByTeamID(teamID string) (*[]models.Player, error) {
	db := r.GetDatabaseInstance()
	var entities []models.Player

	err := db.Where("team_id = ?", teamID).Order("name asc").Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return &entities, nil
}
