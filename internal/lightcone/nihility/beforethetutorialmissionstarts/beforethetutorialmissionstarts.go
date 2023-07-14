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
	mod key.Modifier = "before-tutorial-mission-starts"
)

// Increases the wearer's Effect Hit Rate by 20%.
// When the wearer attacks DEF-reduced enemies, regenerates 4 Energy.

// DM :
// OnAfterAttack : retarget() wtf? byContainBehaviorFLag : stat_defenceDown, includeLimbo
// max : 1
// -> TaskList : modifySPNew (wtf?) add value by some dynamic amt -> SkillTreeParam : ""(nil)
// OnStart : add _Main mod

func init() {
	lightcone.Register(key.BeforetheTutorialMissionStarts, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: addEnergy,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	ehrAmt := 0.15 + 0.05*float64(lc.Imposition)
	energyAmt := 3.0 + float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.EffectHitRate: ehrAmt},
		State:  energyAmt,
	})
}

var triggerFlags = []model.BehaviorFlag{
	model.BehaviorFlag_STAT_DEF_DOWN,
}

func addEnergy(mod *modifier.Instance, e event.AttackEnd) {
	energyAmt := mod.State().(float64)
	// Retarget : filter : has def down behavior flag, Max 1, includeLimbo
	qualified := mod.Engine().Retarget(info.Retarget{
		Targets: e.Targets,
		Filter: func(target key.TargetID) bool {
			// returns true if target has triggerFlags, otherwise false
			return mod.Engine().HasBehaviorFlag(target, triggerFlags...)
		},
		Max:          1,
		IncludeLimbo: true,
	})
	if len(qualified) > 0 {
		mod.Engine().ModifyEnergy(mod.Owner(), energyAmt)
	}
}
