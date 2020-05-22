// How does one tell the difference between an empty value, e.g. 0, "", or false
// and the absence of information when parsing JSON strings? Use pointers.
// Thanks to the following blog for this code practice:
// https://vsupalov.com/go-json-omitempty/
//
// The Go playground version is here: https://play.golang.org/p/4A5bY_Uc2tn

package main

import (
	"encoding/json"
	"fmt"
)

type Numbers struct {
	One *int
	Two *int
	Zer *int
}

func main() {
	jb := []byte(`
	{
		"one": 1,
		"zer": null,
		"two": 0
	}
	`)

	var num Numbers

	json.Unmarshal(jb, &num)

	fmt.Printf("%+v", num)

	if num.One != nil {
		fmt.Println("One:", *num.One)
	}

	if num.Two != nil {
		fmt.Println("Two:", *num.Two)
	}
	if num.Zer != nil {
		fmt.Println("Zer:", *num.Zer)
	}

}

func (ns Numbers) String() string {
	var s string

	s += fmt.Sprintf("One: %v\nTwo: %v\nZer: %v\n\n", ns.One, ns.Two, ns.Zer)

	return s
}
