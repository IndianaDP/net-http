package db

func (d *db) IsUrlExists(userId string, uuid string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(SELECT 1 FROM urls u JOIN user_urls us ON u.id = us.url_id WHERE us.user_id = $1 AND u.uuid = $2);`

	err := d.conn.QueryRow(query, userId, uuid).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
