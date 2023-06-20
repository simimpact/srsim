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
	QPQCheck key.Modifier = "quid-pro-quo-check"
)

//At the start of the wearer's turn, regenerates 8 Energy for a randomly chosen ally
//(excluding the wearer) whose current Energy is lower than 50%.
func init() { 
	lightcone.Register(key.QuidProQuo, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	//OnPhase1. checker. refill 1 char's energy randomly. (condition : <50%)
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
		State:  6 + 2 * float64(lc.Imposition),
	})
}

func randomlyAddEnergy(mod *modifier.ModifierInstance) {
	allyList := mod.Engine().Characters()
	amt := 	mod.State().(float64)
	var validAllyList []key.TargetID
	
	for _, char := range allyList{
		//check if energy is <50%.
		if (mod.Engine().Energy(char) < mod.Engine().Stats(char).MaxEnergy() / 2) {
			validAllyList = append(validAllyList, char)
		}
	}
	//randomly choose 1 char to add energy to from validAllyList.
	chosenOne := validAllyList[mod.Engine().Rand().Intn(len(validAllyList))]
	mod.Engine().ModifyEnergy(chosenOne, amt) 
}