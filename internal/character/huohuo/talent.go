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
)

func init() {
	modifier.Register(TalentBuff, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnPhase1: RunTalent,
			OnBeforeAction: func(mod *modifier.Instance, e event.ActionStart) {
				if e.AttackType == model.AttackType_ULT {
					RunTalent(mod)
				}
			},
		},
	})
}

func (c *char) TalentHeal(target key.TargetID) {
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

func (c *char) TalentActionStartListener(e event.ActionStart) {
	if c.TalentRound == 0 {
		return
	}
	c.TalentRound--
	if c.TalentRound == 0 {
		targets := c.engine.Characters()
		for _, target := range targets {
			c.engine.RemoveModifier(target, TalentBuff)
		}
	}
}
