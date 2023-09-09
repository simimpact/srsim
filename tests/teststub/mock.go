package teststub

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type mockTurnManager struct {
	t            *testing.T
	turnSequence []key.TargetID
}

type TurnCommand struct {
	Next key.TargetID
	Av   float64
}

func newMockManager(t *testing.T) *mockTurnManager {
	return &mockTurnManager{
		t:            t,
		turnSequence: nil,
	}
}

func (m *mockTurnManager) queue(ids ...key.TargetID) {
	m.turnSequence = append(m.turnSequence, ids...)
}

func (m *mockTurnManager) TotalAV() float64 {
	return 0
}

func (m *mockTurnManager) AddTargets(ids ...key.TargetID) {
}

func (m *mockTurnManager) RemoveTarget(id key.TargetID) {
}

// mockTurnManager StartTurn shouldn't be without any TargetID in the sequence, but the corresponding
// error message is being temporarily removed due to this being an asynch thread that may be called
// after test ends, causing improper logging panics
func (m *mockTurnManager) StartTurn() (key.TargetID, float64, []event.TurnStatus, error) {
	if len(m.turnSequence) == 0 {
		return 1, 100000, nil, nil
	}
	if m.turnSequence[0] == -1 {
		return 1, 100000, nil, nil
	}
	tgt := m.turnSequence[0]
	m.turnSequence = m.turnSequence[1:]
	return tgt, 0, nil, nil
}

func (m *mockTurnManager) TurnOrder() []key.TargetID {
	return m.turnSequence
}

func (m *mockTurnManager) ResetTurn() error {
	return nil
}

func (m *mockTurnManager) SetGauge(data info.ModifyAttribute) error {
	return nil
}

func (m *mockTurnManager) ModifyGaugeNormalized(data info.ModifyAttribute) error {
	return nil
}

func (m *mockTurnManager) ModifyGaugeAV(data info.ModifyAttribute) error {
	return nil
}

func (m *mockTurnManager) SetCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
}

func (m *mockTurnManager) ModifyCurrentGaugeCost(data info.ModifyCurrentGaugeCost) {
}
