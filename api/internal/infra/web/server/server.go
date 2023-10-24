package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

var TokenAuth *jwtauth.JWTAuth

type Server struct {
	Client *mongo.Client
}

func NewServer(client *mongo.Client) *Server {
	return &Server{Client: client}
}

func (server *Server) Start() chi.Router {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:  []string{"*"},
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:          300,
	})

	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(corsMiddleware.Handler)
	router.Use(middleware.WithValue("jwt", TokenAuth))
	router.Use(middleware.WithValue("JwtExperesIn", 3000))

	router.Group(func(r chi.Router) {
		server.UserRoutes(r)
		server.BoardRoutes(r)
		server.ListRoutes(r)
		server.CardRoutes(r)
	})

	return router
}
