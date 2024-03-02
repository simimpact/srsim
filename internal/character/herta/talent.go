package herta

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent = "herta-talent"
)

func init() {

}

func (c *char) initTalent() {
	c.engine.Events().HPChange.Subscribe(c.talentListener)

}

var (
	hertaCountInsert  = 0
	hertaCount        = 0
	hertaCountATK     = 0
	passiveCooldown   = 0
	scoringHertaCount = 0
)

func (c *char) talentListener(e event.HPChange) {
	// if herta insert count = 1, set herta atk count to 0, herta insert count to 0
	if hertaCountInsert == 1 {
		hertaCountATK = 0
		hertaCountInsert = 0
	}

	if e.NewHPRatio <= 0.5 {
		if c.engine.IsEnemy(e.Target) && passiveCooldown == 0 && !c.engine.HasBehaviorFlag(c.id, model.BehaviorFlag_STAT_CTRL) {
			if len(c.engine.Enemies()) > 0 {
				c.engine.Events().AttackEnd.Subscribe(c.talentAfterAttackListener)
				hertaCount += 1
			}

		}
	}

}

func (c *char) talentAfterAttackListener(e event.AttackEnd) {
	if hertaCountATK == 0 && hertaCount > 0 && c.engine.IsCharacter(e.Attacker) && !c.passiveFlag {
		if len(c.engine.Enemies()) > 0 {
			hertaCountATK = 1
			hertaCountInsert = 1
			c.talentInsertAttack()
		}
	}
}

func (c *char) talentInsertAttack() {
	c.passiveFlag = true
	hertaCountInsert = 0
	for hertaCount > 0 && len(c.engine.Enemies()) > 0 {
		hertaCount -= 1
		scoringHertaCount += 1

		if c.info.Eidolon >= 2 {

		}
	}

}
