package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func (rd *Redis) Connection() error {

	dsn := fmt.Sprintf("%s:%s", rd.Host, rd.Port)

	client := redis.NewClient(
		&redis.Options{
			Addr: dsn,
			DB:   rd.Database,
		},
	)

	ctx := context.Background()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return err
	}

	rd.Client = client

	return nil
}
