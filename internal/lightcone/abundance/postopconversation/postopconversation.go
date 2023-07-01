package postopconversation

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

// Increases the wearer's Energy Regeneration Rate by 8% and
// increases Outgoing Healing when they use their Ultimate by 12%.
const (
	PostOpErrNCheck key.Modifier = "post-op-conversation-err-and-check"
	PostOpHealBuff  key.Modifier = "post-op-conversation-heal-buff"
)

func init() {
	lightcone.Register(key.PostOpConversation, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})

	// Implement checker here
	modifier.Register(PostOpErrNCheck, modifier.Config{
		Listeners: modifier.Listeners{
			// NOTE : DM uses OnBeforeDealHeal instead of OnBeforeAction. Might need change.
			OnBeforeAction: buffHealsOnUlt,
			OnAfterAction:  removeHealBuff,
		},
	})
	// The actual buff modifier goes here
	modifier.Register(PostOpHealBuff, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// OnStart : Add ERR Buff + Checker
	errAmt := 0.06 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   PostOpErrNCheck,
		Source: owner,
		Stats:  info.PropMap{prop.EnergyRegen: errAmt},
		State:  0.09 + 0.03*float64(lc.Imposition), // heal buff amt passed as state
	})
}

func buffHealsOnUlt(mod *modifier.Instance, e event.ActionStartEvent) {
	amt := mod.State().(float64)
	// NOTE : DM said onbeforeheal(unlike cornucopia which uses OnBeforeAction)
	// Once OnBeforeDealHeal has AttackType prop, need to change this.
	if e.AttackType == model.AttackType_ULT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   PostOpHealBuff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.HealBoost: amt},
		})
	}
}

// remove buff after each "action"
func removeHealBuff(mod *modifier.Instance, e event.ActionEvent) {
	mod.Engine().RemoveModifier(mod.Owner(), PostOpHealBuff)
}
