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
	for iter := 0; iter < width; iter++ {
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
	// The problem is that infos channel can't be filled above for loop because
	// there's no one there to take from the channel (below). But code below won't
	// happen because we're asking them to wg.Wait(). 
	// Channels are a way to signal to wait, but Done will need to be added.
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
