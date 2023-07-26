package march7th

import "github.com/simimpact/srsim/pkg/engine/info"

const (
	Talent = "march7th-talent"
)

func init() {

}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
	})
}

func (c *char) talentCounterListener() {

}
