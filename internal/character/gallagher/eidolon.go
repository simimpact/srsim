package gallagher

import (
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2 = "gallagher-e2"
)

func init() {
	modifier.Register(E2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		Duration:   2,
	})
}
