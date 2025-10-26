package domain

import "time"

type CrawlerLog struct {
	LogId          int       `db:"LogId"`
	ScheduleId     int       `db:"ScheduleId"`
	CryptoName     string    `db:"CryptoName"`
	StatusDownload bool      `db:"StatusDownload"`
	CreatedAt      time.Time `db:"CreatedAt"`
}