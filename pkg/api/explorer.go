package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"server_explorer/pkg/models"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	mainFolder string            = "/homestorage"
	images     map[string]string = map[string]string{".jpg": "", ".png": ""}
)

func Site(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", nil)
}
func GetExplorer(c *gin.Context) {

	newPath, err := url.QueryUnescape(c.Request.URL.String())
	if err != nil {
		fmt.Println(err)
	}
	if newPath == "/" {
		newPath = "/client"
	}
	page := getData(newPath)
	c.HTML(http.StatusOK, "template.html", gin.H{
		"Page": page,
	})
}

func getData(newPath string) models.Page {
	var page models.Page
	replacedPath := strings.Replace(newPath, "/client", mainFolder, -1)
	if replacedPath != mainFolder {
		page.NeedUp = true

	}
	page.Path = newPath

	files, err := ioutil.ReadDir(replacedPath)
	if err != nil {
		fmt.Println("err", err) //log.Fatal(err)
	}
	page.UpLink = upLink(page.Path)
	newId := 1
	var isImage bool
	for _, file := range files {
		if _, ok := images[strings.ToLower(filepath.Ext(newPath+string(os.PathSeparator)+file.Name()))]; ok {
			isImage = true
		}
		if file.IsDir() {
			page.Files = append(page.Files, models.File{Id: newId, Name: file.Name(), Size: file.Size(), IsFolder: file.IsDir(), Link: newPath + string(os.PathSeparator) + file.Name(), IsImage: isImage})
		} else {
			link := "/api/getFile?path=" + newPath + string(os.PathSeparator) + file.Name()
			page.Files = append(page.Files, models.File{Id: newId, Name: file.Name(), Size: file.Size(), IsFolder: file.IsDir(), Link: link, IsImage: isImage})
		}
		newId++
	}
	return page
}
func upLink(link string) string {
	links := strings.Split(link, "/")
	return strings.Join(links[:len(links)-1], "/")
}
