package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type JsonResponse struct {
	Val string
}

func main() {
	// 1. 注册路由
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &JsonResponse{Val: "index page response"})
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &JsonResponse{Val: "get a pong"})
	})

	// 2. 在一个端口启动监听服务
	if err := e.Start(":8081"); err != nil {
		fmt.Println(err)
	}

	// 3. 接口调用测试
}
