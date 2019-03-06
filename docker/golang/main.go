package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getTodoHandler(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)

	todo := getTodo(id)

	c.JSON(200, todo)
}

func getUserHandler(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)

	todo := getUser(id)

	c.JSON(200, todo)
}

func main() {
	r := gin.Default()

	r.GET("/todo/:id", getTodoHandler)
	r.GET("/user/:id", getUserHandler)

	r.Run()
}
