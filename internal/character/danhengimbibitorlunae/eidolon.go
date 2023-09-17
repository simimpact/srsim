package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E6 = "danhengimbibitorlunae-e6"
)

func init() {
	modifier.Register(E6, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		MaxCount:   3,
		Listeners: modifier.Listeners{
			OnBeforeAttack: E6OnBeforeAttack,
			OnAfterAttack:  E6OnAfterAttack,
		},
		CountAddWhenStack: 1,
	})
}

func (c *char) E6ActionEndListener(e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT && e.Owner != c.id {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6,
			Source: c.id,
		})
	}
}

func E6OnBeforeAttack(mod *modifier.Instance, e event.AttackStart) {
	mod.SetProperty(prop.ImaginaryPEN, mod.Count()*0.2)
}
func E6OnAfterAttack(mod *modifier.Instance, e event.AttackEnd) {
	mod.RemoveSelf()
}
