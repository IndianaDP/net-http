package db

import (
	"database/sql"

	"github.com/IndianaDP/net-http/internal/app/db/models"
	_ "github.com/lib/pq"
)

type db struct {
	conn *sql.DB
}
type DB interface {
	Close()
	InsertURL(originalURL string, uuid string) error
	GetURL(uuid string) (string, error)
	URLSCount() (string, error)
	IsUrlExists(originalURL string) (bool, error)
	GetURLs() ([]models.Store, error)
}

func NewDB(connStr string) (DB, error) {

	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	result := &db{conn: conn}

	if err = result.createStoreTable(); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *db) createStoreTable() error {
	query := `CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		original_url TEXT NOT NULL,
		uuid TEXT NOT NULL UNIQUE,
		created_at TIMESTAMPTZ DEFAULT now()
	);`

	_, err := d.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *db) Close() {
	d.conn.Close()
}
