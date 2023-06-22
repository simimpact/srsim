package pkg

import (
	"fmt"
	"github.com/simimpact/srsim/pkg/engine/logging"
	"github.com/simimpact/srsim/pkg/internal/testcfg"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Logging(t *testing.T) {
	log := logging.NewDefaultLogger()
	_, err := simulation.RunWithLog(log, testcfg.TestConfigTwoElites(), testcfg.StandardTestEval(), 0)
	assert.Nil(t, err)
	assert.NotNil(t, logging.Singleton)
	res := log.Flush()
	assert.NotEqual(t, 0, len(res))
	assert.Equal(t, 0, len(log.Flush()))

	fmt.Println(res)
}
