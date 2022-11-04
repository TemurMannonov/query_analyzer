package api

import (
	"errors"
	"strconv"
	"strings"

	"github.com/TemurMannonov/query_analyzer/api/models"
	dbModels "github.com/TemurMannonov/query_analyzer/storage/models"
	"github.com/gofiber/fiber/v2"
)

// @Router /queries [get]
// @Summary Get queries
// @Description API for get  queries
// @Tags query
// @Accept json
// @Produce json
// @Param find query models.GetQueriesRequest false "filters"
// @Success 200 {object} models.GetQueriesResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *Server) GetQueries(c *fiber.Ctx) error {
	req, err := validateQueriesRequest(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	result, err := h.storage.GetList(&dbModels.GetQueriesParams{
		Limit:      req.Limit,
		Page:       req.Page,
		SortByTime: req.SortByTime,
		Type:       req.Type,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	response := parseQueriesResult(result)

	return c.Status(fiber.StatusOK).JSON(response)
}

func parseQueriesResult(result *dbModels.GetQueriesResult) *models.GetQueriesResponse {
	response := models.GetQueriesResponse{
		Count:   result.Count,
		Queries: make([]*models.Query, 0),
	}

	for _, query := range result.Queries {
		response.Queries = append(response.Queries, &models.Query{
			QueryID:       query.QueryID,
			Query:         query.Query,
			Calls:         query.Calls,
			TotalExecTime: query.TotalExecTime,
			MinExecTime:   query.MinExecTime,
			MaxExecTime:   query.MaxExecTime,
			MeanExecTime:  query.MeanExecTime,
		})
	}

	return &response
}

func validateQueriesRequest(c *fiber.Ctx) (*models.GetQueriesRequest, error) {
	var (
		err        error
		limit      int64  = 10     // default value for limit
		page       int64  = 1      // default value for page
		sortByTime string = "desc" // default value for sort by time
	)

	if c.Query("limit") != "" {
		limit, err = strconv.ParseInt(c.Query("limit"), 10, 32)
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.ParseInt(c.Query("page"), 10, 32)
		if err != nil {
			return nil, err
		}
	}

	if c.Query("sort_by_time") != "" {
		sortByTime = strings.ToLower(c.Query("sort_by_time"))

		if sortByTime != "asc" && sortByTime != "desc" {
			return nil, errors.New("incorrect value for sorb_by_time param")
		}
	}

	queryType := strings.ToLower(c.Query("type"))
	if queryType != "" && queryType != "select" && queryType != "insert" &&
		queryType != "update" && queryType != "delete" {
		return nil, errors.New("incorrect value for type param")
	}

	return &models.GetQueriesRequest{
		Page:       int32(page),
		Limit:      int32(limit),
		SortByTime: sortByTime,
		Type:       queryType,
	}, nil
}
