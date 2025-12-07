package redis

import (
	"github.com/IndianaDP/net-http/internal/app/db/models"
)

func (r *rdb) List() ([]models.Store, error) {
	var results []models.Store
	keys, err := r.cli.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		val, err := r.cli.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		results = append(results, models.Store{
			UUID: key,
			URL:  val,
		})
	}
	return results, nil
}
