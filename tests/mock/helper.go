package mock

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	model "github.com/simimpact/srsim/pkg/model"
)

func NewMockEngineWithEvents(ctrl *gomock.Controller) *MockEngine {
	m := NewMockEngine(ctrl)
	events := &event.System{}
	m.EXPECT().Events().Return(events).AnyTimes()
	return m
}

func NewEmptyStats(target key.TargetID) *info.Stats {
	attr := new(info.Attributes)
	mods := &info.ModifierState{
		Props:     info.NewPropMap(),
		DebuffRES: info.NewDebuffRESMap(),
		Weakness:  info.NewWeaknessMap(),
		Counts:    make(map[model.StatusType]int),
		Flags:     nil,
		Modifiers: nil,
	}
	return info.NewStats(target, attr, mods)
}

func NewEmptyStatsWithAttr(target key.TargetID, attr *info.Attributes) *info.Stats {
	mods := &info.ModifierState{
		Props:     info.NewPropMap(),
		DebuffRES: info.NewDebuffRESMap(),
		Weakness:  info.NewWeaknessMap(),
		Counts:    make(map[model.StatusType]int),
		Flags:     nil,
		Modifiers: nil,
	}
	return info.NewStats(target, attr, mods)
}
