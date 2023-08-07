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
	memories   = "memories-of-the-past"
	memoriesCD = "memories-of-the-past-cd"
)

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

	modifier.Register(memoriesCD, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	energyAmt := 3.0 + 1.0*float64(lc.Imposition)
	beAmt := 0.21 + 0.07*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   memories,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: beAmt},
	})

	// Resets CD on all turn end
	engine.Events().TurnEnd.Subscribe(func(e event.TurnEnd) {
		engine.RemoveModifier(owner, memoriesCD)
	})

	// add energy on all attack end by lc holder (for ults etc.)
	engine.Events().AttackEnd.Subscribe(func(event event.AttackEnd) {
		if engine.HasModifier(owner, memoriesCD) || event.Attacker != owner {
			return
		}
		engine.ModifyEnergy(info.ModifyAttribute{
			Key:    memories,
			Target: owner,
			Source: owner,
			Amount: energyAmt,
		})
		// enter cd
		engine.AddModifier(owner, info.Modifier{
			Name:     memoriesCD,
			Source:   owner,
			Duration: 1,
		})
	})
}
