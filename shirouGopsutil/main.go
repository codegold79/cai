package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	cpuInfo, _ := cpu.Info()
	var infoStr string

	for i := range cpuInfo {
		infoStr += cpuInfo[i].String()
	}

	fmt.Println(infoStr)
	fmt.Println(len(cpuInfo))
}
