package global

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/hook"
)

// Use this to add any global hooks/game logic

func init() {
	hook.RegisterStartupHook("EnergyOnDeath", EnergyOnDeath)
}

// When a target dies, give 10 energy to the killer (that can be scaled with ERR)
func EnergyOnDeath(engine engine.Engine) error {
	engine.Events().TargetDeath.Subscribe(func(event event.TargetDeath) {
		engine.ModifyEnergy(event.Killer, 10.0)
	})
	return nil
}
