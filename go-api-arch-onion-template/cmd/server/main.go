package main

import (
	"go-api-arch-onion-template/infrastructure/database"
	"go-api-arch-onion-template/infrastructure/web"
	"go-api-arch-onion-template/pkg/logger"
)

func main() {
	db, err := database.NewDatabaseSQLFactory(database.InstanceMySQL)
	if err != nil {
		logger.Fatal(err.Error())
	}
	web.StartWebServer(web.InstanceGin, db)
}
