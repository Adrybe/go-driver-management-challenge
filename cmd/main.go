package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	//port = ":" + port

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "message : pong")
	})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//log.Println("Server started on port " + port)
	r.Run(":" + port)
	//http.ListenAndServe(port, r)
}
