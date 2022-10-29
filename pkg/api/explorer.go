package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server_explorer/pkg/models"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	mainFolder string = "/homestorage"
	path       string = mainFolder
)

func GetExplorer(c *gin.Context) {
	if len(c.Request.URL.Query()["path"]) > 0 {
		newPath := c.Request.URL.Query()["path"][0]
		if newPath != "/" {
			path += "/" + newPath
		} else {
			arr := strings.Split(path, "/")
			path = ""
			for i := 0; i < len(arr)-1; i++ {
				path += arr[i] + "/"
			}
			path = strings.TrimSuffix(path, "/")
		}
	}

	page := getData(path)

	c.HTML(http.StatusOK, "template.tmpl", gin.H{
		"Page": page,
	})
}

func getData(level string) models.Page {
	var page models.Page
	fmt.Println(level)

	if level != mainFolder {
		page.NeedUp = true

	}
	page.Path = level

	files, err := ioutil.ReadDir(level)
	if err != nil {
		log.Fatal(err)
	}
	newId := 1
	for _, file := range files {
		//path[lvl] = append(path[lvl], models.Object{Id: newId, Prev: id + 1, Path: path[lvl-1][id].Path + "/" + file.Name()})
		page.Files = append(page.Files, models.File{Id: newId, Name: file.Name(), Size: file.Size(), IsFolder: file.IsDir()})

		newId++
	}
	return page
}
