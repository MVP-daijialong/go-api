package common

import (
	"net"
	"net/http"
	"strings"
)

// getClientIP 获取客户端 IP 地址
func GetClientIP(r *http.Request) string {
	// 尝试从 X-Forwarded-For 头获取
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// X-Forwarded-For 可能包含多个 IP 地址，取第一个非空的 IP 地址
		ips := strings.Split(xff, ",")
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			if ip != "" {
				return ip
			}
		}
	}

	// 尝试从 X-Real-IP 头获取
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return xri
	}

	// 从 RemoteAddr 获取
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
