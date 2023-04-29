// Package core provides core interface definitions and glues together
// all the pieces of the combat system
package core

type TaskHandler interface {
}

type Logger interface {
}

type EventHandler interface {
}

type TargetHandler interface {
	TargetCount()
}

type Core struct {
	CurrentAV int //StateCount represents the counter for the current state; advance by 1 each time
}

func New() (*Core, error) {
	return nil, nil
}

// AdvanceAV advances av by x units
func (c *Core) AdvanceAV(x int) {
	c.CurrentAV += x
}
