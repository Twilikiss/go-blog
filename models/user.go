package models

import "time"

type User struct {
	Uid      int
	UserName string
	Passwd   string
	Avatar   string
	CreateAt time.Time
	UpdateAt time.Time
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
