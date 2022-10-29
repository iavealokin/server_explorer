package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {

	auth, err := c.Cookie("Auth")
	fmt.Println("It's cookie")
	if auth != "ok" || err != nil {
		c.Redirect(301, "/login")
	}
}
