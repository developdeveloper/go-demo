package main

import (
	"22-access-rdb-sql/session"
	"fmt"
)

func main() {
	db := session.GetMysqlDB()
	stmt, _ := db.Prepare("INSERT INTO user(`passport`, `password`,`nickname`) VALUES(?, ?, ?)")
	result, _ := stmt.Exec("lisi", "abc123", "李四")
	fmt.Println(result.LastInsertId())
}
