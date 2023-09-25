package hook

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent = "hook-talent"
)

func init() {

}

func (c *char) initTalent() {

}

// TODO: Split into 2 functions? 1 for actual proc and one for damage across indiscriminate amount of units?
func (c *char) talentProc(target key.TargetID) {
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Talent,
		Target: c.id,
		Source: c.id,
		Amount: 5,
	})

	//Actual talent damage
	c.engine.Attack(info.Attack{
		Key:        Talent,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_PURSUED,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: talent[c.info.TalentLevelIndex()],
		},
	})

	if c.info.Eidolon >= 4 {
		c.applySkillBurn(c.engine.AdjacentTo(target))
	}

	if c.info.Traces["101"] {
		c.engine.Heal(info.Heal{
			Key:     Talent,
			Targets: []key.TargetID{c.id},
			Source:  c.id,
			BaseHeal: info.HealMap{
				model.HealFormula_BY_HEALER_MAX_HP: 0.05,
			},
		})
	}

}
