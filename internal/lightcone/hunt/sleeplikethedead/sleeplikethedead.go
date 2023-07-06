package sleeplikethedead

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
	Check    key.Modifier = "sleep-like-the-dead-check"
	Buff     key.Modifier = "sleep-like-the-dead-buff"
	Cooldown key.Modifier = "sleep-like-the-dead-cooldown"
)

// Increases the wearer's CRIT DMG by 30/35/40/45/50%. When the wearer's Basic
// ATK or Skill does not result in a CRIT Hit, increases their CRIT Rate by
// 36/42/48/54/60% for 1 turn(s). This effect can only trigger once every
// 3 turn(s).
func init() {
	lightcone.Register(key.SleepLiketheDead, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHit: onAfterHit,
		},
	})

	modifier.Register(Buff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
	})

	modifier.Register(Cooldown, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	critDmg := 0.25 + 0.05*float64(lc.Imposition)
	critRate := 0.3 + 0.06*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats: info.PropMap{
			prop.CritDMG: critDmg,
		},
		State: critRate,
	})
}

func onAfterHit(mod *modifier.Instance, e event.HitEnd) {
	if !e.IsCrit && !mod.Engine().HasModifier(mod.Owner(), Cooldown) {
		if e.AttackType == model.AttackType_NORMAL || e.AttackType == model.AttackType_SKILL {
			mod.Engine().AddModifier(mod.Owner(), info.Modifier{
				Name:   Buff,
				Source: mod.Owner(),
				Stats: info.PropMap{
					prop.CritChance: mod.State().(float64),
				},
				Duration: 1,
			})

			mod.Engine().AddModifier(mod.Owner(), info.Modifier{
				Name:     Cooldown,
				Source:   mod.Owner(),
				Duration: 2,
			})
		}
	}
}
