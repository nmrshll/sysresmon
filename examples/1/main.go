package main

import (
	"fmt"
	"sysresmon"
	"time"
)

func main() {
	var latestSample, previousSample sysresmon.CPUSample
	for {
		previousSample = latestSample
		time.Sleep(1 * time.Second)
		latestSample = sysresmon.GetCPUSample()

		idleTicks := float64(latestSample.Idle - previousSample.Idle)
		totalTicks := float64(latestSample.Total - previousSample.Total)
		cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

		fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
	}
}
