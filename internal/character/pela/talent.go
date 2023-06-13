package pela

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Talent   key.Modifier = "pela-talent"
	TalentCD key.Modifier = "pela-talent-cd"
)

func init() {
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: talentAfterAttack,
		},
	})
}

func talentAfterAttack(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	if mod.Engine().HasModifier(mod.Owner(), TalentCD) {
		return
	}
	char, _ := mod.Engine().CharacterInfo(mod.Owner())
	mod.Engine().ModifyEnergy(mod.Owner(), talent[char.AbilityLevel.Talent-1])
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:            TalentCD,
		Source:          mod.Owner(),
		Duration:        1,
		TickImmediately: true,
	})
}
