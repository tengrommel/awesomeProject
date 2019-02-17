package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name   string
		value  string
		double int
		status int
		err    string
	}{
		{name: "double of two", value: "2", status: http.StatusOK},
	}
	for _, tc := range tt {
		req, err := http.NewRequest("GET", "localhost:8080/double?v="+tc.value, nil)
		if err != nil {
			t.Fatalf("could not createc request: %v", err)
		}
		rec := httptest.NewRecorder()
		doubleHandler(rec, req)

		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.status {
			t.Errorf("expected status %v; got %v", tc.status, res.StatusCode)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read a response: %v", err)
		}

		d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
		if err != nil {
			t.Fatalf("expected and integer; got %s", b)
		}
		if d != 4 {
			t.Fatalf("expected double to be 4; got %v", d)
		}
	}

}
