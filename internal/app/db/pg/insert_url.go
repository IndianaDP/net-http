package db

func (d *db) InsertURL(originalURL string, uuid string) error {
	var id string
	query := `INSERT INTO urls (original_url, uuid) VALUES ($1, $2) RETURNING id;`

	err := d.conn.QueryRow(query, originalURL, uuid).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
