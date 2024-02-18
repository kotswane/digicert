package main

import (
	"bytes"
	"digicert/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bookId = ""

func TestCreate(t *testing.T) {

	router := setupRouter()

	book := dto.CreateBook{
		Author:      "Moeketsi Kotswane",
		Publication: "DigiCert BooKStore",
		ISBN:        "7896-96963-25444",
	}

	jsonValue, _ := json.Marshal(book)

	t.Log(string(jsonValue))
	router.POST("/api/v1/books/add", ctrl.Create)
	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/api/v1/books/add", bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	t.Log("jsonValue: ", req.Body)
	if err != nil {
		t.Log(err.Error())
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	t.Log(w.Body.String())
	//assert.Equal(t, "pong", w.Body.String())
}

/*
func TestRead(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/read", nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetById(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	var endpoint = fmt.Sprintf("http://localhost:%s/api/v1/books/getbyid/", apiPortNumber)
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestUpdate(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	var endpoint = fmt.Sprintf("http://localhost:%s/api/v1/books/update/", apiPortNumber)
	req, _ := http.NewRequest(http.MethodPut, endpoint, nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestDelete(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	var endpoint = fmt.Sprintf("http://%s:%s/api/v1/books/delete/", ,apiPortNumber)
	req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}*/
