package data

import (
	"reflect"
	"testing"
	"time"
)

func TestUser_CreateSession(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
	}
	tests := []struct {
		name        string
		fields      fields
		wantSession Session
		wantErr     bool
	}{
		{
			name: "",
			// users
			fields: fields{
				Id:        10,
				Uuid:      "",
				Name:      "",
				Email:     "meado",
				Password:  "",
				CreatedAt: time.Time{},
			},
			wantSession: Session{
				Id:        1,
				Uuid:      "session0001",
				Email:     "meado",
				UserId:    10,
				CreatedAt: time.Now(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
			}
			gotSession, err := user.CreateSession()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.CreateSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSession.Email != tt.wantSession.Email || gotSession.UserId != tt.wantSession.UserId {
				t.Errorf("User.CreateSession() = %v, want %v", gotSession, tt.wantSession)
			} else {
				t.Logf("User.CreateSession() = %v, want %v", gotSession, tt.wantSession)
			}
		})
	}
}

func TestUser_Session(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
	}
	tests := []struct {
		name        string
		fields      fields
		wantSession Session
		wantErr     bool
	}{
		{
			name: "",
			fields: fields{
				Id:        10,
				Uuid:      "",
				Name:      "",
				Email:     "",
				Password:  "",
				CreatedAt: time.Time{},
			},
			wantSession: Session{
				Id:        150,
				Uuid:      "xxx10",
				Email:     "tseno",
				UserId:    10,
				CreatedAt: time.Date(2018, 06, 19, 12, 42, 10, 0, time.FixedZone("", 0)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
			}
			gotSession, err := user.Session()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Session() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSession, tt.wantSession) {
				t.Errorf("User.Session() = %v, want %v", gotSession, tt.wantSession)
			} else {
				t.Logf("User.Session() = %v, want %v", gotSession, tt.wantSession)
			}
		})
	}
}

func TestSession_Check(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Email     string
		UserId    int
		CreatedAt time.Time
	}
	tests := []struct {
		name      string
		fields    fields
		wantValid bool
		wantErr   bool
	}{
		{
			name: "",
			fields: fields{
				Id:        0,
				Uuid:      "53b0436b-89e7-0060-5b95-f393944d774a",
				Email:     "",
				UserId:    0,
				CreatedAt: time.Time{},
			},
			wantValid: true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session := &Session{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Email:     tt.fields.Email,
				UserId:    tt.fields.UserId,
				CreatedAt: tt.fields.CreatedAt,
			}
			gotValid, err := session.Check()
			if (err != nil) != tt.wantErr {
				t.Errorf("Session.Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValid != tt.wantValid {
				t.Errorf("Session.Check() = %v, want %v", gotValid, tt.wantValid)
			} else {
				t.Logf("Session.Check() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func TestSession_DeleteByUUID(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Email     string
		UserId    int
		CreatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				Id:        0,
				Uuid:      "22b45978-d48f-402c-4a50-137edbc8b56b",
				Email:     "",
				UserId:    0,
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session := &Session{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Email:     tt.fields.Email,
				UserId:    tt.fields.UserId,
				CreatedAt: tt.fields.CreatedAt,
			}
			if err := session.DeleteByUUID(); (err != nil) != tt.wantErr {
				t.Errorf("Session.DeleteByUUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSession_User(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Email     string
		UserId    int
		CreatedAt time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		wantUser User
		wantErr  bool
	}{
		{
			name: "",
			fields: fields{
				Id:        0,
				Uuid:      "",
				Email:     "",
				UserId:    10,
				CreatedAt: time.Time{},
			},
			wantUser: User{
				Id:        10,
				Uuid:      "xxx10",
				Name:      "tseno",
				Email:     "tseno",
				CreatedAt: time.Date(2018, 06, 18, 21, 16, 49, 0, time.FixedZone("", 0)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session := &Session{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Email:     tt.fields.Email,
				UserId:    tt.fields.UserId,
				CreatedAt: tt.fields.CreatedAt,
			}
			gotUser, err := session.User()
			if (err != nil) != tt.wantErr {
				t.Errorf("Session.User() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("Session.User() = %v, want %v", gotUser, tt.wantUser)
			} else {
				t.Logf("Session.User() = %v, want %v", gotUser, tt.wantUser)

			}
		})
	}
}

func TestSessionDeleteAll(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SessionDeleteAll(); (err != nil) != tt.wantErr {
				t.Errorf("SessionDeleteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Create(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				Name:     "tanaka",
				Email:    "tanaka@email.com",
				Password: "tanakapass",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
			}
			if err := user.Create(); (err != nil) != tt.wantErr {
				t.Errorf("User.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				Id:    30,
				Name:  "takahashi2",
				Email: "takahashi2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
			}
			if err := user.Update(); (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDeleteAll(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{

		{
			name: "",
			// 外部キー制約のため削除できない
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UserDeleteAll(); (err != nil) != tt.wantErr {
				t.Errorf("UserDeleteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsers(t *testing.T) {
	tests := []struct {
		name      string
		wantUsers []User
		wantErr   bool
	}{
		{
			name:      "",
			wantUsers: nil,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsers, err := Users()
			if (err != nil) != tt.wantErr {
				t.Errorf("Users() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// １行以上戻ってくればOK
			if len(gotUsers) <= 0 {
				t.Errorf("Users() = %v, want %v", gotUsers, tt.wantUsers)
			} else {
				// 全ての行を出力する
				for i, thread := range gotUsers {
					t.Logf("%v %v", i, thread)
				}
			}
		})
	}
}

func TestUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name     string
		args     args
		wantUser User
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				email: "yamada",
			},
			wantUser: User{
				Id:        20,
				Uuid:      "xxx20",
				Name:      "yamada",
				Email:     "yamada",
				Password:  "a",
				CreatedAt: time.Date(2018, 06, 18, 21, 16, 49, 0, time.FixedZone("", 0)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := UserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("UserByEmail() = %v, want %v", gotUser, tt.wantUser)
			} else {
				t.Logf("UserByEmail() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUserByUUID(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name     string
		args     args
		wantUser User
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				uuid: "xxx20",
			},
			wantUser: User{
				Id:        20,
				Uuid:      "xxx20",
				Name:      "yamada",
				Email:     "yamada",
				Password:  "a",
				CreatedAt: time.Date(2018, 06, 18, 21, 16, 49, 0, time.FixedZone("", 0)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := UserByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("UserByUUID() = %v, want %v", gotUser, tt.wantUser)
			} else {
				t.Logf("UserByUUID() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
