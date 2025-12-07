package redis

func (r *rdb) Set(key string, value string) (string, error) {
	if err := r.cli.Set(ctx, key, value, 0).Err(); err != nil {
		return "", err
	}
	return "", nil
}
