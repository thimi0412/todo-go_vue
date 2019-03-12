package main

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Todo : todoの内容
type Todo struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Context string `json:"context"`
}

func getTodo(id int) (Todo, error) {

	db := gormConnect()
	defer db.Close()

	todo := Todo{}
	todo.ID = id

	if err := db.First(&todo).Error; gorm.IsRecordNotFoundError(err) {
		return todo, errors.New("Record is not found")
	}

	db.First(&todo)

	return todo, nil
}
