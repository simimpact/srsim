package collapsingsky

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
	Check key.Modifier = "collapsing-sky"
)

// Increases the wearer's Basic ATK and Skill DMG by 20%.
func init() {
	lightcone.Register(key.CollapsingSky, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: onBeforeHit,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		State:  0.15 + 0.05*float64(lc.Imposition),
	})
}

func onBeforeHit(mod *modifier.Instance, e event.HitStartEvent) {
	dmgBonus := mod.State().(float64)

	if e.Hit.AttackType == model.AttackType_NORMAL || e.Hit.AttackType == model.AttackType_SKILL {
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, dmgBonus)
	}
}
