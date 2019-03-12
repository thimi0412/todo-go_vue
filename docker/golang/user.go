package main

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User : userのテーブルcolumn
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getUser(id int) (User, error) {

	db := gormConnect()
	defer db.Close()

	user := User{}
	user.ID = id

	if err := db.First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return user, errors.New("Record is not found")
	}

	db.First(&user)

	return user, nil
}

func registerUser(email string, password string) User {
	db := gormConnect()
	defer db.Close()

	user := User{}
	user.Email = email
	user.Password = password
	db.Create(&user)

	return user
}
