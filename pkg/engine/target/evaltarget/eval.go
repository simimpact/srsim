package evaltarget

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Info struct {
	Source      key.TargetID
	Evaluator   key.TargetEvaluator
	TargetType  model.TargetType
	SourceClass info.TargetClass
}

func Evaluate(engine engine.Engine, i Info) (key.TargetID, error) {
	eval, ok := evaluators[i.Evaluator]
	if !ok {
		// so it's target id
		return validateTarget(engine, i)
	}

	targets, err := candidates(engine, i)
	if err != nil || len(targets) == 0 {
		return -1, fmt.Errorf("unknown error trying to get targets %w, %v", err, targets)
	}

	if len(targets) == 1 {
		return targets[0], nil
	}
	return eval(engine, targets)
}

// Note: neutral is treated like characters (IE lightning-lord)
func candidates(engine engine.Engine, i Info) ([]key.TargetID, error) {
	switch i.TargetType {
	case model.TargetType_ALLIES:
		if i.SourceClass == info.ClassEnemy {
			return engine.Enemies(), nil
		}
		return engine.Characters(), nil
	case model.TargetType_ENEMIES:
		if i.SourceClass == info.ClassEnemy {
			return engine.Characters(), nil
		}
		return engine.Enemies(), nil
	case model.TargetType_SELF:
		return []key.TargetID{i.Source}, nil
	default:
		return nil, fmt.Errorf("unknown TargetType: %v", i.TargetType)
	}
}

func validateTarget(engine engine.Engine, i Info) (key.TargetID, error) {
	target := key.TargetID(i.Evaluator)
	if !engine.IsValid(target) {
		return -1, fmt.Errorf("invalid target id: %d", target)
	}
	if !engine.IsAlive(target) {
		return -1, fmt.Errorf("target %d is dead", target)
	}

	switch i.TargetType {
	case model.TargetType_ALLIES:
		if !engine.IsCharacter(target) {
			return -1, fmt.Errorf("target %d is not a character", target)
		}
	case model.TargetType_ENEMIES:
		if !engine.IsEnemy(target) {
			return -1, fmt.Errorf("target %d is not an enemy", target)
		}
	case model.TargetType_SELF:
		if target != i.Source {
			return -1, fmt.Errorf("target %d is not the same as the source (%d)", target, i.Source)
		}
	default:
		return -1, fmt.Errorf("unknown TargetType: %v", i.TargetType)
	}
	return target, nil
}
