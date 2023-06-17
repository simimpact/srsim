package enemy

import (
	"github.com/simimpact/srsim/pkg/engine/target"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) ExecuteAction(id key.TargetID) (target.ExecutableAction, error) {
	return target.ExecutableAction{
		Execute: func() {},
	}, nil
}
