package tingyun

import (
	"math"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	AtkBuff     = "tingyun-atk-buff"
	Benediction = "tingyun-benediction"
	ProcSkill   = "tingyun-benediction-proc-skill"
)

type skillState struct {
	tingEidolon       int
	e2flag            bool
	skillPursuedMltp  float64
	talentPursuedMltp float64
}

func init() {
	modifier.Register(AtkBuff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
		Listeners: modifier.Listeners{
			OnRemove: removeBenedictionSelf,
		},
	})

	modifier.Register(Benediction, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction:  doE1OnUlt,         // E1
			OnPhase1:       removeE2CD,        // E2 cooldown
			OnTriggerDeath: doE2,              // E2
			OnRemove:       removeAtkBuffSelf, // remove on destroy
			OnAfterAttack:  doProcSkill,       // Skill's Pursued
		},
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// check A2
	if c.info.Traces["101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
			Stats:  info.PropMap{prop.SPDPercent: 0.2},
		})
	}

	// calculate Atk Buff
	allyAtk := c.engine.Stats(target).GetProperty(prop.ATKBase)
	allyAtk *= skillAllyAtk[c.info.SkillLevelIndex()]
	tingAtk := c.engine.Stats(target).ATK()
	tingAtk *= skillTingAtk[c.info.SkillLevelIndex()]

	atkAmount := math.Min(allyAtk, tingAtk)

	// remove Atk Buff from all allies
	//// Assuming removing a modifier triggers OnDestroy:
	//// This Atk Buff modifier has an OnDestroy listener that removes Benediction on the Owner when triggered,
	//// which itself has an OnDestroy listener that removes the Atk Buff on the Owner.
	for _, trg := range c.engine.Characters() {
		c.engine.RemoveModifier(trg, AtkBuff)
	}

	// add new Atk Buff to target based on E1 (Spd Buff on ally Ult)
	c.engine.AddModifier(target, info.Modifier{
		Name:   AtkBuff,
		Source: c.id,
		Stats:  info.PropMap{prop.ATKFlat: atkAmount},
	})

	// remove Benediction from all allies
	//// Because of the removal earlier, there should be no Benediction on any ally at this point,
	//// so the OnDestroy cannot be triggered and the Atk Buff on the Owner won't be removed.
	//// Essentially, this is a redundant step.
	for _, trg := range c.engine.Characters() {
		c.engine.RemoveModifier(trg, Benediction)
	}
	// add Benediction, with effects based on Tingyun's Eidolon
	c.engine.AddModifier(target, info.Modifier{
		Name:   Benediction,
		Source: c.id,
		State: &skillState{
			tingEidolon:       c.info.Eidolon,
			e2flag:            false,
			skillPursuedMltp:  skillPursued[c.info.SkillLevelIndex()],
			talentPursuedMltp: talent[c.info.TalentLevelIndex()],
		},
	})

	c.engine.Events().AttackEnd.Subscribe(func(event event.AttackEnd) {
		c.engine.AddModifier(target, info.Modifier{
			Name:   TingyunNormalMonitor,
			Source: c.id,
		})
	})
}

func doProcSkill(mod *modifier.Instance, e event.AttackEnd) {
	st := mod.State().(*skillState)
	trg := mod.Engine().Retarget(info.Retarget{
		Targets: e.Targets,
		Max:     1,
	})
	mod.Engine().Attack(info.Attack{
		Key:        ProcSkill,
		Targets:    trg,
		Source:     mod.Owner(),
		AttackType: model.AttackType_PURSUED,
		DamageType: model.DamageType_THUNDER,
		BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: st.skillPursuedMltp},
	})
}

func removeBenedictionSelf(mod *modifier.Instance) {
	mod.Engine().RemoveModifier(mod.Owner(), Benediction)
}

func removeAtkBuffSelf(mod *modifier.Instance) {
	mod.Engine().RemoveModifier(mod.Owner(), AtkBuff)
}
