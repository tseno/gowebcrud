package main

import (
	"net/http"
	"github.com/tseno/gowebcrud/data"
)

func login(writer http.ResponseWriter, request *http.Request) {
	// ログイン画面の表示
	t := parseTemplateFiles("login.layout", "public.navber", "login")
	t.Execute(writer, nil)
}

func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	// 入力された値を、構造体に入れる
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	// ユーザの作成
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	// ログイン画面に遷移する
	http.Redirect(writer, request, "/login", 302)
}

// 認証
// POST /authenticate
func authenticate(writer http.ResponseWriter, request *http.Request) {
	// フォームの内容をパースする
	err := request.ParseForm()
	// 構造体Userを作成して、入力されたemailをキーにDBからデータを入れる。
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}

	// ハッシュ化されたpasswordと比較する
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		// DBのsessionテーブルにinsertする
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		// セッションクッキーの作成。有効期限は設定しないことで、ブラウザが終了する際に自動的に消去される。
		// HttpOnlyで、HTTP,HTTPSに限定される。Javascript等からはアクセスできない。
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	}

}

func logout(writer http.ResponseWriter, request *http.Request) {
	// クッキーを取得する
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		// セッションの削除
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
