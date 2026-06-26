package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPLayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(http.MethodPost, player))
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(http.MethodPost, player))
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(http.MethodPost, player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newScoreRequest(http.MethodGet, player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
