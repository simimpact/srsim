package character

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
)

type CharInstance interface {
	Attack(engine engine.Engine, target key.TargetID)
	Skill(engine engine.Engine, target key.TargetID)
}

type singleBurst interface {
	Burst(engine engine.Engine, target key.TargetID)
}

type multiBurst interface {
	BurstAttack(engine engine.Engine, target key.TargetID)
	BurstSkill(engine engine.Engine, target key.TargetID)
}
