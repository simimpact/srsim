package loop

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
	loop = "loop"
)

// DESC : Increases DMG dealt from its wearer to Slowed enemies by 24%.

func init() {
	lightcone.Register(key.Loop, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(loop, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: boostDmgOnSlowed,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	dmgBoostAmt := 0.18 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   loop,
		Source: owner,
		State:  dmgBoostAmt,
	})
}

func boostDmgOnSlowed(mod *modifier.Instance, e event.HitStart) {
	amt := mod.State().(float64)
	if mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_SPEED_DOWN) {
		e.Hit.Attacker.AddProperty(loop, prop.AllDamagePercent, amt)
	}
}
