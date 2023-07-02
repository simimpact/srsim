package basic

import (
	"testing"

	"github.com/simimpact/srsim/testframe/eventchecker/battlestart"
	"github.com/simimpact/srsim/testframe/eventchecker/termination"
	"github.com/simimpact/srsim/testframe/teststub"
	"github.com/stretchr/testify/suite"
)

type BasicTest struct {
	teststub.Stub
}

func TestBasicTest(t *testing.T) {
	suite.Run(t, new(BasicTest))
}

func (t *BasicTest) Test_Framework() {
	t.StartSimulation()
	t.Expect(battlestart.ExpectFor())
	t.Expect(termination.ExpectFor())
}
