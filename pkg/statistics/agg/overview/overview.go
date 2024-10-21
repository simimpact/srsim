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
	iters              uint32
	dpc                *calc.Sample
	totalDamageDealt   calc.StreamStats
	totalDamageTaken   calc.StreamStats
	totalAV            calc.StreamStats
	cumDmgDealtByCycle []*calc.Sample
	cumDmgTakenByCycle []*calc.Sample
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
		iters: cfg.Settings.Iterations,
		dpc:   newSample(cfg.Settings.Iterations),
		totalDamageTaken: calc.StreamStats{
			Count: 0,
			Total: 0,
			Min:   0,
			Max:   0,
		},
		totalDamageDealt: calc.StreamStats{
			Count: 0,
			Total: 0,
			Min:   0,
			Max:   0,
		},
		totalAV: calc.StreamStats{
			Count: 0,
			Total: 0,
			Min:   0,
			Max:   0,
		},
		cumDmgDealtByCycle: make([]*calc.Sample, cfg.Settings.CycleLimit),
		cumDmgTakenByCycle: make([]*calc.Sample, cfg.Settings.CycleLimit),
	}
	for i := range out.cumDmgDealtByCycle {
		out.cumDmgDealtByCycle[i] = newSample(cfg.Settings.Iterations)
	}
	for i := range out.cumDmgTakenByCycle {
		out.cumDmgTakenByCycle[i] = newSample(cfg.Settings.Iterations)
	}
	return &out, nil
}

// TODO: push looping/summation to StatsCollector for peformance boost
func (b *buffer) Add(result *model.IterationResult) {
	b.totalDamageDealt.Add(result.TotalDamageDealt)
	b.totalDamageTaken.Add(result.TotalDamageTaken)
	b.totalAV.Add(result.TotalAv)
	b.dpc.Xs = append(b.dpc.Xs, result.TotalDamageDealt*100/result.TotalAv)

	var last float64
	for i, v := range result.CumulativeDamageDealtByCycle {
		for i >= len(b.cumDmgDealtByCycle) {
			b.cumDmgDealtByCycle = append(b.cumDmgDealtByCycle, newSample(b.iters))
		}
		b.cumDmgDealtByCycle[i].Xs = append(b.cumDmgDealtByCycle[i].Xs, v-last)
		last = v
	}
	last = 0
	for i, v := range result.CumulativeDamageTakenByCycle {
		for i >= len(b.cumDmgTakenByCycle) {
			b.cumDmgTakenByCycle = append(b.cumDmgTakenByCycle, newSample(b.iters))
		}
		b.cumDmgTakenByCycle[i].Xs = append(b.cumDmgTakenByCycle[i].Xs, v-last)
		last = v
	}
}

func (b *buffer) Flush(result *model.Statistics) {
	result.TotalDamageDealt = agg.ToDescriptiveStats(&b.totalDamageDealt)
	result.TotalDamageTaken = agg.ToDescriptiveStats(&b.totalDamageTaken)
	result.TotalAv = agg.ToDescriptiveStats(&b.totalAV)
	result.TotalDamageDealtPerCycle = agg.ToOverviewStats(b.dpc)
	for i := range b.cumDmgDealtByCycle {
		result.DamageDealtByCycle = append(result.DamageDealtByCycle, agg.ToOverviewStats(b.cumDmgDealtByCycle[i]))
	}
	for i := range b.cumDmgTakenByCycle {
		result.DamageTakenByCycle = append(result.DamageTakenByCycle, agg.ToOverviewStats(b.cumDmgTakenByCycle[i]))
	}
}
