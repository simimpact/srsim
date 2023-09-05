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

func dealWeaknessBreakDamage(engine engine.Engine, attackKey key.Attack, characterId, enemyId key.TargetID, damageType model.DamageType, damageMultiplier float64) {
	engine.Attack(info.Attack{
		Key:        attackKey,
		Source:     characterId,
		Targets:    []key.TargetID{enemyId},
		AttackType: model.AttackType_ELEMENT_DAMAGE,
		DamageType: damageType,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: damageMultiplier,
		},
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}

func ApplyWeaknessBreakEffects(engine engine.Engine, characterId, enemyId key.TargetID) {
	characterInfo, err := engine.CharacterInfo(characterId)

	if err != nil {
		panic("incorrect target id for character")
	}

	enemyInfo, err := engine.EnemyInfo(enemyId)

	if err != nil {
		panic("incorrect target id for enemy")
	}

	damageType := characterInfo.Element
	maxStanceMultiplier := 0.5 + enemyInfo.MaxStance/120

	switch damageType {
	case model.DamageType_PHYSICAL:
		dealWeaknessBreakDamage(engine, WeaknessBreakPhysical, characterId, enemyId, damageType, 2*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakPhysical,
			Target: enemyId,
			Source: characterId,
			Amount: 0.25,
		})

		enemyMaxHP := engine.Stats(enemyId).MaxHP()
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

		engine.AddModifier(enemyId, info.Modifier{
			Name:   BreakBleed,
			Source: characterId,
			State: BreakBleedState{
				BaseDamageValue: bleedBaseDmg,
			},
			Duration: 2,
			Chance:   1.5,
		})
	case model.DamageType_FIRE:
		dealWeaknessBreakDamage(engine, WeaknessBreakFire, characterId, enemyId, damageType, 2*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakFire,
			Target: enemyId,
			Source: characterId,
			Amount: 0.25,
		})

		engine.AddModifier(enemyId, info.Modifier{
			Name:     BreakBurn,
			Source:   characterId,
			Duration: 2,
			Chance:   1.5,
		})
	case model.DamageType_ICE:
		dealWeaknessBreakDamage(engine, WeaknessBreakIce, characterId, enemyId, damageType, maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakIce,
			Target: enemyId,
			Source: characterId,
			Amount: 0.25,
		})

		engine.AddModifier(enemyId, info.Modifier{
			Name:     BreakFreeze,
			Source:   characterId,
			Duration: 1,
			Chance:   1.5,
		})
	case model.DamageType_THUNDER:
		dealWeaknessBreakDamage(engine, WeaknessBreakThunder, characterId, enemyId, damageType, maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakThunder,
			Target: enemyId,
			Source: characterId,
			Amount: 0.25,
		})

		engine.AddModifier(enemyId, info.Modifier{
			Name:     BreakShock,
			Source:   characterId,
			Duration: 2,
			Chance:   1.5,
		})
	case model.DamageType_WIND:
		dealWeaknessBreakDamage(engine, WeaknessBreakWind, characterId, enemyId, damageType, 1.5*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakWind,
			Target: enemyId,
			Source: characterId,
			Amount: 0.25,
		})

		var windShearStacksCount float64

		if /* enemyInfo.IsElite */ true {
			windShearStacksCount = 3
		} else {
			windShearStacksCount = 1
		}

		engine.AddModifier(enemyId, info.Modifier{
			Name:     BreakWindShear,
			Source:   characterId,
			Duration: 2,
			Count:    windShearStacksCount,
			Chance:   1.5,
		})
	case model.DamageType_QUANTUM:
		dealWeaknessBreakDamage(engine, WeaknessBreakQuantum, characterId, enemyId, damageType, 0.5*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakQuantum,
			Target: enemyId,
			Source: characterId,
			Amount: 0.25,
		})

		engine.AddModifier(enemyId, info.Modifier{
			Name:   BreakEntanglement,
			Source: characterId,
			State: BreakEntanglementState{
				DelayRatio:                0.2 * (1 + engine.Stats(characterId).BreakEffect()),
				HitsTakenCount:            0,
				TargetMaxStanceMultiplier: maxStanceMultiplier,
			},
			Duration: 1,
			Chance:   1.5,
		})
	case model.DamageType_IMAGINARY:
		dealWeaknessBreakDamage(engine, WeaknessBreakImaginary, characterId, enemyId, damageType, 0.5*maxStanceMultiplier)

		engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    WeaknessBreakImaginary,
			Target: enemyId,
			Source: characterId,
			Amount: 0.25,
		})

		engine.AddModifier(enemyId, info.Modifier{
			Name:   BreakImprisonment,
			Source: characterId,
			State: BreakImprisonState{
				DelayRatio: 0.2 * (1 + engine.Stats(characterId).BreakEffect()),
			},
			Duration: 1,
			Chance:   1.5,
		})
	}
}
