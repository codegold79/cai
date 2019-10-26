package main

import (
	"fmt"
)

// I feel like mammal is a useless interface because I can just as easily not use
// some of the methods in the primate interface, which uses all of mammal's interface.
type mammal interface {
	eat()
	poop()
	sleep()
}

type primate interface {
	mammal
	useThumbs()
}
type myStr string

type identifier struct {
	name    myStr
	species myStr
}

func main() {

	ix := identifier{
		name:    myStr("spot"),
		species: myStr("dog"),
	}

	iy := identifier{
		name:    myStr("dave"),
		species: myStr("human"),
	}

	iz := identifier{

		name:    myStr("aria"),
		species: myStr("hawk"),
	}

	ix.eat()
	ix.sleep()
	ix.poop()

	iy.eat()
	iy.sleep()
	iy.poop()
	iy.useThumbs()

	iz.eat()
	iz.sleep()
}

func (id identifier) eat() {
	fmt.Printf("%v is eating because zir is a %v\n", id.name, id.species)
}

func (id identifier) poop() {
	fmt.Printf("%v is pooping because zir is a %v\n", id.name, id.species)
}

func (id identifier) sleep() {
	fmt.Printf("%v is sleeping because zir is a %v\n", id.name, id.species)
}

func (id identifier) useThumbs() {
	fmt.Printf("%v is using zir thumbs because zir is a %v\n", id.name, id.species)
}
