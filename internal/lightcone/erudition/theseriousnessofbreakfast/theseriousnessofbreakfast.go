package theseriousnessofbreakfast

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Check key.Modifier = "the-seriousness-of-breakfast"
	Buff  key.Modifier = "the-seriousness-of-breakfast-buff"
)

//Increases the wearer's DMG by 12/15/18/21/24%.
//For every defeated enemy, the wearer's ATK increases by 4/5/6/7/8%, stacking up to 3 time(s).

func init() {
	lightcone.Register(key.TheSeriousnessofBreakfast, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: onTriggerDeath,
		},
	})

	modifier.Register(Buff, modifier.Config{
		StatusType:        model.StatusType_STATUS_BUFF,
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          3,
		CountAddWhenStack: 1,
		Listeners: modifier.Listeners{
			OnAdd: onAdd,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	dmgBonus := 0.09 + 0.03*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats:  info.PropMap{prop.AllDamagePercent: dmgBonus},
		State:  0.03 + 0.01*float64(lc.Imposition),
	})
}

func onTriggerDeath(mod *modifier.ModifierInstance, target key.TargetID) {
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   Buff,
		Source: mod.Owner(),
	})
}

func onAdd(mod *modifier.ModifierInstance) {
	mod.AddProperty(prop.ATKPercent, mod.Count()*mod.State().(float64))
}
