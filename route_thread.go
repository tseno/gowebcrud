package main

import (
	"net/http"
	"gowebcrud/data"
	"fmt"
)

func newThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		// ログインしていなければ、ログイン画面に遷移する
		http.Redirect(writer, request, "/login", 302)
	} else {
		// ログアウトボタン付きの、新スレッド作成画面を表示する
		generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

func createThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		// ログインしていなければ、ログイン画面に遷移する
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		// スレッド内容の取得
		topic := request.PostFormValue("topic")
		// スレッドの作成
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}

}

func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	// DBからスレッドを取得する
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		error_message(writer, request, "Cannot read thread")
	} else {
		_, err := session(writer,request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

func postThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		// ログインしていなければ、ログイン画面に遷移する
		http.Redirect(writer, request, "/login", 302)
	} else {
		// リクエストをパースする
		err = request.ParseForm()
		if err != nil {
			danger(err,"Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err,"Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(writer, request, "Cannot read thread")
		}
		// 投稿を保存する
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer,request,url,302)
	}

}
