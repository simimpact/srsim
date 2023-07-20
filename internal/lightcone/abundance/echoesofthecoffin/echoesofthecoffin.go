package echoesofthecoffin

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
	echoes               = "echoes-of-the-coffin"
	spdBuff key.Modifier = "ecoes-of-the-coffin-spd"
)

type state struct {
	energyAmt float64
	speedAmt  float64
}

// Increases the wearer's ATK by 24%. After the wearer uses an attack,
// for each different enemy target the wearer hits, regenerates 3 Energy.
// Each attack can regenerate Energy up to 3 time(s) this way.
// After the wearer uses their Ultimate, all allies gain 12 SPD for 1 turn.

func init() {
	lightcone.Register(key.EchoesoftheCoffin, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	modifier.Register(echoes, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: addEnergyPerEnemyHit,
			OnAfterAction: buffAllySpdOnUlt,
		},
	})
	modifier.Register(spdBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	modState := state{
		energyAmt: 2.5 + 0.5*float64(lc.Imposition),
		speedAmt:  10.0 + 2.0*float64(lc.Imposition),
	}
	atkAmt := 0.20 + 0.04*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   echoes,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: atkAmt},
		State:  &modState,
	})
}

// add energy to lc holder based on # of targets hit. max 3.
func addEnergyPerEnemyHit(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(*state)
	enemyHit := float64(len(e.Targets))
	if enemyHit > 3 {
		enemyHit = 3
	}
	mod.Engine().ModifyEnergy(info.ModifyAttribute{
		Key:    echoes,
		Target: mod.Owner(),
		Source: mod.Owner(),
		Amount: state.energyAmt * enemyHit,
	})
}

func buffAllySpdOnUlt(mod *modifier.Instance, e event.ActionEnd) {
	// if action is not ult, return early
	if e.AttackType != model.AttackType_ULT {
		return
	}
	// if action is ult, add flat spd buff for all allies for 1 turn
	state := mod.State().(*state)
	for _, ally := range mod.Engine().Characters() {
		mod.Engine().AddModifier(ally, info.Modifier{
			Name:     spdBuff,
			Source:   mod.Owner(),
			Stats:    info.PropMap{prop.SPDFlat: state.speedAmt},
			Duration: 1,
		})
	}
}
