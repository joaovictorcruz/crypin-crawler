package domain

import "time"

type CrawlerSchedule struct {
	ScheduleId    int       `db:"ScheduleId"`
	StatusId      int       `db:"StatusId"`
	BaseUrl       string    `db:"BaseUrl"`
	Parameters    []byte    `db:"Parameters"` 
	LastExecution time.Time `db:"LastExecution"`
	NextExecution time.Time `db:"NextExecution"`
	CrawlerName   string    `db:"CrawlerName"`
}
