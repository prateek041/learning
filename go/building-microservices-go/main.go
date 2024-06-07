package main

import (
	"io"
	"log"
	"net/http"
)

type testingThing struct {
	messsage string
}

func (t *testingThing) Read(p []byte) (n int, err error) {
	log.Println("Hello this is me")
	return 0, io.EOF
}

func main() {
	testObj := testingThing{messsage: "hello world"}
	io.ReadAll(&testObj)
	log.Println("I was run")

	// HandleFunc registers a [pattern] handler with the [DefaultServeMux]
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Print("Hello world")
	})

	// ListenAndServe listens on the TCP address [addr] (":9090") and calls the [Serve] method on the handler. Since we have not provided any handler here, it defaults to [DefaultServeMux], where [DefaultServeMux] is a [ServeMux]
	http.ListenAndServe(":9090", nil)
}
