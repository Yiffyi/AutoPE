package wmi

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/go-ole/go-ole"
	"github.com/rs/zerolog/log"
)

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
