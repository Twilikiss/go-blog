package main

import (
	"fmt"
	"time"
)

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

func main() {
	var u = &UserInfo{}
	if u == nil {
		fmt.Println("u为nil")
	} else {
		fmt.Println("u不为nil")
	}

	var data01 = make([]int, 10)
	var data02 = new([]int)
	if data01 == nil {
		fmt.Printf("%v is nil.\n", data01)
	}
	if data02 == nil {
		fmt.Printf("%v is nil.\n", data02)
	}
	fmt.Println(data01)
	fmt.Println(data02)
}
