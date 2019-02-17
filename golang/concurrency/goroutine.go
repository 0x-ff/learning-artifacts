package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var salutation string

func main() {
	salutation = "hello"
	wg.Add(1)
	go sayHello()
	wg.Wait()
	fmt.Println(salutation)
}

func sayHello() {
	defer wg.Done()
	salutation = "welcome"
}
