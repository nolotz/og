package ptrval_test

import (
	"testing"

	"github.com/nolotz/og/ptrval"
	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	t.Parallel()

	v := "test"
	vp := &v

	actual := ptrval.Value(vp)

	assert.Equal(t, "test", actual)
}

func TestNilValue(t *testing.T) {
	t.Parallel()

	var nilValue *string

	actual := ptrval.Value(nilValue)

	assert.IsType(t, "", actual)
}

func TestValueOK(t *testing.T) {
	t.Parallel()

	v := "test"
	vp := &v

	actual, ok := ptrval.ValueOK(vp)

	assert.True(t, ok)
	assert.Equal(t, "test", actual)
}

func TestValueNotOK(t *testing.T) {
	t.Parallel()

	var nilValue *string

	actual, ok := ptrval.ValueOK(nilValue)

	assert.False(t, ok)
	assert.IsType(t, "", actual)
}

func TestPtr(t *testing.T) {
	t.Parallel()

	value := "test"

	actual := ptrval.Pointer(value)

	assert.Equal(t, &value, actual)
}

func TestZero(t *testing.T) {
	t.Parallel()

	actual := ptrval.Zero[string]()

	assert.IsType(t, "", actual)
}
