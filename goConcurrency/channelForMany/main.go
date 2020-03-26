// Use an unbuffered channel process numbers loaded by two go routines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(2)

	// Create an unbuffered channel
	ch := make(chan int)

	// Receive letters and numbers from channel and print.
	// Note: receive must happen first in this case else it
	// will deadlock.
	go receiveOnChannel(ch)

	// Spin up goroutine that puts 10 letters into the channel.
	go randomLowNumber(10, ch)
	// Spin up goroutine that puts 10 numbers into the channel.
	go randomHighNumber(10, ch)

	// Wait for the letters to be loaded so that the channel can be closed
	wg.Wait()

	// Close the channel so receiveOnChannel() stops range.
	close(ch)
}

func randomHighNumber(qty int, ch chan int) {
	for i := 0; i < qty; i++ {
		ch <- rand.Intn(99) + 10000
	}
	wg.Done()
}

func randomLowNumber(qty int, ch chan int) {
	for i := 0; i < qty; i++ {
		ch <- rand.Intn(99)
	}
	wg.Done()
}

func receiveOnChannel(ch chan int) {
	for c := range ch {
		fmt.Println(c)
	}
}
