package wmi

import (
	"fmt"
	"runtime"
	"sync"

	"errors"

	"github.com/go-ole/go-ole"
	"github.com/rs/zerolog/log"
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

func CoInitialize() error {
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
			return ErrAlreadyInitialized
		default:
			return err
		}
	}
	return nil
}

func swbmBaseCloser(p *SWbemBase) {
	p.close()
}

type SWbemBase struct {
	m sync.Mutex
	i *ole.IDispatch

	typeName string
}

func (s *SWbemBase) init(i *ole.IDispatch, typeName string) {
	s.i = i
	s.typeName = typeName

	runtime.SetFinalizer(s, swbmBaseCloser)
}

func (s *SWbemBase) close() {
	log.Debug().Str("type", s.typeName).Str("i", fmt.Sprint(s.i)).Msg("SWbemBase close")
	if s.i != nil {
		s.i.Release()
		s.i = nil
	}
}

func GetMethodInParam(class *SWbemObject, methodName string) (*SWbemObject, error) {
	methodsRaw, err := class.i.GetProperty("Methods_")
	if err != nil {
		return nil, err
	}
	// defer methodsRaw.Clear()
	methods := methodsRaw.ToIDispatch()
	defer methods.Release()

	methodRaw, err := methods.CallMethod("Item", methodName)
	if err != nil {
		return nil, fmt.Errorf("could not get method: %v", err)
	}
	method := methodRaw.ToIDispatch()
	defer method.Release()

	methodInParamRaw, err := method.GetProperty("InParameters")
	if err != nil {
		return nil, fmt.Errorf("could not get method.InParameters: %v", err)
	}
	// defer formatInParamRaw.Clear()
	methodInParam := methodInParamRaw.ToIDispatch()
	defer methodInParam.Release()

	instMethodInParamRaw, err := methodInParam.CallMethod("SpawnInstance_")
	if err != nil {
		return nil, fmt.Errorf("method.InParameters.SpawnInstance_ error: %v", err)
	}
	// defer instmethodInParamRaw.Clear()
	instMethodInParam := instMethodInParamRaw.ToIDispatch()

	return newSWbemObject(instMethodInParam), nil
}
