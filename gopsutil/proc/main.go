package main

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	ctx := context.Background()
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return
	}
	for _, p := range processes {
		name, _ := p.NameWithContext(ctx)
		status, _ := p.StatusWithContext(ctx)
		cpuPercent, _ := p.CPUPercentWithContext(ctx)
		memoryInfo, _ := p.MemoryInfoWithContext(ctx)
		connectionsWithContext, _ := p.ConnectionsWithContext(ctx)
		numCtxSwitches, _ := p.NumCtxSwitchesWithContext(ctx)
		fmt.Println(p.Pid, name, status, cpuPercent, memoryInfo, numCtxSwitches, connectionsWithContext)
	}
}
