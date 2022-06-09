package dcontext

import (
	"context"
	"time"
)

// New returns a detached copy of parent but keeps all values
func New(parent context.Context) context.Context {
	return &detached{
		Context: parent,
	}
}

var _ context.Context = (*detached)(nil)

type detached struct {
	context.Context
}

func (d detached) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (d detached) Done() <-chan struct{} {
	return nil
}
