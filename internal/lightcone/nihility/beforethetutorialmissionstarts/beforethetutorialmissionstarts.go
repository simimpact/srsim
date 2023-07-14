package beforethetutorialmissionstarts

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
	mod key.Modifier = "before-tutorial-mission-starts"
)

// Increases the wearer's Effect Hit Rate by 20%.
// When the wearer attacks DEF-reduced enemies, regenerates 4 Energy.

// DM :
// OnAfterAttack : retarget() wtf? byContainBehaviorFLag : stat_defenceDown, includeLimbo
// max : 1
// -> TaskList : modifySPNew (wtf?) add value by some dynamic amt -> SkillTreeParam : ""(nil)
// OnStart : add _Main mod

func init() {
	lightcone.Register(key.BeforetheTutorialMissionStarts, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: addEnergy,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}

func addEnergy(mod *modifier.Instance, e event.AttackEnd) {

}
