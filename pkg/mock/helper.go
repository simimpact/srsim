package mock

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func NewMockEngineWithEvents(ctrl *gomock.Controller) *MockEngine {
	m := NewMockEngine(ctrl)
	events := &event.System{}
	m.EXPECT().Events().Return(events).AnyTimes()
	return m
}

func NewEmptyStats(target key.TargetID) *info.Stats {
	attr := info.Attributes{}
	mods := info.ModifierState{
		Props:     info.NewPropMap(),
		DebuffRES: info.NewDebuffRESMap(),
	}
	return info.NewStats(target, attr, mods)
}
