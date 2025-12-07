package redis

import (
	"context"

	"github.com/IndianaDP/net-http/internal/app/db/models"
	"github.com/redis/go-redis/v9"
)

type rdb struct {
	cli *redis.Client
}

type RDB interface {
	Close()
	Set(key string, value string) error
	Get(key string) (string, error)
	Count() (string, error)
	Exists(key string) (bool, error)
	List() ([]models.Store, error)
}

var ctx = context.Background()

func NewRedisClient(dbNo int, addr string, pass string) (RDB, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       dbNo,
	})

	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &rdb{cli: cli}, nil
}

func (r *rdb) Close() {
	r.cli.Close()
}
