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

type healRecorder struct {
	onCooldown    bool
	recordedHeals float64
	damageMult    float64
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
			OnAfterDealHeal: recordHealAmt,
		},
		CanModifySnapshot: true,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// initialize healRecorder struct.
	modState := healRecorder{
		onCooldown:    false,
		recordedHeals: 0.0,
		damageMult:    0.30 + 0.06*float64(lc.Imposition),
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

	// event subscriber to atkEnd by all chars -> bypass if atker is enemy.
	engine.Events().AttackEnd.Subscribe(func(e event.AttackEnd) {
		// fetch modifier instance attached to lc owner (HOW?)
		mod := engine.GetModifiers(owner, time)[0]
		// apply extra dmg if attacker is not enemy and not on cd
		if (engine.IsCharacter(e.Attacker) && !mod.State.(*healRecorder).onCooldown) {
			// perform attack, reset dmgAmt, and put mod on CD
			applyExtraDmg(mod, e)
		}
	})
}

// take struct pointer, modify onCooldown value
func refreshCD(mod *modifier.Instance) {
	state := mod.State().(*healRecorder)
	state.onCooldown = false
}

// if onCooldown = 1, Retarget(1 target), add dmg type pursued, byPureDamage, ele same as holder
func applyExtraDmg(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(*healRecorder)
	dmgAmt := mod.State().(*healRecorder).recordedHeals
	if !state.onCooldown {
		chosenEnemy := mod.Engine().Retarget(info.Retarget{
			// targets are enemies hit by this atk.
			Targets: e.Targets,
			Max:     1,
		})
		// get lc holder's info to fetch their element
		holderInfo, _ := mod.Engine().CharacterInfo(mod.Owner())
		mod.Engine().Attack(info.Attack{
			Key:          time,
			Targets:      chosenEnemy,
			Source:       mod.Owner(),
			AttackType:   model.AttackType_PURSUED,
			DamageType:   holderInfo.Element,
			DamageValue:  dmgAmt,
			AsPureDamage: true,
			// NOTE : might need to later change BaseDmg field to optional later
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: 0.0,
			},
			// this attack shouldn't call onHit listeners right?
			UseSnapshot: true,
		})
		// after apply added dmg, reset cd and recordedHeals
		state.onCooldown = true
		state.recordedHeals = 0.0
	}
}

// take struct pointer, modify recordedHeals value.
func recordHealAmt(mod *modifier.Instance, e event.HealEnd) {
	state := mod.State().(*healRecorder)
	// sum all recorded heals
	state.recordedHeals += e.HealAmount
}
