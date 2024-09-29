package luocha

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	abyssFlower       = "luocha-abyss-flower"
	Field             = "luocha-field"
	FieldHeal         = "luocha-field-heal"
	A4                = "luocha-a4"
	Insert            = "luocha-insert"
	InsertMark        = "luocha-insert-mark"
	DisableInsertMark = "luocha-disable-insert-mark"
	E1                = "luocha-e1"
	E4                = "luocha-e4"
)

type state struct {
	talentPer  float64
	talentFlat float64
}

func (c *char) init() {
	modifier.Register(abyssFlower, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          2,
		CountAddWhenStack: 1,
		Listeners: modifier.Listeners{
			OnAdd: checkStacks,
		},
	})

	modifier.Register(InsertMark, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: doInsert,
		},
	})

	modifier.Register(DisableInsertMark, modifier.Config{
		Listeners: modifier.Listeners{},
	})

	modifier.Register(Field, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnAdd:    addSubMods,
			OnRemove: removeSubMods,
		},
	})

	modifier.Register(FieldHeal, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: doFieldHeal,
		},
	})

	modifier.Register(E1, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E4, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_FATIGUE},
	})
}

func checkStacks(mod *modifier.Instance) {
	if mod.Count() >= 2 {
		mod.Engine().AddModifier(mod.Source(), info.Modifier{
			Name:   InsertMark,
			Source: mod.Owner(),
		})
	}
}

func doInsert(mod *modifier.Instance) {
	mod.Engine().RemoveModifier(mod.Source(), DisableInsertMark)

	luo := mod.Owner()
	ci, _ := mod.Engine().CharacterInfo(luo)
	mod.Engine().InsertAbility(info.Insert{
		Key:      Insert,
		Priority: info.CharBuffSelf,
		Execute: func() {
			mod.Engine().AddModifier(mod.Source(), info.Modifier{
				Name:   Field,
				Source: luo,
				State: state{
					talentPer:  talentPer[ci.TalentLevelIndex()],
					talentFlat: talentFlat[ci.TalentLevelIndex()],
				},
			})
		},
		Source:     luo,
		AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
	})

	mod.Engine().RemoveModifier(mod.Source(), abyssFlower)
	mod.Engine().RemoveModifier(mod.Source(), InsertMark)
}

func addSubMods(mod *modifier.Instance) {
	// apply sub modifiers as normal modifiers
	st := mod.State().(state)
	ci, _ := mod.Engine().CharacterInfo(mod.Owner())

	for _, trg := range mod.Engine().Characters() {
		// Talent and A4 heal
		mod.Engine().AddModifier(trg, info.Modifier{
			Name:   FieldHeal,
			Source: mod.Owner(),
			State: state{
				talentPer:  st.talentPer,
				talentFlat: st.talentFlat,
			},
		})

		// E1
		mod.Engine().AddModifier(trg, info.Modifier{
			Name:   E1,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.ATKPercent: 0.2},
		})
	}

	// E4
	if ci.Eidolon >= 4 {
		for _, trg := range mod.Engine().Enemies() {
			mod.Engine().AddModifier(trg, info.Modifier{
				Name:   E4,
				Source: mod.Owner(),
				Stats:  info.PropMap{prop.Fatigue: 0.12},
			})
		}
	}
}

func removeSubMods(mod *modifier.Instance) {
	// remove sub modifiers with a workaround
	ci, _ := mod.Engine().CharacterInfo(mod.Owner())

	for _, trg := range mod.Engine().Characters() {
		mod.Engine().RemoveModifierFromSource(trg, mod.Source(), FieldHeal)
		if ci.Eidolon >= 1 {
			mod.Engine().RemoveModifierFromSource(trg, mod.Source(), E1)
		}
	}

	if ci.Eidolon >= 4 {
		for _, trg := range mod.Engine().Enemies() {
			mod.Engine().RemoveModifierFromSource(trg, mod.Source(), E4)
		}
	}
}

func doFieldHeal(mod *modifier.Instance, e event.AttackEnd) {
	st := mod.State().(state)
	// heal self
	mod.Engine().Heal(info.Heal{
		Key:     FieldHeal,
		Targets: []key.TargetID{mod.Owner()},
		Source:  mod.Source(),
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_ATK: st.talentPer,
		},
		HealValue: st.talentFlat,
	})

	// heal other allies (A4)
	ci, _ := mod.Engine().CharacterInfo(mod.Source())
	if ci.Traces["102"] {
		mod.Engine().Heal(info.Heal{
			Key: FieldHeal,
			Targets: mod.Engine().Retarget(info.Retarget{
				Targets: mod.Engine().Characters(),
				Filter: func(target key.TargetID) bool {
					return target != mod.Owner()
				},
			}),
			Source: mod.Source(),
			BaseHeal: info.HealMap{
				model.HealFormula_BY_HEALER_ATK: 0.07,
			},
			HealValue: 93,
		})
	}
}
