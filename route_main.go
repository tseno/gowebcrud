package main

import (
	"net/http"
	"github.com/tseno/gowebcrud/data"
)

func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	// 一覧の取得
	threads, err := data.Threads()
	if err != nil {
		// データが取れなかった場合、エラー表示
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		if err != nil {
			// セッションが取得できなかった場合、ログインボタンありの一覧画面表示
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			// セッションが取得できた場合、ログアウトボタンありの一覧画面表示
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
