package database

import (
	"context"
	"fmt"
	"go-skeleton/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	logger *logger.Logger
	Cache  *redis.Client
	Host   string
	Port   string
}

func NewRedis(
	l *logger.Logger,
	Host string,
	Port string,
) *Redis {
	return &Redis{
		logger: l,
		Host:   Host,
		Port:   Port,
	}
}

func (m *Redis) Boot() {
	dsn := "%s:%s"
	dsn = fmt.Sprintf(
		dsn,
		m.Host,
		m.Port,
	)

	redis := redis.NewClient(
		&redis.Options{
			Addr: dsn,
		},
	)

	ctx := context.Background()

	_, err := redis.Ping(ctx).Result()

	if err != nil {
		m.logger.Critical(err)
	}

	m.Cache = redis
}
