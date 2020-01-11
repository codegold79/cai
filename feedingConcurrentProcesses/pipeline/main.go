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
	// configuration status, place at top in a var block
	var (
		width = 2
	)
	// ssc is strings channel
	ssc := make(chan string)
	go func() {
		// .
		for _, s := range data() {
			ssc <- s
		}
	}()

	infos := make(chan *Info)
	// ..

	var wg sync.WaitGroup
	wg.Add(width)
	// changed i to iter to show it is out of scope of the anonymous function.
	for iter := 0; iter < width; iter++ {
		// don't be tempted to add to wait group here. it migth be cleared by go routine
		// called, then throw done, before having the chance to add the next. This is why
		// you should only use Add once.
		go func(n int) {
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
