package infra

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"

	"github.com/joaovictorcruz/crypin-crawler/internal/config"
)

func NewDatabase(cfg config.DatabaseConfig) *sqlx.DB {
	db, err := sqlx.Connect("mysql", cfg.DSN())
	if err != nil {
		log.Fatalf("Falha ao conectar com o banco: %v", err)
	}
	log.Println("Conexão com banco feita com sucesso!!")
	return db
}

var Module = fx.Provide(
	config.LoadDatabaseConfig,
	NewDatabase,
)
