package data

import (
	"reflect"
	"testing"
	"time"
)

// postgresqlを起動する
// postgres -D /usr/local/var/postgres
// createdb gocrud
// psql -f data/setup.sql -d gocrud

func TestThread_CreatedAtDate(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Topic     string
		UserId    int
		CreatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "",
			fields: fields{
				Id:        0,
				Uuid:      "",
				Topic:     "",
				UserId:    0,
				CreatedAt: time.Date(2001, 1, 31, 0, 0, 0, 0, time.Local),
			},
			want: "Jan 31, 2001 at 12:00am",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thread := &Thread{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Topic:     tt.fields.Topic,
				UserId:    tt.fields.UserId,
				CreatedAt: tt.fields.CreatedAt,
			}
			if got := thread.CreatedAtDate(); got != tt.want {
				t.Errorf("Thread.CreatedAtDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_CreatedAtDate(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Body      string
		UserId    int
		ThreadId  int
		CreatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "",
			fields: fields{
				Id:        0,
				Uuid:      "",
				Body:      "",
				UserId:    0,
				ThreadId:  0,
				CreatedAt: time.Date(2001, 1, 31, 0, 0, 0, 0, time.Local),
			},
			want: "Jan 31, 2001 at 12:00am",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post := &Post{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Body:      tt.fields.Body,
				UserId:    tt.fields.UserId,
				ThreadId:  tt.fields.ThreadId,
				CreatedAt: tt.fields.CreatedAt,
			}
			if got := post.CreatedAtDate(); got != tt.want {
				t.Errorf("Post.CreatedAtDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThread_NumReplies(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Topic     string
		UserId    int
		CreatedAt time.Time
	}
	tests := []struct {
		name      string
		fields    fields
		wantCount int
	}{
		{
			name: "",
			fields: fields{
				Id:        10,
				Uuid:      "",
				Topic:     "",
				UserId:    0,
				CreatedAt: time.Time{},
			},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thread := &Thread{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Topic:     tt.fields.Topic,
				UserId:    tt.fields.UserId,
				CreatedAt: tt.fields.CreatedAt,
			}
			if gotCount := thread.NumReplies(); gotCount != tt.wantCount {
				t.Errorf("Thread.NumReplies() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestThread_Posts(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Topic     string
		UserId    int
		CreatedAt time.Time
	}
	tests := []struct {
		name      string
		fields    fields
		wantPosts []Post
		wantErr   bool
	}{
		{
			name: "",
			fields: fields{
				Id:        10,
				Uuid:      "",
				Topic:     "",
				UserId:    0,
				CreatedAt: time.Time{},
			},
			wantPosts: []Post{{
				Id:        3,
				Uuid:      "100",
				Body:      "テストです。",
				UserId:    10,
				ThreadId:  10,
				CreatedAt: time.Date(2018, 06, 18, 21, 15, 10, 0, time.FixedZone("", 0)),
			}, {
				Id:        4,
				Uuid:      "101",
				Body:      "テストです。",
				UserId:    20,
				ThreadId:  10,
				CreatedAt: time.Date(2018, 06, 18, 21, 15, 10, 0, time.FixedZone("", 0)),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thread := &Thread{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Topic:     tt.fields.Topic,
				UserId:    tt.fields.UserId,
				CreatedAt: tt.fields.CreatedAt,
			}
			gotPosts, err := thread.Posts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Thread.Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPosts, tt.wantPosts) {
				t.Errorf("Thread.Posts() = %v, want %v", gotPosts, tt.wantPosts)
			}
		})
	}
}

func TestUser_createThread(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
	}
	type args struct {
		topic string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantConv Thread
		wantErr  bool
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
			args: args{
				topic: "テストだよ！",
			},
			wantConv: Thread{
				Id:        0,
				Uuid:      "",
				Topic:     "テストだよ！",
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
			gotConv, err := user.createThread(tt.args.topic)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.createThread() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotConv.UserId != tt.wantConv.UserId || gotConv.Topic != tt.wantConv.Topic {
				t.Errorf("User.createThread() = %v, want %v", gotConv, tt.wantConv)
			} else {
				t.Logf("User.createThread() = %v, want %v", gotConv, tt.wantConv)
			}
		})
	}
}

func TestUser_CreatePost(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
	}
	type args struct {
		conv Thread
		body string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantPost Post
		wantErr  bool
	}{
		{
			name: "",
			// ユーザ
			fields: fields{
				Id:        10,
				Uuid:      "",
				Name:      "",
				Email:     "",
				Password:  "",
				CreatedAt: time.Time{},
			},
			args: args{
				conv: Thread{
					Id:        10,
					Uuid:      "",
					Topic:     "",
					UserId:    0,
					CreatedAt: time.Time{},
				},
				body: "どーも、返信してみました😃",
			},
			wantPost: Post{
				Id:        0,
				Uuid:      "",
				Body:      "どーも、返信してみました😃",
				UserId:    10,
				ThreadId:  10,
				CreatedAt: time.Time{},
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
			gotPost, err := user.CreatePost(tt.args.conv, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPost.ThreadId != tt.wantPost.ThreadId || gotPost.UserId != tt.wantPost.UserId || gotPost.Body != tt.wantPost.Body {
				t.Errorf("User.CreatePost() = %v, want %v", gotPost, tt.wantPost)
			} else {
				t.Logf("User.CreatePost() = %v, want %v", gotPost, tt.wantPost)
			}
		})
	}
}

func TestThreads(t *testing.T) {
	tests := []struct {
		name        string
		wantThreads []Thread
		wantErr     bool
	}{
		{
			name:        "",
			wantThreads: nil,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotThreads, err := Threads()
			if (err != nil) != tt.wantErr {
				t.Errorf("Threads() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// １行以上戻ってくればOK
			if len(gotThreads) <= 0 {
				t.Errorf("Threads() = %v, want %v", gotThreads, tt.wantThreads)
			} else {
				// 全ての行を出力する
				for i, thread := range gotThreads {
					t.Logf("%v %v",i,thread)
				}
			}
		})
	}
}

func TestThreadByUUID(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name     string
		args     args
		wantConv Thread
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				uuid: "ttt10",
			},
			wantConv: Thread{
				Id:        10,
				Uuid:      "ttt10",
				Topic:     "こんにちは",
				UserId:    10,
				CreatedAt: time.Date(2018, 06, 18, 21, 18, 34, 0, time.FixedZone("", 0)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConv, err := ThreadByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThreadByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConv, tt.wantConv) {
				t.Errorf("ThreadByUUID() = %v, want %v", gotConv, tt.wantConv)
			} else {
				t.Logf("ThreadByUUID = %v, want %v", gotConv, tt.wantConv)
			}
		})
	}
}

func TestThread_User(t *testing.T) {
	type fields struct {
		Id        int
		Uuid      string
		Topic     string
		UserId    int
		CreatedAt time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		wantUser User
	}{
		{
			name: "",
			fields: fields{
				Id:        0,
				Uuid:      "",
				Topic:     "",
				UserId:    20,
				CreatedAt: time.Time{},
			},
			wantUser: User{
				Id:        20,
				Uuid:      "xxx20",
				Name:      "yamada",
				Email:     "yamada",
				//Password:  "a",
				CreatedAt: time.Date(2018, 06, 18, 21, 16, 49, 0, time.FixedZone("", 0)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thread := &Thread{
				Id:        tt.fields.Id,
				Uuid:      tt.fields.Uuid,
				Topic:     tt.fields.Topic,
				UserId:    tt.fields.UserId,
				CreatedAt: tt.fields.CreatedAt,
			}
			if gotUser := thread.User(); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("Thread.User() = %v, want %v", gotUser, tt.wantUser)
			} else  {
				t.Logf("Thread.User() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
