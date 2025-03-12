package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloHandlerResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id"`
}

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
	message := &helloHandlerResponse{
		Message: "Hello this is server resopnse",
		Author:  "Prateek Singh",
		Date:    "2024-08-21",
		Id:      69,
	}

	response, err := json.Marshal(message)
	if err != nil {
		panic("Oops")
	}

	fmt.Fprint(rw, string(response))
}
