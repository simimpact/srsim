package tingyun

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1 = "tingyun-e1"
	E2 = "tingyun-e2"
)

func init() {
	modifier.Register(E1, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		TickMoment: modifier.ModifierPhase1End,
	})
}

func doE1OnUlt(mod *modifier.Instance, e event.ActionEnd) {
	st := mod.State().(*skillState)
	// bypass
	if st.tingEidolon < 1 {
		return
	}
	if e.AttackType == model.AttackType_ULT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   E1,
			Source: mod.Source(),
			Stats:  info.PropMap{prop.SPDPercent: 0.2},
		})
	}
}

func doE2(mod *modifier.Instance, target key.TargetID) {
	st := mod.State().(*skillState)
	// bypass
	if st.tingEidolon < 2 || st.e2flag {
		return
	}
	mod.Engine().ModifyEnergy(info.ModifyAttribute{
		Key:    E2,
		Target: mod.Owner(),
		Source: mod.Source(),
		Amount: 5.0,
	})
	st.e2flag = true
}

// not fully correct, this should trigger with OnAllowAction instead of OnPhase1
func removeE2CD(mod *modifier.Instance) {
	st := mod.State().(*skillState)
	st.e2flag = false
}
