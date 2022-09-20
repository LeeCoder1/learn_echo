package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := try01buildEcho()

	addr := ":8081"
	if err := e.Start(addr); err != nil {
		fmt.Println(err)
	}
}

func try01buildEcho() *echo.Echo {
	e := echo.New()
	// 管理员进行鉴权
	//adminGroup := e.Group("/admin", try01MiddlewareAuth)
	//{
	//	adminGroup.GET("greet", try01Check)
	//}
	// 普通用户不鉴权
	userGroup := e.Group("/user")
	{
		// 需要在前面增加/, greet2路由找不到
		//userGroup.GET("greet2", try01Check)
		userGroup.GET("/greet2", try01Check)
		userGroup.GET("greet", try01Check)
	}

	//e.GET("/user/greet2", try01Check)

	return e
}

type try01JsonResponse struct {
	Data interface{}
}

func try01Check(c echo.Context) error {
	err := c.JSON(http.StatusOK, try01JsonResponse{"welcome home, administrator~"})
	return err
}

// middleware方法，接收一个请求处理函数，返回一个请求处理函数
func try01MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 1. 执行auth方法
		fmt.Println("执行auth方法")

		name := c.QueryParam("name")
		password := c.QueryParam("password")
		if name == "" || password == "" {
			err := fmt.Errorf("auth鉴权信息为空")
			return c.JSON(http.StatusOK, try01JsonResponse{Data: err.Error()})
		}
		if !(name == "zqb" && password == "12345678") {
			err := fmt.Errorf("auth鉴权失败")
			return c.JSON(http.StatusOK, try01JsonResponse{Data: err.Error()})
		}
		// 鉴权成功

		// 2. 执行下一个处理函数
		fmt.Println("执行next函数")
		err := next(c)
		fmt.Println("next函数执行完毕")
		fmt.Println("auth 返回")
		return err
	}
}
