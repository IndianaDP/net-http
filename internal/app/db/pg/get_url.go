package db

import "database/sql"

func (d *db) GetURL(uuid string) (string, error) {
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
