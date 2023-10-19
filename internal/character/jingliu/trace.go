package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A6 = "jingliu-a6"
)

func init() {
	modifier.Register(A6, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAfterAction: removeWhenEndUlt,
		},
	})
}

func A6Listener(mod *modifier.Instance, e event.ActionStart) {
	cinfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	if cinfo.Traces["103"] && e.AttackType == model.AttackType_ULT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   A6,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllDamagePercent: 0.2},
		})
	}
}
func removeWhenEndUlt(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT {
		mod.RemoveSelf()
	}
}
