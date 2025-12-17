package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/iJoyRide/ctc-esg/data-service/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseService struct {
	db  *sql.DB
	cfg *config.Config
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
	log.Println("[Database] Connection established")

	// Create schema after connection is established
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	if err := d.createSchema(ctx2); err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}

	return nil
}

func (d *DatabaseService) Close() error {
	if d.db == nil {
		return nil
	}
	return d.db.Close()
}
