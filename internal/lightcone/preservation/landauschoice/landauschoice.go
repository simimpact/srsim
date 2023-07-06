package landauschoice

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod key.Modifier = "landaus-choice"
)

// Desc : The wearer is more likely to be attacked, and DMG taken is reduced by 18%.
// DM Listeners : OnStack = aggroAddedRatio + allDamageReduce, OnStart = addModifier

func init() {
	lightcone.Register(key.LandausChoice, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})
	modifier.Register(mod, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	dmgRedAmt := 0.14 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats: info.PropMap{
			prop.AllDamageReduce: dmgRedAmt,
			prop.AggroPercent:    2,
		},
	})
}
