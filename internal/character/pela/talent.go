package pela

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent = "pela-talent"
)

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: afterAttack,
		},
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
	})
}

func afterAttack(mod *modifier.Instance, e event.AttackEnd) {
	char, _ := mod.Engine().CharacterInfo(mod.Owner())
	for _, trg := range e.Targets {
		if mod.Engine().HPRatio(trg) > 0 && mod.Engine().ModifierStatusCount(trg, model.StatusType_STATUS_DEBUFF) >= 1 {
			mod.Engine().ModifyEnergy(info.ModifyAttribute{
				Key:    Talent,
				Target: mod.Owner(),
				Source: mod.Owner(),
				Amount: talent[char.TalentLevelIndex()],
			})
			return
		}
	}
}
