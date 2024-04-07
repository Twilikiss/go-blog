package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (api *apiHandler) TestHeader(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Header的情况如下\n")
	for key, value := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", key, value)
	}
}

func (api *apiHandler) TestIndex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "Elysia"
	indexData.Desc = "学习go语言简易博客"
	data, _ := json.Marshal(indexData)
	_, err := writer.Write(data)
	if err != nil {
		log.Println(err)
	}
}
