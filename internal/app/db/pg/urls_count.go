package db

func (d *db) URLSCount() (string, error) {
	var count string
	query := `SELECT COUNT(*) FROM urls;`

	err := d.conn.QueryRow(query).Scan(&count)
	if err != nil {
		return "", err
	}

	return count, nil
}
