package main

import (
	"go-with-tests/internal/readingfiles"
	"log"
	"os"
)

func main() {
	posts, err := readingfiles.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
}
