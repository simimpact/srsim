package eyesoftheprey

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod key.Modifier = "eyes-of-the-prey"
)

// Increases the wearer's Effect Hit Rate by 20%/25%/30%/35%/40% and increases DoT by 24%/30%/36%/42%/48%.
func init() {
	lightcone.Register(key.EyesofthePrey, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	ehrAmt := 0.15 + 0.05*float64(lc.Imposition)
	dotAmt := 0.18 + 0.06*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.EffectHitRate: ehrAmt, prop.DOTDamagePercent: dotAmt},
	})
}
