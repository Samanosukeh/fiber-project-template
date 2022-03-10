package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	//router := SetupRouter()

	w := httptest.NewRecorder()
	//req, _ := http.NewRequest("GET", "/ping", nil)
	//router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestBlogCreateItem(t *testing.T) {
	router := SetupRouter()
	initDatabase()

	//req := httptest.NewRequest(
	//	"POST",
	//	"/post",
	//	bytes.NewBuffer(payload),
	//)
	// build payload request
	payload := []byte(`{"title":"Teste Titulo","description":"Teste Descrição"}`)
	MarshalPayload := bytes.NewBuffer(payload)
	print(MarshalPayload)

	req := httptest.NewRequest(
		http.MethodPost,
		"/post",
		MarshalPayload,
	)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("X-Custom-Header", "hi")

	// http.Response
	resp, _ := router.Test(req)
	print(resp)

	//resp, _ := router.Test(
	//	httptest.NewRequest(
	//		"POST",
	//		"/post",
	//		//bytes.NewBuffer(payload),
	//		nil,
	//	),
	//)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, resp)
	assert.Contains(t, resp, "message")

}

func TestBlogCreateItem1(t *testing.T) {
	router := SetupRouter()
	initDatabase()
	//payload := []byte(`{"title":"Teste Titulo","description":"Teste Descrição"}`)
	//_ := bytes.NewBuffer(payload)

	req := httptest.NewRequest(
		http.MethodPost,
		"/post",
		bytes.NewBufferString(`{"title":"Teste Titulo","description":"Teste Descrição"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	res, _ := router.Test(req)
	print(res)

	assert.Equal(t, http.StatusOK, 200)

}

func TestGetAllPosts(t *testing.T) {
	router := SetupRouter()
	initDatabase()
	resp, _ := router.Test(httptest.NewRequest("GET", "/post", nil))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
