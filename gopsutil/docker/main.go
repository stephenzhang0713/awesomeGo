package main

import (
	"fmt"
	_ "github.com/shirou/gopsutil/v3"
	"github.com/shirou/gopsutil/v3/docker"
)

func main() {
	dockerStat, err := docker.GetDockerStat()
	if err != nil {
		return
	}
	fmt.Println(dockerStat)
}
