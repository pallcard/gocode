package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init()  {
	Db, err = sql.Open("mysql", "root:lk123456@tcp(203.195.238.249:3306)/test_go")
	if err != nil {
		panic(err.Error())
	}
}