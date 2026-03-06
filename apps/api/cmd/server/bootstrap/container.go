package bootstrap

import (
	"mytodo/apps/api/pkg/cache/redis"
	"mytodo/apps/api/pkg/database/postgres"
)

type Container struct {
	DB    *postgres.DB
	Redis *redis.Client
	Log   Logger
}

func NewContainer(logger Logger) (*Container, error) {

	db, err := postgres.NewPostgresConnection()
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.NewRedisClient()
	if err != nil {
		return nil, err
	}

	return &Container{
		DB:    db,
		Redis: redisClient,
		Log:   logger,
	}, nil
}
