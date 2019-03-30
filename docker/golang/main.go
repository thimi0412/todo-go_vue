package main

import (
	"strconv"

	"github.com/badoux/checkmail"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getTodoHandler(c *gin.Context) {
	h := c.GetHeader("Authorization")

	userID, err := authTokenString(h)
	if err != nil || userID == 0 {
		c.JSON(400, gin.H{
			"messege": err,
		})
		return
	}

	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)

	todo, err := getTodo(id, userID)
	if err != nil {
		c.JSON(400, gin.H{
			"messege": "Not Found Todo!",
		})
		return
	}
	c.JSON(200, todo)
}

func getTodosHander(c *gin.Context) {

	h := c.GetHeader("Authorization")

	userID, err := authTokenString(h)
	if err != nil || userID == 0 {
		c.JSON(400, gin.H{
			"messege": err,
		})
		return
	}

	todos, err := getTodos(userID)
	if err != nil {
		c.JSON(400, gin.H{
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

	h := c.GetHeader("Authorization")

	userID, err := authTokenString(h)
	if err != nil || userID == 0 {
		c.JSON(400, gin.H{
			"messege": err,
		})
		return
	}
	context := c.PostForm("context")
	limitDate := c.PostForm("limit_date")

	todo, err := registerTodo(userID, context, limitDate)
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

func updateTodoHandler(c *gin.Context) {
	h := c.GetHeader("Authorization")

	userID, err := authTokenString(h)
	if err != nil || userID == 0 {
		c.JSON(400, gin.H{
			"messege": err,
		})
		return
	}

	sid := c.PostForm("id")
	ID, _ := strconv.Atoi(sid)
	context := c.PostForm("context")
	limitDate := c.PostForm("limit_date")

	todo, err := updateTodo(userID, ID, context, limitDate)
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
	email := c.PostForm("email")
	passoword := c.PostForm("password")
	user, err := getUser(email, passoword)

	token := createTokenString(user)

	if err != nil {
		c.JSON(400, gin.H{
			"messege": "Not Found User!",
		})
		return
	}
	c.JSON(200, gin.H{
		"messege": "Success!",
		"token":   token,
	})
}

func signUpHandler(c *gin.Context) {
	email := c.PostForm("email")
	passoword := c.PostForm("password")

	err := checkmail.ValidateFormat(email)
	if err != nil {
		c.JSON(400, gin.H{
			"messege": "Invalid email",
		})
		return
	}

	user, err := registerUser(email, passoword)
	if err != nil {
		c.JSON(400, gin.H{
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/todo/:id", getTodoHandler)
	r.GET("/todo", getTodosHander)
	r.POST("/todo", postTodoHandler)
	r.PUT("/todo", updateTodoHandler)
	r.POST("/signin", signInHandler)
	r.POST("/signup", signUpHandler)

	r.Run()
}
