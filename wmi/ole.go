package wmi

import (
	"errors"
	"runtime"
	"sync"

	"github.com/go-ole/go-ole"
)

var (
	// ErrNegativeCounter is returned when the internal counter of a shim drops
	// below zero. This may indicate that Done() has been called more than once
	// for the same object.
	ErrNegativeCounter = errors.New("component object model shim counter has dropped below zero")

	// ErrAlreadyInitialized is returned when a shim finds itself on a thread
	// that has already been initialized. This probably indicates that some
	// previous goroutine failed to lock the OS thread or failed to call
	// CoUninitialize when it should have.
	ErrAlreadyInitialized = errors.New("component object model shim thread has already been initialized")
)

var _com = newComBase()

type comBase struct {
	m    *sync.Mutex
	cond *sync.Cond

	running bool
	done    bool
	initErr chan error
}

func newComBase() *comBase {
	c := &comBase{
		m:       &sync.Mutex{},
		running: false,
		done:    false,
		initErr: make(chan error),
	}

	c.cond = sync.NewCond(c.m)
	return c
}

func (c *comBase) run() {
	c.m.Lock()
	defer c.m.Unlock()

	c.done = false
	if c.running {
		c.initErr <- nil
		return
	} else {
		c.running = true
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
		switch err.(*ole.OleError).Code() {
		case 0x00000001: // S_FALSE
			// Some other goroutine called CoInitialize on this thread
			// before we ended up with it. This probably means the other
			// caller failed to lock the OS thread or failed to call
			// CoUninitialize.

			// We still decrement this thread's initialization counter by
			// calling CoUninitialize here, as recommended by the docs.
			ole.CoUninitialize()

			// Send an error so that shim.Add panics
			c.initErr <- ErrAlreadyInitialized
			c.running = false
			return
		default:
			c.initErr <- err
			c.running = false
			return
		}
	}
	c.initErr <- nil

	for !c.done {
		c.cond.Wait() // c.m is not locked while Wait is waiting
	}
	ole.CoUninitialize()
	c.running = false
}

func (c *comBase) stop() {
	c.done = true
	c.cond.Broadcast()
}

func CoInitialize() error {
	if _com.running {
		return nil
	}

	go _com.run()
	return <-_com.initErr
}

func CoUninitialize() {
	_com.stop()
}
