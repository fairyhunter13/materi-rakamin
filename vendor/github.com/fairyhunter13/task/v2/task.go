package task

import (
	"sync"

	"github.com/panjf2000/ants"
)

// Manager contains all the required tools to manage the task.
type Manager struct {
	wg  *sync.WaitGroup
	opt *Option
}

// NewManager initialize the task manager.
func NewManager(fns ...OptionFunc) *Manager {
	m := new(Manager).Assign(fns...)
	return m
}

func (m *Manager) init() *Manager {
	if m.opt == nil {
		m.opt = NewDefaultOption()
	}
	if m.wg == nil {
		m.wg = new(sync.WaitGroup)
	}
	return m
}

func (m *Manager) recoverPanic(fn ClosureAnonym) {
	defer func() {
		_ = recover()
	}()
	fn()
}

// Assign assigns the functional option to the opt in the Manager.
func (m *Manager) Assign(fns ...OptionFunc) *Manager {
	if m.opt == nil {
		m.opt = NewDefaultOption()
	}
	m.opt.Assign(fns...)
	return m.init()
}

// ClosureAnonym defines the anonymous function for the Run argument.
type ClosureAnonym func()

// Run runs the task in a separate go function.
func (m *Manager) Run(fn ClosureAnonym) *Manager {
	if fn == nil {
		return m
	}

	m.init().wg.Add(1)
	opt := m.opt.Clone()
	_ = ants.Submit(func() {
		defer m.wg.Done()

		if opt.UsePanicHandler {
			m.recoverPanic(fn)
		} else {
			fn()
		}
	})
	return m
}

// Wait blocks the current thread until the wg counter is zero.
func (m *Manager) Wait() *Manager {
	m.init().wg.Wait()
	return m
}
