package finefruit

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// At the start of the battle, immediately regenerates 6/7.5/9/10.5/12 Energy for all allies.
func init() {
	lightcone.Register(key.FineFruit, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.Events().BattleStart.Subscribe(func(event event.BattleStartEvent) {
		for char := range event.CharInfo {
			engine.ModifyEnergy(char, 4.5+1.5*float64(lc.Rank))
		}
	})
}
