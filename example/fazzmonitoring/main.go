package main

import (
	"github.com/payfazz/fazzmonitoring/pkg/fazzmonitoring"
	"time"
)

func main() {
	counters := fazzmonitoring.CounterOpts{
		Name:      "Custom_Counters",
		Help:      "Count every 2 sec",
		Namespace: "loanfazz",
		Subsystem: "kasbon",
	}
	data := fazzmonitoring.NewCounter(counters)

	go func() {
		for {
			data.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

	fazzmonitoring.Enable("9001", nil)
}
