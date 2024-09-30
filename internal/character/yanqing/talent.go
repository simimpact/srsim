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
		Stacking:   modifier.ReplaceBySource,
		Duration:   1,
		Listeners: modifier.Listeners{
			OnAfterBeingHitAll: OnHitRemove,
			OnTriggerDeath:     E6Listener,
		},
	})
}

func (c *char) getTalent(ult bool) info.Modifier {
	mod := info.Modifier{
		Name:   Soulsteel,
		Source: c.id,
		Stats: info.PropMap{
			prop.CritChance: talentCritRate[c.info.TalentLevelIndex()],
			prop.CritDMG:    talentCritDmg[c.info.TalentLevelIndex()],
		},
		TickImmediately: ult,
	}
	if c.info.Traces["102"] {
		mod.DebuffRES = info.DebuffRESMap{model.BehaviorFlag_STAT_CTRL: 0.2}
	}
	if c.info.Eidolon >= 2 {
		mod.Stats.Set(prop.EnergyRegen, 0.1)
	}
	if ult {
		mod.Stats.Modify(prop.CritDMG, ultCritDmg[c.info.UltLevelIndex()])
	}
	return mod
}

var followHits = []float64{0.3, 0.7}

func (c *char) tryFollow(target key.TargetID) {
	if c.engine.HasModifier(c.id, Soulsteel) && c.engine.Rand().Float64() <= talentFollowChance[c.info.TalentLevelIndex()] {
		c.engine.InsertAbility(info.Insert{
			Execute: func() {
				if !c.engine.IsAlive(target) {
					target = c.engine.Retarget(info.Retarget{
						Targets:      c.engine.Enemies(),
						IncludeLimbo: false,
						Max:          1,
						Filter:       func(t key.TargetID) bool { return true },
					})[0]
				}
				for i, hitRatio := range followHits {
					c.engine.Attack(info.Attack{
						Key:          TalentAttack,
						HitIndex:     i,
						Source:       c.id,
						Targets:      []key.TargetID{target},
						DamageType:   model.DamageType_ICE,
						AttackType:   model.AttackType_INSERT,
						BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: talentFollowRate[c.info.TalentLevelIndex()]},
						StanceDamage: 30,
						EnergyGain:   10,
						HitRatio:     hitRatio,
					})
				}
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

func OnHitRemove(mod *modifier.Instance, e event.HitEnd) {
	if e.HPDamage > 0 {
		mod.RemoveSelf()
	}
}
