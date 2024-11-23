package agg

import (
	"math"

	"github.com/aclements/go-moremath/stats"
	"github.com/simimpact/srsim/pkg/model"
)

func ToDescriptiveStats(ss *stats.StreamStats) *model.DescriptiveStats {
	sd := ss.StdDev()
	if math.IsNaN(sd) {
		sd = 0
	}

	mean := ss.Mean()
	return &model.DescriptiveStats{
		Min:  &ss.Min,
		Max:  &ss.Max,
		Mean: &mean,
		SD:   &sd,
	}
}

func ToOverviewStats(input *stats.Sample) *model.OverviewStats {
	input.Sorted = false
	input.Sort()

	dataMin, datamax := input.Bounds()
	std := input.StdDev()
	if math.IsNaN(std) {
		std = 0
	}

	out := model.OverviewStats{
		SD:   &std,
		Min:  &dataMin,
		Max:  &datamax,
		Mean: Ptr(input.Mean()),
		Q1:   Ptr(input.Quantile(0.25)),
		Q2:   Ptr(input.Quantile(0.5)),
		Q3:   Ptr(input.Quantile(0.75)),
		Hist: nil,
	}

	// Scott's normal reference rule
	h := (3.49 * std) / (math.Pow(float64(len(input.Xs)), 1.0/3.0))
	if h == 0.0 || datamax == dataMin {
		hist := make([]uint32, 1)
		hist[0] = uint32(len(input.Xs))
		out.Hist = hist
	} else {
		nbins := int(math.Ceil((datamax - dataMin) / h))
		hist := NewLinearHist(dataMin, datamax, nbins)
		for _, x := range input.Xs {
			hist.Add(x)
		}
		low, bins, high := hist.Counts()
		bins[0] += low
		bins[len(bins)-1] += high
		out.Hist = bins
	}

	return &out
}

// taken from go-moremath. Need to reimplement for proto type compatibility
type LinearHist struct {
	min, max  float64
	delta     float64 // 1/bin width (to avoid division in hot path)
	low, high uint32
	bins      []uint32
}

// NewLinearHist returns an empty histogram with nbins uniformly-sized
// bins spanning [min, max].
func NewLinearHist(histMin, histMax float64, nbins int) *LinearHist {
	delta := float64(nbins) / (histMax - histMin)
	return &LinearHist{histMin, histMax, delta, 0, 0, make([]uint32, nbins)}
}

func (h *LinearHist) bin(x float64) int {
	return int(h.delta * (x - h.min))
}

func (h *LinearHist) Add(x float64) {
	bin := h.bin(x)

	switch {
	case bin < 0:
		h.low++
	case bin >= len(h.bins):
		h.high++
	default:
		h.bins[bin]++
	}
}

func (h *LinearHist) Counts() (uint32, []uint32, uint32) {
	return h.low, h.bins, h.high
}

func (h *LinearHist) BinToValue(bin float64) float64 {
	return h.min + bin/h.delta
}

func Ptr[T any](v T) *T {
	return &v
}
