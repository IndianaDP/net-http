package db

func (d *db) DeleteUserUrl(userID string, urlID string) error {
	var id string
	query := `DELETE FROM user_urls WHERE user_id = $1 AND url_id = $2;`

	err := d.conn.QueryRow(query, userID, urlID).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
