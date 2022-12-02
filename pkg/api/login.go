package api

import (
	"fmt"
	"net/http"
	"server_explorer/pkg/bus"
	"server_explorer/pkg/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	u := &models.User{
		Name:     c.PostForm("user"),
		Password: c.PostForm("password"),
	}
	err := bus.Dispatch(u)
	if err != nil {
		fmt.Println("Err sql", err)
	}
	c.SetCookie("Auth", "ok", 32000, "http://192.168.1.108:8081/", "", false, true)
	c.Redirect(301, "/client")
}

func LoginForm(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}
