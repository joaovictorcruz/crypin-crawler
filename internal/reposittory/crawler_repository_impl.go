package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/joaovictorcruz/crypin-crawler/internal/domain"
)

type crawlerRepositoryImpl struct {
	db *sqlx.DB
}

func NewCrawlerRepository(db *sqlx.DB) CrawlerRepository {
	return &crawlerRepositoryImpl{db: db}
}

func (r *crawlerRepositoryImpl) GetScheduledCrawlers() ([]domain.CrawlerSchedule, error) {
	var crawlers []domain.CrawlerSchedule
	err := r.db.Select(&crawlers, "SELECT * FROM CrawlerSchedule WHERE StatusId = 1")
	return crawlers, err
}

func (r *crawlerRepositoryImpl) UpdateCrawlerStatus(scheduleId int, statusId int) error {
	_, err := r.db.Exec("UPDATE CrawlerSchedule SET StatusId = ? WHERE ScheduleId = ?", statusId, scheduleId)
	return err
}

func (r *crawlerRepositoryImpl) InsertCrawlerLog(log domain.CrawlerLog) error {
	_, err := r.db.NamedExec(`INSERT INTO crawlerLogs (ScheduleId, CryptoName, StatusDownload)
	VALUES (:ScheduleId, :CryptoName, :StatusDownload)`, log)
	return err
}
