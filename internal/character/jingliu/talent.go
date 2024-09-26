package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TalentBuff                 = "jingliu-talent-buff"
	TalentPull                 = "jingliu-talent-pull"
	EnhanceModeBuff            = "jingliu-enhance-mode-buff"
	TalentHPChange             = "jingliu-talent-hp-change"
	Talent          key.Reason = "jingliu-talent"
)

func init() {
	modifier.Register(EnhanceModeBuff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(TalentBuff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAfterAction: removeWhenEndAttack,
		},
	})
}

func (c *char) getMaxStack() int {
	if c.info.Eidolon >= 6 {
		return 4
	}
	return 3
}

func (c *char) gainSyzygy() {
	c.Syzygy += 1
	if c.isEnhanced {
		if c.Syzygy > c.getMaxStack() {
			c.Syzygy = c.getMaxStack()
		}
		return
	}
	if c.Syzygy < 2 {
		return
	}
	// c.Syzygy >= 2 && !c.isEnhanced enter EnhanceMode
	c.engine.InsertAbility(info.Insert{
		Key:      TalentPull,
		Source:   c.id,
		Priority: info.CharInsertAction,
		Execute: func() {
			c.isEnhanced = true
			c.Syzygy = c.getMaxStack() - 1
			c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
				Key:    Talent,
				Target: c.id,
				Source: c.id,
				Amount: -1,
			})

			mod := info.Modifier{
				Name:   EnhanceModeBuff,
				Source: c.id,
				Stats:  info.PropMap{prop.CritChance: talentCritRate[c.info.TalentLevelIndex()]},
			}
			if c.info.Eidolon >= 6 {
				mod.Stats.Set(prop.CritDMG, 0.5)
			}
			if c.info.Traces["101"] {
				mod.DebuffRES = info.DebuffRESMap{model.BehaviorFlag_STAT_CTRL: 0.35}
			}
			c.engine.AddModifier(c.id, mod)
		},
		AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
	})
}

func (c *char) addTalentBuff() {
	// calc hp pool
	hpCount := 0.0
	for _, ally := range c.engine.Characters() {
		if ally == c.id {
			continue
		}
		if c.engine.HPRatio(ally) < 0.04 {
			hpCount += c.engine.Stats(ally).CurrentHP() - 1
		} else {
			hpCount += c.engine.Stats(ally).MaxHP() * 0.04
		}
		c.engine.ModifyHPByRatio(info.ModifyHPByRatio{
			Key:       TalentHPChange,
			Target:    ally,
			Source:    c.id,
			Ratio:     -0.04,
			RatioType: model.ModifyHPRatioType_MAX_HP,
			Floor:     1.0,
		})
	}

	changeRate := 5.4
	maxRate := talentMaxAtkRate[c.info.TalentLevelIndex()]
	if c.info.Eidolon >= 4 {
		changeRate += 0.9
		maxRate += 0.3
	}
	atkValue := hpCount * changeRate
	maxValue := c.engine.Stats(c.id).GetProperty(prop.ATKBase) * maxRate
	if atkValue > maxValue {
		atkValue = maxValue
	}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:   TalentBuff,
		Source: c.id,
		Stats:  info.PropMap{prop.ATKConvert: atkValue},
	})
}

func removeWhenEndAttack(mod *modifier.Instance, e event.ActionEnd) {
	mod.RemoveSelf()
}

// NOTICE: when use last stack of sygyzy and insert an ult in same round, enhance won't quit and ult can affect by trace 103
func (c *char) checkSyzygy(e event.TurnEnd) {
	if c.Syzygy == 0 {
		c.isEnhanced = false
		c.engine.RemoveModifier(c.id, EnhanceModeBuff)
	}
}
