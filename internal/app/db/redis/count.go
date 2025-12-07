package redis

import "strconv"

func (r *rdb) Count() (string, error) {
	keys, err := r.cli.Keys(ctx, "*").Result()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(len(keys)), nil
}
