package timewaitsfornoone

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
	time     key.Modifier = "time-waits-for-no-one"
	extraDmg key.Modifier = "time-waits-for-no-one-extra-damage"
)

// Desc : Increases the wearer's Max HP by 18% and Outgoing Healing by 12%.
// When the wearer heals allies, record the amount of Outgoing Healing.
// When any ally launches an attack, a random attacked enemy takes Additional DMG
// equal to 36% of the recorded Outgoing Healing value.
// The type of this Additional DMG is of the same Type as the wearer's.
// This Additional DMG is not affected by other buffs, and can only occur 1 time per turn.

// Apparent modifiers :
// HP buff and outgoing healing -> on Create. add static
// OnAfterHeal -> record heal amt (pass it to a struct w/ pointer?)
// OnAfterAttack -> pick random 1 enemy, apply x% dmg based on healAmt + element same as holder
// 1 time per turn -> onBeforeTurn : refresh, pass 1 turn to struct or inherent duration?

// Datamine analysis :
// OnPhase1 : refresh 1-time-per-turn dmg adds
// onStack : add _Sub modifier on owner
// OnListenAfterAttack : if cooldown == 1 = call Retarget() : DamageByAttackProperty
// DamageTypeFromAttacker = T, indirect, attacktype pursued, byPureDmg, canTriggerLastKill
// -> set cooldown to 0.
// OnSnapshotCreate : add _Sub modifier (duplicate?)
// _Sub def : OnAfterDealHeal : store heal value
// OnStart : add _Main mod

// Conclusion on impl. :
// create base modifier. add hp and healboost. add listeners :
// NOTE : do we need to add _Sub as separate mod?
// -> OnPhase1 : refresh cooldown
// -> OnAfterAttack : if cooldown = 1 : Retarget(). apply dmg to chosen target. ele = holder ele.
// type pursued (how about the flags?), byPureDmg
// -> OnSnapshotCreate (?)
// OnAfterDealHeal : record heal value (read incessant rain impl.)

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
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}

func refreshCD(mod *modifier.Instance) {

}

func applyExtraDmg(mod *modifier.Instance, e event.AttackEnd) {

}

func recordHealAmt(mod *modifier.Instance, e event.HealEnd) {

}
