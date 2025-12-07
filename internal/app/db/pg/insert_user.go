package db

func (d *db) InsertUser(userID string) (int, error) {
	var id int
	query := `INSERT INTO users (user_id) VALUES ($1) RETURNING id ON CONFLICT DO NOTHING;`

	err := d.conn.QueryRow(query, userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
