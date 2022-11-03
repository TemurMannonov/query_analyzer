package api

import (
	"strconv"

	"github.com/TemurMannonov/query_analyzer/models"
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
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	page, err := strconv.ParseInt(c.Query("page"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	resp, err := h.storage.GetList(&models.GetQueriesRequest{
		Limit:      int32(limit),
		Page:       int32(page),
		SortByTime: c.Query("sort_by_time"),
		Type:       c.Query("type"),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return c.JSON(resp)
}

func validateQueriesRequest(c *fiber.Ctx) error {
	return nil
}
