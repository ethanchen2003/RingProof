package dataservice

import (
	"container/list"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func addUserReport(tx *sql.Tx, username string, spamPhone string) error {
	insertSQL := "INSERT INTO userReport (reporter, phone) select userid , ? from users where username = ?"
	_, err := tx.Exec(insertSQL, spamPhone, username)
	if err != nil {
		return err
	}
	return nil
}

type UserReport struct {
	LastName  string
	FirstName string
	Username  string
	SpamPhone string
}

func getUserReports(db *sql.DB, maxRow int, offset int) *list.List {
	selectSQL := "SELECT lastName, firstName, username, r.phone FROM users u, userReport r " +
		"where u.userId = r.reporter limit ? offset ?"

	stmt, err := db.Prepare(selectSQL)
	if err != nil {
		msg := err.Error()
		panic(msg)
		return list.New()
	}

	results, err := stmt.Query(maxRow, offset)
	if err != nil {
		msg := err.Error()
		panic(msg)
		return list.New()
	}
	return toUserReports(results)
}

func toUserReports(rows *sql.Rows) *list.List {
	list := list.New()
	for rows.Next() {
		userReport, err := toUserReport(rows)
		if err == nil {
			list.PushBack(userReport)
		}
	}
	defer rows.Close()
	return list
}

func toUserReport(rows *sql.Rows) (UserReport, error) {
	var userReport UserReport
	err := rows.Scan(&userReport.LastName, &userReport.FirstName, &userReport.Username, &userReport.SpamPhone)
	if err != nil {
		panic(err.Error())
		return userReport, err
	}
	return userReport, nil
}
