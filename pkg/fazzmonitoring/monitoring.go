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

type CounterOpts prometheus.CounterOpts
type GaugeOpts prometheus.GaugeOpts
type HistogramOpts prometheus.HistogramOpts
type SummaryOpts prometheus.SummaryOpts
type Counter prometheus.Counter
type CounterVec *prometheus.CounterVec
type CounterFunc prometheus.CounterFunc
type Gauge prometheus.Gauge
type GaugeVec *prometheus.GaugeVec
type GaugeFunc prometheus.GaugeFunc
type Histogram prometheus.Histogram
type HistogramVec *prometheus.HistogramVec
type Summary prometheus.Summary
type SummaryVec *prometheus.SummaryVec

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

func NewCounter(opts CounterOpts) Counter {
	counter := promauto.NewCounter(prometheus.CounterOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Subsystem:   opts.Subsystem,
	})

	return counter
}

func NewCounterVec(opts CounterOpts, labelNames []string) CounterVec {
	counter := promauto.NewCounterVec(prometheus.CounterOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Subsystem:   opts.Subsystem,
	}, labelNames)

	return counter
}

func NewCounterFunc(opts CounterOpts, fn func() float64) CounterFunc {
	counter := promauto.NewCounterFunc(prometheus.CounterOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Subsystem:   opts.Subsystem,
	}, fn)

	return counter
}

func NewGauge(opts GaugeOpts) Gauge {
	gauge := promauto.NewGauge(prometheus.GaugeOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Subsystem:   opts.Subsystem,
	})

	return gauge
}

func NewGaugeVec(opts GaugeOpts, labelNames []string) GaugeVec {
	gauge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Subsystem:   opts.Subsystem,
	}, labelNames)

	return gauge
}

func NewGaugeFunc(opts GaugeOpts, fn func() float64) GaugeFunc {
	gauge := promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Subsystem:   opts.Subsystem,
	}, fn)

	return gauge
}

func NewHistogram(opts HistogramOpts) Histogram {
	histogram := promauto.NewHistogram(prometheus.HistogramOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Subsystem:   opts.Subsystem,
		Buckets:     opts.Buckets,
	})

	return histogram
}

func NewHistogramVec(opts HistogramOpts, labelNames []string) HistogramVec {
	histogram := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Subsystem:   opts.Subsystem,
		Buckets:     opts.Buckets,
	}, labelNames)

	return histogram
}

func NewSummary(opts SummaryOpts) Summary {
	summary := promauto.NewSummary(prometheus.SummaryOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Subsystem:   opts.Subsystem,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Objectives:  opts.Objectives,
		MaxAge:      opts.MaxAge,
		AgeBuckets:  opts.AgeBuckets,
		BufCap:      opts.BufCap,
	})

	return summary
}

func NewSummaryVec(opts SummaryOpts, labelNames []string) SummaryVec {
	summary := promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name:        opts.Name,
		Help:        opts.Help,
		Subsystem:   opts.Subsystem,
		Namespace:   opts.Namespace,
		ConstLabels: opts.ConstLabels,
		Objectives:  opts.Objectives,
		MaxAge:      opts.MaxAge,
		AgeBuckets:  opts.AgeBuckets,
		BufCap:      opts.BufCap,
	}, labelNames)

	return summary
}
