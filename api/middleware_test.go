package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_addIPToRequest(t *testing.T) {

	tests := []struct {
		name     string
		req      *http.Request
		expected string
	}{
		{"IP missing case", &http.Request{RemoteAddr: ""}, "unknown"},
		{"correct IP", &http.Request{RemoteAddr: "192.168.0.1:8080"}, "192.168.0.1"},
		{"X-forward-case", &http.Request{Header: getFwdHeader()}, "192.168.0.2"},
	}

	idx := 0
	//Create handler

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check if the context value for ip exits

		val := r.Context().Value(key)
		if val == nil {
			t.Errorf("Context not populated as expected")
		}

		ip, ok := val.(string)
		if !ok {
			t.Errorf("Expecting IP as string  ")
		}
		if tests[idx].expected != ip {
			t.Errorf("Unexpected header from handler ...%s", ip)
		}
		idx++

	})

	// for _, test := range tests {
	// 	ret := getIP(test.req)
	// 	if ret != test.expected {
	// 		t.Errorf("Unexpected return value : %s ", ret)
	// 	}
	// }
	next := addIPToRequest(handler)
	for _, test := range tests {
		next.ServeHTTP(httptest.NewRecorder(), test.req)
	}
	// Test for remoteaddressmissing
}

func getFwdHeader() http.Header {
	h := http.Header{}
	h.Add("X-Forwarded-For", "192.168.0.2")
	return h

}
