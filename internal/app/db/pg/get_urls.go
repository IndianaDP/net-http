package db

import (
	"database/sql"

	"github.com/IndianaDP/net-http/internal/app/db/models"
	_ "github.com/lib/pq"
)

func (d *db) GetURLs() ([]models.Store, error) {

	var results []models.Store
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
		var s models.Store
		if err := rows.Scan(&s.UUID, &s.URL); err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	return results, nil
}
