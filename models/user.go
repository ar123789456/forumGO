package models

import (
	"forum/config"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nicname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserParams struct {
	Nickname string
	Email    string
	Password string
}

func (user *User) CREATE(userInput UserParams) (*User, error) {
	statement, _ := config.DB.Prepare("INSERT INTO user (NicName, Email, Password) VALUES (?, ?, ?)")
	result, err := statement.Exec(userInput.Nickname, userInput.Email, userInput.Password)
	if err == nil {
		id, _ := result.LastInsertId()
		user.Id = int(id)
		user.Nickname = userInput.Nickname
		user.Email = userInput.Email
		user.Password = userInput.Password
		return user, err
	}
	log.Println("Unable to create user:", err)
	return user, err
}

func (user *User) FETCH(nick string) (*User, error) {
	err := config.DB.QueryRow(
		"SELECT ID, NicName, Email, Password FROM user WHERE NicName=?", nick).Scan(
		&user.Id, &user.Nickname, &user.Email, &user.Password)
	return user, err
}

func (user *User) GetUser(ID int) (*User, error) {
	err := config.DB.QueryRow(
		"SELECT ID, NicName, Email FROM user WHERE ID=?", ID).Scan(
		&user.Id, &user.Nickname, &user.Email)
	return user, err
}

// func (user *User) UPDATEuid(UID string, id int) (*User, error) {
// 	// ToDo add category_id
// 	statement, _ := config.DB.Prepare("UPDATE user SET UID = ? WHERE id = ?;")
// 	_, err := statement.Exec(UID, id)
// 	return user, err
// }
