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

var skillHits = []float64{0.15, 0.35, 0.5}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	for _, hitRatio := range skillHits {
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_ICE,
			AttackType: model.AttackType_SKILL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skillDMG[c.info.SkillLevelIndex()],
			},
			StanceDamage: 60.0,
			EnergyGain:   30.0,
			HitRatio:     hitRatio,
		})
	}

	c.engine.EndAttack()

	freezeChance := 0.65

	if c.info.Eidolon >= 1 {
		freezeChance += 0.35
	}

	freezeSucessful, _ := c.engine.AddModifier(target, info.Modifier{
		Name:   common.Freeze,
		Source: c.id,
		State: common.FreezeState{
			DamagePercentage: skillFreezeDMG[c.info.SkillLevelIndex()],
			DamageValue:      0,
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
