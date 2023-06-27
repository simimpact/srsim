package somethingirreplaceable

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
	Check key.Modifier = "something-irreplaceable"
)

type state struct {
	Heal     float64
	DmgBonus float64
}

//Increases the wearer's ATK by 24%.
//When the wearer defeats an enemy or is hit, immediately restores HP equal to 8% of the wearer's ATK.
//At the same time, the wearer's DMG is increased by 24% until the end of their next turn.
//This effect cannot stack and can only trigger 1 time per turn.

func init() {
	lightcone.Register(key.SomethingIrreplaceable, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: onAfterBeingAttacked,
			OnTriggerDeath:       onTriggerDeath,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkBuff := 0.2 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: atkBuff},
		State:  state{Heal: 0.07 + 0.01*float64(lc.Imposition), DmgBonus: 0.2 + 0.04*float64(lc.Imposition)},
	})
}

func onTriggerDeath(mod *modifier.ModifierInstance, target key.TargetID) {
	conditions(mod)
}

func onAfterBeingAttacked(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	conditions(mod)
}

func conditions(mod *modifier.ModifierInstance) {
	heal := mod.State().(*state).Heal
	dmgBonus := mod.State().(*state).DmgBonus

	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Source:   mod.Owner(),
		Duration: 1,
		Stats:    info.PropMap{prop.AllDamagePercent: dmgBonus},
	})

	mod.Engine().Heal(info.Heal{
		Targets:  []key.TargetID{mod.Owner()},
		Source:   mod.Owner(),
		BaseHeal: info.HealMap{model.HealFormula_BY_HEALER_ATK: heal},
	})
}
