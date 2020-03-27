// This sandbox was taken (and modified) from the signal handling code at
// https://github.com/openfaas-incubator/golang-http-template/blob/master/template/golang-http/main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		sig := make(chan os.Signal, 1)

		// Send SIGTERM with `kill -s TERM <pid>`
		signal.Notify(sig, syscall.SIGTERM)

		<-sig

		fmt.Println("SIGTERM received.")
		wg.Done()
	}()

	wg.Wait()
}
