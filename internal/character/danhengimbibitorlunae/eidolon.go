package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E6Effect = "danhengimbibitorlunae-e6effect" // MAvatar_DanHengIL_00_Rank06_ImaginaryPenetrate
)

// E1 increase talent max stack by 4 and gain 1 extra stack when hit
// E2 100% forward and 1 more ult stack
// E4 skill buff 1 more turn
// E6 20% imaginary pen for attack3 when ally use ult, max 3 stack

func init() {
	// imaginary pen for attack3,change by ally ult count
	modifier.Register(E6Effect, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

// count ally ult
func (c *char) E6ActionEndListener(e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT && e.Owner != c.id && c.E6Count < 3 {
		c.E6Count++
	}
}
