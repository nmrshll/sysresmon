package sysstats

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type CPUSample struct {
	Idle  uint64
	Total uint64
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
