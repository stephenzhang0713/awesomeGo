package main

import "time"

type Counter3 struct {
	Website      string
	Start        time.Time
	PageCounters map[string]int
}

func main() {
	var c Counter3
	c.Website = "baidu.com"

	c.PageCounters["/"]++
}
