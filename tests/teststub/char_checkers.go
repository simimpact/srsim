package teststub

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func newCharacter(char *model.Character) *Character {
	res := &Character{
		t:     nil,
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

// Because we use floating points, our Require will only ask that the two values be close enough
func require(expected, actual float64) bool {
	if actual == 0 { // division by 0
		return expected == 0
	}
	diff := expected / actual
	return 0.99 < diff && diff < 1.01
}

func (c *Character) Equal(statKey prop.Property, expected float64) {
	if require(expected, c.Stats().GetProperty(statKey)) {
		return
	}
	c.t.Fatalf("Assert Fail, stat %s, expected %f, actual %f", statKey.String(), expected, c.Stats().GetProperty(statKey))
}

func (c *Character) Greater(statKey prop.Property, expected float64) {
	if expected > c.Stats().GetProperty(statKey) {
		return
	}
	c.t.Fatalf("Assert Fail, stat %s, expected %f, actual %f", statKey.String(), expected, c.Stats().GetProperty(statKey))
}

func (c *Character) Less(statKey prop.Property, expected float64) {
	if expected < c.Stats().GetProperty(statKey) {
		return
	}
	c.t.Fatalf("Assert Fail, stat %s, expected %f, actual %f", statKey.String(), expected, c.Stats().GetProperty(statKey))
}

func (c *Character) AssertEnergy(expected float64) {
	if require(expected, c.Stats().Energy()) {
		return
	}
	c.t.Fatalf("Assert Fail, stat Energy, expected %f, actual %f", expected, c.Stats().Energy())
}

func (c *Character) AssertATK(expected float64) {
	if require(expected, c.Stats().ATK()) {
		return
	}
	c.t.Fatalf("Assert Fail, stat ATK, expected %f, actual %f", expected, c.Stats().ATK())
}

func (c *Character) AssertDEF(expected float64) {
	if require(expected, c.Stats().DEF()) {
		return
	}
	c.t.Fatalf("Assert Fail, stat DEF, expected %f, actual %f", expected, c.Stats().DEF())
}
