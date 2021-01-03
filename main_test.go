package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func er(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occured during test: %v", err)
	}
}

func setUp(t *testing.T) {
	data := user{
		"Johnny",
		22,
		"Camoro Drive",
	}
	err := WriteFile(data)
	er(err, t)
}

func tearDown(t *testing.T) {
	file, _ := os.OpenFile("users.bin", os.O_TRUNC, 0666)
	defer file.Close()
}

func TestCreateUserHandler(t *testing.T) {
	jsonStr := []byte(`{"name": "John Doe", "age": 23, "address": "1 Hampton Drive, CA"}`)
	req, err := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonStr))
	er(err, t)
	rr := httptest.NewRecorder()

	http.HandlerFunc(createUser).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Invalid status code expected %d, got %d instead", http.StatusCreated, status)
	}

	expected := string(`{"message":"User created","status":"success"}`)
	assert.JSONEq(t, expected, rr.Body.String(), "Response body does not match")

	tearDown(t)
}

func TestListUsersHandler(t *testing.T) {
	setUp(t)

	req, err := http.NewRequest("GET", "/users/list", nil)
	er(err, t)
	rr := httptest.NewRecorder()
	http.HandlerFunc(listUsers).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Invalid status code expected %d, got %d instead", http.StatusCreated, status)
	}
	expected := string(`{"status":"success","users":[{"name":"Johnny","age":22,"address":"Camoro Drive"}]}`)
	assert.JSONEq(t, expected, rr.Body.String(), "Response body does not match")

	tearDown(t)
}

func TestCanWriteUserDataToFile(t *testing.T) {
	setUp(t)
	tearDown(t)
}

func TestCanReadBufferFromSavedBinary(t *testing.T) {
	t.Skip()
	jsonStr := []byte(`{"name": "John Doe", "age": 23, "address": "1 Hampton Drive, CA"}`)
	r := bytes.NewReader(jsonStr)
	users, err := handleFileBuffer(r)
	er(err, t)
	assert.Equal(t, 1, len(users), "Returned length mismatched")
}
