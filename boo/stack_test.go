package boo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nolotz/og/boo"
)

type errorErrorCase struct {
	err      error
	expected string
}

func TestStackError(t *testing.T) {
	cases := []errorErrorCase{
		{
			err:      nil,
			expected: "boo_test.TestStackError/stack_test.go@30: err is nil",
		},
		{
			err:      errors.New("test error"),
			expected: "boo_test.TestStackError/stack_test.go@30: test error",
		},
	}

	for _, c := range cases {
		stacked := boo.WithStack(c.err)

		assert.NotNil(t, stacked)
		assert.Equal(t, c.expected, stacked.Error())
		assert.Equal(t, c.err, errors.Unwrap(stacked))
	}
}

