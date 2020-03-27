package main

import (
	"context"
	"fmt"
	"sync"
)

type runner struct {
	name string
}

func main() {
	ctx := context.Background()
	item := "torch"
	blaze := runner{"blaze"}
	once := sync.Once{}

	laps := 3

	for i := 0; i < laps; i++ {
		once.Do(func() {
			blaze.grab(ctx, item) // This shows the power of closures: they can
			// take values from outside its scope, even without having to pass
			// arguments through function parameters.
		})
		once.Do(func() {
			blaze.flip(ctx, item) // Do will only ever run from `once` once,
			// even if we are calling a different function.
		})
		blaze.runLap(ctx, item)
	}

	blaze.drop(ctx, item)
}

func (r runner) flip(ctx context.Context, obj string) {
	fmt.Printf("context: %v, type: %[1]T\n", ctx)
	fmt.Println(r.name + " flips in the air the " + obj)
}
func (r runner) grab(ctx context.Context, obj string) {
	fmt.Printf("context: %v, type: %[1]T\n", ctx)
	fmt.Println(r.name + " grabs the " + obj)
}

func (r runner) runLap(ctx context.Context, obj string) {
	fmt.Printf("context: %v, type: %[1]T\n", ctx)
	fmt.Println(r.name + " runs a lap with a " + obj)
}

func (r runner) drop(ctx context.Context, obj string) {
	fmt.Printf("context: %v, type: %[1]T\n", ctx)
	fmt.Println(r.name + " drops the " + obj)
}
