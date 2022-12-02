package middleware

import (
	"fmt"
	"server_explorer/pkg/api"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {

	auth, err := c.Cookie("Auth")
	fmt.Println("It's cookie")
	if auth != "ok" || err != nil {
		c.Redirect(301, "/login")
	}
}

func Route(c *gin.Context) {
	request := strings.Split(c.Request.URL.String(), "/")

	if !strings.Contains(request[1], "client") {
		return
	}
	api.GetExplorer(c)
}
