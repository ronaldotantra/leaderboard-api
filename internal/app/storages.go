package app

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ronaldotantra/leaderboard-api/internal/logger"
)

type Storages struct {
	DB *sql.DB
}

type StoragesParams struct {
	DBConnString string
}

func SetupStorages(params StoragesParams) (*Storages, error) {
	db, err := setupDatabase(params.DBConnString)
	if err != nil {
		return nil, err
	}

	return &Storages{
		DB: db,
	}, nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Errorf("error opening database connection - %v", err)
		return nil, err
	}

	return db, nil
}
