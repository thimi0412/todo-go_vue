package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getTodoHandler(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)

	todo, err := getTodo(id)
	if err != nil {
		c.JSON(404, gin.H{
			"messege": "Not Found Todo!",
		})
		return
	}
	c.JSON(200, todo)
}

func getTodosHander(c *gin.Context) {
	todos, err := getTodos(1)
	if err != nil {
		c.JSON(404, gin.H{
			"messege": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"messege": "Success!",
		"result":  todos,
	})
}

func postTodoHandler(c *gin.Context) {
	userID := c.PostForm("user_id")
	context := c.PostForm("context")
	limitDate := c.PostForm("limit_date")

	suserID, _ := strconv.Atoi(userID)

	todo, err := registerTodo(suserID, context, limitDate)
	if err != nil {
		c.JSON(404, gin.H{
			"messege": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"messege": "Success!",
		"result":  todo,
	})

}

func signInHandler(c *gin.Context) {
	session := sessions.Default(c)
	email := c.PostForm("email")
	passoword := c.PostForm("password")
	user, err := getUser(email, passoword)
	if err != nil {
		c.JSON(404, gin.H{
			"messege": "Not Found User!",
		})
		return
	}
	session.Set("id", user.ID)
	session.Save()
	c.JSON(200, user)
}

func signUpHandler(c *gin.Context) {
	email := c.PostForm("email")
	passoword := c.PostForm("password")

	user, err := registerUser(email, passoword)
	if err != nil {
		c.JSON(404, gin.H{
			"messege": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"messege": "Success!",
		"result":  user,
	})

}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/todo/:id", getTodoHandler)
	r.GET("/todo", getTodosHander)
	r.POST("/todo", postTodoHandler)
	r.POST("/signin", signInHandler)
	r.POST("/signup", signUpHandler)

	r.Run()
}
