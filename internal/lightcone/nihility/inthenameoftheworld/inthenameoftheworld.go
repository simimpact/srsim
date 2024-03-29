package inthenameoftheworld

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
	world                  = "in-the-name-of-the-world"
	skillBuff key.Modifier = "in-the-name-of-the-world-skill-buff"
)

type state struct {
	dmgNAtkAmt float64
	ehrAmt     float64
}

// Increases the wearer's DMG to debuffed enemies by 24%.
// When the wearer uses their Skill, the Effect Hit Rate for this attack increases by 18%,
// and ATK increases by 24%.

func init() {
	lightcone.Register(key.IntheNameoftheWorld, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(world, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: boostDmgOnDebuffed,
			OnBeforeAction: addSkillBuff,
			OnAfterAction:  removeSkillBuff,
		},
		CanModifySnapshot: true,
	})
	modifier.Register(skillBuff, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	modState := state{
		dmgNAtkAmt: 0.20 + 0.04*float64(lc.Imposition),
		ehrAmt:     0.15 + 0.03*float64(lc.Imposition),
	}
	engine.AddModifier(owner, info.Modifier{
		Name:   world,
		Source: owner,
		State:  &modState,
	})
}

// if debuff > 0, boost hit dmg
func boostDmgOnDebuffed(mod *modifier.Instance, e event.HitStart) {
	state := mod.State().(*state)
	// if hit enemy has debuff, amplify holder damage.
	if mod.Engine().Stats(e.Defender).StatusCount(model.StatusType_STATUS_DEBUFF) > 0 {
		e.Hit.Attacker.AddProperty(world, prop.AllDamagePercent, state.dmgNAtkAmt)
	}
}

// if action == skill, boost ehr + atk
func addSkillBuff(mod *modifier.Instance, e event.ActionStart) {
	state := mod.State().(*state)
	if e.AttackType == model.AttackType_SKILL {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   skillBuff,
			Source: mod.Owner(),
			Stats: info.PropMap{
				prop.EffectHitRate: state.ehrAmt,
				prop.ATKPercent:    state.dmgNAtkAmt,
			},
		})
	}
}

// remove skillBuff mod instances after action
func removeSkillBuff(mod *modifier.Instance, e event.ActionEnd) {
	mod.Engine().RemoveModifier(mod.Owner(), skillBuff)
}
