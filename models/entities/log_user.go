package entities

import "time"

type LogUserEntity struct {
	Id           int
	ActiveUserId int
	Nick         string
	Email        string
	Password     []byte
	IconURL      string
	Roll         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
