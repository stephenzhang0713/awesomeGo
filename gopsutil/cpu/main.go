package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	cpuinfo, _ := cpu.Info()
	cpuinfoJSON, _ := json.MarshalIndent(cpuinfo, "", "  ")
	fmt.Println(string(cpuinfoJSON))
}
