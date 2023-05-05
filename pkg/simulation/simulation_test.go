package simulation

import (
	"bytes"
	"reflect"
	"testing"
)

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
