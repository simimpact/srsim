package combat

import (
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/tests/mock"
	"github.com/stretchr/testify/assert"
)

func TestPerformHitWithShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	attr := mock.NewMockAttribute(mockCtrl)
	shld := mock.NewMockShield(mockCtrl)
	target := mock.NewMockEngine(mockCtrl)
	rdm := rand.New(rand.NewSource(0))
	mgr := New(&event.System{}, attr, shld, target, rdm)

	hit := &info.Hit{
		Attacker:     mock.NewEmptyStats(1),
		Defender:     mock.NewEmptyStats(2),
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: 0.5},
		AttackType:   model.AttackType_NORMAL,
		DamageType:   model.DamageType_ICE,
		StanceDamage: 25,
		HitRatio:     1.0,
	}

	hitStartEmitted := false
	mgr.event.HitStart.Subscribe(func(e event.HitStart) {
		// assert that this is the same hit (event is being emitted)
		assert.Equal(t, hit, e.Hit)
		hitStartEmitted = true
	})
	defer func() { assert.True(t, hitStartEmitted, "HitStart event was never emitted") }()

	// POPULATE STATS
	hit.Attacker.AddProperty(prop.ATKBase, 200)

	// EXPECTED RESULTS
	total := 100.0 // defender has 0 DEF, attacker only has 200 ATK w/ 50% attack, so 100 total
	absorb := 0.0  // TODO: change this to 10
	hpUpdate := total - absorb

	hitEndEmitted := false
	mgr.event.HitEnd.Subscribe(func(e event.HitEnd) {
		hitEndEmitted = true

		assert.Equal(t, total, e.TotalDamage, "overall calculated hit damage")
		assert.Equal(t, absorb, e.ShieldDamage, "damage dealt to shield (subset of total)")
		assert.Equal(t, hpUpdate, e.HPDamage, "damage dealt to HP (subset of total)")
		assert.Equal(t, total, e.BaseDamage, "base == total since no other stats")
		assert.Equal(t, 1.0, e.BonusDamage)
		assert.Equal(t, 1.0, e.DefenceMultiplier)
		assert.Equal(t, 1.0, e.Resistance)
		assert.Equal(t, 1.0, e.Vulnerability)
		assert.Equal(t, 1.0, e.ToughnessMultiplier)
		assert.Equal(t, 1.0, e.Fatigue)
		assert.Equal(t, 1.0, e.AllDamageReduce)
		assert.Equal(t, 1.0, e.CritDamage)
		assert.False(t, e.IsCrit)
		assert.False(t, e.UseSnapshot)
	})
	defer func() { assert.True(t, hitEndEmitted, "HitEnd event was never emitted") }()

	// total passed into shield, it returns the absorbed amount
	shld.EXPECT().
		AbsorbDamage(hit.Defender.ID(), total).
		DoAndReturn(func(id key.TargetID, amt float64) float64 {
			return amt - absorb
		})

	attr.EXPECT().ModifyHPByAmount(hit.Defender.ID(), hit.Attacker.ID(), -hpUpdate, true)
	attr.EXPECT().ModifyStance(hit.Defender.ID(), hit.Attacker.ID(), -hit.StanceDamage)
	attr.EXPECT().ModifyEnergy(hit.Attacker.ID(), 0.0)
	attr.EXPECT().HPRatio(hit.Defender.ID())
	target.EXPECT().IsCharacter(hit.Attacker.ID()).Return(true)

	mgr.performHit(hit)
}
