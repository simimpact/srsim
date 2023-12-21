package overview

import (
	calc "github.com/aclements/go-moremath/stats"

	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/statistics/agg"
)

func init() {
	agg.Register(NewAgg)
}

type buffer struct {
	dpc              *calc.Sample
	totalDamageDealt calc.StreamStats
	totalDamageTaken calc.StreamStats
	totalAV          calc.StreamStats
}

func newSample(itr uint32) *calc.Sample {
	return &calc.Sample{
		Xs:      make([]float64, 0, itr),
		Sorted:  false,
		Weights: nil,
	}
}

func NewAgg(cfg *model.SimConfig) (agg.Aggregator, error) {
	out := buffer{
		dpc:              newSample(cfg.Settings.Iterations),
		totalDamageTaken: calc.StreamStats{},
		totalDamageDealt: calc.StreamStats{},
		totalAV:          calc.StreamStats{},
	}
	return &out, nil
}

// TODO: push looping/summation to StatsCollector for peformance boost
func (b *buffer) Add(result *model.IterationResult) {
	b.totalDamageDealt.Add(result.TotalDamageDealt)
	b.totalDamageTaken.Add(result.TotalDamageTaken)
	b.totalAV.Add(result.TotalAv)
	b.dpc.Xs = append(b.dpc.Xs, result.TotalDamageDealt*100/result.TotalAv)
}

func (b *buffer) Flush(result *model.Statistics) {
	result.TotalDamageDealt = agg.ToDescriptiveStats(&b.totalDamageDealt)
	result.TotalDamageTaken = agg.ToDescriptiveStats(&b.totalDamageTaken)
	result.TotalAv = agg.ToDescriptiveStats(&b.totalAV)
	result.TotalDamageDealtPerCycle = agg.ToOverviewStats(b.dpc)
}
