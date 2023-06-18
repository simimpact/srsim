package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent key.Modifier = "qingque-talent"
)

func init() {
	modifier.Register(Talent, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) talentTurnStartListener(e event.TurnStartEvent) {
	if c.engine.IsCharacter(e.Active) {
		c.drawTile()
	}
	if e.Active == c.id && c.tiles[0] == 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   Talent,
			Source: c.id,
			Stats:  info.PropMap{prop.ATKPercent: talent[c.info.TalentLevelIndex()]},
		})
	}
}
