package key

type TargetID int

// todo: functions to create target ids
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
	// TODO: random id instead to ensure no one depends on its value?
	out := gen.state
	gen.state += 1
	return TargetID(out)
}
