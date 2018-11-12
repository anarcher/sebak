package metrics

import (
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	VersionView = &view.View{
		Measure:     Version,
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
		ConsensusHeightView,
		ConsensusRoundsView,
		ConsensusNumTxsView,
	}
)
