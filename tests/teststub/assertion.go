package teststub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type assertion struct {
	t *testing.T
}

func (a *assertion) Equal(expected, actual float64) {
	if expected == 0 {
		assert.Equal(a.t, expected, actual)
	} else {
		assert.InEpsilon(a.t, expected, actual, epsilon)
	}
}

func (a *assertion) Greater(expected, actual float64) {
	assert.Greater(a.t, expected, actual)
}

func (a *assertion) Less(expected, actual float64) {
	assert.Less(a.t, expected, actual)
}
