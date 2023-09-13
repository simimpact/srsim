package common

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/combat"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	WeaknessBreakPhysical  = "weakness-break-physical"
	WeaknessBreakIce       = "weakness-break-ice"
	WeaknessBreakFire      = "weakness-break-fire"
	WeaknessBreakWind      = "weakness-break-wind"
	WeaknessBreakThunder   = "weakness-break-thunder"
	WeaknessBreakImaginary = "weakness-break-imaginary"
	WeaknessBreakQuantum   = "weakness-break-quantum"
)

func dealWeaknessBreakDamage(engine engine.Engine, attackKey key.Attack, characterID, enemyID key.TargetID, damageType model.DamageType, maxStanceMultiplier, damageMultiplier float64) {
	engine.Attack(info.Attack{
		Key:        attackKey,
		Source:     characterID,
		Targets:    []key.TargetID{enemyID},
		AttackType: model.AttackType_ELEMENT_DAMAGE,
		DamageType: damageType,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: damageMultiplier * maxStanceMultiplier,
		},
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}

func modifyGaugeAfterWeaknessBreak(engine engine.Engine, reasonKey string, characterID, enemyID key.TargetID) {
	engine.ModifyGaugeNormalized(info.ModifyAttribute{
		Key:    key.Reason(reasonKey),
		Target: enemyID,
		Source: characterID,
		Amount: 0.25,
	})
}

func applyWeaknessBreakPhysical(engine engine.Engine, characterID, enemyID key.TargetID, characterInfo info.Character, maxStanceMultiplier float64) {
	dealWeaknessBreakDamage(engine, WeaknessBreakPhysical, characterID, enemyID, characterInfo.Element, maxStanceMultiplier, 2)
	modifyGaugeAfterWeaknessBreak(engine, WeaknessBreakPhysical, characterID, enemyID)

	enemyMaxHP := engine.Stats(enemyID).MaxHP()
	maxBleedBaseDmg := 2 * maxStanceMultiplier * combat.BreakBaseDamage[characterInfo.Level]
	var bleedBaseDmg float64

	if /* enemyInfo.IsElite */ true {
		bleedBaseDmg = 0.07 * enemyMaxHP
	} else {
		bleedBaseDmg = 0.16 * enemyMaxHP
	}

	if bleedBaseDmg > maxBleedBaseDmg {
		bleedBaseDmg = maxBleedBaseDmg
	}

	engine.AddModifier(enemyID, info.Modifier{
		Name:   BreakBleed,
		Source: characterID,
		State: &BreakBleedState{
			BaseDamageValue: bleedBaseDmg,
		},
		Duration: 2,
		Chance:   1.5,
	})
}

func applyWeaknessBreakFire(engine engine.Engine, characterID, enemyID key.TargetID, characterInfo info.Character, maxStanceMultiplier float64) {
	dealWeaknessBreakDamage(engine, WeaknessBreakFire, characterID, enemyID, characterInfo.Element, maxStanceMultiplier, 1)
	modifyGaugeAfterWeaknessBreak(engine, WeaknessBreakFire, characterID, enemyID)

	engine.AddModifier(enemyID, info.Modifier{
		Name:     BreakBurn,
		Source:   characterID,
		Duration: 2,
		Chance:   1.5,
	})
}

func applyWeaknessBreakIce(engine engine.Engine, characterID, enemyID key.TargetID, characterInfo info.Character, maxStanceMultiplier float64) {
	dealWeaknessBreakDamage(engine, WeaknessBreakIce, characterID, enemyID, characterInfo.Element, maxStanceMultiplier, 1)
	modifyGaugeAfterWeaknessBreak(engine, WeaknessBreakIce, characterID, enemyID)

	engine.AddModifier(enemyID, info.Modifier{
		Name:     BreakFreeze,
		Source:   characterID,
		Duration: 1,
		Chance:   1.5,
	})
}

func applyWeaknessBreakThunder(engine engine.Engine, characterID, enemyID key.TargetID, characterInfo info.Character, maxStanceMultiplier float64) {
	dealWeaknessBreakDamage(engine, WeaknessBreakThunder, characterID, enemyID, characterInfo.Element, maxStanceMultiplier, 1)
	modifyGaugeAfterWeaknessBreak(engine, WeaknessBreakThunder, characterID, enemyID)

	engine.AddModifier(enemyID, info.Modifier{
		Name:     BreakShock,
		Source:   characterID,
		Duration: 2,
		Chance:   1.5,
	})
}

func applyWeaknessBreakWind(engine engine.Engine, characterID, enemyID key.TargetID, characterInfo info.Character, maxStanceMultiplier float64) {
	dealWeaknessBreakDamage(engine, WeaknessBreakWind, characterID, enemyID, characterInfo.Element, maxStanceMultiplier, 1.5)
	modifyGaugeAfterWeaknessBreak(engine, WeaknessBreakWind, characterID, enemyID)

	var windShearStacksCount float64

	if /* enemyInfo.IsElite */ true {
		windShearStacksCount = 3
	} else {
		windShearStacksCount = 1
	}

	engine.AddModifier(enemyID, info.Modifier{
		Name:     BreakWindShear,
		Source:   characterID,
		Duration: 2,
		Count:    windShearStacksCount,
		Chance:   1.5,
	})
}

func applyWeaknessBreakQuantum(engine engine.Engine, characterID, enemyID key.TargetID, characterInfo info.Character, maxStanceMultiplier float64) {
	dealWeaknessBreakDamage(engine, WeaknessBreakQuantum, characterID, enemyID, characterInfo.Element, maxStanceMultiplier, 0.5)
	modifyGaugeAfterWeaknessBreak(engine, WeaknessBreakQuantum, characterID, enemyID)

	engine.AddModifier(enemyID, info.Modifier{
		Name:   BreakEntanglement,
		Source: characterID,
		State: &BreakEntanglementState{
			HitsTakenCount:            0,
			TargetMaxStanceMultiplier: maxStanceMultiplier,
		},
		Duration: 1,
		Chance:   1.5,
	})
}

func applyWeaknessBreakImaginary(engine engine.Engine, characterID, enemyID key.TargetID, characterInfo info.Character, maxStanceMultiplier float64) {
	dealWeaknessBreakDamage(engine, WeaknessBreakImaginary, characterID, enemyID, characterInfo.Element, maxStanceMultiplier, 0.5)
	modifyGaugeAfterWeaknessBreak(engine, WeaknessBreakImaginary, characterID, enemyID)

	engine.AddModifier(enemyID, info.Modifier{
		Name:   BreakImprisonment,
		Source: characterID,
		State: &BreakImprisonState{
			DelayRatio: 0.2 * (1 + engine.Stats(characterID).BreakEffect()),
		},
		Duration: 1,
		Chance:   1.5,
	})
}

func ApplyWeaknessBreakEffects(engine engine.Engine, characterID, enemyID key.TargetID) {
	characterInfo, err := engine.CharacterInfo(characterID)

	if err != nil {
		panic("incorrect target id for character")
	}

	enemyInfo, err := engine.EnemyInfo(enemyID)

	if err != nil {
		panic("incorrect target id for enemy")
	}

	damageType := characterInfo.Element
	maxStanceMultiplier := 0.5 + enemyInfo.MaxStance/120

	switch damageType {
	case model.DamageType_PHYSICAL:
		applyWeaknessBreakPhysical(engine, characterID, enemyID, characterInfo, maxStanceMultiplier)
	case model.DamageType_FIRE:
		applyWeaknessBreakFire(engine, characterID, enemyID, characterInfo, maxStanceMultiplier)
	case model.DamageType_ICE:
		applyWeaknessBreakIce(engine, characterID, enemyID, characterInfo, maxStanceMultiplier)
	case model.DamageType_THUNDER:
		applyWeaknessBreakThunder(engine, characterID, enemyID, characterInfo, maxStanceMultiplier)
	case model.DamageType_WIND:
		applyWeaknessBreakWind(engine, characterID, enemyID, characterInfo, maxStanceMultiplier)
	case model.DamageType_QUANTUM:
		applyWeaknessBreakQuantum(engine, characterID, enemyID, characterInfo, maxStanceMultiplier)
	case model.DamageType_IMAGINARY:
		applyWeaknessBreakImaginary(engine, characterID, enemyID, characterInfo, maxStanceMultiplier)
	}
}
