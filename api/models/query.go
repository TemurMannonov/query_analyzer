package models

type Query struct {
	QueryID       int64   `json:"queryid"`
	Query         string  `json:"query"`
	Calls         int32   `json:"calls"`
	TotalExecTime float64 `json:"total_exec_time"`
	MinExecTime   float64 `json:"min_exec_time"`
	MaxExecTime   float64 `json:"max_exec_time"`
	MeanExecTime  float64 `json:"mean_exec_time"`
}

type GetQueriesRequest struct {
	Limit      int32  `json:"limit" default:"10"`
	Page       int32  `json:"page" default:"1"`
	SortByTime string `json:"sort_by_time" enums:"asc,desc" default:"desc"`
	Type       string `json:"type" enums:"select,insert,update,delete"`
}

type GetQueriesResponse struct {
	Queries []*Query `json:"queries"`
	Count   int32    `json:"count"`
}
