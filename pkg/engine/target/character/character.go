package character

import (
	"github.com/simimpact/srsim/pkg/key"
)

type CharInstance interface {
	Attack(target key.TargetID)
	Skill(target key.TargetID)
}

type SingleBurst interface {
	Burst(target key.TargetID)
}

type MultiBurst interface {
	BurstAttack(target key.TargetID)
	BurstSkill(target key.TargetID)
}
