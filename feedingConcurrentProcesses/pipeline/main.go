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

func newInfo(n int, s string) *Info {
	return &Info{
		routine: n,
		s:       strings.ToUpper(s),
	}
}

func main() {
	var (
		width = 2
	)

	infos := orchestrate(width, data())

	for i := range infos {
		fmt.Println(i)
	}
}

func orchestrate(width int, d []string) <-chan *Info {
	ssc := produce(data())
	infos := consume(width, ssc)
	return infos
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

func produce(d []string) <-chan string {
	ssc := make(chan string)
	go func() {
		defer close(ssc)

		for _, s := range data() {
			ssc <- s
		}
	}()

	return ssc
}

func consume(width int, ssc <-chan string) <-chan *Info {
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
					infos <- newInfo(n, s)
				}
			}(iter)
		}

		wg.Wait()
	}()

	return infos
}
