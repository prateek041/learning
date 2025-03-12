package main

// Read more in Obsidian Resources/go/sync

import (
	"fmt"
	"sync"
)

func main() {
	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan bool)
	fmt.Println("STARING THE FIRST LOOP")
	for i := 0; i < 10; i++ {
		go func() {
			var once sync.Once
			fmt.Println("Starting to execute")
			once.Do(onceBody)
			fmt.Println("Executed")
			done <- true
			fmt.Println("Executed")
		}()
	}

	fmt.Println("Out of the first loop")
	for i := 0; i < 10; i++ {
		fmt.Printf("iteration number %d \n", i)
		fmt.Println("read")
		<-done
	}
}
