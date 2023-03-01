package entities

import "time"

type TBLUserEntity struct {
	UserId    int
	Nickname  string
	Email     string
	Password  []byte
	IconURL   string
	Roll      string
	Revision  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
