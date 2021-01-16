package main

import (
	"22-access-rdb-sql/session"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var str string

	db := session.GetMysqlDB()
	row := db.QueryRow("SELECT 'hello go-mysql'")

	db.Begin()

	if row.Scan(&str) == nil {
		fmt.Println(str)
	}
}
