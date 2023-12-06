package auth

import (
	"context"

	"github.com/incompetent-developer/restgo-struct/redis"
)

// New is create authtication client with redis
func (client *Client) New(rd redis.Redis) error {
	ctx := context.Background()

	/*
		Ping to redis
	*/
	if _, err := rd.Client.Ping(ctx).Result(); err != nil {
		return err
	}

	client.Redis = rd

	return nil
}
