package cmd

import (
	"gitlab.com/M.darvish/funtory/api/routes"
	"gitlab.com/M.darvish/funtory/internal/app/handler"
	"gitlab.com/M.darvish/funtory/internal/database"
	"go.uber.org/dig"
	"net/http"
)

// NewDig service provider...
func NewDig() *dig.Container {
	container := dig.New()

	if err := container.Provide(http.NewServeMux); err != nil {
		panic(err)
	}
	if err := container.Provide(routes.NewAPIRouteHandler); err != nil {
		panic(err)
	}
	if err := container.Provide(database.NewDatabase); err != nil {
		panic(err)
	}
	if err := container.Provide(handler.NewHealthHandler); err != nil {
		panic(err)
	}

	return container
}
