package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Info struct {
	routine int
	s       string
}

func main() {
	var (
		width = 2
	)

	ssc := make(chan string)
	go func() {
		defer close(ssc)

		for _, s := range data() {
			ssc <- s
		}
	}()

	infos := make(chan *Info)

	go func() {
		defer close(infos)

		var wg sync.WaitGroup
		wg.Add(width)

		for iter := 0; iter < width; iter++ {
			go func(n int) {
				defer wg.Done()

				for s := range ssc {
					time.Sleep(time.Second * 2)
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
