package loop

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	loop key.Modifier = "loop"
)

// DESC : Increases DMG dealt from its wearer to Slowed enemies by 24%.

// DM Listeners :
// OnBeforeHitALl : if defender has STAT_SpeedDown, ModifyDamageRatio

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

}

func boostDmgOnSlowed(mod *modifier.Instance, e event.HitStart) {

}