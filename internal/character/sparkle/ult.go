package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Cipher = "sparkle-cipher"
	Ult    = "sparkle-ult"
)

func init() {
	modifier.Register(Cipher, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	spRecover := 4
	if c.info.Eidolon >= 4 {
		spRecover += 1
	}

	c.engine.ModifySP(info.ModifySP{
		Source: c.id,
		Key:    Ult,
		Amount: spRecover,
	})

	buffdur := 2
	atkBuff := 0.0
	if c.info.Eidolon >= 1 {
		buffdur += 1
		atkBuff = 0.4
	}
	for _, char := range c.engine.Characters() {
		c.engine.AddModifier(char, info.Modifier{
			Name:     Cipher,
			Source:   c.id,
			Duration: buffdur,
			Stats: info.PropMap{
				prop.ATKPercent: atkBuff,
			},
		})
	}

	if c.info.Eidolon >= 6 {
		alliesWithCritBuff := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Characters(),
			Filter: func(target key.TargetID) bool {
				return c.engine.HasModifier(target, Dreamdiver) || c.engine.HasModifier(target, DreamdiverExtension)
			},
			Max:          1,
			IncludeLimbo: false,
		})
		if len(alliesWithCritBuff) > 0 {
			c.spreadBuff()
		}
	}

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Ult,
		Source: c.id,
		Target: c.id,
		Amount: 5,
	})
}

func (c *char) spreadBuff() {
	sparkle := c.engine.Stats(c.id)
	sparkleCdmg := sparkle.GetProperty(prop.CritDMG)
	// We check for e6 above
	proportion := skillCdmgScaling[c.info.SkillLevelIndex()] + 0.3

	for _, char := range c.engine.Characters() {
		if c.engine.HasModifier(char, Dreamdiver) {
			c.engine.AddModifier(char, info.Modifier{
				Name:     Dreamdiver,
				Source:   c.id,
				Duration: 1,
			})
		} else if c.engine.HasModifier(char, DreamdiverExtension) {
			c.engine.AddModifier(char, info.Modifier{
				Name:     DreamdiverExtension,
				Source:   c.id,
				Duration: 1,
				Stats: info.PropMap{
					prop.CritDMG: sparkleCdmg * proportion,
				},
			})
		}
	}
}
