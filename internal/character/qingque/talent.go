package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent key.Modifier = "qingque-talent"
)

type talentState struct {
	hiddenHandAtkBoost float64
}

func init() {
	modifier.Register(Talent, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) talentTurnStartListener(e event.TurnStartEvent) {
	if c.engine.IsCharacter(e.Active) {
		c.drawTile()
	}
	if e.Active == c.id {
		if c.tiles[0] == 4 {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   Talent,
				Source: c.id,
				State: talentState{
					hiddenHandAtkBoost: talent[c.info.AbilityLevel.Talent],
				},
			})
		}
	}
}

func (c *char) swap(pos1 int, pos2 int) {
	temp := c.tiles[pos1]
	tempSuit := c.suits[pos1]
	c.tiles[pos1] = c.tiles[pos2]
	c.suits[pos1] = c.suits[pos2]
	c.tiles[pos2] = temp
	c.suits[pos2] = tempSuit
}

func (c *char) drawTile() {
	s1, s2, s3 := c.tiles[0], c.tiles[1], c.tiles[2]
	startingTiles := s1 + s2 + s3
	drawn := c.engine.Rand().Intn(3)
	if c.tiles[drawn] == 0 {
		toUse := c.engine.Rand().Intn(len(c.unusedSuits))
		last := (len(c.unusedSuits) - 1)
		c.suits[drawn] = c.unusedSuits[toUse]
		c.unusedSuits[toUse] = c.unusedSuits[last]
		c.unusedSuits = c.unusedSuits[:last]
	}
	c.tiles[drawn] += 1
	if c.tiles[0] >= c.tiles[1] && c.tiles[1] >= c.tiles[2] {

	} else if c.tiles[1] > c.tiles[0] {
		c.swap(0, 1)
	} else if c.tiles[2] > c.tiles[0] {
		c.swap(0, 2)
	} else {
		c.swap(1, 2)
	}
	if startingTiles == 4 {
		c.discardTile()
	}
}
func (c *char) discardTile() {
	if c.tiles[2] != 0 {
		if c.tiles[1] == c.tiles[2] {
			if c.tiles[2] == c.tiles[0] {
				switch c.engine.Rand().Intn(3) {
				case 0:
					c.suits[0] = c.suits[2]
				case 1:
					c.suits[1] = c.suits[2]
				}
			} else if c.engine.Rand().Intn(2) == 0 {
				c.suits[1] = c.suits[2]
			}
		}
		c.tiles[2] -= 1
		if c.tiles[2] == 0 {
			c.unusedSuits = c.unusedSuits[:(len(c.unusedSuits) + 1)]
			c.unusedSuits[len(c.unusedSuits)-1] = c.suits[2]
			c.suits[2] = ""
		}
	} else if c.tiles[1] != 0 {
		if c.tiles[0] == c.tiles[1] && c.engine.Rand().Intn(2) == 0 {
			c.suits[0] = c.suits[1]
		}
		c.tiles[1] -= 1
		if c.tiles[1] == 0 {
			c.unusedSuits = c.unusedSuits[:(len(c.unusedSuits) + 1)]
			c.unusedSuits[len(c.unusedSuits)-1] = c.suits[1]
			c.suits[1] = ""
		}
	} else if c.tiles[0] != 0 {
		c.tiles[0] -= 1
		if c.tiles[0] == 0 {
			c.unusedSuits = c.unusedSuits[:(len(c.unusedSuits) + 1)]
			c.unusedSuits[len(c.unusedSuits)-1] = c.suits[0]
			c.suits[0] = ""
		}
	} else {
		// This should really never happen
	}
}
