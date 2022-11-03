package api

import (
	"github.com/TemurMannonov/query_analyzer/config"
	"github.com/TemurMannonov/query_analyzer/models"
	"github.com/TemurMannonov/query_analyzer/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/TemurMannonov/query_analyzer/api/docs" // for swagger
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type Server struct {
	cfg     *config.Config
	storage storage.DBRepositoryI
	router  *fiber.App
}

// @title Swagger Database Query API
// @version 1.0
// @description This is a api documentation for getting database queries.
func NewServer(cfg *config.Config, strg storage.DBRepositoryI) *Server {
	server := &Server{
		cfg:     cfg,
		storage: strg,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Get("/queries", server.GetQueries)

	server.router = app
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Listen(address)
}

func errorResponse(err error) models.ResponseError {
	return models.ResponseError{Error: err.Error()}
}
