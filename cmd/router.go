package main

import (
	"go-redis-crypto/internal/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initRouter() *echo.Echo {
	r := echo.New()
	initRoutes(r)

	/* MIDDLEWARES */
	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))
	//r.Use(middleware.CORS()) //For Swagger
	return r
}

func initRoutes(r *echo.Echo) {
	r.GET("/test", handlers.Test)
	r.GET("/cryptop10", handlers.Top10)
	r.GET("/fiat-cur-price/:id", handlers.FiatCurPrice)
}
