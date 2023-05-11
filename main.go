package main

import (
	"fmt"
	"net/http"
	"playground/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

var (
	ethdevXstatsLabels     = []string{"port_id", "port_name", "prefix", "xstat_name"}
	ethdevXstatsMetricName = "ethdev_xstats"
)

func main() {
	reg := prometheus.DefaultRegisterer
	namespace := "example"

	vec := metrics.NewCounterVec(reg, prometheus.CounterOpts{
		Name:      ethdevXstatsMetricName,
		Help:      "Help text for " + ethdevXstatsMetricName,
		Namespace: namespace,
	}, ethdevXstatsLabels)

	portId := "1"
	portName := "eth0"
	prefix := "dpdk"
	statName := "example"
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		vec.WithLabelValues(portId, portName, prefix, statName).Add(float64(i))
	}
}
