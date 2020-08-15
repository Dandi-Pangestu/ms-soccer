package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ms-soccer/service/shared/database/mongodb"
	"ms-soccer/service/shared/database/mysql"
	"ms-soccer/service/shared/database/postgres"
	log "ms-soccer/service/shared/log/app"
)

const (
	DevEnvMode  = "development"
	ProdEnvMode = "production"
	TestEnvMode = "testing"
)

func InitConfig() {
	viper.AddConfigPath("shared/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func InitServer() (*gin.Engine, *http.Server) {
	host := viper.GetString("application.server.host")
	port := viper.GetInt("application.server.port")
	env := viper.GetString("application.env")

	switch env {
	case TestEnvMode:
		gin.SetMode(gin.TestMode)
	case ProdEnvMode:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}

	return r, srv
}

func Start(srv *http.Server) {
	log.Info(nil, "Server is being start")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(&logrus.Fields{"error": err.Error()}, "Error while start server")
	}
}

func Shutdown(srv *http.Server, ctx context.Context) {
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(&logrus.Fields{"error": err.Error()}, "Error while shutdown server")
	}
}

func InitAppLogger(hooks ...logrus.Hook) {
	log.Init(func(l *logrus.Logger) {
		for _, hook := range hooks {
			l.AddHook(hook)
		}
	})
}

func InitMysqlDatabase(database string) *gorm.DB {
	connStr := mysql.ConnectionString(database)
	db, err := gorm.Open("mysql", connStr)

	if err != nil {
		log.Panic(&logrus.Fields{"error": err.Error()}, "Error while init mysql database")
	}

	db.LogMode(true)

	return db
}

func InitPostgresDatabase(database string) *gorm.DB {
	connStr := postgres.ConnectionString(database)
	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		log.Panic(&logrus.Fields{"error": err.Error()}, "Error while init postgres database")
	}

	db.LogMode(true)

	return db
}

func InitMongoDbDatabase(database string) *mongo.Database {
	dbname := viper.GetString(fmt.Sprintf("application.resources.database.%s.dbname", database))
	connStr := mongodb.ConnectionString(database)

	clientOptions := options.Client()
	clientOptions.ApplyURI(connStr)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Panic(&logrus.Fields{"error": err.Error()}, "Error while ini mongodb client")
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Panic(&logrus.Fields{"error": err.Error()}, "Error while connect to mongodb")
	}

	return client.Database(dbname)
}
