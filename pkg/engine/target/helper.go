package target

import (
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func AdjacentTo(targets []key.TargetID, target key.TargetID) []key.TargetID {
	for i, t := range targets {
		if t != target {
			continue
		}

		out := make([]key.TargetID, 0, 3)
		if i != 0 {
			out = append(out, targets[i-1])
		}
		out = append(out, t)
		if i != len(targets)-1 {
			out = append(out, targets[i+1])
		}
		return out
	}
	return nil
}

func Retarget(rand *rand.Rand, attr attribute.Manager, data info.Retarget) []key.TargetID {
	// if filter is empty, add in bypass func.
	if data.Filter == nil {
		data.Filter = func(key.TargetID) bool { return true }
	}
	// merged removal of filtered targets and targets in limbo into 1 loop.
	i := 0
	for _, target := range data.Targets {
		if (data.IncludeLimbo || attr.HPRatio(target) > 0) && data.Filter(target) {
			data.Targets[i] = target
			i++
		}
	}
	// truncate to safely remove unqualified targets.
	data.Targets = data.Targets[:i]

	// shuffle data.Targets IF data.DisableRandom is false
	if !data.DisableRandom {
		rand.Shuffle(len(data.Targets), func(i, j int) {
			data.Targets[i], data.Targets[j] = data.Targets[j], data.Targets[i]
		})
	}

	// truncate if data.Max specified and len(data.Targets) > data.Max
	if data.Max > 0 && len(data.Targets) > data.Max {
		data.Targets = data.Targets[:data.Max]
	}

	// return filtered list
	return data.Targets
}
