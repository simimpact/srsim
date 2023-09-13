package global

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/hook"
	"github.com/simimpact/srsim/pkg/engine/info"
)

// Use this to add any global hooks/game logic

func init() {
	hook.RegisterStartupHook("EnergyOnDeath", EnergyOnDeath)
	hook.RegisterStartupHook("DamageAndDebuffOnWeaknessBreak", DamageAndDebuffOnWeaknessBreak)
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

// When enemy suffers weakness break, deal damage to it and apply debuffs
func DamageAndDebuffOnWeaknessBreak(engine engine.Engine) error {
	engine.Events().StanceBreak.Subscribe(func(event event.StanceBreak) {
		common.ApplyWeaknessBreakEffects(engine, event.Source, event.Target)
	})
	return nil
}
