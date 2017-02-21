# sysresmon
Go library to access system stats.

Gets samples of CPU and memory usage, and provides helpers to start monitoring CPU and memory usage over time.


## Usage examples

Print average CPU usage each second (examples/3)

```go
sampler := sysresmon.NewSampler().StartSampling().Aggregate()
for {
	select {
	case aggregate := <-sampler.AggregateChan:
		fmt.Printf("CPU usage is %f%%\n", aggregate.CPUUsage)
	}
}
```

Start displaying CPU and memory usage values every second (examples/2)

```go
sampler := sysresmon.NewSampler().StartSampling()
var latestSample, previousSample sysresmon.CPUSample
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
```

Print average CPU usage over 1 second (examples/1)

```go
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
```

Further docuentation is available [on godoc](https://godoc.org/github.com/n-marshall/sysresmon)

## Requirements

- Linux (tested only on ubuntu 16.10, manually)

## Credits

- Half of this code was copied on [https://bitbucket.org/bertimus9/systemstat/](https://bitbucket.org/bertimus9/systemstat/)