package sysstats

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type CPUSample struct {
	Idle  uint64
	Total uint64
	Time  time.Time
}

type MemSample struct {
	Buffers   uint64
	Cached    uint64
	MemTotal  uint64
	MemUsed   uint64
	MemFree   uint64
	SwapTotal uint64
	SwapUsed  uint64
	SwapFree  uint64
	Time      time.Time
}

func GetCPUSample() (sample CPUSample) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				sample.Total += val // tally up all the numbers to get total ticks
				if i == 4 {         // idle is the 5th field in the cpu line
					sample.Idle = val
				}
			}
			return
		}
	}
	return
}

func GetMemSample() (samp MemSample) {
	want := map[string]bool{
		"Buffers:":   true,
		"Cached:":    true,
		"MemTotal:":  true,
		"MemFree:":   true,
		"MemUsed:":   true,
		"SwapTotal:": true,
		"SwapFree:":  true,
		"SwapUsed:":  true}

	// read in whole meminfo file with cpu usage information ;"/proc/meminfo"
	contents, err := ioutil.ReadFile("/proc/meminfo")
	samp.Time = time.Now()
	if err != nil {
		return
	}

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Fields(string(line))
		fieldName := fields[0]

		_, ok := want[fieldName]
		if ok && len(fields) == 3 {
			val, numerr := strconv.ParseUint(fields[1], 10, 64)
			if numerr != nil {
				return
			}
			switch fieldName {
			case "Buffers:":
				samp.Buffers = val
			case "Cached:":
				samp.Cached = val
			case "MemTotal:":
				samp.MemTotal = val
			case "MemFree:":
				samp.MemFree = val
			case "SwapTotal:":
				samp.SwapTotal = val
			case "SwapFree:":
				samp.SwapFree = val
			}
		}
	}
	samp.MemUsed = samp.MemTotal - samp.MemFree
	samp.SwapUsed = samp.SwapTotal - samp.SwapFree
	return
}
