package models

import "time"

type User struct {
	UserId       uint64
	UserName     string
	UserEmail    string
	UserPassword string
	CreatedAt    time.Time
}
