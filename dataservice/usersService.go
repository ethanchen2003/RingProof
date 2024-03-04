package dataservice

import (
	//	"container/list"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserId    int64
	LastName  string
	FirstName string
	Bod       time.Time
	Gender    string
	Phone     string
	Email     string
	Username  string
	Password  string
}

func addUser(tx *sql.Tx, user User) (int64, error) {
	insertSQL := "INSERT INTO users (lastName, firstName, bod, gender, phone, email, username, password) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := tx.Exec(insertSQL, user.LastName, user.FirstName, user.Bod, user.Gender, user.Phone, user.Email, user.Username, user.Password)

	if err != nil {
		return 0, err
	}

	userId, _ := res.LastInsertId()
	return userId, nil
}

func getUser(db *sql.DB, username string) (User, error) {
	var user User
	sql := "SELECT userId, lastName, firstName, bod, gender, phone, email, username, password FROM users WHERE username = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return user, err
	}

	rows, err := stmt.Query(username)
	if err != nil {
		return user, err
	}

	if rows.Next() {
		err = rows.Scan(&user.UserId, &user.LastName, &user.FirstName, &user.Bod, &user.Gender,
			&user.Phone, &user.Email, &user.Username, &user.Password)
	}
	return user, err
}
