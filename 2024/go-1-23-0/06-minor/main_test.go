package main_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestMe(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hi", func(w http.ResponseWriter, req *http.Request) {
		// New function: slices.Repeat + http.Request.Pattern
		v, _ := json.Marshal(slices.Repeat([]string{req.Pattern}, 2))
		w.Write(v)
	})

	server := httptest.NewServer(mux)

	req, err := http.NewRequest(http.MethodGet, server.URL+"/hi", nil)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("do failed: %v", err)
	}
	defer resp.Body.Close()

	val, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("do failed: %v", err)
	}

	valStr := string(val)
	if valStr != `["GET /hi","GET /hi"]` {
		t.Fatalf("invalid value: %v", valStr)
	}
}
