package carvethemoonweavetheclouds

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
	carveCheck key.Modifier = "carve-the-moon-weave-the-cloud"
	carveAtk   key.Modifier = "carve-the-moon-weave-the-cloud-atk"
	carveCDmg  key.Modifier = "carve-the-moon-weave-the-cloud-crit-dmg"
	carveERR   key.Modifier = "carve-the-moon-weave-the-cloud-err"
)

type singleBuff struct {
	name       key.Modifier
	statsField info.PropMap
}
type state struct {
	currentBuff int
	buffList    []singleBuff
}

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
			OnBeforeDying: removeBuffsOnDeath,
			OnPhase1:      applyTeamBuffRandomly,
		},
	})
	modifier.Register(carveAtk, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
	modifier.Register(carveCDmg, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
	modifier.Register(carveERR, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkAmt := 0.075 + 0.025*float64(lc.Imposition)
	cDmgAmt := 0.09 + 0.03*float64(lc.Imposition)
	errAmt := 0.045 + 0.015*float64(lc.Imposition)
	modState := state{
		currentBuff: 0,
		buffList: []singleBuff{
			{
				name:       carveAtk,
				statsField: info.PropMap{prop.ATKPercent: atkAmt},
			},
			{
				name:       carveCDmg,
				statsField: info.PropMap{prop.CritDMG: cDmgAmt},
			},
			{
				name:       carveERR,
				statsField: info.PropMap{prop.EnergyRegen: errAmt},
			},
		},
	}

	engine.AddModifier(owner, info.Modifier{
		Name:   carveCheck,
		Source: owner,
		State:  &modState,
	})

	// apply random team buff once onBattleStart.
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		randomlyApplyTeamBuff(&modState, engine, modState.buffList, owner)
	})
}

func removeBuffsOnDeath(mod *modifier.Instance) {
	state := mod.State().(*state)
	currentBuff := state.buffList[state.currentBuff].name
	removeBuffs(mod.Engine().Characters(), mod.Source(), currentBuff, mod.Engine())
}

func applyTeamBuffRandomly(mod *modifier.Instance) {
	// remove currently applied buff
	state := mod.State().(*state)
	currentBuff := state.buffList[state.currentBuff].name
	removeBuffs(mod.Engine().Characters(), mod.Source(), currentBuff, mod.Engine())

	// make possible buffs list
	validBuffs := make([]singleBuff, 0, 3)
	for _, buff := range state.buffList {
		if buff.name != currentBuff {
			validBuffs = append(validBuffs, buff)
		}
	}

	randomlyApplyTeamBuff(state, mod.Engine(), validBuffs, mod.Source())
}

func removeBuffs(characters []key.TargetID, source key.TargetID, currentBuff key.Modifier, engine engine.Engine) {
	for _, char := range characters {
		engine.RemoveModifierFromSource(char, source, currentBuff)
	}
}

func randomlyApplyTeamBuff(state *state, engine engine.Engine, validBuffs []singleBuff, source key.TargetID) {
	// randomly pick between valid buffs.
	chosenBuff := engine.Rand().Intn(len(validBuffs))

	// add picked buff to all chars
	for _, char := range engine.Characters() {
		engine.AddModifier(char, info.Modifier{
			Name:   validBuffs[chosenBuff].name,
			Source: source,
			Stats:  validBuffs[chosenBuff].statsField,
		})
	}
	// track current applied buff
	state.currentBuff = chosenBuff
}
