package sampo

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	WindTornDagger key.Modifier = "sampo_talent"
)

func init() {
	modifier.Register(WindTornDagger, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHit: onAfterHit,
		},
	})
}

// Sampo's attacks have a 65% base chance to inflict Wind Shear for 3 turn(s).
// Enemies inflicted with Wind Shear will take Wind DoT equal to 20% of Sampo's ATK at the beginning of each turn. Wind Shear can stack up to 5 time(s).
// Tree01 add 1 duration
// Rank06 adds DamagePercentegeAdd
func onAfterHit(mod *modifier.Instance, e event.HitEnd) {
	char, _ := mod.Engine().CharacterInfo(e.Attacker)
	duration := 3
	AddWindShearTalent(char, mod.Engine(), e.Attacker, e.Defender, duration, 0.65)
}

func AddWindShearTalent(char info.Character, engine engine.Engine, owner, target key.TargetID, duration int, chance float64) {
	if char.Traces["101"] {
		duration += 1
	}

	damagePercentage := talent[char.TalentLevelIndex()]
	if char.Eidolon >= 6 {
		//+15% dmg multiplier
		damagePercentage += 0.15
	}

	engine.AddModifier(target, info.Modifier{
		Name:   common.WindShear,
		Source: owner,
		State: common.WindShearState{
			DamagePercentage: damagePercentage,
		},
		Chance:   chance,
		Duration: duration,
		MaxCount: 5,
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   WindTornDagger,
		Source: c.id,
	})
}
