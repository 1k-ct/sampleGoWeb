package main

import (
	"net/http"

	user "./service"
	"github.com/1k-ct/sampleGoWeb/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var s user.Service
	todos, _ := s.GetAll()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world go gin")
		c.String(http.StatusOK, "", todos)
	})
	db.Init()
	server.Init()
	defer db.Close()
}
