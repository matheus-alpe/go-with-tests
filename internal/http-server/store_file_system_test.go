package httpserver

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanup()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertDeepEqual(t, got, want)

		got = store.GetLeague()
		assertDeepEqual(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanup()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		got, _ := store.GetPlayerScore("Chris")
		want := 33

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanup()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		store.RecordWin("Chris")

		got, _ := store.GetPlayerScore("Chris")
		want := 34

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanup()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		store.RecordWin("Pepper")

		got, _ := store.GetPlayerScore("Pepper")
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanup := createTempFile(t, "")
		defer cleanup()

		_, err := NewFileSystemStore(database)
		assertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanup()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertDeepEqual(t, got, want)

		got = store.GetLeague()
		assertDeepEqual(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpFile.Write([]byte(initialData))

	cleanup := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, cleanup
}
