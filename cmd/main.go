package main

import (
	"go-with-tests/internal/dependencyinjection"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dependencyinjection.Greet(w, "word")
		dependencyinjection.Greet(os.Stdout, "Eloise <<<\n")
	})))
}
