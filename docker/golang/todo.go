package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Todo : todoの内容
type Todo struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Context     string     `json:"context"`
	LimitDate   *time.Time `json:"limit_date"`
	InsertDate  *time.Time `json:"insert_date"`
	UpdatedDate *time.Time `json:"updated_date"`
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

func getTodos(userID int) ([]Todo, error) {
	db := gormConnect()
	defer db.Close()

	todos := []Todo{}

	if err := db.Find(&todos, "user_id=?", userID).Error; gorm.IsRecordNotFoundError(err) {
		return todos, err
	}

	db.Find(&todos, "user_id=?", userID)

	return todos, nil
}

func registerTodo(userID int, context string, limitDate string) (Todo, error) {
	db := gormConnect()
	defer db.Close()

	timeformat := "2006-01-02 15:04:05"

	t, err := time.Parse(timeformat, limitDate)
	if err != nil {
		panic(err)
	}

	now := time.Now()

	todo := Todo{}
	todo.UserID = userID
	todo.Context = context
	todo.LimitDate = &t
	todo.InsertDate = &now
	todo.UpdatedDate = &now

	db.Create(&todo)
	return todo, nil
}
