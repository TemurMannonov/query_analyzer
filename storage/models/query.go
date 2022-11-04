package models

type Query struct {
	QueryID       int64
	Query         string
	Calls         int32
	TotalExecTime float64
	MinExecTime   float64
	MaxExecTime   float64
	MeanExecTime  float64
}

type GetQueriesParams struct {
	Limit      int32
	Page       int32
	SortByTime string
	Type       string
}

type GetQueriesResult struct {
	Queries []*Query
	Count   int32
}
