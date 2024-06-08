package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prateek041/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "test-app", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodByeHandler := handlers.NewGoodBye(l)

	sm := http.NewServeMux()

	sm.Handle("/", helloHandler)
	sm.Handle("/goodbye", goodByeHandler)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Server started
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
