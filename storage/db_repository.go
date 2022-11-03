package storage

import "github.com/TemurMannonov/query_analyzer/models"

type DBRepositoryI interface {
	GetList(request *models.GetQueriesRequest) (*models.GetQueriesResponse, error)
}
