package metrics

import (
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	VersionView = &view.View{
		Name:        "version",
		Measure:     Version,
		Description: "The version of this node",
		Aggregation: view.LastValue(),
		TagKeys:     []tag.Key{KeyVersion, GitCommit, GoVersion},
	}

	ConsensusHeightView = &view.View{
		Measure:     ConsensusHeight,
		Aggregation: view.LastValue(),
	}
	ConsensusRoundsView = &view.View{
		Measure:     ConsensusRounds,
		Aggregation: view.LastValue(),
	}
	ConsensusNumTxsView = &view.View{
		Measure:     ConsensusNumTxs,
		Aggregation: view.LastValue(),
	}
)

var (
	Views = []*view.View{
		VersionView,
	}
)
