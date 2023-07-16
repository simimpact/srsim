package beforethetutorialmissionstarts

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	name = "before-tutorial-mission-starts"
)

// Increases the wearer's Effect Hit Rate by 20%.
// When the wearer attacks DEF-reduced enemies, regenerates 4 Energy.

func init() {
	lightcone.Register(key.BeforetheTutorialMissionStarts, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(name, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: addEnergy,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	ehrAmt := 0.15 + 0.05*float64(lc.Imposition)
	energyAmt := 3.0 + float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   name,
		Source: owner,
		Stats:  info.PropMap{prop.EffectHitRate: ehrAmt},
		State:  energyAmt,
	})
}

func addEnergy(mod *modifier.Instance, e event.AttackEnd) {
	energyAmt := mod.State().(float64)
	if hasDefReducedTarget(mod.Engine(), e.Targets) {
		mod.Engine().ModifyEnergy(info.ModifyAttribute{
			Key:    name,
			Target: mod.Owner(),
			Source: mod.Owner(),
			Amount: energyAmt,
		})
	}
}

// replace hacky use of Retarget. now can exit once an instance of def down enemy is found.
func hasDefReducedTarget(engine engine.Engine, targets []key.TargetID) bool {
	for _, target := range targets {
		if engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DEF_DOWN) {
			return true
		}
	}
	return false
}
