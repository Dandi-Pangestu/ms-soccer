package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ms-soccer/service/service-team-player/models"
	"ms-soccer/service/service-team-player/repository"
	"ms-soccer/service/service-team-player/router"
	"ms-soccer/service/shared/config"
	"ms-soccer/service/shared/domains"
	hooks "ms-soccer/service/shared/log"
	log "ms-soccer/service/shared/log/app"
	"ms-soccer/service/shared/middleware"

	"github.com/spf13/viper"
)

var (
	serviceName string
	env         string
)

func init() {
	config.InitConfig()

	serviceName = "service-team-player"
	env = viper.GetString("application.env")

	config.InitAppLogger(hooks.NewEnvFieldHook(serviceName, env))
}

func closeAllInstance(fn func()) {
	log.Info(nil, "Closing all instance")

	fn()

	log.Info(nil, "All instance are closed")
}

func main() {
	var auth domains.Auth

	r, srv := config.InitServer()
	r.Use(middleware.CORS())

	db := config.InitMysqlDatabase(serviceName)
	db.AutoMigrate(&models.Team{}, &models.Player{})
	repoModule := repository.Load(db, &auth)
	if repoModule == nil {
		log.Panic(nil, "Error while load repositories")
	}

	router.Configure(&router.ConfigParam{
		R:           r,
		ServiceName: serviceName,
		Auth:        &auth,
		RepoModule:  repoModule,
	})

	closeInstance := func() {
		if db != nil {
			_ = db.Close()
		}
	}

	go config.Start(srv)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info(nil, "Shutdown server")

	closeAllInstance(closeInstance)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	config.Shutdown(srv, ctx)
	<-ctx.Done()
	log.Info(nil, "Server has shutdown")
}
