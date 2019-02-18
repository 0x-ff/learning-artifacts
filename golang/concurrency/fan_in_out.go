package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	primeFinder := func(
		done <-chan interface{},
		randIntStream <-chan int,
	) <-chan int {
		primeStream := make(chan int)
		go func() {
			defer close(primeStream)
			for v := range randIntStream {
				select {
				case <-done:
					return
				case primeStream <- v:
				}
			}
		}()
		return primeStream
	}
	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	toInt := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}
	take := func(
		done <-chan interface{},
		valueStream <-chan int,
		num int,
	) <-chan int {
		takeStream := make(chan int)
		go func() {
			defer close(takeStream)
			for i := num; i > 0 || i == -1; {
				if i != -1 {
					i--
				}
				select {
				case <-done:
					return
				case takeStream <- (<-valueStream):
				}
			}
		}()
		return takeStream
	}
	fanIn := func(
		done <-chan interface{},
		channels ...<-chan int,
	) <-chan int {
		var wg sync.WaitGroup
		multiplexedStream := make(chan int)

		multiplex := func(c <-chan int) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	rand := func() interface{} { return rand.Intn(50000000) }
	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders\n", numFinders)
	finders := make([]<-chan int, numFinders)
	fmt.Println("Primes:")

	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
