/*
 * Use a channel to wait, instead of WaitGroup.
 */
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	wait := make(chan struct{})

	m := map[int]int{}

	go func() {
		for i := 0; i < 10; i++ {
			m[rand.Intn(1000000)] = i
		}
		close(wait)
	}()

	<-wait
	for k, v := range m {
		fmt.Printf("key: %v, val: %d\n", k, v)
	}
}
