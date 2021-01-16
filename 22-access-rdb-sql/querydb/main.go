package main

import (
	"22-access-rdb-sql/session"
	"fmt"
	"time"
)

type User struct {
	id        int
	passport  string
	password  string
	nickname  string
	createdAt time.Time
}

func scan() {
	db := session.GetMysqlDB()
	rows, _ := db.Query("SELECT * FROM user")
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.id, &user.passport, &user.password, &user.nickname, &user.createdAt)
		fmt.Println(user)
	}
}

func single() {
	var user User
	db := session.GetMysqlDB()
	db.QueryRow("SELECT * FROM user WHERE id = 1").Scan(&user.id, &user.passport, &user.password, &user.nickname, &user.createdAt)
	fmt.Println(user)
}

func prepare() {
	db := session.GetMysqlDB()
	stmt, _ := db.Prepare("SELECT * FROM user WHERE id <> ?")
	defer stmt.Close()
	rows, _ := stmt.Query(0)
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.id, &user.passport, &user.password, &user.nickname, &user.createdAt)
		fmt.Println(user)
	}
}

func main() {
	scan()
	single()
	prepare()
}
