package modifier

import "github.com/simimpact/srsim/pkg/model"

var (
	behaviorFlagImmunities = map[model.BehaviorFlag][]model.BehaviorFlag{
		model.BehaviorFlag_ENDURANCE: {model.BehaviorFlag_STAT_CTRL},
	}

	statusTypeImmmunities = map[model.BehaviorFlag][]model.StatusType{
		model.BehaviorFlag_RESIST_DEBUFF: {model.StatusType_STATUS_DEBUFF},
	}
)
