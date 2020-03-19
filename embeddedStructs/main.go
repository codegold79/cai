package main

import (
	"fmt"
)

type animal struct {
	mammal
}

type mammal struct {
	hedgehog
}

type hedgehog struct {
	color  string
	age    int
	living bool
	name   string
}

func main() {
	rufus := animal{
		mammal{
			hedgehog{
				color:  "blue",
				age:    3,
				living: true,
				name:   "Rufus",
			},
		},
	}
	rufus.breath("nitrogen and oxygen")
	rufus.eat("kitten kibbles")
	rufus.sleep(10)
}

func (h hedgehog) eat(food string) {
	fmt.Println(h.name, "ate", food)
}
func (m mammal) eat(food string) {
	fmt.Println(m.name, "ate", food)
}

func (h hedgehog) sleep(hours int) {
	fmt.Println(h.name, "slept for", hours, "hours")
}

func (h hedgehog) breath(gas string) {
	fmt.Println(h.name, "breathes", gas)
}
