package httpserver

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
	GetLeague() League
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	value, found := i.store[name]
	return value, found
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() (league League) {
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return
}
