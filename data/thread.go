package data

import "time"

type Thread struct {
	Id			int
	Uuid 		string
	Topic		string
	UserId 		int
	CreatedAt	time.Time
}


type Post struct {
	Id 			int
	Uuid 		string
	Body 		string
	UserId 		int
	ThreadId	int		// 親id
	CreatedAt 	time.Time
}

func (thread* Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// NumRepliesは、thread_idをキーに返信の数を返す関数です。
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// thread_idのすべてのpostを取りだす
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thrrad_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return 
		}
	}
}

func (user *User) createThread(topic string) (conv Thread, err error) {
	statement := "INSERT INTO threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id. uuid, topic, user_id,created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}