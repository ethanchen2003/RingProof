package dataservice

import (
	"container/list"
)

/**
* this method will add user data to database user table.
* It is used to process user register.
 */
func AddUser(user User) int64 {
	//open database connection
	db := connection()
	var userId int64 = -1

	//start database transaction
	tx, err := db.Begin()
	if err != nil {
		println(err.Error())
		return userId
	}

	userId, err = addUser(tx, user)

	if err != nil {
		//rollback if error
		tx.Rollback()
	} else {
		//commit transaction
		err = tx.Commit()
	}
	return userId
}

/*
*  user report a spam phone
*  the service will insert spam phone to spamphone table or increase reportNumber by one.
*  Then insert data to userReport table.
 */
func ReportSpamPhone(username string, spamPhone string) {
	//open database connection
	db := connection()

	//start database transaction
	tx, err := db.Begin()
	if err != nil {
		println(err.Error())
		return
	}

	_, err = reportPhone(tx, spamPhone)

	if err != nil {
		//rollback if error
		tx.Rollback()
		return
	}

	err = addUserReport(tx, username, spamPhone)

	if err != nil {
		//rollback if error
		tx.Rollback()
		return
	}

	//commit transaction
	err = tx.Commit()
}

func FindSpamPhoneByPrefix(phonePrefix string) *list.List {
	//open database connection
	db := connection()
	return findSpamPhoneByPrefix(db, phonePrefix)
}

func GetAllSpamPhonePagination(maxRow int, offset int) *list.List {
	//open database connection
	db := connection()
	return getAllSpamPhonePagination(db, maxRow, offset)
}

func GetUser(username string) (User, error) {
	//open database connection
	db := connection()
	return getUser(db, username)
}

func GetUserReports(maxRow int, offset int) *list.List {
	//open database connection
	db := connection()
	return getUserReports(db, maxRow, offset)
}
