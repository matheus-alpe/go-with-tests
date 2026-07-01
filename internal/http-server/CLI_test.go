package httpserver

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}
		cli := NewCLI(playerStore, in)

		cli.PlayPoker()
		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}
		cli := NewCLI(playerStore, in)

		cli.PlayPoker()
		assertPlayerWin(t, playerStore, "Cleo")
	})
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore, want string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := store.winCalls[0]

	if got != want {
		t.Errorf("didn't record correct winner, got %q want %q", got, want)
	}
}
