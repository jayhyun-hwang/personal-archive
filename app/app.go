package app

import (
	"github.com/jaeyo/personal-archive/controllers"
	"github.com/jaeyo/personal-archive/internal"
	"github.com/jaeyo/personal-archive/repositories"
	"github.com/jaeyo/personal-archive/services"
)

type App struct {
	controllers.ControllerModule
	services.ServiceModule
	repositories.RepositoryModule
	internal.InternalModule
}

func New() *App {
	app := &App{}

	app.ControllerModule = controllers.NewControllerModule(app)
	app.ServiceModule = services.NewServiceModule(app)
	app.RepositoryModule = repositories.NewRepositoryModule(app)
	app.InternalModule = internal.NewInternalModule()

	return app
}
