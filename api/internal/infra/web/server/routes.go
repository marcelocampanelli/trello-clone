package server

import "github.com/go-chi/chi/v5"

func (server *Server) UserRoutes(router chi.Router) {
	handler := InitializeUserHandler(server.Client)

	router.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/", handler.Create)
		r.Put("/{id}", handler.Update)
	})
}
