package main

import (
	"go-with-tests/internal/pointersanderrors"
	"log"
)

func main() {
	_, err := pointersanderrors.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
