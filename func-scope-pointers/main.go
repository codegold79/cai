/*
Functions, scope, and pointers.

Because of scope, functions cannot change the original values of variable
instances before they were passed to them as arguments. This is true for first
class functions that have been passed around as arguments. In order to be able
to change the original value, use pointers.
*/

package main

import (
	"fmt"
)

func main() {
	// Function that works a variable's instance (copy of value).
	fn := func(name string) {
		fmt.Println("original name in fn is", name)
		name = "bob"
		fmt.Println("copy of name var (staying inside this scope) is now", name)
	}

	// Function that works with a pointer address value.
	fnp := func(pname *string) {
		fmt.Println("original name in fnp is", *pname)
		*pname = "cat"
		fmt.Println("pointer points to name var", *pname)
	}

	runAFunc(fn, fnp)
}

func runAFunc(fn func(string), fnp func(*string)) {
	name := "alice"

	// This function cannot change the original name due to the variable's scope
	// and the fact that its value is copied while inside the function.
	fn(name)
	fmt.Println("Name is", name)

	fmt.Println()

	// Use a pointer to get the original name to change.
	fnp(&name)
	fmt.Println("Name is", name)
}
