package shadowedbynight

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
	check = "shadowed_by_night"
	buff  = "shadowed_by_night_spd_buff"
)

type state struct {
	spdAmt float64
	flag   bool
}

// Increases the wearer's Break Effect by 28/35/42/49/56%.
// When entering battle or after dealing Break DMG, increases SPD by 8/9/10/11/12%, lasting for 2 turn(s).
// This effect can only trigger once per turn.
func init() {
	lightcone.Register(key.ShadowedbyNight, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(check, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1:       resetFlag,
			OnBeforeHitAll: applySpdBuff,
		},
	})

	modifier.Register(buff, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_BUFF,
		CanDispel:     true,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	beAmt := 0.21 + 0.07*float64(lc.Imposition)
	spdAmt := 0.07 + 0.01*float64(lc.Imposition)
	state := state{
		spdAmt: spdAmt,
		flag:   false,
	}
	engine.AddModifier(owner, info.Modifier{
		Name:   check,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: beAmt},
		State:  &state,
	})

	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		engine.AddModifier(owner, info.Modifier{
			Name:     buff,
			Source:   owner,
			Stats:    info.PropMap{prop.SPDPercent: spdAmt},
			Duration: 2,
		})
	})
}

func resetFlag(mod *modifier.Instance) {
	st := mod.State().(*state)
	st.flag = false
}

func applySpdBuff(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType != model.AttackType_ELEMENT_DAMAGE {
		return
	}
	st := mod.State().(*state)
	if !st.flag {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     buff,
			Source:   mod.Owner(),
			Stats:    info.PropMap{prop.SPDPercent: st.spdAmt},
			Duration: 2,
		})
	}
	st.flag = true
}
