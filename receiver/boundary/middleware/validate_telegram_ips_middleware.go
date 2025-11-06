package middleware

import (
	"net"
	"net/http"
)

var telegramNets = []string{
	"149.154.160.0/20",
	"91.108.4.0/22",
}

func isTelegramIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	for _, cidr := range telegramNets {
		_, netblock, _ := net.ParseCIDR(cidr)
		if netblock.Contains(parsedIP) {
			return true
		}
	}
	return false
}

func ValidateTelegramIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ip, _, err := net.SplitHostPort(r.RemoteAddr)
		//if err != nil || !isTelegramIP(ip) {
		//	http.Error(w, "Acesso não autorizado", http.StatusForbidden)
		//	return
		//}

		next.ServeHTTP(w, r)
	})
}
