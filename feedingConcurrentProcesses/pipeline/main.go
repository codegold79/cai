package main

import (
	"fmt"
	"strings"
	"sync"
)

type Info struct {
	routine int
	s       string
}

func main() {
	var (
		width = 2
	)

	ssc := make(chan string) // Producer sends messages through the ssc.
	go func() {
		defer close(ssc)

		for _, s := range data() {
			ssc <- s
		}
	}()


	infos := make(chan *Info)

	go func() { // consumer
		defer close(infos)

		var wg sync.WaitGroup
		wg.Add(width)

		for iter := 0; iter < width; iter++ {
			go func(n int) { // digester
				defer wg.Done()

				for s := range ssc {
					i := &Info{
						routine: n,
						s:       strings.ToUpper(s),
					}
					infos <- i
				}
			}(iter)
		}

		wg.Wait()
	}()

	for i := range infos {
		fmt.Println(i)
	}
}

func data() []string {
	return []string{
		"this",
		"then",
		"that",
		"and",
		"the",
		"other",
		"thing",
	}
}
