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
	ss := data()

	var infos []*Info

	var wg sync.WaitGroup
	wg.Add(len(ss))

	for i, s := range ss {
		go func(n int, v string) {
			defer wg.Done()
			i := &Info{
				routine: n,
				s:       strings.ToUpper(v),
			}
			infos = append(infos, i)
		}(i, s)
	}

	wg.Wait()
	for _, i := range infos {
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
