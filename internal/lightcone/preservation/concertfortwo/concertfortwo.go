package concertfortwo

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
	Check = "concert-for-two"
	Buff  = "concert-for-two-buff"
)

type state struct {
	shieldCount int
	dmgPerStack float64
}

// Increases the wearer's DEF by 16/20/24/28/32%.
// For every on-field character that has a Shield, the DMG dealt by the wearer increases by 4/5/6/7/8%.

func init() {
	lightcone.Register(key.ConcertforTwo, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: listenShield,
		},
	})

	modifier.Register(Buff, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	state := state{
		shieldCount: 0,
		dmgPerStack: 0.03 + 0.01*float64(lc.Imposition),
	}

	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats:  info.PropMap{prop.DEFPercent: 0.12 + 0.04*float64(lc.Imposition)},
		State:  &state,
	})
}

func listenShield(mod *modifier.Instance) {
	st := mod.State().(*state)

	mod.Engine().Events().ShieldAdded.Subscribe(func(event event.ShieldAdded) {
		if !mod.Engine().IsCharacter(event.Info.Target) {
			return
		}
		st.shieldCount++
		mod.Engine().AddModifier(mod.Source(), info.Modifier{
			Name:   Buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllDamagePercent: st.dmgPerStack * float64(st.shieldCount)},
		})
	})

	mod.Engine().Events().ShieldRemoved.Subscribe(func(event event.ShieldRemoved) {
		if !mod.Engine().IsCharacter(event.Target) {
			return
		}
		st.shieldCount--
		if st.shieldCount == 0 {
			mod.Engine().RemoveModifier(mod.Source(), Buff)
		}
	})
}
