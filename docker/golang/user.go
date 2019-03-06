package main

// User : userのテーブルcolumn
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUser(id int) User {

	db := gormConnect()
	defer db.Close()

	user := User{}
	user.ID = id
	db.First(&user)

	return user
}
