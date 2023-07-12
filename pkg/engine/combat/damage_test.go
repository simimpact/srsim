package combat

import (
	"math/rand"
	"testing"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/tests/mock"
	"github.com/stretchr/testify/assert"
)

func TestBaseDamageByATK(t *testing.T) {
	hit := &info.Hit{
		Key:        "tst",
		HitIndex:   0,
		Attacker:   mock.NewEmptyStats(1),
		Defender:   mock.NewEmptyStats(2),
		BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: 0.5},
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_FIRE,
		HitRatio:   1.0,
	}

	hit.Attacker.AddProperty(prop.ATKBase, 100.0)
	result := baseDamage(hit)

	assert.Equal(t, 50.0, result)
}

func TestCritCheckNormal(t *testing.T) {
	rdm := rand.New(rand.NewSource(1))
	hit := &info.Hit{
		Key:        "tst",
		HitIndex:   0,
		Attacker:   mock.NewEmptyStats(1),
		Defender:   mock.NewEmptyStats(2),
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{},
		HitRatio:   1.0,
	}

	hit.Attacker.AddProperty(prop.CritChance, 1.0)
	assert.True(t, crit(hit, rdm), "is a crit with 100% chance")
}

func TestCritCheckPureDamage(t *testing.T) {
	rdm := rand.New(rand.NewSource(1))
	hit := &info.Hit{
		Key:          "tst",
		HitIndex:     0,
		Attacker:     mock.NewEmptyStats(1),
		Defender:     mock.NewEmptyStats(2),
		AttackType:   model.AttackType_NORMAL,
		DamageType:   model.DamageType_FIRE,
		BaseDamage:   info.DamageMap{},
		HitRatio:     1.0,
		AsPureDamage: true,
	}

	hit.Attacker.AddProperty(prop.CritChance, 1.0)
	assert.False(t, crit(hit, rdm), "is never a crit")
}
