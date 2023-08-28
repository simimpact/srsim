package seele

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2Check key.Modifier = "seele-a2-check"
	A2Aggro key.Modifier = "seele-a2-aggro"
)

// A2: When current HP percentage is 50% or lower, reduces the chance of being attacked by enemies

func init() {
	modifier.Register(A2Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnHPChange: reduceAggro,
		},
	})
	modifier.Register(A2Aggro, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) initTraces() {
	if c.info.Traces["101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2Check,
			Source: c.id,
		})
	}
}

func reduceAggro(mod *modifier.Instance, e event.HPChange) {
	if mod.Engine().HPRatio(mod.Owner()) <= 0.5 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   A2Aggro,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AggroPercent: -0.5},
		})
	} else {
		mod.Engine().RemoveModifier(mod.Owner(), A2Aggro)
	}
}

// advances forward based on if current turn is an insert or not.
func (c *char) advanceForward(key key.Reason, isInsert bool) {
	// if not insert : modifygaugecost -20%
	// if insert : modifynormalized -20%
	if isInsert {
		c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    key,
			Source: c.id,
			Target: c.id,
			Amount: -0.2,
		})
	} else {
		c.engine.ModifyCurrentGaugeCost(info.ModifyCurrentGaugeCost{
			Key:    key,
			Source: c.id,
			Amount: -0.2,
		})
	}
}
