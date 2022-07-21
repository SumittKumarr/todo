package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID       string `db:"id"     json:"id"`
	Name     string `db:"name"     json:"name"`
	Password string `db:"password" json:"password"`
}

type Task struct {
	ID          string    `db:"id"     json:"id"`
	Name        string    `db:"name"     json:"name"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	IsCompleted bool      `db:"is_completed" json:"isCompleted"`
	ArchivedAt  time.Time `db:"archived_at"  json:"archivedAt"`
	UserId      string    `db:"user_id"      json:"userId"`
}

type Session struct {
	ID         string       `db:"id"     json:"id"`
	CreatedAt  time.Time    `db:"created_at" json:"createdAt"`
	ExpiryTime time.Time    `db:"expiry_time" json:"expiryTime"`
	ArchivedAt sql.NullTime `db:"archived_at"  json:"archivedAt"`
	UserId     string       `db:"user_id"      json:"userId"`
}
