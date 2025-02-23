package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Dreamdiver          = "dreamdiver"
	DreamdiverExtension = "dreamdiver-extension"
	SparkleSkill        = "sparkle-skill"
)

func init() {
	modifier.Register(Dreamdiver, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		CanDispel:  true,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd:    addActualBuff,
			OnRemove: A4Extend,
		},
		Duration: 1,
	})

	modifier.Register(DreamdiverExtension, modifier.Config{
		Stacking:   modifier.Replace,
		CanDispel:  true,
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   1,
		TickMoment: modifier.ModifierPhase1End,
	})
}

type SkillBuffState struct {
	cdmgBuff float64
}

// TODO: Adjust skill crit scaling to use Base/Convert versions of crit buffs when appropriate. (once implemented)
func (c *char) Skill(target key.TargetID, state info.ActionState) {
	if target != c.id {
		c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    SparkleSkill,
			Target: target,
			Source: c.id,
			Amount: -0.5,
		})
	}

	sparkle := c.engine.Stats(c.id)
	sparkleCdmg := sparkle.GetProperty(prop.CritDMG)
	proportion := skillCdmgScaling[c.info.SkillLevelIndex()]
	if c.info.Eidolon >= 6 {
		proportion += 0.3
	}
	c.engine.AddModifier(target, info.Modifier{
		Name:     Dreamdiver,
		Source:   c.id,
		Duration: 1,
		State: SkillBuffState{
			cdmgBuff: proportion*sparkleCdmg + skillFlatCdmg[c.info.SkillLevelIndex()],
		},
	})

	// At e6, when using skill sparkle should add skill buff to all teammates with cipher
	if c.info.Eidolon >= 6 {
		targets := make([]key.TargetID, 0, 4)
		for _, char := range c.engine.Characters() {
			if c.engine.HasModifier(char, Cipher) {
				targets = append(targets, char)
			}
		}

		for _, char := range targets {
			c.engine.AddModifier(char, info.Modifier{
				Name:     Dreamdiver,
				Source:   c.id,
				Duration: 1,
				State: SkillBuffState{
					cdmgBuff: proportion*sparkleCdmg + skillFlatCdmg[c.info.SkillLevelIndex()],
				},
			})
		}
	}

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    SparkleSkill,
		Source: c.id,
		Target: c.id,
		Amount: 30,
	})
}

func addActualBuff(mod *modifier.Instance) {
	mod.Engine().RemoveModifier(mod.Owner(), DreamdiverExtension)
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   DreamdiverExtension,
		Source: mod.Source(),
		Stats: info.PropMap{
			prop.CritDMG: mod.State().(SkillBuffState).cdmgBuff,
		},
	})
}

func A4Extend(mod *modifier.Instance) {
	sparkleinfo, _ := mod.Engine().CharacterInfo(mod.Source())
	if sparkleinfo.Traces["102"] {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   DreamdiverExtension,
			Source: mod.Source(),
			Stats: info.PropMap{
				prop.CritDMG: mod.State().(SkillBuffState).cdmgBuff,
			},
		})
	}
}
