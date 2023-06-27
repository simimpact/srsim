package wearewildfire

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
	mod     key.Modifier = "wearewildfire"
	modHeal key.Modifier = "wearewildfire-heal"
)

// At the start of the battle, the DMG dealt to all allies decreases by
// 8%/10%/12%/14%/16% for 5 turn(s). At the same time, immediately restores HP
// to all allies equal to 30%/35%/40%/45%/50% of the respective HP difference
// between the characters' Max HP and current HP.
func init() {
	lightcone.Register(key.WeAreWildfire, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(mod, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		TickMoment: modifier.ModifierPhase1End,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// team DMG RES
	amt := 0.06 + 0.02*float64(lc.Imposition)
	amtHeal := 0.25 + 0.05*float64(lc.Imposition)

	dmgmod := info.Modifier{
		Name:     mod,
		Source:   owner,
		Stats:    info.PropMap{prop.AllDamageReduce: amt},
		Duration: 5,
	}

    // TODO: recheck when multiple waves support is added
	engine.Events().BattleStart.Subscribe(func(event event.BattleStartEvent) {
		for char := range event.CharInfo {
			engine.AddModifier(char, dmgmod)
		}

		engine.Heal(info.Heal{
			Targets: engine.Characters(),
			Source:  owner,
			BaseHeal: info.HealMap{
				model.HealFormula_BY_TARGET_LOST_HP: amtHeal,
			},
		})
	})
}
