package memoriesofthepast

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
	memories = "memories-of-the-past"
)

type state struct {
	onCooldown bool
	energyAmt  float64
}

// Increases the wearer's Break Effect by 28%. When the wearer attacks,
// additionally regenerates 4 Energy. This effect can only be triggered 1 time per turn.

func init() {
	lightcone.Register(key.MemoriesofthePast, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})
	modifier.Register(memories, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	modState := state{
		onCooldown: false,
		energyAmt:  3.0 + 1.0*float64(lc.Imposition),
	}
	beAmt := 0.21 + 0.07*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   memories,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: beAmt},
	})

	// REPLACE : currently uses mod listener. change it to subscriber.
	// NOTE : might want to change cooldown to a separate Mod.

	// Resets CD on all turn end
	engine.Events().TurnEnd.Subscribe(func(e event.TurnEnd) {
		modState.onCooldown = false
	})

	// add energy on all attack end (for ults etc.)
	engine.Events().AttackEnd.Subscribe(func(event event.AttackEnd) {
		if modState.onCooldown {
			return
		}
		engine.ModifyEnergy(info.ModifyAttribute{
			Key:    memories,
			Target: owner,
			Source: owner,
			Amount: modState.energyAmt,
		})
		// enter cd
		modState.onCooldown = true
	})
}
