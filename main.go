package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"flag"
	"log"
	"io"
	"os"
	"path/filepath"
)

func main() {
	port := flag.String("port", ":8080", "Server running port.")
	logFile := flag.String("log", "tinyUrl.log", "Server log file. Only work in release mode.")
	flag.Parse()

	if gin.Mode() == gin.ReleaseMode {
		os.MkdirAll(filepath.Dir(*logFile), os.ModePerm)
		f, err := os.Create(*logFile)
		checkErr(err)
		log.SetOutput(io.MultiWriter(os.Stdout, f))
	}

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
