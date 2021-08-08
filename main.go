package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

type User struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type LoginForm struct {
	UserName   string `json:"username" binding:"required,min=3,max=7"`
	Password   string `json:"password" binding:"required,len=8"`
	RePassword string `json:"re_password" binding:"required,len=8,eqfield=Password"`
}
type RegisterForm struct {
	UserName string `json:"username" binding:"required,min=3,max=7"`
	Password string `json:"password" binding:"required,len=8"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Sex      uint32 `json:"sex" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

var trans ut.Translator

func main() {
	if err := InitializeTrans(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/hello", hello)
	r.GET("/test", test)
	r.GET("/test2", test2)
	r.GET("user/", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, gin.H{
				"code":    200,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    0,
			"message": "success",
		})
	})
	// 登录
	r.POST("login", loginHandler)
	// 注册
	r.POST("register", registerHandler)
	r.Run()
}

func test2(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "test2",
		"code":    200,
		"data": map[string]interface{}{
			"name": "wls",
			"age":  20,
		},
	})
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "test",
	})
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func InitializeTrans() (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("json")
			return name
		})

		zhT := zh.New()
		uni := ut.New(zhT, zhT)

		trans, _ = uni.GetTranslator("zh")
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
		return
	}
	return
}

// registerHandler  注册
func registerHandler(c *gin.Context) {
	var r RegisterForm
	if err := c.ShouldBindJSON(&r); err != nil {
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(200, gin.H{
				"code": 40010,
				"msg":  "注册失败",
				"err":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 40004,
			"msg":  "注册失败，请检测参数",
			"err":  removeTopStruct(err.Translate(trans)),
		})
		return
	}
	// 注册成功
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": r,
	})
}

// loginHandler 登录
func loginHandler(c *gin.Context) {
	var l LoginForm
	if err := c.ShouldBindJSON(&l); err != nil {
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(200, gin.H{
				"code": 40010,
				"msg":  "登录失败",
				"err":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 40004,
			"msg":  "登录失败，请检测参数",
			"err":  removeTopStruct(err.Translate(trans)),
		})
		return
	}
	// 登录成功
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": l.UserName,
	})
}

func removeTopStruct(fields validator.ValidationErrorsTranslations) validator.ValidationErrorsTranslations {
	r := make(validator.ValidationErrorsTranslations)
	for f, v := range fields {
		r[f[strings.Index(f, ".")+1:]] = v
	}
	return r
}
