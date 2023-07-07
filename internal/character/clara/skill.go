package clara

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
    c.engine.Attack(info.Attack{
        Source: c.id,
        Targets: c.engine.Enemies(),
        DamageType: model.DamageType_PHYSICAL,
        AttackType: model.AttackType_SKILL,
        BaseDamage: info.DamageMap{

        },
		StanceDamage: 30.0,
		EnergyGain:   30,
    })
    // TODO: remove modifier
    // onStart check contains revenge modifier
    // MAvatar_Klara_00_BPSkill_Revenge
    state.EndAttack()
}
