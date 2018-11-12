package metrics

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	Version = stats.Int64("version", "Version information ", stats.UnitDimensionless)

	ConsensusHeight                    = stats.Int64("consensus/height", "Height of the consensus", stats.UnitDimensionless)
	ConsensusRounds                    = stats.Int64("consensus/rounds", "Round of the consensus", stats.UnitDimensionless)
	ConsensusNumTxs                    = stats.Int64("consensus/numtxs", "Total number txs of consensus", stats.UnitDimensionless)
	ConsensusValidators                = stats.Int64("consensus/validators", "Number of validators", stats.UnitDimensionless)
	ConsensusMissingValidators         = stats.Int64("consensus/missing_validators", "Number of missing validators", stats.UnitDimensionless)
	ConsensusBlockIntervalMilliseconds = stats.Float64("consensus/block_interval_milliseconds", "Interval betwwen this and last block", stats.UnitMilliseconds)

	TxPoolSize = stats.Int64("tx/pool_size", "Size of tx pool", stats.UnitDimensionless)

	SyncHeight               = stats.Int64("sync/height", "Height of sync", stats.UnitDimensionless)
	SyncErrors               = stats.Int64("sync/errors", "Number of errors in sync", stats.UnitDimensionless)
	SyncDurationMilliseconds = stats.Float64("sync/duration_milliseconds", "", stats.UnitMilliseconds)

	APIRequestTotal                = stats.Int64("api/request_total", "Total api request", stats.UnitDimensionless)
	APIRequestErrorsTotal          = stats.Int64("api/request_errors_total", "Total errors api request", stats.UnitDimensionless)
	APIRequestDurationMilliseconds = stats.Int64("api/request_duration_milliseconds", "Duration api request ", stats.UnitMilliseconds)
)

var (
	KeyVersion, _     = tag.NewKey("version")   // with Version
	GitCommit, _      = tag.NewKey("gitcommit") // with Version
	GoVersion, _      = tag.NewKey("goversion") // with Version
	HTTPMethod, _     = tag.NewKey("http_method")
	HTTPStatusCode, _ = tag.NewKey("http_status")
)

var (
	DefaultSizeDistribution = view.Distribution(0, 1024, 2048, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824, 4294967296)
	// [>=0ms, >=25ms, >=50ms, >=75ms, >=100ms, >=200ms, >=400ms, >=600ms, >=800ms, >=1s, >=2s, >=4s, >=6s]
	DefaultLatencyDistribution = view.Distribution(0, 25, 50, 75, 100, 200, 400, 600, 800, 1000, 2000, 4000, 6000)
)
