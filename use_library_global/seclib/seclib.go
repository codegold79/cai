package seclib

import (
	"fmt"

	"github.com/codegold79/cai/use_library_global/projglobal"
)

func CallProjGlobal() {
	if projglobal.Name != "" {
		fmt.Println("call project global variable from seclib and read", projglobal.Name)
	}
}
