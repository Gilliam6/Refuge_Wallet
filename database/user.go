package database

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Login string `json:"Login"`
	Pass  string `json:"Pass"`
}

func AddUser(us User) (int64, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(us.Pass), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	res, err := db.Exec("INSERT INTO Accounts (Login, Pass) VALUES (?, ?)", us.Login, hashedPass)
	if err != nil {
		return 0, fmt.Errorf("Can't add user: %v\n", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Add User fault: %v", err)
	}
	return id, nil

}
