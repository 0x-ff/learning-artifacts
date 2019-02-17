package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var stringStream chan string
	stringStream = make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range stringStream {
			fmt.Printf("Received: %s\n", result)
		}
		fmt.Println("Done receiving!")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(stringStream)
		for _, s := range []string{"a", "b", "c"} {
			select {
			case stringStream <- s:
			}
		}
	}()

	wg.Wait()
	fmt.Println("All done!")
}
