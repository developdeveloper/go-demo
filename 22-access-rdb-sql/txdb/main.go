package main

import (
	"22-access-rdb-sql/session"
)

func main() {
	db := session.GetMysqlDB()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO user(`passport`, `password`,`nickname`) VALUES(?, ?, ?)")
	stmt.Exec("wangwu", "abc123", "王五")
	stmt.Exec("zhaoliu", "abc123", "赵六")
	//tx.Commit()
	tx.Rollback()
}
