package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Println("Server started on port " + port)
	r.Run(":3000")
	//http.ListenAndServe(port, r)
}
