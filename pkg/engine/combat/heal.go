package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) Heal(heal info.Heal) {
	if len(heal.Targets) == 0 || !mgr.attr.IsAlive(heal.Source) {
		return
	}

	for _, t := range heal.Targets {
		source := mgr.attr.Stats(heal.Source)
		target := mgr.attr.Stats(t)

		// make copy of healMap
		baseHeal := make(info.HealMap, len(heal.BaseHeal))
		for k, v := range heal.BaseHeal {
			baseHeal[k] = v
		}

		e := &event.HealStart{
			Target:      target,
			Healer:      source,
			BaseHeal:    baseHeal,
			HealValue:   heal.HealValue,
			UseSnapshot: heal.UseSnapshot,
		}
		mgr.event.HealStart.Emit(e)

		// Get base heal amount
		hpLost := target.MaxHP() - target.HP()
		base := heal.HealValue
		for k, v := range baseHeal {
			switch k {
			case model.HealFormula_BY_HEALER_ATK:
				base += v * source.ATK()
			case model.HealFormula_BY_HEALER_DEF:
				base += v * source.DEF()
			case model.HealFormula_BY_HEALER_MAX_HP:
				base += v * source.MaxHP()
			case model.HealFormula_BY_TARGET_MAX_HP:
				base += v * target.MaxHP()
			case model.HealFormula_BY_TARGET_LOST_HP:
				base += v * hpLost
			}
		}

		// Apply Outgoing Heal Bonus of healer
		// Apply Incoming Heal Bonus of target
		healAmount := base * (1 + source.HealBoost()) * (1 + target.GetProperty(prop.HealTaken))
		overflow := 0.0
		if healAmount+target.HP() > target.MaxHP() {
			overflow = healAmount + target.HP() - target.MaxHP()
			healAmount -= overflow
		}

		// Call ModifyHP to add the new HP to the healed target
		mgr.attr.ModifyHPByAmount(t, heal.Source, healAmount, false)

		mgr.event.HealEnd.Emit(event.HealEnd{
			Target:             t,
			Healer:             heal.Source,
			HealAmount:         healAmount,
			OverflowHealAmount: overflow,
			UseSnapshot:        heal.UseSnapshot,
		})
	}
}
