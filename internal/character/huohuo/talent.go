package huohuo

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TalentBuff key.Modifier = "huohuo-divineprovision"
	TalentHeal key.Heal     = "huohuo-divineprovision-heal"
	A4         key.Modifier = "huohuo-a4"
	A6         key.Reason   = "huohuo-a6"
)

func init() {
	modifier.Register(TalentBuff, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
		Listeners: modifier.Listeners{
			OnPhase1: RunTalent,
			OnBeforeAction: func(mod *modifier.Instance, e event.ActionStart) {
				if e.AttackType == model.AttackType_ULT {
					RunTalent(mod)
				}
			},
			OnLimboWaitHeal: E2OnKill,
		},
	})
}

func (c *char) TalentInit() {
	if c.info.Traces["101"] {
		c.DispelCount = 6
		c.TalentRound = 1
	}
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:      A4,
			Source:    c.id,
			DebuffRES: info.DebuffRESMap{model.BehaviorFlag_STAT_CTRL: 0.35},
		})
	}
	c.engine.Events().ActionStart.Subscribe(c.TalentActionStartListener)
	c.engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if c.id == e.Target {
			c.RemoveBuff()
		}
	})
}

func (c *char) TalentHeal(target key.TargetID) {
	if c.info.Traces["103"] {
		c.engine.ModifyEnergyFixed(info.ModifyAttribute{
			Target: c.id,
			Source: c.id,
			Amount: 1,
			Key:    A6,
		})
	}
	c.engine.Heal(info.Heal{
		Key:     TalentHeal,
		Targets: []key.TargetID{target},
		Source:  c.id,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: TalentRate[c.info.TalentLevelIndex()],
		},
		HealValue: TalentValue[c.info.TalentLevelIndex()],
	})
	if c.engine.ModifierStatusCount(target, model.StatusType_STATUS_DEBUFF) > 0 && c.DispelCount > 0 {
		c.DispelCount--
		c.engine.DispelStatus(target, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}
}

func RunTalent(mod *modifier.Instance) {
	tmp, _ := mod.Engine().CharacterInstance(mod.Source())
	c := tmp.(*char)
	if c.TalentRound == 0 {
		return
	}
	c.TalentHeal(mod.Owner())
	targets := c.engine.Characters()
	for _, target := range targets {
		if c.engine.HPRatio(target) <= 0.5 {
			c.TalentHeal(target)
		}
	}
}

func (c *char) RemoveBuff() {
	for _, target := range c.engine.Characters() {
		c.engine.RemoveModifier(target, TalentBuff)
	}
}

func (c *char) TalentActionStartListener(e event.ActionStart) {
	if c.TalentRound == 0 {
		return
	}
	c.TalentRound--
	if c.TalentRound == 0 {
		c.RemoveBuff()
	}
}
