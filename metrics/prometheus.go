package metrics

import "github.com/prometheus/client_golang/prometheus"

// NewCounter creates a new Counter based on opts and registers it in reg.
func NewCounter(reg prometheus.Registerer, opts prometheus.CounterOpts) prometheus.Counter {
	c := prometheus.NewCounter(opts)
	if exists, existingCol := register(reg, c); exists {
		return existingCol.(prometheus.Counter)
	}
	return c
}

// NewCounterVec creates a new CounterVec based on opts and registers it in reg.
func NewCounterVec(reg prometheus.Registerer, opts prometheus.CounterOpts, labelNames []string) *prometheus.CounterVec {
	c := prometheus.NewCounterVec(opts, labelNames)
	if exists, existingCol := register(reg, c); exists {
		return existingCol.(*prometheus.CounterVec)
	}
	return c
}

// NewGauge creates a new Gauge based on opts and registers it in reg.
func NewGauge(reg prometheus.Registerer, opts prometheus.GaugeOpts) prometheus.Gauge {
	g := prometheus.NewGauge(opts)
	if exists, existingCol := register(reg, g); exists {
		return existingCol.(prometheus.Gauge)
	}
	return g
}

// NewGauge creates a new Gauge based on opts and registers it in reg.
func NewGaugeVec(reg prometheus.Registerer, opts prometheus.GaugeOpts, labelNames []string) *prometheus.GaugeVec {
	g := prometheus.NewGaugeVec(opts, labelNames)
	if exists, existingCol := register(reg, g); exists {
		return existingCol.(*prometheus.GaugeVec)
	}
	return g
}

// NewSummary creates a new Summary based on opts and registers it in reg.
func NewSummary(reg prometheus.Registerer, opts prometheus.SummaryOpts) prometheus.Summary {
	s := prometheus.NewSummary(opts)
	if exists, existingCol := register(reg, s); exists {
		return existingCol.(prometheus.Summary)
	}
	return s
}

// register registers col to be included in reg. If the collector is already
// registered before, the existing one is returned.
func register(reg prometheus.Registerer, col prometheus.Collector) (exists bool, existingCol prometheus.Collector) {
	err := reg.Register(col)
	if err == nil {
		return
	}
	if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
		return true, are.ExistingCollector
	}
	panic(err)
}
