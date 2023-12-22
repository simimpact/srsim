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

type breakInfo struct {
	engine              engine.Engine
	charID              key.TargetID
	charInfo            *info.Character
	enemyID             key.TargetID
	enemyRank           model.EnemyRank
	damageType          model.DamageType
	maxStanceMultiplier float64
}

func dealWeaknessBreakDamage(breakInfoData *breakInfo, attackKey key.Attack, damageMultiplier float64) {
	breakInfoData.engine.Attack(info.Attack{
		Key:        attackKey,
		Source:     breakInfoData.charID,
		Targets:    []key.TargetID{breakInfoData.enemyID},
		AttackType: model.AttackType_ELEMENT_DAMAGE,
		DamageType: breakInfoData.damageType,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: damageMultiplier * breakInfoData.maxStanceMultiplier,
		},
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}

func modifyGaugeAfterWeaknessBreak(engine engine.Engine, reasonKey string, charID, enemyID key.TargetID) {
	engine.ModifyGaugeNormalized(info.ModifyAttribute{
		Key:    key.Reason(reasonKey),
		Target: enemyID,
		Source: charID,
		Amount: 0.25,
	})
}

func applyWeaknessBreakPhysical(breakInfoData *breakInfo) {
	dealWeaknessBreakDamage(breakInfoData, WeaknessBreakPhysical, 2)
	modifyGaugeAfterWeaknessBreak(breakInfoData.engine, WeaknessBreakPhysical, breakInfoData.charID, breakInfoData.enemyID)

	enemyMaxHP := breakInfoData.engine.Stats(breakInfoData.enemyID).MaxHP()
	maxBleedBaseDmg := 2 * breakInfoData.maxStanceMultiplier * combat.BreakBaseDamage[breakInfoData.charInfo.Level]
	var bleedBaseDmg float64

	if breakInfoData.enemyRank.IsElite() {
		bleedBaseDmg = 0.07 * enemyMaxHP
	} else {
		bleedBaseDmg = 0.16 * enemyMaxHP
	}

	if bleedBaseDmg > maxBleedBaseDmg {
		bleedBaseDmg = maxBleedBaseDmg
	}

	breakInfoData.engine.AddModifier(breakInfoData.enemyID, info.Modifier{
		Name:   BreakBleed,
		Source: breakInfoData.charID,
		State: &BreakBleedState{
			BaseDamageValue: bleedBaseDmg,
		},
		Duration: 2,
		Chance:   1.5,
	})
}

func applyWeaknessBreakFire(breakInfoData *breakInfo) {
	dealWeaknessBreakDamage(breakInfoData, WeaknessBreakFire, 2)
	modifyGaugeAfterWeaknessBreak(breakInfoData.engine, WeaknessBreakFire, breakInfoData.charID, breakInfoData.enemyID)

	breakInfoData.engine.AddModifier(breakInfoData.enemyID, info.Modifier{
		Name:     BreakBurn,
		Source:   breakInfoData.charID,
		Duration: 2,
		Chance:   1.5,
		State: &BurnState{
			DamagePercentage:    1,
			DamageValue:         0,
			DEFDamagePercentage: 0,
		},
	})
}

func applyWeaknessBreakIce(breakInfoData *breakInfo) {
	dealWeaknessBreakDamage(breakInfoData, WeaknessBreakIce, 1)
	modifyGaugeAfterWeaknessBreak(breakInfoData.engine, WeaknessBreakIce, breakInfoData.charID, breakInfoData.enemyID)

	breakInfoData.engine.AddModifier(breakInfoData.enemyID, info.Modifier{
		Name:     BreakFreeze,
		Source:   breakInfoData.charID,
		Duration: 1,
		Chance:   1.5,
	})
}

func applyWeaknessBreakThunder(breakInfoData *breakInfo) {
	dealWeaknessBreakDamage(breakInfoData, WeaknessBreakThunder, 1)
	modifyGaugeAfterWeaknessBreak(breakInfoData.engine, WeaknessBreakThunder, breakInfoData.charID, breakInfoData.enemyID)

	breakInfoData.engine.AddModifier(breakInfoData.enemyID, info.Modifier{
		Name:     BreakShock,
		Source:   breakInfoData.charID,
		Duration: 2,
		Chance:   1.5,
	})
}

func applyWeaknessBreakWind(breakInfoData *breakInfo) {
	dealWeaknessBreakDamage(breakInfoData, WeaknessBreakWind, 1.5)
	modifyGaugeAfterWeaknessBreak(breakInfoData.engine, WeaknessBreakWind, breakInfoData.charID, breakInfoData.enemyID)

	var windShearStacksCount float64

	if breakInfoData.enemyRank.IsElite() {
		windShearStacksCount = 3
	} else {
		windShearStacksCount = 1
	}

	breakInfoData.engine.AddModifier(breakInfoData.enemyID, info.Modifier{
		Name:     BreakWindShear,
		Source:   breakInfoData.charID,
		Duration: 2,
		Count:    windShearStacksCount,
		Chance:   1.5,
	})
}

func applyWeaknessBreakQuantum(breakInfoData *breakInfo) {
	dealWeaknessBreakDamage(breakInfoData, WeaknessBreakQuantum, 0.5)
	modifyGaugeAfterWeaknessBreak(breakInfoData.engine, WeaknessBreakQuantum, breakInfoData.charID, breakInfoData.enemyID)

	breakInfoData.engine.AddModifier(breakInfoData.enemyID, info.Modifier{
		Name:   BreakEntanglement,
		Source: breakInfoData.charID,
		State: &BreakEntanglementState{
			HitsTakenCount:            0,
			TargetMaxStanceMultiplier: breakInfoData.maxStanceMultiplier,
		},
		Duration: 1,
		Chance:   1.5,
	})
}

func applyWeaknessBreakImaginary(breakInfoData *breakInfo) {
	dealWeaknessBreakDamage(breakInfoData, WeaknessBreakImaginary, 0.5)
	modifyGaugeAfterWeaknessBreak(breakInfoData.engine, WeaknessBreakImaginary, breakInfoData.charID, breakInfoData.enemyID)

	breakInfoData.engine.AddModifier(breakInfoData.enemyID, info.Modifier{
		Name:     BreakImprisonment,
		Source:   breakInfoData.charID,
		Duration: 1,
		Chance:   1.5,
	})
}

func ApplyWeaknessBreakEffects(engine engine.Engine, charID, enemyID key.TargetID) {
	charInfo, err := engine.CharacterInfo(charID)

	if err != nil {
		panic("incorrect target id for character")
	}

	enemyInfo, err := engine.EnemyInfo(enemyID)

	if err != nil {
		panic("incorrect target id for enemy")
	}

	breakInfoData := &breakInfo{
		engine:              engine,
		charID:              charID,
		charInfo:            &charInfo,
		enemyID:             enemyID,
		enemyRank:           enemyInfo.Rank,
		damageType:          charInfo.Element,
		maxStanceMultiplier: 0.5 + engine.MaxStance(enemyID)/120,
	}

	switch charInfo.Element {
	case model.DamageType_PHYSICAL:
		applyWeaknessBreakPhysical(breakInfoData)
	case model.DamageType_FIRE:
		applyWeaknessBreakFire(breakInfoData)
	case model.DamageType_ICE:
		applyWeaknessBreakIce(breakInfoData)
	case model.DamageType_THUNDER:
		applyWeaknessBreakThunder(breakInfoData)
	case model.DamageType_WIND:
		applyWeaknessBreakWind(breakInfoData)
	case model.DamageType_QUANTUM:
		applyWeaknessBreakQuantum(breakInfoData)
	case model.DamageType_IMAGINARY:
		applyWeaknessBreakImaginary(breakInfoData)
	}
}
