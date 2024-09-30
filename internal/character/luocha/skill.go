package luocha

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill               = "luocha-skill"
	InsertSkill         = "luocha-insert-skill"
	InsertSkillCD       = "luocha-insert-skill-cooldown"
	InsertSkillRetarget = "luocha-insert-skill-retarget"
	InsertSkillMark     = "luocha-insert-skill-mark"
)

// this flag might be a global value, not sure what else it does
var IsInsert bool

func init() {
	modifier.Register(InsertSkill, modifier.Config{})

	modifier.Register(InsertSkillCD, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
	})

	modifier.Register(InsertSkillRetarget, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: applyInsertSkillMark,
		},
	})

	modifier.Register(InsertSkillMark, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: doInsertSkill,
			OnBeforeDying: func(mod *modifier.Instance) {
				mod.Engine().RemoveModifier(mod.Source(), InsertSkillRetarget)
			},
		},
	})
}

func (c *char) initInsertSkill() {
	c.engine.Events().HPChange.Subscribe(c.InsertSkillListener)
	c.engine.Events().InsertEnd.Subscribe(c.onInsertFinish)
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// do A2
	if c.info.Traces["102"] {
		c.engine.DispelStatus(target, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}

	// do E2
	if c.info.Eidolon >= 2 {
		if c.engine.HPRatio(target) < 0.5 {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   E2HealBoost,
				Source: c.id,
			})
		} else {
			c.engine.AddModifier(target, info.Modifier{
				Name:   E2Shield,
				Source: c.id,
			})
		}
	}

	// heal target
	c.engine.Heal(info.Heal{
		Key:     Skill,
		Targets: []key.TargetID{target},
		Source:  c.id,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_ATK: skillPer[c.info.SkillLevelIndex()],
		},
		HealValue: skillFlat[c.info.SkillLevelIndex()],
	})

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Skill,
		Target: c.id,
		Source: c.id,
		Amount: 30,
	})

	// add 1 stack of Abyss Flower if no Field active
	if !c.engine.HasModifier(c.id, Field) {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   AbyssFlower,
			Source: c.id,
		})
	}

	// reset insert flag
	if IsInsert {
		IsInsert = false
	}
}

func (c *char) InsertSkillListener(e event.HPChange) {
	// bypass if hp change is positive
	if e.NewHP > e.OldHP {
		return
	}

	// bypass if on cooldown
	if c.engine.HasModifier(c.id, InsertSkillCD) {
		return
	}

	// bypass if CC'd or unable to act
	if c.engine.HasBehaviorFlag(c.id, model.BehaviorFlag_STAT_CTRL) {
		return
	}
	if c.engine.HasBehaviorFlag(c.id, model.BehaviorFlag_DISABLE_ACTION) {
		return
	}

	trg := c.engine.Retarget(info.Retarget{
		Targets: c.engine.Characters(),
		Filter: func(target key.TargetID) bool {
			// Filter conditions as bypasses
			// bypass if HP at or below 0
			if c.engine.HPRatio(target) <= 0 {
				return false
			}
			// bypass if BattleEventEntity (using workaround)
			if !c.engine.IsCharacter(target) {
				return false
			}
			// bypass if not lowest HP ratio
			if !isMinHPRatio(c.engine, target, c.engine.Characters()) {
				return false
			}
			return true
		},
		Max: 1,
	})

	if c.engine.HPRatio(trg[0]) <= 0.5 {
		// apply another mod that applies another mod...
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   InsertSkillRetarget,
			Source: c.id,
		})
	}
}

func applyInsertSkillMark(mod *modifier.Instance) {
	trg := doRetarget(mod)

	mod.Engine().AddModifier(trg[0], info.Modifier{
		Name:   InsertSkillMark,
		Source: mod.Source(),
	})
}

func doInsertSkill(mod *modifier.Instance) {
	mod.Engine().InsertAbility(info.Insert{
		Key: InsertSkill,
		Execute: func() {
			IsInsert = true
			trg := doRetarget(mod)
			if mod.Engine().HPRatio(trg[0]) <= 0.5 {
				// apply cooldown mod
				mod.Engine().AddModifier(mod.Source(), info.Modifier{
					Name:     InsertSkillCD,
					Source:   mod.Source(),
					Duration: 2,
				})
				// remove mark mod on all allies and retarget mod on source
				for _, marktrg := range mod.Engine().Characters() {
					mod.Engine().RemoveModifier(marktrg, InsertSkillMark)
				}
				mod.Engine().RemoveModifier(mod.Source(), InsertSkillRetarget)
				// do skill as insert

			} else {
				// remove mark mod on all allies and retarget mod on source
				for _, marktrg := range mod.Engine().Characters() {
					mod.Engine().RemoveModifier(marktrg, InsertSkillMark)
				}
				mod.Engine().RemoveModifier(mod.Source(), InsertSkillRetarget)
				IsInsert = false
			}
		},
		Source:     mod.Source(),
		AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
		Priority:   info.CharHealOthers,
	})
}

func (c *char) onInsertFinish(e event.InsertEnd) {
	// check if CC'd or unable to act
	cond1 := c.engine.HasBehaviorFlag(c.id, model.BehaviorFlag_STAT_CTRL)
	cond2 := c.engine.HasBehaviorFlag(c.id, model.BehaviorFlag_DISABLE_ACTION)
	if cond1 || cond2 {
		// remove mark mod and retarget mod from all allies
		for _, trg := range c.engine.Characters() {
			c.engine.RemoveModifier(trg, InsertSkillMark)
			c.engine.RemoveModifier(trg, InsertSkillRetarget)
		}
		IsInsert = false
		if c.engine.HasModifier(c.id, TalentInsertMark) {
			c.engine.RemoveModifier(c.id, TalentInsertMark)
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   DisableTalentInsertMark,
				Source: c.id,
			})
		}
	}
}

// helper function for evaluating whether target is lowest HP ratio out of a list comparetargets
func isMinHPRatio(engine engine.Engine, target key.TargetID, comparetargets []key.TargetID) bool {
	// bypass if empty list
	if len(comparetargets) == 0 {
		return false
	}

	// Start by assuming the first target has the minimum HP ratio
	minTarget := comparetargets[0]
	minHPRatio := engine.HPRatio(minTarget)

	// Loop through all the targets to find the one with the smallest HP ratio
	for _, target := range comparetargets {
		currentHPRatio := engine.HPRatio(target)
		if currentHPRatio < minHPRatio {
			minHPRatio = currentHPRatio
			minTarget = target
		}
	}

	// Compare target's HPRatio with the minimum found
	return target == minTarget
}

// helper function for small retarget (without BattleEvenEntity)
func doRetarget(mod *modifier.Instance) []key.TargetID {
	trg := mod.Engine().Retarget(info.Retarget{
		Targets: mod.Engine().Characters(),
		Filter: func(target key.TargetID) bool {
			// Filter conditions as bypasses
			// bypass if HP at or below 0
			if mod.Engine().HPRatio(target) <= 0 {
				return false
			}
			// no check for BattleEventEntity
			// bypass if not lowest HP ratio
			if !isMinHPRatio(mod.Engine(), target, mod.Engine().Characters()) {
				return false
			}
			return true
		},
		Max: 1,
	})

	return trg
}
