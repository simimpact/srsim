package guinaifen

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E4 = "guinaifen-e4"
)

func init() {
	modifier.Register(E4Listener, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingHitAll: checkE4,
		},
	})
}

func checkE4(mod *modifier.Instance, e event.HitEnd) {
	cond1 := e.AttackType == model.AttackType_DOT
	cond2 := e.DamageType == model.DamageType_FIRE
	// TODO: check if damage is not from split damage
	cond3 := true
	cond4 := e.Attacker == mod.Source()
	if cond1 && cond2 && cond3 && cond4 {
		mod.Engine().ModifyEnergy(info.ModifyAttribute{
			Key:    E4,
			Target: mod.Source(),
			Source: mod.Source(),
			Amount: 2.0,
		})
	}
}
