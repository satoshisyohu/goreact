package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/*")
	r.GET("/", ShowAllBlog)
	r.GET("/show/:id", ShowOneData)
	r.POST("/create", CreateBlog)
	r.GET("/caliculate", Caliculatehome)
	// r.GET("/caliculate/:gram", Caliculate)
	r.POST("/new", New)
	r.POST("/signup", SignUp)
	r.GET("/signup", SignUpHome)
	r.GET("/login", LoginHome)
	r.POST("/login", Login)
	r.GET("/beans", Beans)
	r.POST("/beans/register", Register)

	return r
}
