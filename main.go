package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)
	// デフォルトマルチプレクサを生成する
	mux := http.NewServeMux()
	// publicディレクトリを起点とする。
	files := http.FileServer(http.Dir(config.Static))

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

}

// func hello(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("Hello.html")
// 	t.Execute(w, "Hello World!!")
// }

// func post(w http.ResponseWriter, r *http.Request) {
// 	body := r.PostFormValue("body")

// 	_, err := Db.Exec(
// 		`INSERT INTO POSTS (ID, BODY) VALUES (?)`,
// 		body,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	t, _ := template.ParseFiles("Hello.html")
// 	t.Execute(w, body)
// }
