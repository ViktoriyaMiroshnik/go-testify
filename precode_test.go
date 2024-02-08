package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenMissingCount(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	if responseRecorder.Code == http.StatusBadRequest {
		assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
	}
}

func TestMainHandlerAllCafes(t *testing.T) {

	totalCount := len(cafeList["moscow"])
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusBadRequest {
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.NotEmpty(t, responseRecorder.Body.String())
		assert.Equal(t, responseRecorder.Body.String(), "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент")
		assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
	}
}
