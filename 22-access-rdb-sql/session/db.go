package session

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	// db, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	//db.Ping()
	return db
}

func GetSqlxDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	//db.Ping()
	return db
}
