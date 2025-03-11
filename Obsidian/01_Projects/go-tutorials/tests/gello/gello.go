package main

import (
	"fmt"
	"log"

	"example.com/hello"
)

func main() {
	greeting, err := hello.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)
}
