package global

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/hook"
	"github.com/simimpact/srsim/pkg/engine/info"
)

// Use this to add any global hooks/game logic

func init() {
	hook.RegisterStartupHook("EnergyOnDeath", EnergyOnDeath)
}

// When a target dies, give 10 energy to the killer (that can be scaled with ERR)
func EnergyOnDeath(engine engine.Engine) error {
	engine.Events().TargetDeath.Subscribe(func(event event.TargetDeath) {
		engine.ModifyEnergy(info.ModifyAttribute{
			Key:    "kill",
			Target: event.Killer,
			Source: event.Target,
			Amount: 10.0,
		})
	})
	return nil
}
