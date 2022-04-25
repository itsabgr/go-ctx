package ctx

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type CancelFunc func(err error)

type ctx struct {
	ctx    context.Context
	cancel context.CancelFunc
	err    atomic.Value
}

func (ctx *ctx) String() string {
	return ctx.ctx.(fmt.Stringer).String()
}

func (ctx *ctx) Deadline() (deadline time.Time, ok bool) {
	return ctx.ctx.Deadline()
}

func (ctx *ctx) Done() <-chan struct{} {
	return ctx.ctx.Done()
}

func (ctx *ctx) Err() error {
	err := ctx.err.Load()
	if err == nil {
		return nil
	}
	return err.(error)
}

func (ctx *ctx) Value(key any) any {
	return ctx.ctx.Value(key)
}

func (ctx *ctx) Cancel(err error) {
	if ctx.err.CompareAndSwap(nil, err) {
		ctx.cancel()
	}
}

func WithCancel(parent context.Context) (context.Context, CancelFunc) {
	c, cancel := context.WithCancel(parent)
	ctx := &ctx{ctx: c, cancel: cancel}
	return ctx, func(err error) {
		if err == nil {
			err = context.Canceled
		}
		ctx.Cancel(err)
	}
}
