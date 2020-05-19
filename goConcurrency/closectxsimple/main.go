package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("run program")
	runSay(ctx)

	time.Sleep(1 * time.Minute)

	fmt.Println("stop program")
	cancel()

}

func sayOne() {
	fmt.Println("one")
}

func sayTwo() {
	fmt.Println("two")
}

func sayThree() {
	fmt.Println("three")
}

func runSay(ctx context.Context) {
	go func() {
		for {
			timer := time.NewTicker(2 * time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
				fmt.Println("--------context cancelled one")
				return
			case <-timer.C:
				sayOne()
			}

		}
	}()

	go func() {
		for {
			timer := time.NewTicker(2 * time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
				fmt.Println("----context cancelled two")
				return
			case <-timer.C:
				sayTwo()
			}
		}
	}()

	go func() {
		for {
			timer := time.NewTicker(2 * time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
				fmt.Println("context cancelled three")
				return
			case <-timer.C:
				sayThree()
			}
		}
	}()
}
