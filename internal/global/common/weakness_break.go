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

func dealWeaknessBreakDamage(engine engine.Engine, attackKey key.Attack, characterID, enemyID key.TargetID, damageType model.DamageType, damageMultiplier float64) {
	engine.Attack(info.Attack{
		Key:        attackKey,
		Source:     characterID,
		Targets:    []key.TargetID{enemyID},
		AttackType: model.AttackType_ELEMENT_DAMAGE,
		DamageType: damageType,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: damageMultiplier,
		},
		AsPureDamage: true,
		UseSnapshot:  true,
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
		dealWeaknessBreakDamage(engine, WeaknessBreakPhysical, characterID, enemyID, damageType, 2*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakPhysical,
			Target: enemyID,
			Source: characterID,
			Amount: 0.25,
		})

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
			State: BreakBleedState{
				BaseDamageValue: bleedBaseDmg,
			},
			Duration: 2,
			Chance:   1.5,
		})
	case model.DamageType_FIRE:
		dealWeaknessBreakDamage(engine, WeaknessBreakFire, characterID, enemyID, damageType, 2*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakFire,
			Target: enemyID,
			Source: characterID,
			Amount: 0.25,
		})

		engine.AddModifier(enemyID, info.Modifier{
			Name:     BreakBurn,
			Source:   characterID,
			Duration: 2,
			Chance:   1.5,
		})
	case model.DamageType_ICE:
		dealWeaknessBreakDamage(engine, WeaknessBreakIce, characterID, enemyID, damageType, maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakIce,
			Target: enemyID,
			Source: characterID,
			Amount: 0.25,
		})

		engine.AddModifier(enemyID, info.Modifier{
			Name:     BreakFreeze,
			Source:   characterID,
			Duration: 1,
			Chance:   1.5,
		})
	case model.DamageType_THUNDER:
		dealWeaknessBreakDamage(engine, WeaknessBreakThunder, characterID, enemyID, damageType, maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakThunder,
			Target: enemyID,
			Source: characterID,
			Amount: 0.25,
		})

		engine.AddModifier(enemyID, info.Modifier{
			Name:     BreakShock,
			Source:   characterID,
			Duration: 2,
			Chance:   1.5,
		})
	case model.DamageType_WIND:
		dealWeaknessBreakDamage(engine, WeaknessBreakWind, characterID, enemyID, damageType, 1.5*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakWind,
			Target: enemyID,
			Source: characterID,
			Amount: 0.25,
		})

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
	case model.DamageType_QUANTUM:
		dealWeaknessBreakDamage(engine, WeaknessBreakQuantum, characterID, enemyID, damageType, 0.5*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakQuantum,
			Target: enemyID,
			Source: characterID,
			Amount: 0.25,
		})

		engine.AddModifier(enemyID, info.Modifier{
			Name:   BreakEntanglement,
			Source: characterID,
			State: BreakEntanglementState{
				DelayRatio:                0.2 * (1 + engine.Stats(characterID).BreakEffect()),
				HitsTakenCount:            0,
				TargetMaxStanceMultiplier: maxStanceMultiplier,
			},
			Duration: 1,
			Chance:   1.5,
		})
	case model.DamageType_IMAGINARY:
		dealWeaknessBreakDamage(engine, WeaknessBreakImaginary, characterID, enemyID, damageType, 0.5*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakImaginary,
			Target: enemyID,
			Source: characterID,
			Amount: 0.25,
		})

		engine.AddModifier(enemyID, info.Modifier{
			Name:   BreakImprisonment,
			Source: characterID,
			State: BreakImprisonState{
				DelayRatio: 0.2 * (1 + engine.Stats(characterID).BreakEffect()),
			},
			Duration: 1,
			Chance:   1.5,
		})
	}
}
