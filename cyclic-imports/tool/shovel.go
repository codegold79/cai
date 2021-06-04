package tool

import (
	"fmt"
)

type Color string

type Shovel struct {
	Name string
}

func (s Shovel) Dig(color Color) string {
	return fmt.Sprintln("diggin with a", color, s.Name)
}

func (s Shovel) Cover(color Color) string {
	return fmt.Sprintln("covering with a", color, s.Name)
}
