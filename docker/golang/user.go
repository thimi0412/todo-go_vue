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

func getUser(email string, password string) (User, error) {

	db := gormConnect()
	defer db.Close()

	user := User{}

	if err := db.Where("email = ?", email).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return user, errors.New("Record is not found")
	}
	err := passwordVerify(user.Password, password)
	if err != nil {
		return user, errors.New("Record is not found")
	}

	return user, nil
}

func registerUser(email string, password string) (User, error) {
	db := gormConnect()
	defer db.Close()

	user := User{}
	user.Email = email
	hash, err := passwordHash(password)
	if err != nil {
		return user, err
	}
	user.Password = hash

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
