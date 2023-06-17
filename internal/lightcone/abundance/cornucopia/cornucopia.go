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
	CornucopiaBuff  key.Modifier = "cornucopia-buff"
)

func init() {
	lightcone.Register(key.Cornucopia, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	
	modifier.Register(CornucopiaBuff, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: buffHealsOnSkillUlt,
			OnAfterAction: removeHealBuff,
		},
		Stacking: modifier.Unique,
	})

}

//When the wearer uses their Skill or Ultimate, their Outgoing Healing increases by 12%.
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   CornucopiaBuff,
		Source: owner,
		State:  0.09 + 0.03 * float64(lc.Imposition),
	})
}

func buffHealsOnSkillUlt(mod *modifier.ModifierInstance, e event.ActionEvent) {
	healAmt := mod.State().(float64)
	switch e.AttackType {
	case model.AttackType_SKILL, model.AttackType_ULT :
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     CornucopiaBuff,
			Source:   mod.Owner(),
			Duration: -100,
			Stats:    info.PropMap{prop.HealBoost: healAmt},
		})
	}
}
//is it neccessary to remove a permanent buff after each turn?
func removeHealBuff(mod *modifier.ModifierInstance, e event.ActionEvent)  {
	mod.Engine().RemoveModifier(mod.Owner(), mod.Name())
}