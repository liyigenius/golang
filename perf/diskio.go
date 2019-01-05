package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func generateRes(i int, txt string, db sql.DB) {

	//db := openDB()
	stmt, _ := db.Prepare("INSERT INTO txt(id, text) values(?,?)")

	stmt.Exec(i, "liyi")
	//affect, _ := res.RowsAffected()

	//fmt.Println(affect)

}

func openDB() sql.DB {
	var db *sql.DB
	db, _ = sql.Open("sqlite3", "perf/perfdb.sqlite")

	return *db
}

func main() {
	db1 := openDB()
	for i := 10000; i < 20000; i++ {
		generateRes(i, "123", db1)
	}

	//read & write...

}
