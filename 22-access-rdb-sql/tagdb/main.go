package main

import (
	"22-access-rdb-sql/session"
	"fmt"
	"time"
)

type User struct {
	ID        int       `db:"id"`
	Passport  string    `db:"passport"`
	Password  string    `db:"password"`
	Nickname  string    `db:"nickname"`
	CreatedAt time.Time `db:"create_time"`
}

func main() {
	var total int
	var user User
	var users []User
	var names []string
	db := session.GetSqlxDB()

	db.Get(&total, "SELECT COUNT(*) FROM user")
	fmt.Println(total)

	db.Get(&user, "SELECT * FROM user LIMIT 1")
	fmt.Println(user)

	db.Select(&users, "SELECT * FROM user")
	fmt.Println(users)

	db.Select(&names, "SELECT nickname FROM user")
	fmt.Println(names)

	db.QueryRowx("SELECT * FROM user LIMIT 1").StructScan(&user)
	fmt.Println(user)

	rows, _ := db.Queryx("SELECT * FROM user")
	defer rows.Close()
	for rows.Next() {
		var u User
		rows.StructScan(&u)
		fmt.Println(u)
	}
}
