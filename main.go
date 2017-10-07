package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"flag"
)

func main() {
	port := flag.String("port", ":8080", "Server running port.")
	flag.Parse()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "index.html")
	})
	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/tiny", tinyEndpoint)
		v1.GET("/r/:tinyUrl", restoreEndpoint)
	}
	router.Run(*port)
}
