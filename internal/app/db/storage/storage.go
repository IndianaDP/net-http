package storage

import (
	"fmt"

	"github.com/IndianaDP/net-http/internal/app/config"
	"github.com/IndianaDP/net-http/internal/app/db/models"
	db "github.com/IndianaDP/net-http/internal/app/db/pg"
	rdb "github.com/IndianaDP/net-http/internal/app/db/redis"
)

const (
	PG    = "pg"
	REDIS = "redis"
)

type storage struct {
	pgClient    *db.DB
	redisClient *rdb.RDB
}

type Storage interface {
	InsertUrl(url string, uuid string) (int, error)
	InsertUser(userID string) (int, error)
	InsertUserUrl(userID int, urlID int) error
	Get(uuid string) (string, error)
	Exists(userId string, url string) (bool, error)
	Count() (string, error)
	List() ([]models.Store, error)
	DeleteUrl(urlID string) error
	DeleteUserUrl(userID string, urlID string) error
}

func NewStorage(cfg *config.Values) (Storage, error) {
	switch cfg.StorageType {
	case PG:
		pgClient, _ := db.NewDB(config.BuildDSN(cfg))
		return &storage{
			pgClient: &pgClient,
		}, nil
	case REDIS:
		redisClient, _ := rdb.NewRedisClient(0, cfg.RedisConn.RedisAddr, cfg.RedisConn.RedisPass)
		return &storage{
			redisClient: &redisClient,
		}, nil

	default:
		return nil, fmt.Errorf("unsupported storage type: %s", cfg.StorageType)
	}
}

func (s *storage) InsertUrl(url string, uuid string) (int, error) {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).InsertURL(url, uuid)
	// case s.redisClient != nil:
	// 	return (*s.redisClient).Set(uuid, url)
	default:
		return 0, fmt.Errorf("no storage client available")
	}
}

func (s *storage) InsertUserUrl(userID int, urlID int) error {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).InsertUserUrl(userID, urlID)
	// case s.redisClient != nil:
	// 	return (*s.redisClient).Set(uuid, url)
	default:
		return fmt.Errorf("no storage client available")
	}
}

func (s *storage) InsertUser(userID string) (int, error) {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).InsertUser(userID)
	// case s.redisClient != nil:
	// 	return (*s.redisClient).Set(uuid, url)
	default:
		return 0, fmt.Errorf("no storage client available")
	}
}

func (s *storage) Get(uuid string) (string, error) {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).GetURL(uuid)
	case s.redisClient != nil:
		return (*s.redisClient).Get(uuid)
	default:
		return "", fmt.Errorf("no storage client available")
	}
}

func (s *storage) Exists(userId string, uuid string) (bool, error) {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).IsUrlExists(userId, uuid)
	// case s.redisClient != nil:
	// 	return (*s.redisClient).Exists(userId, uuid)
	default:
		return false, fmt.Errorf("no storage client available")
	}
}

func (s *storage) Count() (string, error) {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).URLSCount()
	case s.redisClient != nil:
		return (*s.redisClient).Count()
	default:
		return "", fmt.Errorf("no storage client available")
	}
}

func (s *storage) List() ([]models.Store, error) {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).GetURLs()
	case s.redisClient != nil:
		return (*s.redisClient).List()
	default:
		return nil, fmt.Errorf("no storage client available")
	}
}

func (s *storage) DeleteUrl(urlID string) error {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).DeleteUrl(urlID)
	// case s.redisClient != nil:
	// 	return (*s.redisClient).List()
	default:
		return fmt.Errorf("no storage client available")
	}
}

func (s *storage) DeleteUserUrl(userID, urlID string) error {
	switch {
	case s.pgClient != nil:
		return (*s.pgClient).DeleteUserUrl(userID, urlID)
	// case s.redisClient != nil:
	// 	return (*s.redisClient).List()
	default:
		return fmt.Errorf("no storage client available")
	}
}
