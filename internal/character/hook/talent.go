package hook

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
)

const (
	Talent = "hook-talent"
)

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAttack: talentProc,
		},
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
	})
}

func talentProc(mod *modifier.Instance, e event.AttackStart) {
	mod.Engine().Attack(info.Attack{})

}
