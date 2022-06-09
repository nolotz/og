package dcontext_test

import (
	"context"
	"testing"

	"github.com/nolotz/og/dcontext"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")

	canceledCtx, cancel := context.WithCancel(ctx)
	cancel()

	ctx = dcontext.New(canceledCtx)

	assert.Equal(t, "value", ctx.Value("key"))
	assert.Nil(t, ctx.Done())

	_, deadlineOK := ctx.Deadline()
	assert.False(t, deadlineOK)
}
