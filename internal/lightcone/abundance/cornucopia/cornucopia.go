package cornucopia

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
	CornucopiaCheck key.Modifier = "cornucopia-check"
	CornucopiaBuff  key.Modifier = "cornucopia-buff"
)

//When the wearer uses their Skill or Ultimate, 
//their Outgoing Healing increases by 12%.

func init() {
	lightcone.Register(key.Cornucopia, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})

	modifier.Register(CornucopiaCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: buffHealOnSkillUlt,
		},
	})

	modifier.Register(CornucopiaBuff, modifier.Config{
		Stacking:   modifier.Unique,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   CornucopiaCheck,
		Source: owner,
		State:  0.09 + 0.03*float64(lc.Imposition),
	})
}

func buffHealOnSkillUlt(mod *modifier.ModifierInstance, e event.ActionEvent) {
	amt := mod.State().(float64)

	if (e.AttackType == model.AttackType_SKILL || e.AttackType == model.AttackType_ULT) {
			mod.Engine().AddModifier(mod.Owner(), info.Modifier{
				Name:     CornucopiaBuff,
				Source:   mod.Owner(),
				Duration: -1,
				Stats:    info.PropMap{prop.HealBoost: amt},
			})
		}
}