package gepard

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2Tracker key.Modifier = "gepard-e2-tracker"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_ICE,
		AttackType: model.AttackType_NORMAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillDMG[c.info.SkillLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   30.0,
	})

	freezeChance := 0.65

	if c.info.Eidolon >= 1 {
		freezeChance += 0.35
	}

	freezeSucessful, _ := c.engine.AddModifier(target, info.Modifier{
		Name:   common.Freeze,
		Source: c.id,
		State: common.FreezeState{
			DamagePercentage: skillFreezeDMG[c.info.SkillLevelIndex()],
		},
		Chance:   freezeChance,
		Duration: 1,
	})

	if c.info.Eidolon >= 2 && freezeSucessful {
		c.engine.AddModifier(target, info.Modifier{
			Name:   E2Tracker,
			Source: c.id,
		})
	}
}
