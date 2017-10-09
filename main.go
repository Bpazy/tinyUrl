package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"flag"
	"log"
	"io"
	"os"
	"path/filepath"
	"github.com/thinkerou/favicon"
)

func main() {
	port := flag.String("port", ":8080", "Server running port.")
	logFile := flag.String("log", "tinyUrl.log", "Server log file. Only work in release mode.")
	flag.Parse()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(favicon.New("./static/"))

	if gin.Mode() == gin.ReleaseMode {
		os.MkdirAll(filepath.Dir(*logFile), os.ModePerm)
		f, err := os.Create(*logFile)
		checkErr(err)

		writer := io.MultiWriter(os.Stdout, f)
		log.SetOutput(writer)
		router.Use(gin.LoggerWithWriter(writer))
	} else {
		router.Use(gin.Logger())
	}

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
	log.Fatal(router.Run(*port))
}
