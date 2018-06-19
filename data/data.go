package data

import (
	"database/sql"
	"log"
	"crypto/rand"
	"fmt"
	"crypto/sha1"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func init() {
	var err error
	//Db, err := sql.Open("sqlite3", "./test.db")
	Db, err = sql.Open("postgres", "dbname=gocrud sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Db)

	return
}

func createUUID() (uuid string) {
	// 16バイトのバイト配列を作成
	u := new([16]byte)
	// ランダムの文字列をuに入れる
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("UUIDの作成に失敗しました。", err)
	}
	// 0x40 をくっつけて、0x7Fでマスクする。
	u[8] = (u[8] | 0x40) & 0x7F
	//
	u[6] = (u[6] | 0xF) & (0x4 << 4)

	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(plaintext string) (crypt string) {

	crypt = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
