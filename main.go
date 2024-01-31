package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	r := gin.Default()
	client := redis.NewClient(&redis.Options{
		Addr:     "master.tuki-cache.29m23s.sae1.cache.amazonaws.com:6379",
		Password: "e5215ccb-a781-400c-a1e8-4cf2c543560d",
		DB:       0,
	})
	r.GET("/test", func(c *gin.Context) {
		ctx := context.Background()

		err := client.Set(ctx, "foo", "bar", 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := client.Get(ctx, "foo").Result()
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"Redis": val,
		})
	})
	r.Run()
}
