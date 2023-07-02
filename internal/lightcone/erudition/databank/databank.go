package databank

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
	DataBank key.Modifier = "data-bank"
)

func init() {
	lightcone.Register(key.DataBank, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})

	modifier.Register(DataBank, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: onBeforeHit,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   DataBank,
		Source: owner,
		State:  0.21 + float64(lc.Imposition)*0.07,
	})
}

func onBeforeHit(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_ULT {
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, mod.State().(float64))
	}
}
