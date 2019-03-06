package main

// Todo : todoの内容
type Todo struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Context string `json:"context"`
}

func getTodo(id int) Todo {

	db := gormConnect()
	defer db.Close()

	todo := Todo{}
	todo.ID = id
	db.First(&todo)

	return todo
}
