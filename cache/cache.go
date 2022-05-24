package cache

import (
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
)

var ctx = context.Background()

type Cache struct {
	Conn *redis.Client
}

func New(add, pass string) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	return &Cache{
		Conn: rdb,
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	if value == nil {
		return nil
	}

	data, err := jsoniter.Marshal(value)
	if err != nil {
		return err
	}

	err = c.Conn.Set(ctx, key, data, time.Second*30).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	val, err := c.Conn.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, redis.Nil
	}
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func GetCache(c *Cache) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := c.Get(ctx.Request.RequestURI)
		if errors.Is(redis.Nil, err) {
			ctx.Header("cache", "not found")
			ctx.Next()
			return
		}

		if err != nil {
			ctx.Header("cache", "is not working")
			ctx.Next()
			return
		}
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.Writer.Write(data)
		ctx.Abort()
	}
}
