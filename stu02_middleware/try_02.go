package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// 测试路由级别的中间件，多个中间件是怎么执行的
func main() {
	e := try02buildEcho()

	addr := ":8081"
	if err := e.Start(addr); err != nil {
		fmt.Println(err)
	}
}

func try02buildEcho() *echo.Echo {
	e := echo.New()
	// 管理员进行鉴权
	//adminGroup := e.Group("/admin", try02MiddlewareAuth)
	//{
	//	adminGroup.GET("greet", try02Check)
	//}
	// 普通用户不鉴权
	userGroup := e.Group("/user", try02m4)
	{
		userGroup.GET("/auth-greet", try02Check, try02m1, try02m2, try02m3)
		userGroup.Use(try02m5)
		userGroup.GET("/greet", try02Check)
	}

	//e.GET("/user/greet2", try02Check)

	fmt.Println("test")

	return e
}

type try02JsonResponse struct {
	Data interface{}
}

func try02Check(c echo.Context) error {
	fmt.Println("try02Check")
	err := c.JSON(http.StatusOK, try02JsonResponse{"welcome home, administrator~"})
	return err
}

// middleware方法，接收一个请求处理函数，返回一个请求处理函数
func try02MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 1. 执行auth方法
		fmt.Println("执行auth方法")

		name := c.QueryParam("name")
		password := c.QueryParam("password")
		if name == "" || password == "" {
			err := fmt.Errorf("auth鉴权信息为空")
			return c.JSON(http.StatusOK, try02JsonResponse{Data: err.Error()})
		}
		if !(name == "zqb" && password == "12345678") {
			err := fmt.Errorf("auth鉴权失败")
			return c.JSON(http.StatusOK, try02JsonResponse{Data: err.Error()})
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

func try02m1(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("try02m1_s")
		err := h(c)
		fmt.Println("try02m1_e")
		return err
	}
}

func try02m2(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("try02m2_s")
		err := h(c)
		fmt.Println("try02m2_e")
		return err
	}
}

func try02m3(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("try02m3_s")
		err := h(c)
		fmt.Println("try02m3_e")
		return err
	}
}

func try02m4(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("try02m4_s")
		err := h(c)
		fmt.Println("try02m4_e")
		return err
	}
}

func try02m5(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("try02m5_s")
		err := h(c)
		fmt.Println("try02m5_e")
		return err
	}
}
