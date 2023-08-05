package march7th

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	MarchE1       = "march7th-e1"
	MarchE2       = "march7th-e2"
	MarchE2Shield = "march7th-e2-shield"
)

func init() {
	modifier.Register(MarchE2Shield, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.Engine().AddShield(MarchE2Shield, info.Shield{
					Source: mod.Source(),
					Target: mod.Owner(),
					BaseShield: info.ShieldMap{
						model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.24,
					},
					ShieldValue: 320,
				})
			},
			OnRemove: func(mod *modifier.Instance) {
				mod.Engine().RemoveShield(MarchE2Shield, mod.Owner())
			},
		},
	})

	modifier.Register(MarchE2, modifier.Config{
		Stacking: modifier.Replace,
		Listeners: modifier.Listeners{
			OnAdd: determineE2Target,
		},
	})
}

func (c *char) initEidolons() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   MarchE2,
		Source: c.id,
	})

}

func (c *char) e1() {
	for _, target := range c.engine.Enemies() {
		if c.engine.HasModifier(target, MarchUltFreeze) {
			c.engine.ModifyEnergy(info.ModifyAttribute{
				Key:    MarchE1,
				Source: c.id,
				Target: target,
				Amount: 6,
			})
		}
	}
}

func determineE2Target(mod *modifier.Instance) {
	lowestHpRatio := 1.0
	for _, Target := range mod.Engine().Characters() {
		if mod.Engine().HPRatio(Target) <= lowestHpRatio {
			lowestHpRatio = mod.Engine().HPRatio(Target)
		}
	}
	E2Target := mod.Engine().Retarget(info.Retarget{
		Targets: mod.Engine().Characters(),
		Max:     1,
		Filter: func(target key.TargetID) bool {
			return mod.Engine().HPRatio(target) == lowestHpRatio
		},
	})[0]
	mod.Engine().AddModifier(E2Target, info.Modifier{
		Name:     MarchE2Shield,
		Source:   mod.Source(),
		Duration: 3,
	})
}
