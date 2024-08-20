package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"

	log.Printf("Server starting on port %v\n", PORT)

	http.HandleFunc("/", helloHandlerFunc)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println("Error starting the server", err)
	}
}

func helloHandlerFunc(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from the Server")
}
