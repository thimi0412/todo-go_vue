package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
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
	id, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(500, gin.H{
			"messege": err,
		})
	} else {
		todo := getUser(id)

		c.JSON(200, todo)
	}

}

func postUserHandler(c *gin.Context) {
	email := c.PostForm("email")
	passoword := c.PostForm("passoword")

	user := registerUser(email, passoword)

	c.JSON(200, gin.H{
		"messege": "Success!",
		"result":  user,
	})

}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todo/:id", getTodoHandler)
	r.GET("/user/:id", getUserHandler)
	r.POST("/user", postUserHandler)

	r.Run()
}
