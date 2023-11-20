package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_registerroutes(t *testing.T) {

	tests := []struct {
		name    string
		pattren string
		method  string
	}{
		{"Ping", "/ping", "GET"},
	}

	//Initlize
	s := &server{addr: ":8080", mux: chi.NewRouter()}
	handler := registerRoutes(s)

	//fmt.Println(handler.(*chi.Mux).Match(chi.NewRouteContext(), "GET", "/ping"))
	for _, test := range tests {
		if !handler.(*chi.Mux).Match(chi.NewRouteContext(), test.method, test.pattren) {
			t.Errorf("Route %s not registered with method : %s", test.pattren, test.method)
		}

	}
}

func Test_pingHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	s := &server{}
	s.HandlePing(resp, req)

	if resp.Result().StatusCode != 200 {
		t.Errorf("Unexpected response code %d", req.Response.StatusCode)
	}

	if resp.Body.String() != "pong" {
		t.Errorf("Unexpected response message %s", resp.Body.String())
	}
}
