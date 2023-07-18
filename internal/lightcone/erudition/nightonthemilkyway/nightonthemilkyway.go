package nightonthemilkyway

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)
const (
	NightontheMilkyWay = "night-on-the-milky-way"
)
type Amts struct {
	dmg float64
	atkB float64
}
func init() {
	lightcone.Register(key.NightontheMilkyWay, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
	//listen for break and before hits 
	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnListenBreak:  onListenBreak,
			OnBeforeHit: onBeforeHit,
		},
	})

	
	modifier.Register(Target, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

//For every enemy on the field, increases the wearer's ATK by 9.0~15.0%, up to 5 stacks.
//When an enemy is inflicted with Weakness Break, the DMG dealt by the wearer increases by 30~50% for 1 turn.
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkAmt := 0.075 + 0.015*float64(lc.Imposition)
	dmgAmt := 0.25 + 0.05*float64(lc.Imposition)
	//Using similar logic to only silence remains where it may not work for trotters - can try an implementation using an
	//onbeforehit mod but i wasnt sure how engine would work with it so i just used this
	engine.Events().EnemiesAdded.Subscribe(func(e event.EnemiesAdded) {
		atkBuff(engine, owner, atkAmt)
	})

	engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		atkBuff(engine, owner, atkAmt)
	})
	engine.AddModifier(owner, info.Modifier{
		Name:   nightonthemilkyway,
		Source: owner,
		State:  Amts{atkB: atkAmt, dmg: dmgAmt},
	})
}

func atkBuff(engine engine.Engine, owner key.TargetID, amt float64){
	//get count of enemies capping at 5
	enemCount := len(engine.Enemies())
	if enemCount > 5 {
		enemCount = 5
	}
	buffAtkAmt := amt * enemCount
	//buff atk 
	engine.AddModifier(owner, info.Modifier{
		Name:   NightontheMilkyWay,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: buffAtkAmt},
	})
}

//get damage bonus 
func onListenBreak(mod *modifier.Instance, target key.TargetID){
	dmgBonus := mod.State().(state).dmg
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   BuffDmgBonus,
		Source: mod.Owner(),
		Stats:  info.PropMap{prop.AllDamagePercent: dmgBonus},
	})
}
