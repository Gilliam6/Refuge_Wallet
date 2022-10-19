package database

import (
	"fmt"

	"RefugeWallet/token"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	U_Id  uint   `json:"U_Id"`
	Login string `json:"Login"`
	Pass  string `json:"Pass"`
}

func VerifyPassword(pass, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}

func LoginCheck(us User) (string, error) {
	var info User
	row := db.QueryRow("SELECT * FROM Accounts WHERE Login = ?", us.Login)

	if err := row.Scan(&info.U_Id, &info.Login, &info.Pass); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Username: %s is not registered\n", us.Login)
		}
		return "", fmt.Errorf("CheckLogin fail: %v\n", err)
	}

	err := VerifyPassword(us.Pass, info.Pass)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(info.U_Id)

	if err != nil {
		return "", err
	}

	return token, nil

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

func GetUserByID(userID uint) (User, error) {
	var info User
	row := db.QueryRow("SELECT * FROM Accounts WHERE U_Id = ?", userID)

	if err := row.Scan(&info.U_Id, &info.Login, &info.Pass); err != nil {
		if err == sql.ErrNoRows {
			return info, fmt.Errorf("U_Id: %v is not registered\n", userID)
		}
		return info, fmt.Errorf("CheckLogin fail: %v\n", err)
	}
	return info, nil
}
