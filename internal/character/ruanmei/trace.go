package ruanmei

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 = "ruanmei-a2"
	A4 = "ruanmei-a4"
	A6 = "ruanmei-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(A4, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: doA4,
		},
	})
	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnPropertyChange: checkA6,
		},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A4,
			Source: c.id,
		})
	}
}

func doA4(mod *modifier.Instance) {
	mod.Engine().ModifyEnergy(info.ModifyAttribute{
		Key:    A4,
		Target: mod.Source(),
		Source: mod.Source(),
		Amount: 5,
	})
}

func checkA6(mod *modifier.Instance) {
	// Calculate how much DMG Bonus should be given
	rm, _ := mod.Engine().CharacterInfo(mod.Owner())
	dmgAmt := skillDmg[rm.SkillLevelIndex()]
	if rm.Traces["103"] {
		dmgAmtA6 := 0.06 * float64(int((mod.OwnerStats().BreakEffect()-1.2)/0.1))
		if dmgAmtA6 > 0.36 {
			dmgAmtA6 = 0.36
		}
		dmgAmt += dmgAmtA6
	}

	// Reapply Overtone's DMG Bonus Buff to all allies with the new amount
	for _, trg := range mod.Engine().Characters() {
		if mod.Engine().HasModifier(trg, OvertoneDmgBuff) {
			mod.Engine().AddModifier(trg, info.Modifier{
				Name:   OvertoneDmgBuff,
				Source: mod.Owner(),
				Stats:  info.PropMap{prop.AllDamagePercent: dmgAmt},
			})
		}
	}
}
