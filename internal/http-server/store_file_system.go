package httpserver

import (
	"encoding/json"
	"io"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func NewFileSystemStore(db io.ReadWriteSeeker) *FileSystemStore {
	return &FileSystemStore{db}
}

func (f *FileSystemStore) GetLeague() League {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemStore) GetPlayerScore(name string) (int, bool) {
	if player := f.GetLeague().Find(name); player != nil {
		return player.Wins, true
	}

	return 0, false
}

func (f *FileSystemStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}
