package repository

import (
	"fmt"

	"github.com/TemurMannonov/query_analyzer/api"
	"github.com/TemurMannonov/query_analyzer/storage/models"

	"github.com/jmoiron/sqlx"
)

type queryRepo struct {
	db *sqlx.DB
}

func NewDBRepository(db *sqlx.DB) api.DBRepositoryI {
	return &queryRepo{
		db: db,
	}
}

func (b *queryRepo) GetList(req *models.GetQueriesParams) (*models.GetQueriesResult, error) {
	result := models.GetQueriesResult{
		Queries: make([]*models.Query, 0),
	}

	offset := (req.Page - 1) * req.Limit
	params := map[string]interface{}{
		"limit":      req.Limit,
		"offset":     offset,
		"query_type": req.Type,
	}

	filter := ""
	if req.Type != "" {
		filter = " WHERE query ILIKE :query_type "
	}

	orderBy := ""
	if req.SortByTime != "" {
		orderBy = fmt.Sprintf(" ORDER BY max_exec_time %s ", req.SortByTime)
	}

	count, err := b.GetListCount(filter, params)
	if err != nil {
		return nil, err
	}
	result.Count = count

	query := `
		SELECT
			queryid,
			query,
			calls,
			total_exec_time,
			min_exec_time,
			max_exec_time,
			mean_exec_time
		FROM pg_stat_statements
		` + filter + orderBy + `
		LIMIT :limit OFFSET :offset
	`

	stmt, err := b.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(params)
	if err != nil {
		return nil, err

	}
	defer rows.Close()

	for rows.Next() {
		var item models.Query

		err := rows.Scan(
			&item.QueryID,
			&item.Query,
			&item.Calls,
			&item.TotalExecTime,
			&item.MinExecTime,
			&item.MaxExecTime,
			&item.MeanExecTime,
		)
		if err != nil {
			return nil, err
		}

		result.Queries = append(result.Queries, &item)
	}

	return &result, nil
}

func (b *queryRepo) GetListCount(filter string, params map[string]interface{}) (int32, error) {
	var count int32
	queryCount := `SELECT count(1) FROM pg_stat_statements ` + filter

	stmtCount, err := b.db.PrepareNamed(queryCount)
	if err != nil {
		return 0, err
	}

	row := stmtCount.QueryRow(params)
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
