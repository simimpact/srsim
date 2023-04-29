package simulation

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/simimpact/srsim/pkg/core"
	"github.com/simimpact/srsim/pkg/model"
)

func TestSimShouldStopAtTurnLimit(t *testing.T) {
	c := &core.Core{
		CurrentAV: 10,
	}
	s := &model.SimulatorSettings{
		AvLimit: 10,
	}
	assert(t, simShouldStop(c, s), true, "stop should be true when turn deadline reached")
}

func assert[K any](t *testing.T, expect, got K, msg string) {
	if !checkEqual(expect, got) {
		t.Errorf("%v: expecting %v got %v", msg, expect, got)
	}
}

func checkEqual(expect, got interface{}) bool {
	if expect == nil || got == nil {
		return expect == got
	}

	exp, eok := expect.([]byte)
	act, aok := got.([]byte)

	if eok && aok {
		return bytes.Equal(exp, act)
	}

	return reflect.DeepEqual(expect, got)
}
