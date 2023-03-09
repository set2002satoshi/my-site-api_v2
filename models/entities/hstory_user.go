package entities

import "time"

type HistoryUserEntity struct {
	Id           int
	ActiveUserId int
	Nick         string
	Email        string
	Password     []byte
	ImgURL       string
	ImgKey       string
	Roll         string
	Revision     int
	CreatedAt    time.Time // active　timeを格納する
	DeletedAt    time.Time // 履歴が作られた時間を表示格納する
}
