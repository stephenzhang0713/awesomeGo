package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
)

func main() {
	stats, err := net.Connections("inet")
	if err != nil {
		return
	}
	for _, connectionStat := range stats {
		connectionStat.Status = ""
	}
	fmt.Println(stats)

	ioCounters, err := net.IOCounters(true)
	if err != nil {
		return
	}
	fmt.Println(ioCounters)

}
