package db

func (d *db) DeleteUrl(urlID string) error {
	var id string
	query := `DELETE FROM urls WHERE id = $1;`

	err := d.conn.QueryRow(query, urlID).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
