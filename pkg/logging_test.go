package pkg

import (
	"github.com/simimpact/srsim/pkg/engine/logging"
	"github.com/simimpact/srsim/pkg/internal/testcfg"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Logging(t *testing.T) {
	_, err := simulation.RunWithLog(testcfg.TestConfigTwoElites(), testcfg.StandardTestEval(), 0)
	assert.Nil(t, err)
	assert.NotNil(t, logging.Singleton)
	logging.PrintToConsole()
}
