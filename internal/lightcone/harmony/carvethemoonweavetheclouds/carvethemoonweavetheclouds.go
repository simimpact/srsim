package carvethemoonweavetheclouds

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	carveCheck key.Modifier = "carve-the-moon-weave-the-cloud"
	carveAtk   key.Modifier = "carve-the-moon-weave-the-cloud-atk"
	carveCDmg  key.Modifier = "carve-the-moon-weave-the-cloud-cdmg"
	carveERR   key.Modifier = "carve-the-moon-weave-the-cloud-err"
)

// At the start of the battle and whenever the wearer's turn begins,
// one of the following effects is applied randomly: All allies' ATK increases by 10%,
// all allies' CRIT DMG increases by 12%, or all allies' Energy Regeneration Rate
// increases by 6%. The applied effect cannot be identical to the last effect applied,
// and will replace the previous effect. The applied effect will be removed when the wearer
// has been knocked down. Effects of the similar type cannot be stacked.

func init() {
	lightcone.Register(key.CarvetheMoonWeavetheClouds, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})
	modifier.Register(carveCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeDying: removeBuffs,
			OnPhase1:      applyTeamBuffRandomly,
		},
	})
	modifier.Register(carveAtk, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(carveCDmg, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(carveERR, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkAmt := 0.075 + 0.025*float64(lc.Imposition)
	cDmgAmt := 0.09 + 0.03*float64(lc.Imposition)
	ErrAmt := 0.045 + 0.015*float64(lc.Imposition)

	// add checker. apply random team buff once onBattleStart.
	engine.AddModifier(owner, info.Modifier{
		Name:   carveCheck,
		Source: owner,
		// TODO : add state struct.
		State: 0.0,
	})
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		for _, char := range engine.Characters() {
			// take prev buff #, remove that instance, randomly add the other 2 buffs
		}
	})
}

func removeBuffs(mod *modifier.Instance) {

}

func applyTeamBuffRandomly(mod *modifier.Instance) {

}
