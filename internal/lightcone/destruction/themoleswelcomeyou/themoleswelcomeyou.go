package themoleswelcomeyou

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TheMolesWelcomeYou key.Modifier = "the-moles-welcome-you"
)

// When the wearer uses Basic ATK, Skill, or Ultimate to attack enemies,
// the wearer gains one stack of Mischievous. Each stack increases the wearer's ATK by 12%.
func init() {
	lightcone.Register(key.TheMolesWelcomeYou, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(TheMolesWelcomeYou, modifier.Config{
		MaxCount:          4, // max is actually 3, doing -1 in stat calc
		CountAddWhenStack: 1,
		Stacking:          modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnAfterAttack: onAfterAttack,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   TheMolesWelcomeYou,
		Source: owner,
		State:  0.12 + 0.03*float64(lc.Ascension),
	})
}

// treat count 1 as no stacks
func onAdd(mod *modifier.ModifierInstance) {
	amt := mod.State().(float64)
	mod.AddProperty(prop.ATKPercent, amt*(mod.Count()-1))
}

// after an attack, add 1 stack
func onAfterAttack(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	if e.AttackType == model.AttackType_NORMAL ||
		e.AttackType == model.AttackType_SKILL ||
		e.AttackType == model.AttackType_ULT {
		mod.Engine().ExtendModifierCount(mod.Owner(), TheMolesWelcomeYou, 1)
	}
}
