package server

import (
	"fmt"
	"server_explorer/pkg/api"
	"server_explorer/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewHttpServer(port int) {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Static("/public", "./public")
	r.Static("/homestorage", "/homestorage")
	r.LoadHTMLFiles("public/auth/login.html")
	r.LoadHTMLGlob("public/templates/*.html")

	r.POST("/login", api.Login)
	r.GET("/login", api.LoginForm)
	r.GET("api/getFile", api.GetFile)
	r.Use(middleware.BasicAuth)
	r.Use(middleware.Route)
	//r.GET("/", api.Site)
	r.Run(fmt.Sprintf(":%v", port))
}
