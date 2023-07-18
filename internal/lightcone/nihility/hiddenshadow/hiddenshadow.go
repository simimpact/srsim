package hiddenshadow

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// DESC : After using Skill, the wearer's next Basic ATK deals
// Additional DMG equal to 60% of ATK to the target enemy.

// DM :
// OnAfterSkillUse : if skill, add _Sub mod
// _Sub def : OnAfterAttack = if flag = 1, retarget Max 1, includeLimbo.
// -> deal extra pursued dmg, param x% of lc holder's ATK.
// OnBeforeSkillUse : set flag to 1.
// OnAfterSkillUse : delete flag definition + modifier(wtf?)
// OnStart : add _Main mod.

func init() {
	lightcone.Register(key.HiddenShadow, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
