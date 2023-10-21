package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Client *mongo.Client
}

func NewServer(client *mongo.Client) *Server {
	return &Server{Client: client}
}

func (server *Server) Start() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Group(func(r chi.Router) {
		server.UserRoutes(r)
	})

	return router
}
