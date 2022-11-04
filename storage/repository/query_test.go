package repository

import (
	"testing"

	"github.com/TemurMannonov/query_analyzer/storage/models"
	"github.com/stretchr/testify/assert"
)

func TestQuery_GetList(t *testing.T) {
	testCases := []struct {
		name  string
		param *models.GetQueriesParams
	}{
		{
			name: "without filter and sort",
			param: &models.GetQueriesParams{
				Limit: 10,
				Page:  1,
			},
		},
		{
			name: "sort by time desc",
			param: &models.GetQueriesParams{
				Limit:      10,
				Page:       1,
				SortByTime: "desc",
			},
		},
		{
			name: "filter by type",
			param: &models.GetQueriesParams{
				Limit: 10,
				Page:  1,
				Type:  "select",
			},
		},
		{
			name: "filter and sorting",
			param: &models.GetQueriesParams{
				Limit:      10,
				Page:       1,
				SortByTime: "asc",
				Type:       "select",
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			res, err := strg.GetList(tc.param)
			assert.NoError(t, err)
			assert.NotEmpty(t, res)
		})
	}
}
