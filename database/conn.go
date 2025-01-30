package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	db *pgxpool.Pool
}

var (
	pgDBInstance *PostgresDB
	pgOnce sync.Once
)

// Create a new connection pool and return the same.
func NewPgDB(ctx context.Context, connString string) (*PostgresDB, error) {
	var initErr error 
	pgOnce.Do(func () {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			initErr = fmt.Errorf("unabe to create connection pool: %w", err)
			return
		}
		pgDBInstance = &PostgresDB{db: db}
	})

	if initErr != nil {
		return nil, initErr
	}
	return pgDBInstance, nil
}

func (pgdb *PostgresDB) ExecuteSQLFile(ctx context.Context,filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	if _, err := pgdb.db.Exec(ctx, string(content)); err != nil {
		return "", err
	}
	return string(content), nil
}

func (pgdb *PostgresDB) Ping(ctx context.Context) error {
	return pgdb.db.Ping(ctx)
}

func (pgdb *PostgresDB) Close() {
	pgdb.db.Close()
}