package infraestructure

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Conn *redis.Client
}

func (r *Redis) Init(addr string) {
	r.Conn = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	pong, err := r.Conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis (", addr, ") - Error: ", err.Error())
	}
	if pong == "PONG" {
		log.Println("Redis LIVE on", addr)
	}
}
