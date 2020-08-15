package repository

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"ms-soccer/service/service-team-player/repository/playerrepo"
	"ms-soccer/service/service-team-player/repository/teamrepo"
	"ms-soccer/service/shared/domains"
	"ms-soccer/service/shared/repository"
)

type RepositoryModule struct {
	PlayerRepo playerrepo.PlayerRepository
	TeamRepo   teamrepo.TeamRepository
}

func Load(db interface{}, auth *domains.Auth) *RepositoryModule {
	reflectDb := reflect.ValueOf(db)
	typeOfDb := fmt.Sprintf("%T", reflectDb.Interface())

	if typeOfDb == "*gorm.DB" {
		dbType := db.(*gorm.DB)
		pagination := repository.NewSqlPagination(dbType)
		baseRepo := repository.NewSqlBaseRepository(dbType, auth, pagination)

		if dbType.Dialect().GetName() == "mysql" {
			return LoadMysqlRepository(baseRepo)
		}
	}

	return nil
}

func LoadMysqlRepository(baseRepo repository.BaseRepository) *RepositoryModule {
	playerRepo := playerrepo.NewMysqlRepositoryImpl(baseRepo)
	teamRepo := teamrepo.NewMysqlRepositoryImpl(baseRepo)

	return &RepositoryModule{
		PlayerRepo: playerRepo,
		TeamRepo:   teamRepo,
	}
}
