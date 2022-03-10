package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePost(t *testing.T) {
	router := SetupRouter()
	initDatabase()
	payload := []byte(`{"title":"Teste Titulo 3","description":"Teste Descrição 3"}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/post",
		bytes.NewBuffer(payload),
	)
	req.Header.Set("Content-Type", "application/json")
	res, _ := router.Test(req)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NotEmpty(t, res)
}

func TestGetAllPosts(t *testing.T) {
	router := SetupRouter()
	initDatabase()
	resp, _ := router.Test(httptest.NewRequest("GET", "/post", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
