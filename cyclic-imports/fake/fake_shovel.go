/*
Package fake tries to be like the cool package, but doesn't do anything.
*/
package fake

import (
	"github.com/codegold79/cai/cyclic-imports/tool"
)

type Shovel struct{}

func (s Shovel) Dig(color tool.Color) string {
	return ""
}

func (s Shovel) Cover(tool.Color) string {
	return ""
}
