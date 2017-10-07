package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DwzStruct struct {
	Id      string `json:"id"`
	TinyUrl string `json:"tinyUrl"`
	LongUrl string `json:"longUrl"`
}

func tinyEndpoint(c *gin.Context) {
	var json *DwzStruct
	if c.BindJSON(&json) == nil {
		if dwz, err := queryDwzWithLongUrl(json.LongUrl); err == nil {
			c.JSON(http.StatusOK, dwz)
			return
		}
		json.Id = newUUID()
		json.TinyUrl = GetRandomString(6)
		saveDwz(json)
		c.JSON(http.StatusOK, json)
	}
}

func restoreEndpoint(c *gin.Context) {
	tinyUrl := c.Param("tinyUrl")
	if dwz, err := queryDwzWithTinyUrl(tinyUrl); err == nil {
		c.Redirect(http.StatusFound, dwz.LongUrl)
	}
}
