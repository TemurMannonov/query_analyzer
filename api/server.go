package api

import (
	"github.com/TemurMannonov/query_analyzer/api/models"
	"github.com/TemurMannonov/query_analyzer/config"
	dbModels "github.com/TemurMannonov/query_analyzer/storage/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/TemurMannonov/query_analyzer/api/docs" // for swagger
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type DBRepositoryI interface {
	GetList(request *dbModels.GetQueriesParams) (*dbModels.GetQueriesResult, error)
}

type Server struct {
	cfg     *config.Config
	storage DBRepositoryI
	Router  *fiber.App
}

// @title Swagger Database Query API
// @version 1.0
// @description This is a api documentation for getting database queries.
func NewServer(cfg *config.Config, strg DBRepositoryI) *Server {
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

	server.Router = app
}

func errorResponse(err error) models.ResponseError {
	return models.ResponseError{Error: err.Error()}
}
