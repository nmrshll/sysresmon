package main

import (
	"fmt"
	"sysstats"
	"time"
)

func main() {
	sampler := sysstats.NewSampler(1 * time.Second).StartSampling().Aggregate()
	for {
		select {
		case aggregate := <-sampler.AggregateChan:
			fmt.Printf("CPU usage is %f%%\n", aggregate.CPUUsage)
		}
	}
}
