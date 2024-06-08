package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prateek041/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "test-app", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodByeHandler := handlers.NewGoodBye(l)

	http.Handle("/", helloHandler)
	http.Handle("/goodbye", goodByeHandler)

	// ListenAndServe listens on the TCP address [addr] (":9090") and calls the [Serve] method on the handler. Since we have not provided any handler here, it defaults to [DefaultServeMux], where [DefaultServeMux] is a [ServeMux]
	http.ListenAndServe(":9090", nil)
}
