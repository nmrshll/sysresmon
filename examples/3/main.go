package main

import (
	"fmt"
	"sysresmon"
)

func main() {
	sampler := sysresmon.NewSampler().StartSampling().Aggregate()
	for {
		select {
		case aggregate := <-sampler.AggregateChan:
			fmt.Printf("CPU usage is %f%%\n", aggregate.CPUUsage)
		}
	}
}
