package simulation

import (
	"runtime/debug"
	"strconv"

	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/statistics/agg"
)

var (
	sha1ver   string
	buildTime string
	modified  bool
)

func init() {
	info, _ := debug.ReadBuildInfo()
	for _, bs := range info.Settings {
		if bs.Key == "vcs.revision" {
			sha1ver = bs.Value
		}
		if bs.Key == "vcs.time" {
			buildTime = bs.Value
		}
		if bs.Key == "vcs.modified" {
			bv, _ := strconv.ParseBool(bs.Value)
			modified = bv
		}
	}
}

func Version() string {
	return sha1ver
}

type Aggregators []agg.Aggregator

func InitializeAggregators(itrs int, cfg *model.SimConfig) (Aggregators, error) {
	aggregators := make(Aggregators, 0, len(agg.Aggregators()))
	for _, aggregator := range agg.Aggregators() {
		a, err := aggregator(cfg)
		if err != nil {
			return nil, err
		}
		aggregators = append(aggregators, a)
	}
	return aggregators, nil
}

func (aggs Aggregators) Add(result *model.IterationResult) {
	for _, a := range aggs {
		a.Add(result)
	}
}

func (aggs Aggregators) Flush() *model.Statistics {
	stats := new(model.Statistics)
	for _, a := range aggs {
		a.Flush(stats)
	}
	return stats
}

// TODO: Creates the base result structure with everything except statistics populated.
func CreateResult(cfg *model.SimConfig, seed int64) *model.SimResult {
	result := &model.SimResult{
		SchemaVersion: &model.Version{Major: "1", Minor: "0"},
		SimVersion:    &sha1ver,
		BuildDate:     buildTime,
		Modified:      &modified,
		Config:        cfg,
		DebugSeed:     strconv.FormatUint(uint64(seed), 10),
	}

	return result
}
