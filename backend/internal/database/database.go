package database

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime time.Duration) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", addr)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
