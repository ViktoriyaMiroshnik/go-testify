package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// здесь нужно создать запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body.String())

	if responseRecorder.Code == http.StatusBadRequest {
		assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
	}

	if responseRecorder.Code != http.StatusBadRequest {
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.NotEmpty(t, responseRecorder.Body.String())
		assert.Equal(t, responseRecorder.Body.String(), "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент")
		assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
	}
}
