package db

import "database/sql"

func (d *db) IsUrlExists(uuid string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE uuid = $1);`

	err := d.conn.QueryRow(query, uuid).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}
