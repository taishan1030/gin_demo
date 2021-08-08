package main

import (
	v1 "gin_demo/login/api/v1"
	"gin_demo/login/model"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := model.InitDb(); err != nil {
		panic(err)
	}

	r := gin.Default()

	v := r.Group("api/v1")
	{
		v.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "pong",
			})
		})
		v.POST("user/register", v1.RegisterHandler)
		v.POST("user/login", v1.LoginHandler)

	}

	r.Run()
}
