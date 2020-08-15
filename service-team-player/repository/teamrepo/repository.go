package teamrepo

import (
	"github.com/jinzhu/gorm"
	"ms-soccer/service/service-team-player/models"
	"ms-soccer/service/shared/domains"
)

type TeamRepository interface {
	FindByID(id string) (*models.Team, error)
	Fetch() (*[]models.Team, error)
	FetchWithFilter(params map[string]interface{}) (*[]models.Team, error)
	FetchWithPage(page int, limit int) *domains.Paginator
	FetchWithPageAndFilter(page int, limit int, params map[string]interface{}) *domains.Paginator
	SaveOrUpdate(entity *models.Team, id *string) (*models.Team, error)
	Delete(entity *models.Team) error
	GetDatabaseInstance() *gorm.DB
}
