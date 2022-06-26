package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	article, err := getArticleByID(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   article.Title,
		"payload": article}, "article.html")
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
