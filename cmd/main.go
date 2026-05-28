package main

import (
	"fmt"
	"go-with-tests/internal/helloworld"
)

func main() {
	fmt.Println("Hello, world")
	fmt.Println(helloworld.Hello("world", ""))
}
