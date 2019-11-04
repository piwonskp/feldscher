package models

type Page struct {
	ID       int
	URL      string `json:"url"`
	Interval int    `json:"interval"`
}

type FetchedPage struct {
	ID        int
	PageID    int
	Response  *string `json:"response"`
	Duration  float64 `json:"duration"`
	CreatedAt int64   `json:"created_at"`
}
