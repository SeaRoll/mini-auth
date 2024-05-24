package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type service struct {
	conn *sql.DB
}

func NewService() *service {
	db, err := sql.Open("sqlite3", "./file.db")
	if err != nil {
		panic("Failed to open sqlite3")
	}
	migrate(db)

	return &service{
		conn: db,
	}
}

func (s *service) Close() error {
	err := s.conn.Close()
	if err != nil {
		log.Println("Failed to close db")
		return err
	}
	log.Println("Closed db")
	return nil
}

func (s *service) WithTX(ctx context.Context, fn func(q Querier) error) error {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	querierInTx := New(tx)
	err = fn(querierInTx)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
