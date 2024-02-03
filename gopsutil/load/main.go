package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/load"
)

func main() {
	avg, err := load.Avg()
	if err != nil {
		return
	}

	misc, err := load.Misc()
	if err != nil {
		return
	}
	fmt.Println(avg)
	fmt.Println(misc)
}
