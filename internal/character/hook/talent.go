package hook

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent = "hook-talent"
	A2     = "hook-a2"
)

func init() {

}

// Main entry for proccing talent normally; Use if only targetting single enemy with attack
func (c *char) talentProc(target key.TargetID) {
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Talent,
		Target: c.id,
		Source: c.id,
		Amount: 5,
	})

	c.talentPursuedDamage(target)

	if c.info.Eidolon >= 4 {
		c.applySkillBurn(c.engine.AdjacentTo(target))
	}

	c.talentHeal()
}

func (c *char) talentPursuedDamage(target key.TargetID) {
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

	c.engine.EndAttack()
}

func (c *char) talentHeal() {
	if c.info.Traces["101"] {
		c.engine.Heal(info.Heal{
			Key:     A2,
			Targets: []key.TargetID{c.id},
			Source:  c.id,
			BaseHeal: info.HealMap{
				model.HealFormula_BY_HEALER_MAX_HP: 0.05,
			},
		})
	}
}
