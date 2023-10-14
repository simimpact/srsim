package yanqing

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A6                = "yanqing-a6"
	Traces            = "yanqing-traces"
	A2     key.Attack = "yanqing-a2"
)

func (c *char) initTraces() {
	modifier.Register(Traces, modifier.Config{
		StatusType: model.StatusType_UNKNOWN_STATUS,
		Listeners: modifier.Listeners{
			OnAfterAttack: A2Listener,
			OnAfterHitAll: A6Listener,
		},
	})
	modifier.Register(A6, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.Refresh,
		Duration:   2,
	})
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Traces,
		Source: c.id,
	})
}
func A2Listener(mod *modifier.Instance, e event.AttackEnd) {
	cinfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	if cinfo.Traces["101"] && mod.Engine().Stats(e.Targets[0]).IsWeakTo(model.DamageType_ICE) {
		mod.Engine().Attack(info.Attack{
			Key:        A2,
			Source:     e.Attacker,
			Targets:    e.Targets,
			DamageType: model.DamageType_ICE,
			AttackType: e.AttackType,
			BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: 0.3},
		})
	}
}
func A6Listener(mod *modifier.Instance, e event.HitEnd) {
	cinfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	if cinfo.Traces["103"] && e.IsCrit {
		mod.Engine().AddModifier(e.Attacker, info.Modifier{
			Name:  A6,
			Stats: info.PropMap{prop.SPDPercent: 0.1},
		})
	}
}
