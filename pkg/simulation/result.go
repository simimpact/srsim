package simulation

import (
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/statistics/agg"
)

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

func (aggs Aggregators) Flush() *model.SimulationStatistics {
	stats := new(model.SimulationStatistics)
	for _, a := range aggs {
		a.Flush(stats)
	}
	return stats
}

// TODO: Creates the base result structure with everything except statistics populated.
func CreateResult(cfg *model.SimConfig) *model.SimulationResult {
	return new(model.SimulationResult)
}
