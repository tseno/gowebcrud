package data

import (
	"testing"
	"log"
)

func TestCreateUUID(t *testing.T) {

	uuid := createUUID()

	if uuid == "" || len(uuid) != 36 {
		log.Fatal(uuid)
	} else {
		log.Println("uuid: ", uuid, " len: ", len(uuid))
	}
}

func TestEncrypt(t *testing.T) {
	crypt := Encrypt("test1")
	if crypt != "b444ac06613fc8d63795be9ad0beaf55011936ac" {
		log.Fatal(crypt)
	} else {
		log.Println("crypt: ", crypt, " len: ", len(crypt))
	}
}


