package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginAuth(c *gin.Context) {
	fmt.Println("我是登录保护中间件,login")
}

func main() {
	r := gin.Default()
	// 全局调用
	//r.Use(LoginAuth)

	r.GET("ping", func(c *gin.Context) {
		fmt.Println("我是ping")
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	r.POST("login", func(c *gin.Context) {
		fmt.Println("我是login")
		c.JSON(200, gin.H{
			"data": "token",
		})
	})
	r.POST("register", func(c *gin.Context) {

	})

	// 登录保护
	r.GET("user/:id", func(c *gin.Context) {
		fmt.Println("我是用户详情接口")
		// 判断是否登录
		c.JSON(200, gin.H{
			"msg": "获取用户详情接口，需要登录保护",
		})
	})

	r.PUT("user/:id", LoginAuth, func(c *gin.Context) {
		fmt.Println("我是更新用户详情接口")
		// 判断是否登录
		c.JSON(200, gin.H{
			"msg": "更新用户详情接口，需要登录保护",
		})
	})

	user := r.Group("user", LoginAuth)
	{
		// 登录保护
		user.GET(":id", func(c *gin.Context) {
			fmt.Println("我是用户详情接口")
			// 判断是否登录
			c.JSON(200, gin.H{
				"msg": "获取用户详情接口，需要登录保护",
			})
		})
		user.PUT(":id", func(c *gin.Context) {
			// 判断是否登录
			c.JSON(200, gin.H{
				"msg": "更新用户详情接口，需要登录保护",
			})
		})
	}

	r.Run()
}
