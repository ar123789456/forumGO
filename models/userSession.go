package models

import (
	"forum/config"
	"log"
	"time"
)

type UserSession struct {
	User_id        int
	UID            string
	ExpirationTime string
}

func (userSession *UserSession) CREATE(UID string, id int) (*UserSession, error) {
	userSession.ExpirationTime = time.Now().Add(120 * time.Minute).Format(TimeFormat)
	extime := userSession.ExpirationTime
	statement, _ := config.DB.Prepare("INSERT INTO user_session (uid, user_id, extime) VALUES (?, ?, ?)")
	_, err := statement.Exec(UID, id, extime)
	if err == nil {
		userSession.UID = UID
		userSession.User_id = id
		return userSession, err
	}
	log.Println("Unable to create userSession:", err)
	return userSession, err
}

func (userSession *UserSession) GET(UID string) (*UserSession, error) {
	err := config.DB.QueryRow(
		"SELECT uid, user_id, extime FROM user_session WHERE uid=?", UID).Scan(
		&userSession.UID, &userSession.User_id, &userSession.ExpirationTime)
	return userSession, err
}

func (userSession *UserSession) DELETE(UID string) error {
	statement, _ := config.DB.Prepare("SELECT uid, user_id, extime FROM user_session WHERE uid=?")
	_, err := statement.Exec(UID)
	return err
}

func (userSession *UserSession) GetUserId(UID string) (int, error) {
	err := config.DB.QueryRow(
		"SELECT uid, user_id, extime FROM user_session WHERE UID=?", UID).Scan(
		&userSession.UID, &userSession.User_id, &userSession.ExpirationTime)
	return userSession.User_id, err
}

func (userSession *UserSession) Reconciliation() (bool, error) {
	u, err := time.Parse(TimeFormat, userSession.ExpirationTime)
	return time.Now().Before(u), err
}
