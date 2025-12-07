package db

func (d *db) InsertUserUrl(userID int, urlID int) error {
	var id string
	query := `INSERT INTO user_urls (user_id, url_id) VALUES ($1, $2) ON CONFLICT (user_id, url_id) DO NOTHING;`

	err := d.conn.QueryRow(query, userID, urlID).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
