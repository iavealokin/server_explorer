package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetFile(c *gin.Context) {
	if len(c.Request.URL.Query()["path"]) > 0 {
		path := c.Request.URL.Query()["path"][0]
		path = strings.Replace(path, "client", "homestorage", -1)
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`filename=%s`, strings.Split(path, "/")[len(strings.Split(path, "/"))-1]))
		http.ServeFile(c.Writer, c.Request, path)
	}
}
