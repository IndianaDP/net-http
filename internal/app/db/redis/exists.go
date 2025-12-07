package redis

func (r *rdb) Exists(key string) (bool, error) {
	count, err := r.cli.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
