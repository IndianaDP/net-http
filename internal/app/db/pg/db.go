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
	InsertURL(originalURL string, uuid string) (int, error)
	InsertUser(userID string) (int, error)
	InsertUserUrl(userID int, urlID int) error
	GetURL(uuid string) (string, error)
	URLSCount() (string, error)
	IsUrlExists(userId string, originalURL string) (bool, error)
	GetURLs() ([]models.Store, error)
	DeleteUrl(urlID string) error
	DeleteUserUrl(userID, urlID string) error
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

	createTables := []func() error{
		result.createStoreTable,
		result.createUsersTable,
		result.createUserUrlsTable,
	}

	for _, create := range createTables {
		if err := create(); err != nil {
			return nil, err
		}
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

func (d *db) createUsersTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,        
    user_id TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now()
	);`

	_, err := d.conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *db) createUserUrlsTable() error {
	query := `CREATE TABLE IF NOT EXISTS user_urls (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    url_id INT NOT NULL REFERENCES urls(id) ON DELETE CASCADE,
    UNIQUE(user_id, url_id)
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
