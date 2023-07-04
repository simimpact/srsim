package modifier

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/tests/mock"
)

func FailOnPanic(t *testing.T) {
	if r := recover(); r != nil {
		t.FailNow()
	}
}

func NewTestManager(t *testing.T) (*Manager, *gomock.Controller) {
	mockCtrl := gomock.NewController(t)
	engine := mock.NewMockEngine(mockCtrl)

	manager := &Manager{
		engine:    engine,
		targets:   make(map[key.TargetID]activeModifiers),
		turnCount: 0,
	}
	return manager, mockCtrl
}

func NewTestManagerWithEvents(t *testing.T) (*Manager, *gomock.Controller) {
	mockCtrl := gomock.NewController(t)
	engine := mock.NewMockEngineWithEvents(mockCtrl)

	manager := &Manager{
		engine:    engine,
		targets:   make(map[key.TargetID]activeModifiers),
		turnCount: 0,
	}
	return manager, mockCtrl
}
