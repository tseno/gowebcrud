package main

import (
	"net/http"
	"html/template"

	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)


var Db *sql.DB

func init() {
	Db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// テーブル作成
	_, err = Db.Exec(
		`CREATE TABLE IF NOT EXISTS "POSTS" ("ID" INTEGER PRIMARY KEY, "BODY" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}
}


func main() {

	p("ChitChat" ,version(), "started at", config.Address)

	mux := http.NewServeMux()

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/hello",hello)
	mux.HandleFunc("/post",post)

	server.ListenAndServe()
	
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Hello.html")
	t.Execute(w, "Hello World!!")
}

func post(w http.ResponseWriter, r *http.Request) {
	body := r.PostFormValue("body")

	_, err := Db.Exec(
		`INSERT INTO POSTS (ID, BODY) VALUES (?)`,
		body,
	)
	if err != nil {
		panic(err)
	}

	t, _ := template.ParseFiles("Hello.html")
	t.Execute(w, body)
}
