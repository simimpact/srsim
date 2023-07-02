package quidproquo

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	QPQCheck key.Modifier = "quid-pro-quo"
)

// At the start of the wearer's turn, regenerates 8 Energy for a randomly chosen ally
// (excluding the wearer) whose current Energy is lower than 50%.
func init() {
	lightcone.Register(key.QuidProQuo, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	// OnPhase1. checker. refill 1 char's energy randomly. (condition : <50% + not LC holder)
	modifier.Register(QPQCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: randomlyAddEnergy,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   QPQCheck,
		Source: owner,
		State:  6 + 2*float64(lc.Imposition),
	})
}

func randomlyAddEnergy(mod *modifier.Instance) {
	allyList := mod.Engine().Characters()
	amt := mod.State().(float64)
	var validAllyList []key.TargetID

	for _, char := range allyList {
		// check if energy is <50% and current char isn't LC's holder.
		if mod.Engine().EnergyRatio(char) < 0.5 && char != mod.Owner() {
			validAllyList = append(validAllyList, char)
		}
	}
	// add in checker to add energy only if validAllyList isn't empty
	if validAllyList != nil {
		// randomly choose 1 char to add energy to from validAllyList.
		chosenOne := validAllyList[mod.Engine().Rand().Intn(len(validAllyList))]
		mod.Engine().ModifyEnergy(chosenOne, amt)
	}
}
