package evaltarget

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
)

type evaluator func(engine engine.Engine, candidates []key.TargetID) (key.TargetID, error)

const (
	First         key.TargetEvaluator = 100 // i hope no one spawns 100+ targets
	LowestHP      key.TargetEvaluator = 101
	LowestHPRatio key.TargetEvaluator = 102
)

var evaluators = map[key.TargetEvaluator]evaluator{
	First:         first,
	LowestHP:      lowestHP,
	LowestHPRatio: lowestHPRatio,
}

func first(engine engine.Engine, candidates []key.TargetID) (key.TargetID, error) {
	return candidates[0], nil
}

func lowestHPRatio(engine engine.Engine, candidates []key.TargetID) (key.TargetID, error) {
	min := candidates[0]
	minHP := engine.HPRatio(candidates[0])
	for _, c := range candidates {
		if hp := engine.HPRatio(c); hp < minHP {
			min, minHP = c, hp
		}
	}
	return min, nil
}

func lowestHP(engine engine.Engine, candidates []key.TargetID) (key.TargetID, error) {
	min := candidates[0]
	minHP := engine.Stats(candidates[0]).CurrentHP()
	for i := 1; i < len(candidates); i++ {
		if hp := engine.Stats(candidates[i]).CurrentHP(); hp < minHP {
			min, minHP = candidates[i], hp
		}
	}
	return min, nil
}
