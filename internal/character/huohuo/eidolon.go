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
	E2 key.Heal     = "huohuo-e2"
	E6 key.Modifier = "huohuo-e6"
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
	if c.E2Count == 2 {
		return false
	}
	c.E2Count++
	c.TalentRound--
	if c.TalentRound == 0 {
		c.RemoveBuff()
	}
	c.engine.Heal(info.Heal{
		Key:      E2,
		Targets:  []key.TargetID{mod.Owner()},
		Source:   c.id,
		BaseHeal: info.HealMap{model.HealFormula_BY_TARGET_MAX_HP: 0.5},
	})
	return true
}

//TODO: E4 max ratio? add to heal rate or finally multiply

func (c *char) E4OnHeal(e *event.HealStart) {

}

func (c *char) E6OnHeal(e *event.HealStart) {
	if c.info.Eidolon < 6 {
		return
	}
	if e.Healer.ID() != c.id {
		return
	}
	c.engine.AddModifier(e.Target.ID(), info.Modifier{
		Name:  E6,
		Stats: info.PropMap{prop.AllDamagePercent: 0.5},
	})
}
