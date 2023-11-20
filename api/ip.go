package api

import (
	"fmt"
	"net/http"
)

func (s *server) HandleIP(w http.ResponseWriter, r *http.Request) {
	// Extract IP from the request
	ip := r.Context().Value(key)
	if ip == nil {
		// Send it to the request
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send it to the request
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Your IP  : %s", ip)))
}
