package redis

func (r *rdb) Get(key string) (string, error) {
	val, err := r.cli.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
