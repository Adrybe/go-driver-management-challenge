package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	port = ":" + port

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "message : pong")
	})

	r.Run()
}
