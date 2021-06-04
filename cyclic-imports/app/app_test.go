package main

import (
	"testing"

	"github.com/codegold79/cai/cyclic-imports/fake"
	"github.com/codegold79/cai/cyclic-imports/tool"
)

func TestProspect(t *testing.T) {
	tests := []struct {
		want bool
	}{
		{true},
	}

	var fakeShovel fake.Shovel
	var color tool.Color

	for _, tt := range tests {
		got := prospect(fakeShovel, color)

		if got != tt.want {
			t.Errorf("got: %t, but want %t", got, tt.want)
		}
	}
}
