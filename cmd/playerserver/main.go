package main

import (
	httpserver "go-with-tests/internal/http-server"
	"log"
	"net/http"
)

func main() {
	server := httpserver.NewPlayerServer(httpserver.NewInMemoryPLayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
