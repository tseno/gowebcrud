package main

import (
	"github.com/mattn/go-sqlite3"
	"database/sql"
)

var db *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
}

func main() {

	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "POSTS" ("ID" INTEGER PRIMARY KEY, "BODY" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}

}
