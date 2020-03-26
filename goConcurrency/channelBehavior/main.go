// This file was inspired by The Behavior Of Channels by William Kennedy
// written October 24, 2017
// https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(6)

	waitForTask()
	waitForResult()
	fanOut()
	selectDrop()
	waitForTasks()
	withTimeout()

	wg.Wait()
}

// Signal With Data - Guarantee - Unbuffered Channels
// Listing 5 (modified) - Scenario 1 - Wait For Task
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch // Receive ("employee waits for task")

		fmt.Println(p)
		// Employee performs work.
		// Employee is done and free to go.

		wg.Done()
	}()

	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

	ch <- "paper 1" // Send data
}

// Signal With Data - Guarantee - Unbuffered Channels
// Listing 6 (modified) - Scenario 2 - Wait for Result
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

		ch <- "paper 2" // Send paper

		// Employee is done and free to go
		wg.Done()
	}()

	p := <-ch
	fmt.Println(p)
}

// Signal With Data - No Guarantee - Buffered Channels >1
// Listing 7 (modified) - Scenario 1 - Fan Out
func fanOut() {
	emps := 20

	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

			ch <- fmt.Sprintf("paper 3 from emp %d", i)
		}(e)
	}
	wg.Done()

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		emps--
	}
}

// Signal With Data - No Guarantee - Buffered Channels >1
// Listing 8 (modified) - Scenario 2 - Drop
func selectDrop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
		wg.Done()
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper 4":
			fmt.Println("manager : send ack")
		default:
			fmt.Println("manager: drop")
		}
	}

	close(ch)
}

// Signal With Data - Delayed Guarantee - Buffered Channel 1
// listing 9 (modified) - Scenario 1 - Wait For Tasks
func waitForTasks() {
	ch := make(chan string, 1)

	go func() {
		for p := range ch {
			fmt.Println("employee : working :", p)
		}
		wg.Done()
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper 5"
	}
	close(ch)
}

// Signal Without Data - Context
// listing 10 (modified) - Scenario 1 - Wait For Tasks
func withTimeout() {
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "paper 6"
		wg.Done()
	}()

	select {
	case p := <-ch:
		fmt.Println("work complete", p)
	case <-ctx.Done():
		fmt.Println("moving on")
	}
}
