// package dummy implements a dummy character for testing purposes
package dummy

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
)

func init() {
	// .RegisterCharFunc(key.DummyCharacter, New)
}

type char struct {
}

func New(engine *engine.Engine, id key.TargetID) (interface{}, error) {
	c := &char{}

	return c, nil
}

func (c *char) Exec(typ key.ActionType) {

}
