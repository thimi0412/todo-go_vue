package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Todo : todoの内容
type Todo struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Context   string     `json:"context"`
	LimitDate *time.Time `json:"limit_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func getTodo(id int, userID int) (Todo, error) {

	db := gormConnect()
	defer db.Close()

	todo := Todo{}

	if err := db.First(&todo, "id=? AND user_id =?", id, userID).Error; gorm.IsRecordNotFoundError(err) {
		return todo, errors.New("Record is not found")
	}

	return todo, nil
}

func getTodos(userID int) ([]Todo, error) {
	db := gormConnect()
	defer db.Close()

	todos := []Todo{}

	if err := db.Find(&todos, "user_id=?", userID).Error; gorm.IsRecordNotFoundError(err) {
		return todos, err
	}

	return todos, nil
}

func registerTodo(userID int, context string, limitDate string) (Todo, error) {
	db := gormConnect()
	defer db.Close()

	timeformat := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Tokyo")

	t, err := time.ParseInLocation(timeformat, limitDate, loc)
	if err != nil {
		panic(err)
	}

	todo := Todo{}
	todo.UserID = userID
	todo.Context = context
	todo.LimitDate = &t

	db.Create(&todo)
	return todo, nil
}

func updateTodo(userID int, ID int, context string, limitDate string) (Todo, error) {
	db := gormConnect()
	defer db.Close()

	timeformat := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Tokyo")

	t, err := time.ParseInLocation(timeformat, limitDate, loc)
	if err != nil {
		panic(err)
	}

	todo := Todo{}
	todo.ID = ID
	todo.UserID = userID

	if err := db.Model(&todo).Updates(Todo{Context: context, LimitDate: &t}).Error; err != nil {
		return todo, err
	}
	return todo, nil
}

func deleteTodo(userID int, ID int) error {
	db := gormConnect()
	defer db.Close()

	todo := Todo{}
	todo.ID = ID
	todo.UserID = userID

	if err := db.Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}
