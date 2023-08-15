package wewillmeetagain

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
	meet = "we-will-meet-again"
)

// After the wearer uses Basic ATK or Skill, deals Additional DMG
// equal to 48% of the wearer's ATK to a random enemy that has been attacked.

// IMPL NOTES :
// OnAfterAttack, if flag = 1, retarget max 1, add pursued attack.
// attack : type same as atker, indirect, pursued, can trigger last kill.
// OnBeforeSkillUse, if = basic atk/skill, set flag to 1.
// OnAfterSkillUse, define flag.

func init() {
	lightcone.Register(key.WeWillMeetAgain, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(meet, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: activateTrigger,
			OnAfterAttack:  addExtraDmgOnTrigger,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	extraDmgAmt := 0.36 + 0.12*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   meet,
		Source: owner,
		State:  extraDmgAmt,
	})
}

func activateTrigger(mod *modifier.Instance, e event.ActionStart) {

}

func addExtraDmgOnTrigger(mod *modifier.Instance, e event.AttackEnd) {

}
