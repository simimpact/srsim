package march7th

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	MarchUltFreeze = "march7th-ultfreeze"
	Ult            = "march7th-ult"
)

func init() {
	modifier.Register(MarchUltFreeze, modifier.Config{
		Duration: 1,
	})
}

var ultHits = []float64{0.25, 0.25, 0.25, 0.25}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	for hitIndex, hitRatio := range ultHits {
		c.ultHit(hitIndex, hitRatio)
	}

	c.engine.EndAttack()

	freezeChance := 0.5
	if c.info.Traces["103"] {
		freezeChance += 0.15
	}

	for _, freezeTarget := range c.engine.Enemies() {
		successfullyApplied, _ := c.engine.AddModifier(freezeTarget, info.Modifier{
			Name:   common.Freeze,
			Source: c.id,
			State: common.FreezeState{
				DamagePercentage: ultFreeze[c.info.UltLevelIndex()],
				DamageValue:      0,
			},
			Chance:   freezeChance,
			Duration: 1,
		})

		if c.info.Eidolon >= 1 {
			if successfullyApplied {
				c.engine.ModifyEnergy(info.ModifyAttribute{
					Key:    Ult,
					Target: c.id,
					Source: c.id,
					Amount: 6,
				})
			}
		}
	}
}

func (c *char) ultHit(hitIndex int, hitRatio float64) {
	hitTargets := make(map[key.TargetID]bool)
	for i := 0; i < 2; i++ {
		targets := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Enemies(),
			Filter: func(target key.TargetID) bool {
				_, beenHit := hitTargets[target]
				return !beenHit
			},
			Max: 2,
		})

		c.engine.Attack(info.Attack{
			Key:          Ult,
			Source:       c.id,
			Targets:      targets,
			HitIndex:     hitIndex,
			HitRatio:     hitRatio / 3,
			StanceDamage: 60,
			EnergyGain:   0,
			DamageType:   model.DamageType_ICE,
			AttackType:   model.AttackType_ULT,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
			},
		})
		// Log the targets as hit in the map
		for _, t := range targets {
			hitTargets[t] = true
		}
	}

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Ult,
		Target: c.id,
		Source: c.id,
		Amount: 5 * 0.25,
	})
}
