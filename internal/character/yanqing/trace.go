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
	A6            = "yanqing-a6"
	A2 key.Attack = "yanqing-a2"
)

func init() {
	modifier.Register(A6, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.Refresh,
		Duration:   2,
	})
}

func (c *char) A2Listener(e event.AttackEnd) {
	if c.info.Traces["101"] && c.engine.Stats(e.Targets[0]).IsWeakTo(model.DamageType_ICE) {
		c.engine.Attack(info.Attack{
			Key:        A2,
			Source:     e.Attacker,
			Targets:    e.Targets,
			DamageType: model.DamageType_ICE,
			AttackType: e.AttackType,
			BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: 0.3},
		})
	}
}

func (c *char) A6Listener(e event.HitEnd) {
	if c.info.Traces["103"] && e.IsCrit {
		c.engine.AddModifier(e.Attacker, info.Modifier{
			Name:   A6,
			Stats:  info.PropMap{prop.SPDPercent: 0.1},
			Source: c.id,
		})
	}
}
