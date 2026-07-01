package main

import (
	"fmt"
	httpserver "go-with-tests/internal/http-server"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to record a win")

	store, cleanup, err := httpserver.NewFileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	game := httpserver.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
