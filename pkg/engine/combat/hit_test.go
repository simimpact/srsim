package combat

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/tests/mock"
)

func TestPerformHitWithShield(t *testing.T) {
	pht := NewPerformHitTester(t)

	hit := &info.Hit{
		Key:          "tst",
		HitIndex:     0,
		Attacker:     mock.NewEmptyStats(1),
		Defender:     mock.NewEmptyStats(2),
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: 0.5},
		AttackType:   model.AttackType_NORMAL,
		DamageType:   model.DamageType_ICE,
		StanceDamage: 25,
		HitRatio:     1.0,
	}

	// POPULATE STATS
	hit.Attacker.AddProperty("tst", prop.ATKBase, 200)

	pht.AssertPerformHit(hit, &ExpectHit{
		TotalDamage:         100.0,
		BaseDamage:          100.0,
		ShieldAbsorb:        10.0,
		HPDamage:            90.0,
		IsAttackerChar:      true,
		DefenceMultiplier:   1.0,
		Resistance:          1.0,
		Vulnerability:       1.0,
		ToughnessMultiplier: 1.0,
		Fatigue:             1.0,
		AllDamageReduce:     1.0,
		CritDamage:          1.0,
		IsCrit:              false,
	})
}

func TestDamageValueNotModifiedByBonus(t *testing.T) {
	pht := NewPerformHitTester(t)

	hit := &info.Hit{
		Key:          "tst",
		HitIndex:     0,
		Attacker:     mock.NewEmptyStats(1),
		Defender:     mock.NewEmptyStats(2),
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: 0},
		DamageValue:  50,
		AttackType:   model.AttackType_NORMAL,
		DamageType:   model.DamageType_ICE,
		StanceDamage: 25,
		HitRatio:     1.0,
	}

	// POPULATE STATS
	hit.Attacker.AddProperty("tst", prop.ATKBase, 200)
	hit.Attacker.AddProperty("tst", prop.AllDamagePercent, 0.10)

	pht.AssertPerformHit(hit, &ExpectHit{
		TotalDamage:         50.0,
		HPDamage:            50.0,
		BaseDamage:          50.0,
		ShieldAbsorb:        0.0,
		IsAttackerChar:      true,
		DefenceMultiplier:   1.0,
		Resistance:          1.0,
		Vulnerability:       1.0,
		ToughnessMultiplier: 1.0,
		Fatigue:             1.0,
		AllDamageReduce:     1.0,
		CritDamage:          1.0,
		IsCrit:              false,
	})
}

func TestDamageBaseModifiedByBonus(t *testing.T) {
	pht := NewPerformHitTester(t)

	hit := &info.Hit{
		Key:          "tst",
		HitIndex:     0,
		Attacker:     mock.NewEmptyStats(1),
		Defender:     mock.NewEmptyStats(2),
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_DEF: 0.5},
		AttackType:   model.AttackType_NORMAL,
		DamageType:   model.DamageType_ICE,
		StanceDamage: 25,
		HitRatio:     1.0,
	}

	// POPULATE STATS
	hit.Attacker.AddProperty("tst", prop.DEFBase, 200)
	hit.Attacker.AddProperty("tst", prop.AllDamagePercent, 0.10)

	pht.AssertPerformHit(hit, &ExpectHit{
		TotalDamage:         110,
		HPDamage:            110,
		BaseDamage:          110,
		ShieldAbsorb:        0.0,
		IsAttackerChar:      true,
		DefenceMultiplier:   1.0,
		Resistance:          1.0,
		Vulnerability:       1.0,
		ToughnessMultiplier: 1.0,
		Fatigue:             1.0,
		AllDamageReduce:     1.0,
		CritDamage:          1.0,
		IsCrit:              false,
	})
}

func TestDamageBreakWithEffect(t *testing.T) {
	pht := NewPerformHitTester(t)

	hit := &info.Hit{
		Key:      "tst",
		HitIndex: 0,
		Attacker: mock.NewEmptyStatsWithAttr(1, &info.Attributes{
			Level:         50,
			BaseStats:     info.PropMap{},
			BaseDebuffRES: info.DebuffRESMap{},
			Weakness:      info.WeaknessMap{},
			HPRatio:       1.0,
			Energy:        50.0,
			MaxEnergy:     50.0,
			Stance:        0,
			MaxStance:     0,
		}),
		Defender:     mock.NewEmptyStats(2),
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_BREAK_DAMAGE: 1},
		AttackType:   model.AttackType_NORMAL,
		DamageType:   model.DamageType_ICE,
		StanceDamage: 25,
		HitRatio:     1.0,
		AsPureDamage: true,
	}

	hit.Attacker.AddProperty("tst", prop.BreakEffect, 0.5)

	// crits do nothing when pure damage
	hit.Attacker.AddProperty("tst", prop.CritChance, 1.0)
	hit.Attacker.AddProperty("tst", prop.CritDMG, 0.5)

	pht.AssertPerformHit(hit, &ExpectHit{
		TotalDamage:         1162.356159,
		HPDamage:            1162.356159,
		BaseDamage:          1162.356159,
		ShieldAbsorb:        0.0,
		IsAttackerChar:      true,
		DefenceMultiplier:   1.0,
		Resistance:          1.0,
		Vulnerability:       1.0,
		ToughnessMultiplier: 1.0,
		Fatigue:             1.0,
		AllDamageReduce:     1.0,
		CritDamage:          1.0,
		IsCrit:              false,
	})
}
