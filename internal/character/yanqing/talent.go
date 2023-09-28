package yanqing

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Soulsteel               = "yanqing-soulsteelsync"
	Talent                  = "yanqing-talent"
	TalentAttack key.Attack = "yanqing-talent-attack"
	TalentInsert key.Insert = "yanqing-talent-insert"
)

func init() {
	modifier.Register(Soulsteel, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.Refresh,
		Duration:   1,
		Listeners: modifier.Listeners{
			OnHPChange: talentOnHit,
		},
	})
}

func (c *char) addTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Soulsteel,
		Source: c.id,
		Stats: info.PropMap{
			prop.CritChance: talentCritRate[c.info.TalentLevelIndex()],
			prop.CritDMG:    talentCritDmg[c.info.TalentLevelIndex()],
		},
	})
}

func (c *char) tryFollow(target key.TargetID) {
	if c.engine.HasModifier(c.id, Soulsteel) && c.engine.Rand().Float64() <= talentFollowChance[c.info.TalentLevelIndex()] {
		c.engine.InsertAbility(info.Insert{
			Execute: func() {
				c.engine.Attack(info.Attack{
					Key:          TalentAttack,
					Source:       c.id,
					Targets:      []key.TargetID{target},
					DamageType:   model.DamageType_ICE,
					AttackType:   model.AttackType_INSERT,
					BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: talentFollowRate[c.info.TalentLevelIndex()]},
					StanceDamage: 30,
					EnergyGain:   10,
				})
				c.engine.AddModifier(target, info.Modifier{
					Name:   common.Freeze,
					Source: c.id,
					State: &common.FreezeState{
						DamagePercentage: talentIceDot[c.info.TalentLevelIndex()],
						DamageValue:      0,
					},
					Chance:   0.65,
					Duration: 1,
				})
			},
			Key:        TalentInsert,
			Source:     c.id,
			Priority:   info.CharInsertAction,
			AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
		})
	}
}

func talentOnHit(mod *modifier.Instance, e event.HPChange) {
	if e.IsHPChangeByDamage && e.NewHP < e.OldHP {
		mod.RemoveSelf()
	}
}
