package v1

import (
	"gin_demo/login/service"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var s service.UserLoginService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
	} else {
		res := s.Login()
		c.JSON(200, res)
	}

}

func RegisterHandler(c *gin.Context) {

	var s service.UserRegisterSerivce
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
	} else {
		res := s.Register()
		c.JSON(200, res)
	}

}
