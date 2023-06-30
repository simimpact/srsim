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
	Buff  key.Modifier = "collapsing-sky-buff"
)

func init() {
	lightcone.Register(key.CollapsingSky, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAttack: onBeforeAttack,
			OnAfterAttack:  onAfterAttack,
		},
	})

	modifier.Register(Buff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		State:  0.15 + 0.05*float64(lc.Imposition),
	})
}

func onBeforeAttack(mod *modifier.ModifierInstance, e event.AttackStartEvent) {
	dmgBonus := mod.State().(float64)

	if e.AttackType == 1 || e.AttackType == 2 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   Buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllDamagePercent: dmgBonus},
		})
	}
}

func onAfterAttack(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	if e.AttackType == 1 || e.AttackType == 2 {
		mod.Engine().RemoveModifier(mod.Owner(), Buff)
	}
}
