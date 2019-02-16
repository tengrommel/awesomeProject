package contextimpl

import (
	"github.com/kataras/iris/core/errors"
	"sync"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int

func (emptyCtx) Deadline() (deadline time.Time, ok bool) { return }
func (emptyCtx) Done() <-chan struct{}                   { return nil }
func (emptyCtx) Err() error                              { return nil }
func (emptyCtx) Value(key interface{}) interface{}       { return nil }

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context { return background }
func TODO() Context       { return todo }

type cancelCtx struct {
	Context
	done chan struct{}
	err  error
	mu   sync.Mutex
}

func (ctx *cancelCtx) Deadline() (deadline time.Time, ok bool) { return ctx.Deadline() }
func (ctx *cancelCtx) Done() <-chan struct{}                   { return ctx.done }
func (ctx *cancelCtx) Err() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}
func (ctx *cancelCtx) Value(key interface{}) interface{} { return ctx.Value(key) }

var Canceled = errors.New("context canceled")

type CancelFunc func()

func WithCanel(parent Context) (Context, CancelFunc) {
	ctx := &cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}
	cancel := func() {
		ctx.mu.Lock()
		defer ctx.mu.Unlock()
		if ctx.err != nil {
			return
		}
		ctx.err = Canceled
		close(ctx.done)
	}
	go func() {
		select {
		case <-parent.Done():
			ctx.mu.Lock()
			if ctx.err != nil {
				return
			}
			ctx.err = parent.Err()
			close(ctx.done)
			ctx.mu.Unlock()
		case <-ctx.Done():
		}
	}()
	return ctx, cancel
}
