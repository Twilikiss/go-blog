package main

import "fmt"

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"` // 文章ID
	Title string `orm:"title" json:"title"`
}

func test01(pid int) SearchResp {
	data01 := SearchResp{
		Pid:   pid,
		Title: "wwww",
	}
	return data01
}

func test02(in SearchResp) {
	in.Title = "sss"
}

func main() {
	data01 := test01(1)
	test02(data01)
	fmt.Println(data01)
}
