package dataservice

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// one time dagabase connection thats used throughout the app
var sqlConn *sql.DB

var cursor = "mysql"

var dbUser = "DB_USER"
var dbPASSWORD = "DB_PASSWORD"
var dbHOST = "DB_HOST"
var dbPORT = "DB_PORT"
var dbSERVER = "DB_SERVER"

var DBUser = "dev"
var DBPassword = "iex.prms"
var DBHost = "localhost"
var DBPort = "3306"
var DBServer = "dev"

type errorInterface interface {
	Error() string
}

func handleError(err errorInterface) {
	if err != nil {
		log.Println(err.Error())
	}
}

func Init(maxConns int) {
	sqlConn = open(maxConns)
}

func Close() {
	if sqlConn != nil {
		sqlConn.Close()
	}
}

func connection() *sql.DB {

	if sqlConn != nil {
		return sqlConn
	}
	sqlConn = open(5)
	return sqlConn
}

func open(maxConns int) *sql.DB {
	log.Println("Connect to mysql...")
	var err error
	var conn *sql.DB

	conn, err = sql.Open(cursor, getConnectionString())
	if maxConns > 0 {
		conn.SetMaxOpenConns(maxConns)
	}

	handleError(err)

	log.Println("Checking the connectivity")
	pingerr := conn.Ping()
	handleError(pingerr)
	return conn
}

func getConnectionString() string {
	var sb string
	sb = getenv(dbUser, DBUser)
	sb += ":"
	sb += getenv(dbPASSWORD, DBPassword)
	sb += "@tcp("
	sb += getenv(dbHOST, DBHost)
	sb += ":"
	sb += getenv(dbPORT, DBPort)
	sb += ")/"
	sb += getenv(dbSERVER, DBServer)
	sb += "?parseTime=true"
	return sb
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
