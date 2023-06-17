package logging

import (
	"github.com/simimpact/srsim/pkg/internal/testcfg"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Logging(t *testing.T) {
	_, err := simulation.Run(testcfg.TestConfigTwoElites(), testcfg.StandardTestEval(), 0)
	assert.Nil(t, err)
}
