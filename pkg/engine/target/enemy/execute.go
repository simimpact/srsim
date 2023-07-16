package enemy

import (
	"github.com/simimpact/srsim/pkg/engine/target"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) ExecuteAction(id key.TargetID) (target.ExecutableAction, error) {
	return target.ExecutableAction{
		Execute:    func() {},
		AttackType: model.AttackType_NORMAL,
		IsInsert:   false,
		SPDelta:    0,
		Key:        "dummy",
	}, nil
}
