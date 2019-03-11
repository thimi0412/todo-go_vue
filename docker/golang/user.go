package main

// User : userのテーブルcolumn
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getUser(id int) User {

	db := gormConnect()
	defer db.Close()

	user := User{}
	user.ID = id
	db.First(&user)

	return user
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
