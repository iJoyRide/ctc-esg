package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/iJoyRide/ctc-esg/data-service/internal/config"
	db "github.com/iJoyRide/ctc-esg/data-service/internal/database/sqlc"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseService struct {
	db      *sql.DB
	cfg     *config.Config
	queries *db.Queries
}

func NewDatabaseService(configuration *config.Config) *DatabaseService {
	return &DatabaseService{cfg: configuration}
}

func (d *DatabaseService) Init() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		d.cfg.Database.User,
		d.cfg.Database.Password,
		"timescale",
		"5432",
		d.cfg.Database.DB,
	)

	dbEngine, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	dbEngine.SetMaxOpenConns(25)
	dbEngine.SetMaxIdleConns(10)
	dbEngine.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := dbEngine.PingContext(ctx); err != nil {
		return err
	}

	d.db = dbEngine
	d.queries = db.New(d.db)
	log.Println("[Database] Connection established")
	return nil
}

func (d *DatabaseService) Close() error {
	if d.db == nil {
		return nil
	}
	return d.db.Close()
}

func (d *DatabaseService) Queries() *db.Queries {
	if d.queries == nil {
		log.Println("[Database] WARNING: queries is nil (did you call Init?)")
	}
	return d.queries
}
