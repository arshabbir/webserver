package api

import (
	"context"
	"net"
	"net/http"
)

type contextType string

var key contextType = "IP"

func addIPToRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := getIP(r)
		// Implement the middleware logic here

		ctx := context.WithValue(r.Context(), key, ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getIP(r *http.Request) string {
	//Check forwded flag
	fwdAddress := r.Header.Get("X-Forwarded-For")
	if len(fwdAddress) != 0 {
		return fwdAddress
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown"
	}

	if len(ip) == 0 {
		return "unknown"
	}

	return ip

}
