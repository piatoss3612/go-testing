package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {
	app := application{}
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.3.1", "", false},
		{"", "", "hello:world", false},
	}

	// create a dummy handler that will use to check context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that the value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, " not present")
		}

		// make sure we got a string back
		_, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
	})

	for _, e := range tests {
		// create the handler to test
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	app := application{}

	ctx := context.WithValue(context.Background(), contextUserKey, "127.0.0.1")

	ip := app.ipFromContext(ctx)

	if !strings.EqualFold("127.0.0.1", ip) {
		t.Errorf("expected 127.0.0.1 but got %s", ip)
	}
}
