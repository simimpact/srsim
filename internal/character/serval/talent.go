package serval

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	ServalTalentListener key.Modifier = "serval-talent-listener"
	ServalTalent         key.Attack   = "serval-talent"
)

func init() {
	modifier.Register(ServalTalentListener, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: onAfterAttack,
		},
	})
}

func onAfterAttack(mod *modifier.Instance, e event.AttackEnd) {
	// for each target in e.Targets, if the target has shock, then add pursued damage
	for _, target := range e.Targets {
		if mod.Engine().HasModifier(target, common.Shock) {
			mod.Engine().Attack(info.Attack{
				Key:        ServalTalent,
				Source:     e.Attacker,
				Targets:    []key.TargetID{target},
				DamageType: model.DamageType_THUNDER,
				AttackType: model.AttackType_PURSUED,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: talent[e.Attacker],
				},
				StanceDamage: 0.0,
				EnergyGain:   0.0,
			})
		}
	}
}
