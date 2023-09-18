package routes

import (
	"github.com/go-chi/chi"
	"gitlab.com/M.darvish/funtory/internal/app/handler"
)

type APIRouteHandler struct {
	Router *chi.Mux
}

func NewAPIRouteHandler(healthHandler *handler.HealthHandler) *APIRouteHandler {

	apiRouterHandler := &APIRouteHandler{
		Router: chi.NewRouter(),
	}

	apiRouterHandler.Router.Get("/health", healthHandler.Health)
	return apiRouterHandler
}
