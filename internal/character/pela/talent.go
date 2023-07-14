package pela

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent   key.Modifier = "pela-talent"
	TalentCD key.Modifier = "pela-talent-cd"
)

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: func(mod *modifier.Instance, e event.AttackEnd) {
				char, _ := mod.Engine().CharacterInfo(mod.Owner())
				for _, trg := range e.Targets {
					if mod.Engine().ModifierStatusCount(trg, model.StatusType_STATUS_DEBUFF) >= 1 {
						mod.Engine().ModifyEnergy(mod.Owner(), talent[char.TalentLevelIndex()])
						return
					}
				}
			},
		},
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
	})
}
