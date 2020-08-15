package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	log "ms-soccer/service/shared/log/app"
)

func ConnectionString(database string) string {
	instance := viper.GetString(fmt.Sprintf("application.resources.database.%s.instance", database))
	port := viper.GetInt(fmt.Sprintf("application.resources.database.%s.port", database))
	dbname := viper.GetString(fmt.Sprintf("application.resources.database.%s.dbname", database))
	username := viper.GetString(fmt.Sprintf("application.resources.database.%s.username", database))
	password := viper.GetString(fmt.Sprintf("application.resources.database.%s.password", database))
	options := viper.GetString(fmt.Sprintf("application.resources.database.%s.options", database))

	if len(instance) == 0 {
		log.Panic(nil, "Database instance is required")
	}

	if port == 0 {
		log.Panic(nil, "Database port is required")
	}

	if len(dbname) == 0 {
		log.Panic(nil, "Database name is required")
	}

	if len(username) == 0 {
		log.Panic(nil, "Database username is required")
	}

	if len(password) == 0 {
		log.Panic(nil, "Database password is required")
	}

	creds := fmt.Sprintf("%s:%s", username, password)
	connectionString := fmt.Sprintf("%s@(%s:%d)/%s", creds, instance, port, dbname)

	if len(options) > 0 {
		connectionString = fmt.Sprintf("%s?%s", connectionString, options)
	}

	return connectionString
}
