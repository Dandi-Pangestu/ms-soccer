package playerrepo

import (
	"github.com/jinzhu/gorm"
	"ms-soccer/service/service-team-player/models"
	"ms-soccer/service/shared/domains"
)

type PlayerRepository interface {
	FindByID(id string) (*models.Player, error)
	Fetch() (*[]models.Player, error)
	FetchWithFilter(params map[string]interface{}) (*[]models.Player, error)
	FetchWithPage(page int, limit int) *domains.Paginator
	FetchWithPageAndFilter(page int, limit int, params map[string]interface{}) *domains.Paginator
	SaveOrUpdate(entity *models.Player, id *string) (*models.Player, error)
	Delete(entity *models.Player) error
	GetDatabaseInstance() *gorm.DB
	FindByTeamIDAndNumber(teamID string, number int) (*models.Player, error)
	FetchByTeamID(teamID string) (*[]models.Player, error)
}
