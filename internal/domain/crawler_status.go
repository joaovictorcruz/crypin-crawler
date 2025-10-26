package domain

type CrawlerStatus struct {
	StatusId    int    `db:"StatusId"`
	StatusName  string `db:"StatusName"`
	Description string `db:"Description"`
}