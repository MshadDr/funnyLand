package routes

import (
	"github.com/go-chi/chi"
	"gitlab.com/M.darvish/funtory/internal/app/handler"
	"gitlab.com/M.darvish/funtory/internal/app/middleware"
	"gitlab.com/M.darvish/funtory/internal/model/repository"
)

type APIRouteHandler struct {
	Router *chi.Mux
}

func NewAPIRouteHandler(healthHandler *handler.HealthHandler,
	userRepo repository.IUser,
	userHandler handler.UserHandler) *APIRouteHandler {

	apiRouterHandler := &APIRouteHandler{
		Router: chi.NewRouter(),
	}

	apiRouterHandler.Router.Get("/health", healthHandler.Health)

	apiRouterHandler.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(users chi.Router) {
			users.Post("/login", userHandler.Login)
			users.Post("/register", userHandler.Register)
			users.Get("/{id}}", middleware.ValidateJwtAuthToken(userRepo, userHandler.ConnectAccount))
		})
	})
	return apiRouterHandler
}
