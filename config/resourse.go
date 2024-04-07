// Package config 将对应的静态资源部署上线
package config

import (
	"log"
	"net/http"
	"os"
)

func LoadStatic() {
	http.HandleFunc("/resource/css/index.css", indexCssHandler)
	http.HandleFunc("/resource/js/index.js", indexJsHandler)
	http.HandleFunc("/resource/images/favicon.ico", faviconHandler)
	http.HandleFunc("/resource/images/logo.png", logoHandler)
}

func indexJsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	data, err := os.ReadFile("./public/resource/js/index.js")
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(data)
}

func faviconHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	data, err := os.ReadFile("./public/resource/images/favicon.ico")
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(data)
}

func logoHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	data, err := os.ReadFile("./public/resource/images/logo.png")
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(data)
}

func indexCssHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/css")
	data, err := os.ReadFile("./public/resource/css/index.css")
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(data)
}
