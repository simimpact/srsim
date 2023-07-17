package teststub

import (
	"testing"
	"time"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type mockManager struct {
	t        *testing.T
	turnPipe chan TurnCommand
}

type TurnCommand struct {
	Next key.TargetID
	Av   float64
}

func newMockManager(t *testing.T, pipe chan TurnCommand) *mockManager {
	return &mockManager{
		t:        t,
		turnPipe: pipe,
	}
}

func (m *mockManager) TotalAV() float64 {
	return 0
}

func (m *mockManager) AddTargets(ids ...key.TargetID) {
}

func (m *mockManager) RemoveTarget(id key.TargetID) {
}

func (m *mockManager) StartTurn() (key.TargetID, float64, []event.TurnStatus, error) {
	select {
	case t := <-m.turnPipe:
		return t.Next, t.Av, nil, nil
	case <-time.After(1 * time.Second):
		LogError(m.t, "mockManager StartTurn did not receive next turn command")
		panic("Test failed, be sure to call NextTurn")
	}
}

func (m *mockManager) ResetTurn() error {
	return nil
}

func (m *mockManager) SetGauge(data info.ModifyAttribute) error {
	return nil
}

func (m *mockManager) ModifyGaugeNormalized(data info.ModifyAttribute) error {
	return nil
}

func (m *mockManager) ModifyGaugeAV(data info.ModifyAttribute) error {
	return nil
}

func (m *mockManager) SetCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
}

func (m *mockManager) ModifyCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
}
