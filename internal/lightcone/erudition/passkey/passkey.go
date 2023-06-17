package passkey

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
	Passkey key.Modifier = "passkey"
)

// After the wearer uses their Skill, additionally regenerates 8/9/10/11/12 Energy.
// This effect can only be triggered 1 time per turn.
func init() {
	lightcone.Register(key.Passkey, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})

	modifier.Register(Passkey, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction: onAfterAction,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	mod := info.Modifier{
		Name:   Passkey,
		Source: owner,
		State:  7.0 + float64(lc.Ascension),
	}

	// At the end of every turn, add this modifier (no-op if already exists)
	engine.Events().TurnEnd.Subscribe(func(event event.TurnEndEvent) {
		engine.AddModifier(owner, mod)
	})

	// start with this modifier added to owner (turn 1 case)
	engine.AddModifier(owner, mod)
}

// after giving energy, remove this modifier so it cannot do it again
func onAfterAction(mod *modifier.ModifierInstance, e event.ActionEvent) {
	mod.Engine().ModifyEnergy(mod.Owner(), mod.State().(float64))
	mod.RemoveSelf()
}
