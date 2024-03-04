package dataservice

import (
	"container/list"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SpamPhone struct {
	Phone        string
	ReportNumber int
}

/**
*
 */
func reportPhone(tx *sql.Tx, phone string) (int64, error) {
	count, err := insertSpamPhone(tx, phone)
	if err == nil {
		return count, nil
	}
	return updateSpamPhone(tx, phone)
}

func insertSpamPhone(tx *sql.Tx, phone string) (int64, error) {
	insertSQL := "INSERT INTO spamPhone (phone, reportNumber) VALUES (?, 1)"
	res, err := tx.Exec(insertSQL, phone)
	if err != nil {
		return 0, err
	}

	count, _ := res.RowsAffected()
	return count, nil
}

func updateSpamPhone(tx *sql.Tx, phone string) (int64, error) {
	updateSQL := "update spamPhone set reportNumber = reportNumber + 1 where phone = ? "
	res, err := tx.Exec(updateSQL, phone)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	return count, nil
}

func findSpamPhoneByPrefix(db *sql.DB, phonePrefix string) *list.List {
	if len(phonePrefix) == 0 {
		return list.New()
	}

	stmt, err := db.Prepare("SELECT phone, reportNumber FROM spamPhone where phone like ?")
	if err != nil {
		msg := err.Error()
		panic(msg)
		return list.New()
	}

	results, err := stmt.Query(phonePrefix + "%")
	if err != nil {
		msg := err.Error()
		panic(msg)
		return list.New()
	}

	return tophones(results)
}

func findSpamPhoneByPrefixNumber(db *sql.DB, phonePrefix string, reportNumber int) *list.List {
	stmt, err := db.Prepare("SELECT phone, reportNumber FROM spamPhone where phone like ? and reportNumber >= ? ")
	if err != nil {
		msg := err.Error()
		panic(msg)
		return list.New()
	}

	results, err := stmt.Query(phonePrefix+"%", reportNumber)
	if err != nil {
		msg := err.Error()
		panic(msg)
		return list.New()
	}
	return tophones(results)
}

func getSpamPhoneByPhone(db *sql.DB, phone string) (SpamPhone, error) {
	stmt, err := db.Prepare("SELECT phone, reportNumber FROM spamPhone where phone = ?")
	var spamPhone SpamPhone
	if err != nil {
		msg := err.Error()
		panic(msg)
		return spamPhone, err
	}

	err = stmt.QueryRow(phone).Scan(&spamPhone.Phone, &spamPhone.ReportNumber)
	if err != nil {
		return spamPhone, err
	}
	return spamPhone, nil
}

func getAllSpamPhonePagination(db *sql.DB, maxRow int, offset int) *list.List {
	stmt, err := db.Prepare("SELECT phone, reportNumber FROM spamPhone limit ? offset ?")
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
	return tophones(results)
}

func tophone(rows *sql.Rows) (SpamPhone, error) {
	var spamPhone SpamPhone
	err := rows.Scan(&spamPhone.Phone, &spamPhone.ReportNumber)
	if err != nil {
		panic(err.Error())
		return spamPhone, err
	}
	return spamPhone, nil
}

func tophones(rows *sql.Rows) *list.List {
	list := list.New()
	for rows.Next() {
		spamPhone, err := tophone(rows)
		if err == nil {
			list.PushBack(spamPhone)
		}
	}
	defer rows.Close()
	return list
}
