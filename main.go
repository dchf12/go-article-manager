package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", showIndexPage)
	r.GET("/article/view/:article_id", getArticle())
	r.Run()
}
