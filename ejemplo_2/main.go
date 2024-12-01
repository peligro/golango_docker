package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"estado":  "ok",
			"mensaje": "Hola desde Docker",
		})
	})

	router.Run(":8080")

}

//go get -u github.com/gin-gonic/gin
//go clean -i github.com/gin-gonic/gin

//docker build . -t ejemplo-go-gin-docker
//docker run -e PORT=8080 -p 8080:8080 ejemplo-go-gin-docker
