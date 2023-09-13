package march7th

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult = "march7th-ult"
	E1  = "march7th-e1"
)

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

	frozenEnemiesCount := 0.0
	for _, freezeTarget := range c.engine.Enemies() {
		successfullyApplied, _ := c.engine.AddModifier(freezeTarget, info.Modifier{
			Name:   common.Freeze,
			Source: c.id,
			State: &common.FreezeState{
				DamagePercentage: ultFreeze[c.info.UltLevelIndex()],
				DamageValue:      0,
			},
			Chance:   freezeChance,
			Duration: 1,
		})

		if successfullyApplied {
			frozenEnemiesCount++
		}
	}
	if c.info.Eidolon >= 1 {
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    E1,
			Target: c.id,
			Source: c.id,
			Amount: 6 * frozenEnemiesCount,
		})
	}
}

func (c *char) ultHit(hitIndex int, hitRatio float64) {
	hitTargets := make(map[key.TargetID]bool, len(c.engine.Enemies()))
	for i := 0; i < 3; i++ {
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
			HitRatio:     hitRatio,
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
		Amount: 5 * hitRatio,
	})
}
