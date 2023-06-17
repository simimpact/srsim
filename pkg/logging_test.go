package pkg

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/internal/testcfg"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Logging(t *testing.T) {
	handler.InitLogger()
	_, err := simulation.Run(testcfg.TestConfigTwoElites(), testcfg.StandardTestEval(), 0)
	assert.Nil(t, err)
	assert.NotNil(t, handler.Singleton)
	handler.Singleton.PrintToConsole()
}
