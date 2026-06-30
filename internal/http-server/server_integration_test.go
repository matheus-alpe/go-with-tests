package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanup := createTempFile(t, "[]")
	defer cleanup()

	store, err := NewFileSystemStore(database)
	assertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))
	server.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))
	server.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newPlayersRequest(http.MethodGet, player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)
		got := decodeFromResponse[League](t, response.Body)
		want := League{{"Pepper", 3}}
		assertDeepEqual(t, got, want)
	})
}
