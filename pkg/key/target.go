package key

import (
	"encoding/json"
	"fmt"
)

type TargetID int

type TargetIDGenerator struct {
	state int
}

func NewTargetIDGenerator() *TargetIDGenerator {
	return &TargetIDGenerator{
		state: 1, // start at 1 to ensure 0 is always invalid
	}
}

// Generate a new TargetID
func (gen *TargetIDGenerator) New() TargetID {
	out := gen.state
	gen.state += 1
	return TargetID(out)
}

func (t TargetID) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(t))
}
