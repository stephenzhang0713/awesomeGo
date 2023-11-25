package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"log/slog"
)

// ProcessState 是监控进程状态的函数。
type ProcessState struct {
	CpuPercent float64  `json:"CpuPercent"` // cpu占用
	MemPercent float32  `json:"MemPercent"` // 内存占用
	CpuTime    float64  `json:"CpuTime"`    // cpu持续时间
	Status     string   `json:"Status"`     // 进程状态
	Ports      []uint32 `json:"Ports"`      // 进程端口
	Pid        int32    `json:"Pid"`        // 进程pid
	Name       string   `json:"Name"`       // 进程名称
	PortStatus string   `json:"PortStatus"`
}

func main() {
	processes, _ := process.Processes()
	for _, proc := range processes {
		pid := proc.Pid
		procConnections, err := proc.Connections()
		if err != nil {
			slog.Error("failed to get connections for pid %d: %v", pid, err)
			continue
		}

		for _, conn := range procConnections {
			fmt.Println("port status: ", conn.Status)
		}
	}
}
