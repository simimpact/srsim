package simulation

import (
	"github.com/simimpact/srsim/pkg/core"
	"github.com/simimpact/srsim/pkg/model"
)

type Simulation struct {
	c *core.Core
}

type Config struct {
}

func New() (*Simulation, error) {
	return nil, nil
}

func (s *Simulation) createCore() error {
	var err error
	s.c, err = core.New()
	if err != nil {
		return err
	}
	return nil
}

// addTarget adds a new target to the simulation
func (s *Simulation) addTarget(t core.Target) error {
	return nil
}

func (s *Simulation) Run() (*model.SimulationResult, error) {

	return nil, nil
}

func simShouldStop(c *core.Core, s *model.SimulatorSettings) bool {
	if s.TtkMode {
		//TODO: check if all enemies dead
		return true
	}
	//otherwise end if
	return c.CurrentAV >= int(s.AvLimit)
}
