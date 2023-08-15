package patienceisallyouneed

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
	patience key.Modifier = "patience-is-all-you-need"
	spdBuff  key.Modifier = "patience-is-all-you-need-spd-buff"
	erode                 = "patience-is-all-you-need-erode"
)

// Increases DMG dealt by the wearer by 24%. After every attack launched by wearer,
// their SPD increases by 4.8%, stacking up to 3 times. If the wearer hits an enemy target
// that is not afflicted by Erode, there is a 100% base chance to inflict Erode to the target.
// Enemies afflicted with Erode are also considered to be Shocked and will receive
// Lightning DoT at the start of each turn equal to 60% of the wearer's ATK, lasting for 1 turn(s).

func init() {
	lightcone.Register(key.PatienceIsAllYouNeed, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(patience, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHit:    inflictErode,
			OnAfterAttack: addSpeedBuff,
		},
	})
	modifier.Register(spdBuff, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_SPEED_UP,
		},
		Stacking: modifier.ReplaceBySource,
	})
	modifier.Register(erode, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: addDotDmg,
		},
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT_ELECTRIC,
		},
		Stacking: modifier.ReplaceBySource,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}

func addSpeedBuff(mod *modifier.Instance, e event.AttackEnd) {

}

func inflictErode(mod *modifier.Instance, e event.HitEnd) {

}

func addDotDmg(mod *modifier.Instance) {

}
