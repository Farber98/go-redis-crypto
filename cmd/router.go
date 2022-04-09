package main

import (
	"go-redis-crypto/internal/handlers"
	"go-redis-crypto/internal/infraestructure"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initRouter(redis *infraestructure.Redis) *echo.Echo {
	r := echo.New()
	initRoutes(r, redis)

	/* MIDDLEWARES */
	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))
	//r.Use(middleware.CORS()) //For Swagger
	return r
}

func initRoutes(r *echo.Echo, redis *infraestructure.Redis) {

	cryptoHandler := &handlers.CryptoHandler{Redis: redis}
	r.GET("/test", cryptoHandler.Test)
	r.GET("/cryptop10", cryptoHandler.Top10)
	r.GET("/fiat-cur-price/:id", cryptoHandler.FiatCurPrice)
	r.GET("/trending24h", cryptoHandler.Trending24h)
}
