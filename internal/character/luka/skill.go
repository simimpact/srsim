package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.e1Check(target)

	if c.info.Eidolon >= 2 && c.engine.Stats(target).IsWeakTo(model.DamageType_PHYSICAL) {
		c.fightingSpirit += 1
	}

}
