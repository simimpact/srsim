package ruanmei

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SkillMain               = "ruanmei-skill"
	OvertoneDmgBuff         = "ruanmei-skill-overtone-dmg-buff"
	OvertoneBreakEfficiency = "ruanmei-skill-overtone-weakness-break-efficiency"
)

func init() {
	modifier.Register(SkillMain, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		TickMoment: modifier.ModifierPhase1End,
		Listeners: modifier.Listeners{
			OnAdd:    addOvertone,
			OnRemove: removeOvertone,
		},
	})
	modifier.Register(OvertoneDmgBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(OvertoneBreakEfficiency, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   SkillMain,
		Source: c.id,
	})
}

// Overtone is summarized from 4 "sub" mods with 2 being purely for display to only 2 "sub" mods
func addOvertone(mod *modifier.Instance) {
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
	for _, trg := range mod.Engine().Characters() {
		mod.Engine().AddModifier(trg, info.Modifier{
			Name:   OvertoneDmgBuff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllDamagePercent: dmgAmt},
		})
		mod.Engine().AddModifier(trg, info.Modifier{
			Name:   OvertoneBreakEfficiency,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllStanceDMGPercent: 0.5},
		})
	}
}

func removeOvertone(mod *modifier.Instance) {
	for _, trg := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(trg, OvertoneDmgBuff)
		mod.Engine().RemoveModifier(trg, OvertoneBreakEfficiency)
	}
}

// Technique doing Insert with autocast Skill (without consuming SP)
