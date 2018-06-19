package data

import (
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func Test_createUUID(t *testing.T) {
	tests := []struct {
		name     string
		wantUuid string
	}{
		{
			name:     "",
			wantUuid: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUuid := createUUID()
			fmt.Println(gotUuid)
			// 空では無いことを確認する
			if len(gotUuid) <= 0 {
				t.Errorf("createUUID() = %v, want %v", gotUuid, tt.wantUuid)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	type args struct {
		plaintext string
	}
	tests := []struct {
		name      string
		args      args
		wantCrypt string
	}{
		{
			name: "",
			args: args{
				plaintext: "test1",
			},
			wantCrypt: "b444ac06613fc8d63795be9ad0beaf55011936ac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCrypt := Encrypt(tt.args.plaintext); gotCrypt != tt.wantCrypt {
				t.Errorf("Encrypt() = %v, want %v", gotCrypt, tt.wantCrypt)
			}
		})
	}
}
