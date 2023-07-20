package inthenameoftheworld

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
	world key.Modifier = "in-the-name-of-the-world"
)

type state struct {
	dmgNAtkAmt float64
	ehrAmt     float64
}

// Increases the wearer's DMG to debuffed enemies by 24%.
// When the wearer uses their Skill, the Effect Hit Rate for this attack increases by 18%,
// and ATK increases by 24%.

// DM :
// OnBeforeHitALl : compare ByStatusCount >= 1, ModifyDmgRatio -> AllDmgTypeAddedRatio
// OnBeforeSkillUse : Skill -> add _Sub
// OnAfterSkillUse : remove _Sub
// _Sub : AttackAddedRatio + StatusProbabilityBase buffs

func init() {
	lightcone.Register(key.IntheNameoftheWorld, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(world, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: boostDmgOnDebuffed,
			OnBeforeAction: boostEhrNAtkOnSkill,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	modState := state{
		dmgNAtkAmt: 0.20 + 0.04*float64(lc.Imposition),
		ehrAmt:     0.15 + 0.03*float64(lc.Imposition),
	}
	engine.AddModifier(owner, info.Modifier{
		Name:   world,
		Source: owner,
		State:  &modState,
	})
}

// if debuff > 0, boost hit dmg
func boostDmgOnDebuffed(mod *modifier.Instance, e event.HitStart) {

}

// if action == skill, boost ehr + atk
func boostEhrNAtkOnSkill(mod *modifier.Instance, e event.ActionStart) {

}