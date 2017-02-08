package main

import (
	"fmt"
	"sysstats"
)

func main() {
	sampler := sysstats.NewSampler().StartSampling().Aggregate()
	for {
		select {
		case aggregate := <-sampler.AggregateChan:
			fmt.Printf("CPU usage is %f%%\n", aggregate.CPUUsage)
		}
	}
}
