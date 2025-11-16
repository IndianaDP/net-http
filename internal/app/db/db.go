package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type db struct {
	conn *sql.DB
}

type store struct {
	UUID string
	URL  string
}

type DB interface {
	Close()
	InsertURLIntoDB(originalURL string, uuid string) (string, error)
	GetURLFromDB(uuid string) (string, error)
	StoreCount() (string, error)
	IsUrlExists(originalURL string) (string, error)
	StoredUrls() ([]store, error)
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

func (d *db) Close() {
	d.conn.Close()
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

func (d *db) InsertURLIntoDB(originalURL string, uuid string) (string, error) {
	var id string
	query := `INSERT INTO urls (original_url, uuid) VALUES ($1, $2) RETURNING id;`

	err := d.conn.QueryRow(query, originalURL, uuid).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (d *db) GetURLFromDB(uuid string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE uuid = $1;`

	err := d.conn.QueryRow(query, uuid).Scan(&originalURL)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return originalURL, nil
}

func (d *db) StoreCount() (string, error) {
	var count string
	query := `SELECT COUNT(*) FROM urls;`

	err := d.conn.QueryRow(query).Scan(&count)
	if err != nil {
		return "", err
	}

	return count, nil
}

func (d *db) IsUrlExists(originalURL string) (string, error) {
	var id string
	query := `SELECT uuid FROM urls WHERE original_url = $1;`

	err := d.conn.QueryRow(query, originalURL).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return id, nil
}

func (d *db) StoredUrls() ([]store, error) {

	var results []store
	query := `SELECT uuid, original_url FROM urls;`

	rows, err := d.conn.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s store
		if err := rows.Scan(&s.UUID, &s.URL); err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	return results, nil
}
