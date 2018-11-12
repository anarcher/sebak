package metrics

import (
	"context"
	"runtime"

	"boscoin.io/sebak/lib/version"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
)

func RecordVersion(pctx context.Context) error {
	if pctx == nil {
		pctx = context.Background()
	}
	ctx, err := tag.New(pctx,
		tag.Insert(KeyVersion, version.Version),
		tag.Insert(GitCommit, version.GitCommit),
		tag.Insert(BuildDate, version.BuildDate),
		tag.Insert(GoVersion, runtime.Version()),
	)
	if err != nil {
		log.Error("record version error", "err", err)
		return err
	}
	stats.Record(ctx, Version.M(1))
	return nil
}
