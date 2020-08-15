package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"ms-soccer/service/service-team-player/httphandler"
	"ms-soccer/service/service-team-player/repository"
	"ms-soccer/service/shared/domains"
	"ms-soccer/service/shared/middleware"
)

type ConfigParam struct {
	R           *gin.Engine
	ServiceName string
	Auth        *domains.Auth
	RepoModule  *repository.RepositoryModule
}

func Configure(c *ConfigParam) {
	jwtSecretKey := viper.GetString(fmt.Sprintf("application.resources.jwt.%s.secret-key", c.ServiceName))
	authable := middleware.NewJWTService(jwtSecretKey)
	authMiddleware := middleware.Auth(c.Auth, authable)

	v1 := c.R.Group("/v1")
	{
		teamRouter := v1.Group("/teams")
		{
			V1TeamRouter(teamRouter, authable, authMiddleware, c.RepoModule)
		}
	}
}

func V1TeamRouter(r *gin.RouterGroup, authable middleware.Authable, authMiddleware gin.HandlerFunc,
	repoModule *repository.RepositoryModule) {

	handler := httphandler.NewTeamHandler(repoModule.TeamRepo, repoModule.PlayerRepo)

	r.POST("/", handler.Store)
	r.GET("/", handler.GetList)
	r.POST("/:id/players", handler.StorePlayer)
	r.GET("/:id/players", handler.GetListPlayer)
}
