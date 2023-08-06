package march7th

import (
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
	targets := c.engine.Enemies()

	for i, hitRatio := range ultHits {
		c.engine.Attack(info.Attack{
			Key:          Ult,
			Source:       c.id,
			Targets:      targets,
			HitIndex:     i,
			HitRatio:     hitRatio,
			StanceDamage: 60,
			EnergyGain:   5,
			DamageType:   model.DamageType_ICE,
			AttackType:   model.AttackType_ULT,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
			},
		})
	}

	freezeChance := 0.5
	if c.info.Traces["103"] {
		freezeChance += 0.15
	}

	for _, freezeTarget := range targets {
		c.engine.AddModifier(freezeTarget, info.Modifier{
			Name:     MarchUltFreeze,
			Source:   c.id,
			Chance:   freezeChance,
			Duration: 1,
		})
	}
	c.e1()
}
