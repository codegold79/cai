package main

import (
	"fmt"
)

type object struct {
	name string
	numb int
}

func main() {
	var radio = object{
		name: "fm",
		numb: 1,
	}

	alter(&radio)
	fmt.Println(radio)
}

func alter(r *object) {
	r.numb = 2
}
