package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/simimpact/srsim/tests/mock"
	"github.com/stretchr/testify/assert"
)

// helper to build sim with necessary mocks
func NewSim(t *testing.T, seed int64) (*simulation.Simulation, *mock.MockAttribute) {
	mockCtrl := gomock.NewController(t)
	attr := mock.NewMockAttribute(mockCtrl)
	sim := &simulation.Simulation{
		Attr:   attr,
		Random: rand.New(rand.NewSource(seed)),
	}
	return sim, attr
}

// test var : Targets list
func TestRetargetEmptyTargets(t *testing.T) {
	sim, attr := NewSim(t, 1)

	// for all targets, return full health in this test
	attr.EXPECT().HPRatio(gomock.Any()).Return(1.0).AnyTimes()

	result := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{},
		Filter:  func(target key.TargetID) bool { return target == 2 },
		Max:     3,
	})

	assert.Equal(t, []key.TargetID{}, result)
}

func TestRetargetNoFilterFunc(t *testing.T) {
	sim, attr := NewSim(t, 100)

	// for all targets, return full health in this test
	attr.EXPECT().HPRatio(gomock.Any()).Return(1.0).AnyTimes()

	result := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{1, 2, 3, 4},
	})

	// need to use ElementsMatch since order not guaranteed with random
	assert.ElementsMatch(t, []key.TargetID{1, 2, 3, 4}, result)
}

func TestRetargetEmptyFiltered(t *testing.T) {
	sim, attr := NewSim(t, 1)

	// for all targets, return full health in this test
	attr.EXPECT().HPRatio(gomock.Any()).Return(1.0).AnyTimes()

	result := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{1, 2, 3, 4},
		Filter:  func(target key.TargetID) bool { return false },
		Max:     3,
	})

	assert.Equal(t, []key.TargetID{}, result)
}

// test variable : Max
func TestRetargetMaxZero(t *testing.T) {
	sim, attr := NewSim(t, 1)

	// for all targets, return full health in this test
	attr.EXPECT().HPRatio(gomock.Any()).Return(1.0).AnyTimes()

	result := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{1, 2, 3},
		Filter:  func(target key.TargetID) bool { return target == 2 },
		// Max:     3, // TODO: remove this line to see the bug
	})

	assert.Equal(t, []key.TargetID{2}, result)
}

func TestRetargetMaxHigh(t *testing.T) {
	sim, attr := NewSim(t, 1)

	// for all targets, return full health in this test
	attr.EXPECT().HPRatio(gomock.Any()).Return(1.0).AnyTimes()

	result := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{1, 2, 3},
		Filter:  func(target key.TargetID) bool { return target == 2 },
		Max:     100, // TODO: uncomment this line to see the bug
	})

	assert.Equal(t, []key.TargetID{2}, result)
}

func TestRetargetMaxLow(t *testing.T) {
	sim, attr := NewSim(t, 1)

	// for all targets, return full health in this test
	attr.EXPECT().HPRatio(gomock.Any()).Return(1.0).AnyTimes()

	result1 := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{1, 2, 3, 4, 5},
		Filter:  func(target key.TargetID) bool { return true },
		Max:     2,
	})
	result2 := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{1, 2, 3, 4, 5},
		Filter:  func(target key.TargetID) bool { return true },
		Max:     3,
	})
	result3 := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{1, 2, 3, 4, 5},
		Filter:  func(target key.TargetID) bool { return true },
		Max:     4,
	})

	assert.Exactly(t, []key.TargetID{3, 1}, result1)
	assert.Exactly(t, []key.TargetID{2, 4, 1}, result2)
	assert.Exactly(t, []key.TargetID{5, 3, 4, 2}, result3)
}

// test var : no random
func TestRetargetNoRandom(t *testing.T) {
	sim, attr := NewSim(t, 1)

	// for all targets, return full health in this test
	attr.EXPECT().HPRatio(gomock.Any()).Return(1.0).AnyTimes()

	result := sim.Retarget(info.Retarget{
		Targets:       []key.TargetID{1, 2, 3, 4},
		Filter:        func(target key.TargetID) bool { return true },
		Max:           3,
		DisableRandom: true,
	})

	assert.Exactly(t, []key.TargetID{1, 2, 3}, result) // expect exact match w/ order
}

// test variable : Ignore Limbo
func TestRetargetIgnoresLimbo(t *testing.T) {
	sim, attr := NewSim(t, 124)
	targets := []key.TargetID{1, 2, 3}

	// target 1, 0 HP
	// target 2, 0.5 HP
	// target 3, 1.0 HP
	attr.EXPECT().HPRatio(targets[0]).Return(0.0)
	attr.EXPECT().HPRatio(targets[1]).Return(0.5)
	attr.EXPECT().HPRatio(targets[2]).Return(1.0)

	result := sim.Retarget(info.Retarget{
		Targets: []key.TargetID{1, 2, 3},
		Filter:  func(target key.TargetID) bool { return true },
		// Max:     3,
	})

	// need to use ElementsMatch since order not guaranteed with random
	assert.ElementsMatch(t, []key.TargetID{2, 3}, result)
}
