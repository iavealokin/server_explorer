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
	r.LoadHTMLFiles("public/auth/login.html")
	r.LoadHTMLGlob("public/templates/*")

	r.POST("/login", api.Login)
	r.GET("/login", api.LoginForm)
	r.GET("/main", middleware.BasicAuth, api.GetExplorer)
	r.Run(fmt.Sprintf(":%v", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
