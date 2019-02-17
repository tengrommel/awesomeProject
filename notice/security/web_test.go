package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	HealthCheckHandler(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHealthCheckHandler2(t *testing.T) {
	reqData := struct {
		Info string `json:"info"`
	}{Info: "P123451"}
	reqBody, _ := json.Marshal(reqData)
	fmt.Println("input:", string(reqBody))
	req := httptest.NewRequest(http.MethodPost, "/health-check", bytes.NewReader(reqBody))
	req.Header.Set("userid", "wdt")
	req.Header.Set("commpay", "brk")
	rr := httptest.NewRecorder()
	HealthCheckHandler(rr, req)
	result := rr.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(body))
	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status 200", result.StatusCode)
	}
}
