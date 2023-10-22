package server

import "github.com/go-chi/chi/v5"

func (server *Server) UserRoutes(router chi.Router) {
	handler := InitializeUserHandler(server.Client)

	router.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/", handler.Create)
		r.Put("/{id}", handler.Update)
	})
}

func (server *Server) BoardRoutes(router chi.Router) {
	handler := InitializeBoardHandler(server.Client)

	router.Route("/api/v1/boards", func(r chi.Router) {
		r.Get("/user/{userID}", handler.FindAll)
		r.Post("/", handler.Create)
		r.Get("/{id}", handler.FindByID)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
}

func (server *Server) ListRoutes(router chi.Router) {
	handler := InitializeListHandler(server.Client)

	router.Route("/api/v1/lists", func(r chi.Router) {
		r.Get("/board/{boardID}", handler.FindAll)
		r.Post("/", handler.Create)
		r.Get("/{id}", handler.FindByID)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
}

func (server *Server) CardRoutes(router chi.Router) {
	handler := InitializeCardHandler(server.Client)

	router.Route("/api/v1/cards", func(r chi.Router) {
		r.Get("/list/{listID}", handler.FindAllByList)
		r.Post("/", handler.Create)
		r.Get("/{id}", handler.FindByID)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
}
