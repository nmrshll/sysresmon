package sysstats

import (
	"log"
	"time"
)

type Sampler struct {
	Rate          time.Duration
	SampleSetChan chan SampleSet
	AggregateChan chan Aggregate
}

type SampleSet struct {
	CPUSample CPUSample
	MemSample MemSample
}

type Aggregate struct {
	CPUUsage float64
	MemUsage float64
}

func NewSampler(rate time.Duration) *Sampler {
	return &Sampler{Rate: rate}
}

func (s *Sampler) StartSampling() *Sampler {
	s.SampleSetChan = make(chan SampleSet)
	go func() {
		for {
			var ss SampleSet
			ss.CPUSample = GetCPUSample()
			ss.MemSample = GetMemSample()
			s.SampleSetChan <- ss
			time.Sleep(s.Rate)
		}
	}()
	return s
}

func (s *Sampler) Aggregate() *Sampler {
	if s.SampleSetChan == nil {
		log.Fatalf("Before using Aggregate() you need to StartSampling()")
	}
	s.AggregateChan = make(chan Aggregate)

	var latestSample, previousSample CPUSample
	go func() {
		for {
			select {
			case sampleSet := <-s.SampleSetChan:
				previousSample = latestSample
				latestSample = sampleSet.CPUSample

				idleTicks := float64(latestSample.Idle - previousSample.Idle)
				totalTicks := float64(latestSample.Total - previousSample.Total)
				cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

				s.AggregateChan <- Aggregate{CPUUsage: cpuUsage}
			}
		}
	}()
	return s
}
