package fazzmonitoring

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func Enable(port string, fn func() error) error {
	if fn != nil {
		err := fn()
		if err != nil {
			return err
		}
	}
	http.Handle("/metrics", promhttp.Handler())
	// start an http server using the mux server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("Error Listen: ", err)
		return errors.New("error when creating new http server")
	}
	return nil
}

func NewCounter(name string, desc string) prometheus.Counter {
	counter := promauto.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: desc,
	})

	return counter
}
