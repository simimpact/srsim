package blade

import (
	"math"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	UltHPChange key.Reason = "blade-ult-hp-change"
	UltPrimary  key.Attack = "blade-ult-primary"
	UltAdjacent key.Attack = "blade-ult-adjacent"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.SetHP(info.ModifyAttribute{
		Key:    UltHPChange,
		Target: c.id,
		Source: c.id,
		Amount: c.engine.Stats(c.id).MaxHP() * 0.5,
	})

	hpTally := math.Min(c.hpLoss, 0.9*c.engine.Stats(c.id).MaxHP())

	// TODO: Seperate Tally?
	e1TallyMod := 0.0

	if c.info.Eidolon >= 1 {
		e1TallyMod = 1.5
	}

	// Primary Target
	c.engine.Attack(info.Attack{
		Key:        UltPrimary,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_WIND,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK:    ultSingleAtk[c.info.UltLevelIndex()],
			model.DamageFormula_BY_MAX_HP: ultSingleHP[c.info.UltLevelIndex()],
		},
		DamageValue:  hpTally * (ultSingleTally[c.info.UltLevelIndex()] + e1TallyMod),
		StanceDamage: 60.0,
		EnergyGain:   5.0,
	})

	// Adjacent Targets
	c.engine.Attack(info.Attack{
		Key:        UltAdjacent,
		Source:     c.id,
		Targets:    c.engine.AdjacentTo(target),
		DamageType: model.DamageType_WIND,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK:    ultBlastAtk[c.info.UltLevelIndex()],
			model.DamageFormula_BY_MAX_HP: ultBlastHP[c.info.UltLevelIndex()],
		},
		DamageValue:  hpTally * ultBlastTally[c.info.UltLevelIndex()],
		StanceDamage: 60.0,
	})

	c.hpLoss = 0
}
