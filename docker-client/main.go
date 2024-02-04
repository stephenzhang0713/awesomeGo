package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func calculateCPUPercentUnix(previousCPU, previousSystem uint64, v *types.StatsJSON) float64 {
	cpuDelta := float64(v.CPUStats.CPUUsage.TotalUsage) - float64(previousCPU)
	systemDelta := float64(v.CPUStats.SystemUsage) - float64(previousSystem)
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		return (cpuDelta / systemDelta) * float64(len(v.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return 0.0
}

func calculateMemoryPercent(v *types.StatsJSON) float64 {
	// MemoryStats.Limit gives the total memory available (memory limit)
	// MemoryStats.Usage gives the current memory usage
	if v.MemoryStats.Limit != 0 {
		return float64(v.MemoryStats.Usage) / float64(v.MemoryStats.Limit) * 100.0
	}
	return 0.0
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Error creating Docker client: %s\n", err)
		return
	}

	ctx := context.Background()
	containerID := "058495f77aeb313dbf1b1c7204d35f1c9161dc00ea60c672c3680e668d50a595" // Replace with your container ID

	for {
		stats, err := cli.ContainerStats(ctx, containerID, false)
		if err != nil {
			fmt.Printf("Error getting container stats: %s\n", err)
			time.Sleep(10 * time.Second) // Wait before retrying
			continue
		}

		var v *types.StatsJSON
		decoder := json.NewDecoder(stats.Body)
		err = decoder.Decode(&v)
		stats.Body.Close() // Close the Body when the decoding is complete
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error decoding container stats: %s\n", err)
			}
			// If EOF or other error, wait before retrying
			time.Sleep(10 * time.Second)
			continue
		}

		memoryPercent := calculateMemoryPercent(v)
		fmt.Printf("Memory Usage: %.2f%%\n", memoryPercent)

		time.Sleep(10 * time.Second) // Update interval
	}
}
