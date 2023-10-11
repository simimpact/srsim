package combat

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/tests/mock"
	"github.com/stretchr/testify/assert"
)

const TOLERANCE = 0.000001

type ModifyAttributeMatcher struct {
	t      *testing.T
	Target key.TargetID
	Source key.TargetID
	Amount float64
}

func (m ModifyAttributeMatcher) Matches(x interface{}) bool {
	arg, ok := x.(info.ModifyAttribute)
	if !ok {
		return false
	}

	assert.Equal(m.t, m.Target, arg.Target)
	assert.Equal(m.t, m.Source, arg.Source)
	assert.InDelta(m.t, m.Amount, arg.Amount, TOLERANCE)
	return true
}

func (m ModifyAttributeMatcher) String() string {
	return fmt.Sprintf(
		"{ModifyAttribute - Target:%d Source:%d Amount:%f}", m.Target, m.Source, m.Amount)
}

type PerformHitTester struct {
	t         *testing.T
	mockCtrl  *gomock.Controller
	Attribute *mock.MockAttribute
	Shield    *mock.MockShield
	Engine    *mock.MockEngine
}

func NewPerformHitTester(t *testing.T) *PerformHitTester {
	mockCtrl := gomock.NewController(t)

	return &PerformHitTester{
		t:         t,
		mockCtrl:  mockCtrl,
		Attribute: mock.NewMockAttribute(mockCtrl),
		Shield:    mock.NewMockShield(mockCtrl),
		Engine:    mock.NewMockEngine(mockCtrl),
	}
}

func (pht *PerformHitTester) Finish() {
	pht.mockCtrl.Finish()
}

type ExpectHit struct {
	TotalDamage         float64
	ShieldAbsorb        float64
	HPDamage            float64
	IsAttackerChar      bool
	BaseDamage          float64
	DefenceMultiplier   float64
	Resistance          float64
	Vulnerability       float64
	ToughnessMultiplier float64
	Fatigue             float64
	AllDamageReduce     float64
	CritDamage          float64
	IsCrit              bool
}

func (pht *PerformHitTester) AssertPerformHit(hit *info.Hit, expect *ExpectHit) {
	rdm := rand.New(rand.NewSource(0))
	mgr := New(&event.System{}, pht.Attribute, pht.Shield, pht.Engine, rdm)

	hitStartEmitted := false
	mgr.event.HitStart.Subscribe(func(e event.HitStart) {
		// assert that this is the same hit (event is being emitted)
		assert.Equal(pht.t, hit, e.Hit)
		hitStartEmitted = true
	})

	hitEndEmitted := false
	mgr.event.HitEnd.Subscribe(func(e event.HitEnd) {
		hitEndEmitted = true

		assert.InDelta(pht.t, expect.TotalDamage, e.TotalDamage, TOLERANCE)
		assert.InDelta(pht.t, expect.ShieldAbsorb, e.ShieldDamage, TOLERANCE)
		assert.InDelta(pht.t, expect.HPDamage, e.HPDamage, TOLERANCE)
		assert.InDelta(pht.t, expect.BaseDamage, e.BaseDamage, TOLERANCE)
		assert.InDelta(pht.t, expect.DefenceMultiplier, e.DefenceMultiplier, TOLERANCE)
		assert.InDelta(pht.t, expect.Resistance, e.Resistance, TOLERANCE)
		assert.InDelta(pht.t, expect.Vulnerability, e.Vulnerability, TOLERANCE)
		assert.InDelta(pht.t, expect.ToughnessMultiplier, e.ToughnessMultiplier, TOLERANCE)
		assert.InDelta(pht.t, expect.Fatigue, e.Fatigue, TOLERANCE)
		assert.InDelta(pht.t, expect.AllDamageReduce, e.AllDamageReduce, TOLERANCE)
		assert.InDelta(pht.t, expect.CritDamage, e.CritDamage, TOLERANCE)
		assert.Equal(pht.t, expect.IsCrit, e.IsCrit)
	})

	pht.Shield.EXPECT().
		AbsorbDamage(hit.Defender.ID(), gomock.Any()).
		DoAndReturn(func(id key.TargetID, amt float64) float64 {
			return amt - expect.ShieldAbsorb
		})

	modify := ModifyAttributeMatcher{
		t:      pht.t,
		Target: hit.Defender.ID(),
		Source: hit.Attacker.ID(),
		Amount: 0,
	}

	ratio := 1.0
	if hit.HitRatio > 0 {
		ratio = hit.HitRatio
	}

	modify.Amount = -expect.HPDamage
	pht.Attribute.EXPECT().ModifyHPByAmount(modify, true)

	modify.Amount = -hit.StanceDamage * ratio
	pht.Attribute.EXPECT().ModifyStance(modify)

	if expect.IsAttackerChar {
		modify.Target = hit.Attacker.ID()
	}
	modify.Amount = hit.EnergyGain * ratio
	pht.Attribute.EXPECT().ModifyEnergy(modify)

	pht.Attribute.EXPECT().HPRatio(hit.Defender.ID())
	pht.Engine.EXPECT().IsCharacter(hit.Attacker.ID()).Return(expect.IsAttackerChar)

	mgr.performHit(hit)
	assert.True(pht.t, hitStartEmitted, "HitStart event was never emitted")
	assert.True(pht.t, hitEndEmitted, "HitEnd event was never emitted")
}
