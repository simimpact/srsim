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
	time key.Modifier = "time-waits-for-no-one"
	// key for extra dmg attack, not for modifier.
	extraDmg key.Attack = "time-waits-for-no-one-extra-damage"
)

type healRecorder struct {
	cooldown    int
	lastHealAmt float64
}

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
	modifier.Register(time, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1:        refreshCD,
			OnAfterAttack:   applyExtraDmg,
			OnAfterDealHeal: recordHealAmt,
		},
		CanModifySnapshot: true,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// initialize healRecorder struct.
	modState := healRecorder{
		cooldown:    1,
		lastHealAmt: 0,
	}
	// add in the HP + out. heal buffs. add struct pointer as state
	hpBuffAmt := 0.15 + 0.03*float64(lc.Imposition)
	outHealAmt := 0.10 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   time,
		Source: owner,
		Stats: info.PropMap{
			prop.HPPercent: hpBuffAmt,
			prop.HealBoost: outHealAmt,
		},
		State: &modState,
	})
}

// take struct pointer, modify cooldown value
func refreshCD(mod *modifier.Instance) {
	state := mod.State().(*healRecorder)
	state.cooldown = 1
}

// if cooldown = 1, Retarget(1 target), add dmg type pursued, byPureDamage, ele same as holder
func applyExtraDmg(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(*healRecorder)
	dmgAmt := mod.State().(*healRecorder).lastHealAmt
	if state.cooldown == 1 {
		validTargets := e.Targets
		chosenEnemy := mod.Engine().Retarget(info.Retarget{
			// targets are enemies hit by this atk. NOTE : confirm if onLimbo is included.
			Targets: validTargets,
			Max:     1,
		})
		// get lc holder's element
		holderInfo, _ := mod.Engine().CharacterInfo(mod.Owner())
		dmgType := holderInfo.Element
		mod.Engine().Attack(info.Attack{
			Key:          extraDmg,
			Targets:      chosenEnemy,
			Source:       mod.Owner(),
			AttackType:   model.AttackType_PURSUED,
			DamageType:   dmgType,
			DamageValue:  dmgAmt,
			AsPureDamage: true,
			// NOTE : might need to later change BaseDmg field to optional later
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: 0.0,
			},
			// this attack shouldn't call onHit listeners right?
			UseSnapshot: true,
		})
		// after apply added dmg, set cd to 0
		state.cooldown = 0
	}
}

// take struct pointer, modify lastHealAmt value.
func recordHealAmt(mod *modifier.Instance, e event.HealEnd) {
	state := mod.State().(*healRecorder)
	state.lastHealAmt = e.HealAmount
}
