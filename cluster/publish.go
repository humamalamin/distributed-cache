package cluster

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/humamalamin/distributed-cache/cache"
)

type Publish struct {
	Cache *cache.Cache
	redis *redis.Client
}

var redisClient *redis.Client

func InitRedis(c *cache.Cache) Publish {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return Publish{Cache: c, redis: redisClient}
}

func (p *Publish) SetKeyWithReplication(ctx context.Context, key, value string) error {
	p.Cache.Set(key, value)

	err := p.redis.Publish(ctx, "cache-channel", fmt.Sprintf("SET %s %s", key, value)).Err()
	if err != nil {
		return err
	}

	log.Printf("Published change for key %s with value %s", key, value)
	return nil
}
