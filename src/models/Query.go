package models

type Query struct {
	SQL            string `json:"sql"`
	TrackTotalHits bool   `json:"track_total_hits"`
	SQLMode        string `json:"sql_mode"`
}

var RequestData struct {
	SearchType string       `json:"search_type"`
	Query      QueryDetails `json:"query"`
	From       int          `json:"from"`
	MaxResults int          `json:"max_results"`
}

type QueryDetails struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}
