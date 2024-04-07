package main

import (
	"fmt"
	"go-blog/log"
	"strconv"
	"strings"
)

func main() {
	path := "/p/1000000.html"
	log.Info("get the path:", path)
	cIdStr := strings.TrimPrefix(path, "/p/")
	cIdF, found := strings.CutSuffix(cIdStr, ".html")
	if !found {
		log.Error("failed to find suffix:[.html]")
		return
	}
	cId, err := strconv.Atoi(cIdF)
	if err != nil {
		log.Error("category disable to get the param from path:", err)
		return
	}
	fmt.Println(cId)
}
