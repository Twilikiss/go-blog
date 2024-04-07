package utils

import (
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
)

func GetClientIp(r *http.Request) string {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	return ip
}
