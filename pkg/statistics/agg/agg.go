package agg

import (
	"sync"

	"github.com/simimpact/srsim/pkg/model"
)

type Aggregator interface {
	Add(result *model.IterationResult)
	// TODO: Merge(other Aggregator) Aggregator for multi-threaded aggregations (optional optimization)
	Flush(result *model.Statistics)
}

type NewAggFunc func(cfg *model.SimConfig) (Aggregator, error)

var (
	mu          sync.Mutex
	aggregators []NewAggFunc
)

func Register(f NewAggFunc) {
	mu.Lock()
	defer mu.Unlock()
	aggregators = append(aggregators, f)
}

func Aggregators() []NewAggFunc {
	return aggregators
}
