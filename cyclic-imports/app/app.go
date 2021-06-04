package main

import (
	"fmt"

	"github.com/codegold79/cai/cyclic-imports/tool"
)

func main() {
	equipment := tool.Shovel{Name: "shovel"}
	var color tool.Color = "indigo"
	prospect(equipment, color)
}

type shoveler interface {
	Dig(tool.Color) string
	Cover(tool.Color) string
}

func prospect(equipment shoveler, color tool.Color) bool {
	fmt.Println(equipment.Dig(color))
	fmt.Println(equipment.Cover(color))
	return true
}
