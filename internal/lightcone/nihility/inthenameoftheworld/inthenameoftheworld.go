package inthenameoftheworld

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	world key.Modifier = "in-the-name-of-the-world"
)

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
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
