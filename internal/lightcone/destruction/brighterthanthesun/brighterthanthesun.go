package brighterthanthesun

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
	brighter    key.Modifier = "brighter-than-the-sun"
	dragonsCall key.Modifier = "dragons-call"
)

// Increases the wearer's CRIT Rate by 18%. When the wearer uses their Basic ATK,
// they will gain 1 stack of Dragon's Call, lasting for 2 turns.
// Each stack of Dragon's Call increases the wearer's ATK by 18% and
// Energy Regeneration Rate by 6%. Dragon's Call can be stacked up to 2 times.

func init() {
	lightcone.Register(key.BrighterThantheSun, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})
	modifier.Register(brighter, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: dragonsCallOnBasic,
		},
	})
	modifier.Register(dragonsCall, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          2,
		CountAddWhenStack: 1,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// add flat crit rate at stats.
	engine.AddModifier(owner, info.Modifier{
		Name:   brighter,
		Source: owner,
		// TODO : create state struct for atk and err buff.
	})
}

func dragonsCallOnBasic(mod *modifier.Instance, e event.ActionStart) {

}
