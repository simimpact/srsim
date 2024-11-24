package geniusesrepose

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	checker key.Modifier = "geniuses-repose-checker"
	buff    key.Modifier = "geniuses-repose-crit-dmg"
)

// DESC : Increases the wearer's ATK by 16%.
// When the wearer defeats an enemy, the wearer's CRIT DMG increases by 24% for 3 turn(s).

func init() {
	lightcone.Register(key.GeniusesRepose, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
	modifier.Register(checker, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: boostCDmg,
		},
	})
	modifier.Register(buff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkAmt := 0.12 + 0.04*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   checker,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: atkAmt},
		State:  0.18 + 0.06*float64(lc.Imposition),
	})
}

// enemy killed -> add CDmg buff, 3 turns
func boostCDmg(mod *modifier.Instance, target key.TargetID) {
	cDmgAmt := mod.State().(float64)

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     buff,
		Source:   mod.Owner(),
		Duration: 3,
		Stats:    info.PropMap{prop.CritDMG: cDmgAmt},
	})
}
