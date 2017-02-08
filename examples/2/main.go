package main

import (
	"fmt"
	"sysstats"
)

func main() {
	sampler := sysstats.NewSampler().StartSampling()
	var latestSample, previousSample sysstats.CPUSample
	for {
		select {
		case sampleSet := <-sampler.SampleSetChan:
			previousSample = latestSample
			latestSample = sampleSet.CPUSample

			idleTicks := float64(latestSample.Idle - previousSample.Idle)
			totalTicks := float64(latestSample.Total - previousSample.Total)
			cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

			fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
		}
	}
}
