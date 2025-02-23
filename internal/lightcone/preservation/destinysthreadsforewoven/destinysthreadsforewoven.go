package destinysthreadsforewoven

import (
	"math"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	check = "destinys-threads-forewoven"
	buff  = "destinys-threads-forewoven-dmg-buff"
)

type state struct {
	dmgPer float64
	dmgMax float64
}

// Increases the wearer's Effect RES by 12%.
// For every 100 of DEF the wearer has, increases the wearer's DMG dealt by 0.8%,
// up to a maximum DMG increase of 32%.

func init() {
	lightcone.Register(key.DestinysThreadsForewoven, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:            onCheck,
			OnPropertyChange: onCheck,
		},
	})

	modifier.Register(buff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		CanDispel:  true,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	effres := 0.1 + 0.02*float64(lc.Imposition)
	dmgPer := 0.007 + 0.001*float64(lc.Imposition)
	dmgMax := 0.28 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   check,
		Source: owner,
		Stats:  info.PropMap{prop.EffectRES: effres},
		State: state{
			dmgPer: dmgPer,
			dmgMax: dmgMax,
		},
	})
}

func onCheck(mod *modifier.Instance) {
	state := mod.State().(state)
	def := mod.OwnerStats().DEF()
	dmg := math.Floor(def/100) * state.dmgPer
	if dmg > state.dmgMax {
		dmg = state.dmgMax
	}

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   buff,
		Source: mod.Owner(),
		Stats:  info.PropMap{prop.AllDamagePercent: dmg},
	})
}
