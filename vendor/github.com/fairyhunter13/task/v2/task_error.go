package task

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants"
)

const (
	channelClosed uint64 = 1
	channelOpen   uint64 = 0
)

// ErrorManager contains Manager but especially for error function.
type ErrorManager struct {
	chErr    chan error
	isClosed uint64
	wg       *sync.WaitGroup
	opt      *Option
}

// NewErrorManager initialize the new error manager.
func NewErrorManager(fns ...OptionFunc) *ErrorManager {
	em := (&ErrorManager{}).Assign(fns...)
	return em
}

func (em *ErrorManager) init() *ErrorManager {
	if em.opt == nil {
		em.opt = NewDefaultOption()
	}
	if em.chErr == nil || em.isChannelClosed() {
		em.chErr = make(chan error, em.opt.BufferSize)
		atomic.StoreUint64(&em.isClosed, channelOpen)
	}
	if em.wg == nil {
		em.wg = new(sync.WaitGroup)
	}
	return em
}

func (em *ErrorManager) isChannelClosed() bool {
	return atomic.LoadUint64(&em.isClosed) == channelClosed
}

func (em *ErrorManager) recoverPanic(fn ClosureErr) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			switch valRec := rec.(type) {
			case error:
				err = valRec
			default:
				err = fmt.Errorf("%v", valRec)
			}
		}
	}()
	err = fn()
	return
}

func (em *ErrorManager) close() {
	close(em.chErr)
	atomic.StoreUint64(&em.isClosed, channelClosed)
}

// Assign assigns the OptionFunc to the opt in the error manager.
func (em *ErrorManager) Assign(fns ...OptionFunc) *ErrorManager {
	if em.opt == nil {
		em.opt = NewDefaultOption()
	}
	em.opt.Assign(fns...)
	return em.init()
}

// ClosureErr defines closure that returns error.
type ClosureErr func() (err error)

// Run runs the closure error function.
func (em *ErrorManager) Run(fn ClosureErr) *ErrorManager {
	if fn == nil {
		return em
	}

	em.init().wg.Add(1)
	opt := em.opt.Clone()
	_ = ants.Submit(func() {
		defer em.wg.Done()

		if opt.UsePanicHandler {
			em.chErr <- em.recoverPanic(fn)
		} else {
			em.chErr <- fn()
		}
	})
	return em
}

// ErrChan returns the receiving error channel of this error manager.
func (em *ErrorManager) ErrChan() <-chan error {
	return em.init().chErr
}

// WaitClose wait all go routines to complete and close the channel in the separate go routine.
func (em *ErrorManager) WaitClose() *ErrorManager {
	em.init()
	_ = ants.Submit(func() {
		em.wg.Wait()
		em.close()
	})
	return em
}

// Error returns the first error from the Run execution of the fn closure.
func (em *ErrorManager) Error() (err error) {
	em.WaitClose()

	for errTemp := range em.chErr {
		if err != nil {
			continue
		}

		err = errTemp
	}
	return
}
