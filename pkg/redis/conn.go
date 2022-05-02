package redis

import (
	"github.com/rueian/rueidis"

	"rest-company/config"
)

// NewRedisClient Returns new redis client
func NewRedisClient(cfg *config.Config) (rueidis.Client, error) {
	return rueidis.NewClient(rueidis.ClientOption{
		Password:         cfg.Redis.Password,
		InitAddress:      cfg.Redis.Hosts,
		SelectDB:         cfg.Redis.DB,
		BlockingPoolSize: cfg.Redis.PoolSize,
	})
}
