package main

import (
	"github.com/kzolatech/go-monitoring/pkg/fazzmonitoring"
	"time"
)

func main() {
	temp := fazzmonitoring.NewCounter("Custom_Counter", "Count every 2 sec")
	go func() {
		for {
			temp.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

	fazzmonitoring.Enable("9001", nil)
}
