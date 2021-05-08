package main

import (
	"fmt"

	"github.com/codegold79/cai/use_library_global/projglobal"
	"github.com/codegold79/cai/use_library_global/seclib"
)

func main() {
	projglobal.Name = "Alice"
	fmt.Println("set project global variable in main.go to", projglobal.Name)

	seclib.CallProjGlobal()
}
