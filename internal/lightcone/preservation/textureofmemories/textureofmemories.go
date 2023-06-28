package textureofmemories

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
	mod           key.Modifier = "texture"
	modshield     key.Shield   = "texture-shield"
	modcd         key.Modifier = "texture-cooldown"
	modshieldbuff key.Modifier = "texture-shield-buff"
)

type State struct {
	shieldAmt float64
	dmgRes    float64
}

// Increases the wearer's Effect RES by 8%/10%/12%/14%/16%. If the wearer is
// attacked and has no Shield, they gain a Shield equal to
// 16%/20%/24%/28%/32% of their Max HP for 2 turn(s). This effect can only be
// triggered once every 3 turn(s). If the wearer has a Shield when attacked,
// the DMG they receive decreases by 12%/15%/18%/21%/24%.
func init() {
	lightcone.Register(key.TextureofMemories, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeBeingHitAll:  onBeforeBeingHitAll,
			OnAfterBeingAttacked: onAfterBeingAttacked,
		},
	})

	modifier.Register(modshieldbuff, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:    shieldBuffOnAdd,
			OnRemove: shieldBuffOnRemove,
		},
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(modcd, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// wearer Effect RES
	effResAmt := 0.06 + 0.02*float64(lc.Imposition) // probably can omit from state
	shieldAmt := 0.12 + 0.04*float64(lc.Imposition)
	dmgRes := 0.09 + 0.03*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.EffectRES: effResAmt},
		State:  State{shieldAmt, dmgRes},
	})
}

// wearer has Shield before attacked
func onBeforeBeingHitAll(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	if mod.Engine().IsShielded(mod.Owner()) {
		state := mod.State().(State)
		e.Hit.Defender.AddProperty(prop.AllDamageReduce, state.dmgRes)
	}
}

// wearer doesn't have Shield after attacked
func onAfterBeingAttacked(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	isShielded := mod.Engine().IsShielded(mod.Owner())
	isOnCd := mod.Engine().HasModifier(mod.Owner(), modcd)

	// no shield + lc effect is off-cd
	if !isShielded && !isOnCd {
		// apply shield as buff with 2t duration
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     modshieldbuff,
			Source:   mod.Owner(),
			Duration: 2,
		})

		// apply cd modifier
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     modcd,
			Source:   mod.Owner(),
			Duration: 3,
		})
	}
}

func shieldBuffOnAdd(mod *modifier.ModifierInstance) {
	state := mod.State().(State)
	// apply shield
	mod.Engine().AddShield(modshield, info.Shield{
		Source:     mod.Owner(),
		Target:     mod.Owner(),
		BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_MAX_HP: state.shieldAmt},
	})
}

func shieldBuffOnRemove(mod *modifier.ModifierInstance) {
	mod.Engine().RemoveShield(modshield, mod.Owner())
}
