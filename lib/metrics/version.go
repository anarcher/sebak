package metrics

import (
	"runtime"

	"boscoin.io/sebak/lib/version"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Version = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "version",
		Namespace: Namespace,
		Help:      "version information of the node",
	}, []string{
		"version",
		"git_commit",
		"go_version",
	})
)

func init() {
	prometheus.MustRegister(Version)
}

func SetVersion() {
	goVersion := runtime.Version()
	Version.WithLabelValues(version.Version, version.GitCommit, goVersion).Set(1)
}
