package beforedawn

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

//Increases the wearer's CRIT DMG by x%. Increases the wearer's Skill and Ultimate DMG by x%.
//After the wearer uses their Skill or Ultimate, they gain Somnus Corpus.
//Upon triggering a follow-up attack, Somnus Corpus will be consumed and the follow-up attack DMG increases by x%.
const (
	BeforeDawn   key.Modifier = "before_dawn"
	SomnusCorpus key.Modifier = "somnus_corpus"
)

type somnusState struct {
	amt  float64
	used bool
}

//
func init() {
	lightcone.Register(key.BeforeDawn, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})

	modifier.Register(BeforeDawn, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit:   onBeforeHit,
			OnAfterAction: onAfterAction,
		},
	})

	modifier.Register(SomnusCorpus, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnBeforeHit:   onBeforeHitSomnus,
			OnAfterAttack: onAfterAttack,
		},
	})
}

//Add crit dmg modifier
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.30 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   BeforeDawn,
		Source: owner,
		Stats:  info.PropMap{prop.CritDMG: amt},
		State:  float64(lc.Imposition),
	})

}

//BeforeHit if its ult or skill add the dmg% to that hit
func onBeforeHit(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	if e.Hit.AttackType == model.AttackType_ULT ||
		e.Hit.AttackType == model.AttackType_SKILL {
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.15+0.03*mod.State().(float64))
	}
}

//Beforehit if its follow and it has the SomnusCorpusMod add the dmg% to that hit and change used to true
func onBeforeHitSomnus(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	state := mod.State().(*somnusState)
	if e.Hit.AttackType == model.AttackType_INSERT {
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, state.amt)
		state.used = true
	}
}

//after attack if SomnusCorpMod is used, remove self.
func onAfterAttack(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	state := mod.State().(*somnusState)
	if state.used {
		mod.RemoveSelf()
	}
}

//AfterAction if its ult or skill add the SomnusCorpusMod
func onAfterAction(mod *modifier.ModifierInstance, e event.ActionEvent) {
	if e.AttackType == model.AttackType_ULT ||
		e.AttackType == model.AttackType_SKILL {
		amt := mod.State().(float64)
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   SomnusCorpus,
			Source: mod.Owner(),
			State: &somnusState{
				amt: 0.40 + 0.08*amt,
			},
		})
	}
}
