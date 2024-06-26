package main

import (
	"net/http"
	"net/http/httptest"

	"bytes"
	"encoding/json"
	"testing"

	models "github.com/Laeeqdev/urlShortner/API/Models"
	"github.com/stretchr/testify/assert"
)

func TestShortenEndpoint(t *testing.T) {
	payload := models.ShortUrlRequest{"https://www.infracloud.io/cloud-native-open-source-contributions/"}
	jsonPayload, _ := json.Marshal(payload)
	request, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonPayload))
	response := httptest.NewRecorder()
	InitializeApp().MyRouter().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
