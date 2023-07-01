package meshingcogs

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const MeshingCogs key.Modifier = "meshing_cogs"

type cogsState struct {
	imposition int
	charged    bool
}

func init() {
	lightcone.Register(key.MeshingCogs, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})

	modifier.Register(MeshingCogs, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack:        tryEnergyRegen,
			OnAfterBeingAttacked: tryEnergyRegen,
		},
	})
}

func tryEnergyRegen(mod *modifier.Instance, _ event.AttackEndEvent) {
	state := mod.State().(*cogsState)
	if !state.charged {
		return
	}
	mod.Engine().ModifyEnergy(mod.Owner(), float64(3+state.imposition))
	state.charged = false
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	state := cogsState{
		imposition: lc.Imposition,
		charged:    true,
	}
	engine.AddModifier(owner, info.Modifier{
		Name:   MeshingCogs,
		Source: owner,
		State:  &state,
	})

	engine.Events().TurnEnd.Subscribe(func(e event.TurnEnd) {
		state.charged = true
	})
}
