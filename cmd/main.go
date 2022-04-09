package main

import (
	"go-redis-crypto/internal/infraestructure"
	"go-redis-crypto/internal/services"
)

func main() {
	redis := &infraestructure.Redis{}
	redis.Init("localhost:6379")
	cryptoService := &services.CryptoService{RedisConn: redis.Conn}
	r := initRouter(redis, cryptoService)
	r.Debug = true
	r.Logger.Fatal(r.Start(":1525"))

}
