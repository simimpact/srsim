package disciple

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	check  = "longevous-disciple"
	crbuff = "longevous-disciple-cr-buff"
)

// 2pc: Increases Max HP by 12%.
// 4pc: When the wearer is hit or has their HP consumed by an ally or themselves,
//      their CRIT Rate increases by 8% for 2 turn(s) and up to 2 stacks.

// TO-DO: in onHPChange, assumes e.Target means source; check whether this is correct

func init() {
	relic.Register(key.BandOfSizzlingThunder, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.HPPercent: 0.12},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   check,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(check, modifier.Config{
		Listeners: modifier.Listeners{
			OnHPChange:           onHPChange,
			OnAfterBeingAttacked: onAfterBeingAttacked,
		},
	})
	modifier.Register(crbuff, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		StatusType:        model.StatusType_STATUS_BUFF,
		MaxCount:          2,
		CountAddWhenStack: 1,
		Duration:          2,
		Listeners: modifier.Listeners{
			OnAdd: onAdd,
		},
	})
}

func onHPChange(mod *modifier.Instance, e event.HPChange) {
	// source is friendly, not HP change by damage, hp change is negative
	if mod.Engine().IsCharacter(e.Target) && !e.IsHPChangeByDamage && e.NewHP < e.OldHP {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   crbuff,
			Source: mod.Owner(),
		})
	}
}

func onAfterBeingAttacked(mod *modifier.Instance, e event.AttackEnd) {
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   crbuff,
		Source: mod.Owner(),
	})
}

func onAdd(mod *modifier.Instance) {
	mod.AddProperty(prop.ATKPercent, 0.08*mod.Count())
}
