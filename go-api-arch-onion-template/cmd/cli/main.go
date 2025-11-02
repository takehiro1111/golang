package main

import (
	"go-api-arch-onion-template/infrastructure/database/repository"
	"go-api-arch-onion-template/userinterface/cli/action"
	"go-api-arch-onion-template/userinterface/cli/command"
	"os"
	"sort"

	"github.com/urfave/cli/v2"

	"go-api-arch-onion-template/application/service"
	"go-api-arch-onion-template/infrastructure/database"
	"go-api-arch-onion-template/pkg/logger"
)

func main() {
	db, err := database.NewDatabaseSQLFactory(database.InstanceMySQL)
	if err != nil {
		logger.Fatal(err.Error())
	}

	albumRepository := repository.NewAlbumRepository(db)
	albumService := service.NewAlbumService(albumRepository)
	albumAction := action.NewAlbumAction(albumService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryAction := action.NewCategoryAction(categoryService)

	app := &cli.App{}
	command.SetAlbumCommand(app, albumAction)
	command.SetCategoryCommand(app, categoryAction)
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))
	err = app.Run(os.Args)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
