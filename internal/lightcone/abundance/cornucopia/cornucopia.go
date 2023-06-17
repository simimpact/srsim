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

func init() {
	lightcone.Register(key.Cornucopia, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	//Implement checker here
	modifier.Register(CornucopiaCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: buffHealsOnSkillUlt,
			OnAfterAction: removeHealBuff,
		},
	})
	//The actual buff modifier goes here
	modifier.Register(CornucopiaBuff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

//When the wearer uses their Skill or Ultimate, their Outgoing Healing increases by 12%(S1)
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	//checker goes here
	engine.AddModifier(owner, info.Modifier{ 
		Name:   CornucopiaCheck,
		Source: owner,
		State:  0.09 + 0.03 * float64(lc.Imposition),
	})
}
//add buff only on skill and ult actions
func buffHealsOnSkillUlt(mod *modifier.ModifierInstance, e event.ActionEvent) {
	healAmt := mod.State().(float64)
	switch e.AttackType {
	case model.AttackType_SKILL, model.AttackType_ULT :
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     CornucopiaBuff,
			Source:   mod.Owner(),
			Stats:    info.PropMap{prop.HealBoost: healAmt},
		})
	}
}
//remove buff after each "action"
func removeHealBuff(mod *modifier.ModifierInstance, e event.ActionEvent)  {
	mod.Engine().RemoveModifier(mod.Owner(), CornucopiaBuff)
}