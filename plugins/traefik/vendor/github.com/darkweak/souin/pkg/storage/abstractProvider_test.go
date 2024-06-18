package storage

import (
	"testing"

	"github.com/HobMartin/souin/errors"
	"github.com/HobMartin/souin/tests"
)

func TestInitializeProvider(t *testing.T) {
	c := tests.MockConfiguration(tests.BaseConfiguration)
	p := InitializeProvider(c)
	err := p.Init()
	if nil != err {
		errors.GenerateError(t, "Init shouldn't crash")
	}
}
