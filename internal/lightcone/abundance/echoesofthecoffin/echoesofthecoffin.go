package echoesofthecoffin

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Increases the wearer's ATK by 24%. After the wearer uses an attack,
// for each different enemy target the wearer hits, regenerates 3 Energy.
// Each attack can regenerate Energy up to 3 time(s) this way.
// After the wearer uses their Ultimate, all allies gain 12 SPD for 1 turn.

// DM :
// OnAfterAttack : retarget on atktargetlist max 3. set enemyCount. modifySPNew. enemyCount to 0
// OnAfterSkillUse : if Ult, addMod spdBuff,

// impl :
// OnAfterAttack : forEach hit enemies, add energy to holder. max 3
// OnAfterAction : if action == ult, buff allies spd.
// concl : 2 mods, 1 checker 1 ally spd buff.

func init() {
	lightcone.Register(key.EchoesoftheCoffin, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
