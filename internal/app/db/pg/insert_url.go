package db

func (d *db) InsertURL(originalURL string, uuid string) (int, error) {
	var id int
	query := `INSERT INTO urls (original_url, uuid) VALUES ($1, $2) RETURNING id;`

	err := d.conn.QueryRow(query, originalURL, uuid).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
