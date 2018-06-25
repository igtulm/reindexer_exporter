package main

import "github.com/prometheus/client_golang/prometheus"

var (
	//Version of Reindexer Exporter is set during build.
	Version string
	//Revision of Reindexer Exporter is set during build.
	Revision string
	//Branch of Reindexer Exporter is set during build.
	Branch string
	//BuildDate of Reindexer Exporter is set during build.
	BuildDate string
)

//BuildInfo is a metric with a constant '1' value labeled by version, revision, branch and build date on which the reindexer_exporter was built
var BuildInfo *prometheus.GaugeVec

func init() {
	BuildInfo = newBuildInfo()
}

func newBuildInfo() *prometheus.GaugeVec {
	metric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "reindexer_exporter_build_info",
			Help: "A metric with a constant '1' value labeled by version, revision, branch and build date on which the reindexer_exporter was built.",
		},
		[]string{"version", "revision", "branch", "builddate"},
	)
	metric.WithLabelValues(Version, Revision, Branch, BuildDate).Set(1)
	return metric
}
