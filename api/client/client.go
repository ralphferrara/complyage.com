package client

import (
	"net"
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		parts := strings.Split(ip, ",")
		ip = strings.TrimSpace(parts[0])
	}

	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}

	if ip == "" {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		} else {
			ip = host
		}
	}

	// Handle ::1 for local dev
	if ip == "::1" {
		return "127.0.0.1"
	}

	return ip
}
