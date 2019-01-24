package main

import (
	"net/http"
	"time"

	"github.com/chneau/limiter"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	var limitEvent limiter.Limiter = limiter.New(10)
	e.GET("/high", func(c echo.Context) error {
		limitEvent.Execute(func() {
			for index := 0; index < 100; index++ {
				time.Sleep(1 * time.Second)
			}
		})
		return c.String(http.StatusOK, "success")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
