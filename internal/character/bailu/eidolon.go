package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// E4 : Every healing provided by the Skill makes the recipient deal 10% more DMG for 2 turn(s).
//      This effect can stack up to 3 time(s).
// E6 : Bailu can heal allies who received a killing blow 1 more time(s) in a single battle.

const (
	E2 key.Modifier = "bailu-e2"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) initEidolons() {
	// add mods/set values here
}
