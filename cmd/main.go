package main

import "go-redis-crypto/internal/infraestructure"

func main() {
	redis := &infraestructure.Redis{}
	redis.Init("localhost:6379")
	r := initRouter(redis)
	r.Debug = true
	r.Logger.Fatal(r.Start(":1525"))

}
