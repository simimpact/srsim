package timewaitsfornoone

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
	time = "time-waits-for-no-one"
)

// Desc : Increases the wearer's Max HP by 18% and Outgoing Healing by 12%.
// When the wearer heals allies, record the amount of Outgoing Healing.
// When any ally launches an attack, a random attacked enemy takes Additional DMG
// equal to 36% of the recorded Outgoing Healing value.
// The type of this Additional DMG is of the same Type as the wearer's.
// This Additional DMG is not affected by other buffs, and can only occur 1 time per turn.

func init() {
	lightcone.Register(key.TimeWaitsforNoOne, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	// vacated for switch to event subscription
	modifier.Register(time, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// modifier now only added the hp and healboost buffs
	engine.AddModifier(owner, info.Modifier{
		Name:   time,
		Source: owner,
		Stats: info.PropMap{
			prop.HPPercent: 0.15 + 0.03*float64(lc.Imposition),
			prop.HealBoost: 0.10 + 0.02*float64(lc.Imposition),
		},
	})

	extraDmgMult := 0.30 + 0.06*float64(lc.Imposition)
	healAmt := 0.0

	// event subscriber to atkEnd by all chars -> bypass if atker is enemy.
	engine.Events().AttackEnd.Subscribe(func(e event.AttackEnd) {
		// fetch modifier instance attached to lc owner
		mod := engine.GetModifiers(owner, time)[0]
		if engine.IsCharacter(e.Attacker) && !engine.HasModifier(owner, cd) {
			// perform attack, reset dmgAmt, and put mod on CD
			applyExtraDmg(engine, e.Targets, mod.Source, healAmt *  extraDmgMult)
		}
	})
}

func refreshCD(mod *modifier.Instance) {
	state := mod.State().(*healRecorder)
	state.onCooldown = false
}

// if onCooldown = 1, Retarget(1 target), add dmg type pursued, byPureDamage, ele same as holder
func applyExtraDmg(engine engine.Engine, targets []key.TargetID, source key.TargetID, dmgAmt float64) {
	chosenEnemy := engine.Retarget(info.Retarget{
		Targets: targets,
		Max:     1,
	})
	// get lc holder's info to fetch their element
	holderInfo, _ := engine.CharacterInfo(source)
	engine.Attack(info.Attack{
		Key:          time,
		Targets:      chosenEnemy,
		Source:       source,
		AttackType:   model.AttackType_PURSUED,
		DamageType:   holderInfo.Element,
		DamageValue:  dmgAmt,
		AsPureDamage: true,
		// NOTE : might need to later change BaseDmg field to optional later
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.0,
		},
		UseSnapshot: true,
	})
	// after apply added dmg, reset cd and recordedHeals
	state.onCooldown = true
	state.recordedHeals = 0.0
}

func recordHealAmt(mod *modifier.Instance, e event.HealEnd) {
	state := mod.State().(*healRecorder)
	// sum all recorded heals
	state.recordedHeals += e.HealAmount
}
