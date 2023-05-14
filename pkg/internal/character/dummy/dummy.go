// package dummy implements a dummy character for testing purposes
package dummy

import (
	"github.com/simimpact/srsim/pkg/engine/simulation"
	"github.com/simimpact/srsim/pkg/engine/system"
	"github.com/simimpact/srsim/pkg/key"
)

func init() {
	simulation.RegisterCharFunc(key.DummyCharacter, New)
}

type char struct {

}

func New(sys *system.CharacterServices, id key.TargetID) (simulation.Target, error) {
	c := &char{}

	return c, nil
}

func (c *char) Exec(typ key.ActionType) {

}
