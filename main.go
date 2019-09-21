package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/liujiangwei/econf/api/etcdtree"
	"log"
	"net/http"
)
func main() {
	e := echo.New()
	e.Use(middleware.CORS())


	e.GET("/test", func(context echo.Context) error {
		return context.String(http.StatusOK, "this is test")
	})
	e.GET("/etcd/search", etcdtree.List)

	log.Println(e.Start(":8081"))
}
