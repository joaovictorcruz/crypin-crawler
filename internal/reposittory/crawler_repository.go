package repository

import "github.com/joaovictorcruz/crypin-crawler/internal/domain"

type CrawlerRepository interface {
	GetScheduledCrawlers() ([]domain.CrawlerSchedule, error)
	UpdateCrawlerStatus(scheduleId int, statusId int) error
	InsertCrawlerLog(log domain.CrawlerLog) error
}