package teststub

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func newCharacter(char *model.Character) *Character {
	res := &Character{
		model: char,
		key:   0,
		stat:  nil,
	}
	return res
}

func (c *Character) Stats() *info.Stats {
	return c.stat(c.key)
}

func (c *Character) ID() key.TargetID {
	return c.key
}
