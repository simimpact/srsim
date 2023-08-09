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
	errAmt := 0.045 + 0.015*float64(lc.Imposition)
	// init state. populate with default vals.
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

	// add checker.
	engine.AddModifier(owner, info.Modifier{
		Name:   carveCheck,
		Source: owner,
		State:  &modState,
	})

	// apply random team buff once onBattleStart.
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		// randomly pick between 3 available buffs.
		chosenBuff := engine.Rand().Intn(len(modState.buffList))

		// add picked buff to all chars
		for _, char := range engine.Characters() {
			engine.AddModifier(char, info.Modifier{
				Name:   modState.buffList[chosenBuff].name,
				Source: owner,
				Stats:  modState.buffList[chosenBuff].statsField,
			})
		}
		// track current applied buff
		modState.currentBuff = chosenBuff
	})
}

func removeBuffsOnDeath(mod *modifier.Instance) {
	removeBuffs(mod.Engine().Characters(), mod.Source(), mod.Engine())
}

// REWRITE : DRY : make remove buffs its own function.
func removeBuffs(characters []key.TargetID, source key.TargetID, engine engine.Engine) {
	for _, char := range characters {
		engine.RemoveModifierFromSource(char, source, carveAtk)
		engine.RemoveModifierFromSource(char, source, carveCDmg)
		engine.RemoveModifierFromSource(char, source, carveERR)
	}
}

func applyTeamBuffRandomly(mod *modifier.Instance) {
	// remove currently applied buff
	removeBuffs(mod.Engine().Characters(), mod.Source(), mod.Engine())

	// create modstate slice (- previously applied buff)
	state := mod.State().(*state)
	var validBuffs []singleBuff
	for _, buff := range state.buffList {
		if buff.name != state.buffList[state.currentBuff].name {
			validBuffs = append(validBuffs, buff)
		}
	}

	// pick between qualified buffs
	chosenBuff := mod.Engine().Rand().Intn(len(validBuffs))

	// apply chosen buff to all chars
	for _, char := range mod.Engine().Characters() {
		mod.Engine().AddModifier(char, info.Modifier{
			Name:   state.buffList[chosenBuff].name,
			Source: mod.Owner(),
			Stats:  state.buffList[chosenBuff].statsField,
		})
	}
	// track current applied buff
	state.currentBuff = chosenBuff
}
