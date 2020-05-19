package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*

Questions:
- If I want to protect two separate fields in a struct, should I make two sync.RWMutexes? Or protect them with one?
- Is it ok to let go routines slam shut because the main one did?

*/
func main() {
	stor := new()
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("run program")
	stor.run(ctx)

	time.Sleep(15 * time.Second)
	fmt.Println("stop program")
	cancel()

	// Give other go routines time to stop.
	time.Sleep(1 * time.Second)
}

type storage struct {
	datmux sync.RWMutex
	data   map[time.Time]time.Time
	imux   sync.RWMutex
	i      int
}

func new() *storage {
	s := storage{
		data: map[time.Time]time.Time{},
		i:    0,
	}
	return &s
}

func (s *storage) run(ctx context.Context) {
	go func() {
		for {
			timer := time.NewTicker(2 * time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
				fmt.Println("x context cancelled 1")
				return
			case <-timer.C:
				s.add()
			}

		}
	}()

	go func() {
		for {
			timer := time.NewTicker(2 * time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
				fmt.Println("x context cancelled 2")
				return
			case <-timer.C:
				s.rem()
			}
		}
	}()

	go func() {
		for {
			timer := time.NewTicker(2 * time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
				fmt.Println("x context cancelled 3")
				return
			case <-timer.C:
				s.get()
			}
		}
	}()
}

func (s *storage) get() time.Time {
	var t time.Time

	fmt.Println("# get key and counter")

	s.datmux.RLock()
	for k := range s.data {
		fmt.Println("# got ", k)
	}
	s.datmux.RUnlock()

	s.imux.RLock()
	fmt.Println("# i is", s.i)
	s.imux.RUnlock()

	return t
}

// Incrementally add a key and value to map.
func (s *storage) add() {
	fmt.Println("+ add key")

	t := time.Now()

	s.datmux.Lock()
	s.data[t] = t
	s.datmux.Unlock()

	fmt.Println("+ key added", t)

	s.imux.RLock()
	s.i++
	s.imux.RUnlock()

	fmt.Println("+ i++")
}

// Remove a random key from map.
func (s *storage) rem() {
	fmt.Println("- remove key")

	s.datmux.Lock()
	for k := range s.data {
		fmt.Println("- removing ", k)
		delete(s.data, k)
		break
	}
	s.datmux.Unlock()

	fmt.Println("- remove unlocked")
}
