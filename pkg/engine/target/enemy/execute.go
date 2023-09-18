package enemy

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/target"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) ExecuteAction(id key.TargetID, isInsert bool) (target.ExecutableAction, error) {
	enemy := mgr.instances[id]
	primaryTarget := chooseTarget(mgr.engine)

	return target.ExecutableAction{
		Execute: func() {
			enemy.Action(primaryTarget, actionState{
				engine:   mgr.engine,
				isInsert: isInsert,
			})
		},
		AttackType: model.AttackType_NORMAL,
		IsInsert:   isInsert,
		SPDelta:    0,
		Key:        mgr.info[id].Key.String(),
	}, nil
}

// choose the primary target using aggro to determine the probabilities for each character
// being a target
func chooseTarget(engine engine.Engine) key.TargetID {
	charIDs := engine.Characters()
	charAggro := make([]float64, len(charIDs))

	var total float64
	for i, c := range charIDs {
		total += engine.Stats(c).Aggro()
		charAggro[i] = total
	}

	choice := engine.Rand().Float64() * total
	for i, aggro := range charAggro {
		if aggro < choice {
			return charIDs[i]
		}
	}

	// grab last if nothing returns, should never happen
	return charIDs[len(charIDs)-1]
}

type actionState struct {
	engine   engine.Engine
	isInsert bool
}

func (a actionState) IsInsert() bool {
	return a.isInsert
}

func (a actionState) EndAttack() {
	a.engine.EndAttack()
}
