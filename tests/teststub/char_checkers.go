package teststub

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/stretchr/testify/assert"
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
const epsilon = 0.001

func (c *Character) Equal(statKey prop.Property, expected float64) {
	actual := c.Stats().GetProperty(statKey)
	assert.InEpsilonf(c.t, expected, actual, epsilon,
		"char %s, stat %s, expected %f, actual %f", c.model.Key, statKey.String(), expected, actual)
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
