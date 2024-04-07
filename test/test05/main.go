package main

import (
	"fmt"
	"github.com/thinkeridea/go-extend/exnet"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// var r *http.Request
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	fmt.Println(ip)
}
