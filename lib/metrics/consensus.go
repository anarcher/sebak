package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ConsensusHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: Consensus,
		Name:      "height",
		Help:      "Height of the consensus",
	})
	ConsensusRounds = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: Consensus,
		Name:      "rounds",
		Help:      "Rounds of this consensus",
	})
)

func init() {
	prometheus.MustRegister(ConsensusHeight)
	prometheus.MustRegister(ConsensusRounds)
}

func SetConsensusHeight(height uint64) {
	ConsensusHeight.Set(float64(height))
}

func SetConsensusRounds(round int) {
	ConsensusRounds.Set(float64(round))
}
