package ruanmei

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult           = "ruanmei-ult"
	UltResPenAlly = "ruanmei-ult-res-pen-ally"
	UltBuffAlly   = "ruanmei-ult-buff-ally"
	UltDebuff     = "ruanmei-thanatopium-rebloom"
	UltDebuffCD   = "ruanmei-thanatopium-rebloom-cooldown"
)

func init() {
	// 2 mods summarized to 1 mod that represents both logic and visual/Buff status type
	modifier.Register(Ult, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		TickMoment: modifier.ModifierPhase1End,
		Listeners: modifier.Listeners{
			OnAdd:    addResPen,
			OnRemove: removeUltMods,
		},
	})
	modifier.Register(UltResPenAlly, modifier.Config{
		Stacking: modifier.Refresh,
	})
	// Purely for visual/Buff status type
	modifier.Register(UltBuffAlly, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(UltDebuff, modifier.Config{
		Listeners: modifier.Listeners{
			// OnAllowAction: removeReset
			// OnListenTurnEnd: removeSelf
			// OnHPChange: removeCDAndSelf
			OnEndBreak: doUltImprint,
			// OnDispel: removeCD,
			// OnBreakExtendAnim: doUltImprintWithoutRemove (??)
		},
	})
	modifier.Register(UltDebuffCD, modifier.Config{
		Listeners: modifier.Listeners{
			OnLimboWaitHeal: func(mod *modifier.Instance) bool {
				mod.RemoveSelf()
				return false
			},
			OnEndBreak: func(mod *modifier.Instance) {
				mod.RemoveSelf()
			},
		},
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	durationUlt := 2
	if c.info.Eidolon >= 6 {
		durationUlt += 1
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:     Ult,
		Source:   c.id,
		Duration: durationUlt,
		State:    ultResPen[c.info.UltLevelIndex()],
	})
	if c.info.Eidolon >= 1 {
		for _, trg := range c.engine.Characters() {
			c.engine.AddModifier(trg, info.Modifier{
				Name:   E1,
				Source: c.id,
			})
		}
	}
	// Subscribe to CharactersAdded for E1 and AttackEnd for Debuff
}

// Need to implement AllDamagePEN
func addResPen(mod *modifier.Instance) {
	for _, trg := range mod.Engine().Characters() {
		if trg == mod.Owner() {
			mod.AddProperty(prop.IcePEN, mod.State().(float64))
		} else {
			mod.Engine().AddModifier(trg, info.Modifier{
				Name:   UltResPenAlly,
				Source: mod.Owner(),
				Stats:  info.PropMap{prop.IcePEN: mod.State().(float64)},
			})
		}
	}
}

func removeUltMods(mod *modifier.Instance) {
	for _, trg := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(trg, UltResPenAlly)
		mod.Engine().RemoveModifier(trg, UltBuffAlly)
		mod.Engine().RemoveModifier(trg, E1)
	}
}

func doUltImprint(mod *modifier.Instance) {

}
