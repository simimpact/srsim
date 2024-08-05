package huohuo

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2       key.Heal     = "huohuo-e2"
	E2Boost  key.Reason   = "huohuo-e2-heal-boost"
	E2Insert key.Insert   = "huohuo-e2-insert"
	E6       key.Modifier = "huohuo-e6"
)

func init() {
	modifier.Register(E6, modifier.Config{
		Stacking:   modifier.Replace,
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func E2OnKill(mod *modifier.Instance) bool {
	tmp, _ := mod.Engine().CharacterInstance(mod.Source())
	c := tmp.(*char)
	if c.info.Eidolon < 2 {
		return false
	}
	if c.E2ReviveCount == 0 {
		return false
	}
	c.E2ReviveCount--
	c.TalentRound--
	if c.TalentRound == 0 {
		c.RemoveBuff()
	}
	c.engine.InsertAbility(info.Insert{
		Key: E2Insert,
		Execute: func() {
			c.engine.Heal(info.Heal{
				Key:      E2,
				Targets:  []key.TargetID{mod.Owner()},
				Source:   c.id,
				BaseHeal: info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: 0.5},
			})
		},
		Source:     c.id,
		AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
		Priority:   info.CharReviveOthers,
	})
	return true
}

func (c *char) E4OnHeal(e *event.HealStart) {
	if c.info.Eidolon < 4 {
		return
	}
	e.Healer.AddProperty(E2Boost, prop.HealBoost, 0.8*(1-e.Target.CurrentHPRatio()))
}

func (c *char) E6OnHeal(e *event.HealStart) {
	if c.info.Eidolon < 6 {
		return
	}
	if e.Healer.ID() != c.id {
		return
	}
	c.engine.AddModifier(e.Target.ID(), info.Modifier{
		Source: c.id,
		Name:   E6,
		Stats:  info.PropMap{prop.AllDamagePercent: 0.5},
	})
}
