package wmi

import (
	"fmt"
	"runtime"
	"sync"

	"errors"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
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

type SWbemProperty struct {
	SWbemBase

	IsArray bool
	IsLocal bool
	Name    string
	Origin  string
}

func newSWbemProperty(i *ole.IDispatch) (*SWbemProperty, error) {
	// comshim.Add(1)

	isArray, err := i.GetProperty("IsArray")
	if err != nil {
		return nil, err
	}
	defer isArray.Clear()

	isLocal, err := i.GetProperty("IsLocal")
	if err != nil {
		return nil, err
	}
	defer isLocal.Clear()

	name, err := i.GetProperty("Name")
	if err != nil {
		return nil, err
	}
	defer name.Clear()

	origin, err := i.GetProperty("Origin")
	if err != nil {
		return nil, err
	}
	defer origin.Clear()

	p := &SWbemProperty{IsArray: isArray.Val > 0, IsLocal: isLocal.Val > 0, Name: name.ToString(), Origin: origin.ToString()}
	p.init(i, "SWbemProperty")
	return p, nil
}

func (s *SWbemProperty) ValueGet() (interface{}, error) {
	s.m.Lock()
	defer s.m.Unlock()

	valueRaw, err := s.i.GetProperty("Value")
	if err != nil {
		return nil, err
	}
	defer valueRaw.Clear()

	value := valueRaw.Value()
	return value, nil
}

func (s *SWbemProperty) ValuePut(v interface{}) error {
	s.m.Lock()
	defer s.m.Unlock()

	r, err := s.i.PutProperty("Value", v)
	defer r.Clear()

	return err
}

type SWbemObject struct {
	SWbemBase
}

func newSWbemObject(i *ole.IDispatch) *SWbemObject {
	// comshim.Add(1)
	obj := &SWbemObject{}
	obj.init(i, "SWbemObject")
	return obj
}

func (s *SWbemObject) ExecMethod_(methodName string, inParam *SWbemObject) (*SWbemObject, error) {
	s.m.Lock()
	defer s.m.Unlock()

	var objRaw *ole.VARIANT
	var err error
	if inParam != nil {
		objRaw, err = s.i.CallMethod("ExecMethod_", methodName, inParam.i)
	} else {
		objRaw, err = s.i.CallMethod("ExecMethod_", methodName)
	}

	if err != nil {
		return nil, err
	}
	return newSWbemObject(objRaw.ToIDispatch()), nil
}

func (s *SWbemObject) CallMethod(methodName string, args ...interface{}) (uint32, error) {
	s.m.Lock()
	defer s.m.Unlock()

	var objRaw *ole.VARIANT
	var err error
	objRaw, err = s.i.CallMethod(methodName, args...)
	if err != nil {
		return 0, err
	}

	resultInt := objRaw.Value().(uint32)
	return resultInt, nil
}

func (s *SWbemObject) PropertyGet(propertyName string) (*SWbemProperty, error) {
	s.m.Lock()
	defer s.m.Unlock()

	propsRaw, err := s.i.GetProperty("Properties_")
	if err != nil {
		return nil, err
	}
	props := propsRaw.ToIDispatch()
	defer props.Release()

	pRaw, err := props.CallMethod("Item", propertyName)
	if err != nil {
		return nil, err
	}

	return newSWbemProperty(pRaw.ToIDispatch())
}

func (s *SWbemObject) PropertyMustGetValue(propertyName string) interface{} {
	prop, err := s.PropertyGet(propertyName)
	if err != nil {
		panic(err)
	}

	v, err := prop.ValueGet()
	if err != nil {
		panic(err)
	}
	return v
}

func (s *SWbemObject) PropertyPutValue(propertyName string, value interface{}) error {
	prop, err := s.PropertyGet(propertyName)
	if err != nil {
		return err
	}

	err = prop.ValuePut(value)
	if err != nil {
		return err
	}
	return nil
}

func (s *SWbemObject) GetObjectText() (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	objTextRaw, err := s.i.CallMethod("GetObjectText_")
	if err != nil {
		return "", err
	}
	defer objTextRaw.Clear()

	return objTextRaw.ToString(), nil
}

func (s *SWbemObject) String() string {
	objText, err := s.GetObjectText()
	if err != nil {
		return fmt.Sprintf("{ERROR:%s}", err.Error())
	} else {
		return objText
	}
}

type SWbemObjectSet struct {
	SWbemBase
}

func newSWbemObjectSet(i *ole.IDispatch) *SWbemObjectSet {
	// comshim.Add(1)
	obj := &SWbemObjectSet{}
	obj.init(i, "SWbemObjectSet")
	return obj
}

func (s *SWbemObjectSet) ForEach(f func(v *SWbemObject) error) error {
	s.m.Lock()
	defer s.m.Unlock()

	return oleutil.ForEach(s.i, func(v *ole.VARIANT) error {
		return f(newSWbemObject(v.ToIDispatch()))
	})
}

func (s *SWbemObjectSet) Count() (int64, error) {
	s.m.Lock()
	defer s.m.Unlock()

	result, err := s.i.GetProperty("Count")
	if err != nil {
		return 0, err
	}
	defer result.Clear()

	return result.Val, nil
}

func (s *SWbemObjectSet) ToSlice() ([]*SWbemObject, error) {
	s.m.Lock()
	defer s.m.Unlock()

	result := make([]*SWbemObject, 0)
	err := oleutil.ForEach(s.i, func(v *ole.VARIANT) error {
		result = append(result, newSWbemObject(v.ToIDispatch()))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, err
}

type SWbemServices struct {
	SWbemBase
}

func newSWbemServices(i *ole.IDispatch) *SWbemServices {
	// comshim.Add(1)
	obj := &SWbemServices{}
	obj.init(i, "SWbemServices")
	return obj
}

func (s *SWbemServices) ExecQuery(query string) (*SWbemObjectSet, error) {
	s.m.Lock()
	defer s.m.Unlock()

	objSetRaw, err := s.i.CallMethod("ExecQuery", query)
	if err != nil {
		return nil, err
	}
	// defer objSetRaw.Clear()

	objSet := objSetRaw.ToIDispatch()
	return newSWbemObjectSet(objSet), nil
}

func (s *SWbemServices) InstancesOf(class string) (*SWbemObjectSet, error) {
	s.m.Lock()
	defer s.m.Unlock()

	objSetRaw, err := s.i.CallMethod("InstancesOf", class)
	if err != nil {
		return nil, err
	}
	// defer objSetRaw.Clear()

	objSet := objSetRaw.ToIDispatch()
	return newSWbemObjectSet(objSet), nil
}

func (s *SWbemServices) Get(objectPath string) (*SWbemObject, error) {
	s.m.Lock()
	defer s.m.Unlock()

	objRaw, err := s.i.CallMethod("Get", objectPath)
	if err != nil {
		return nil, err
	}

	return newSWbemObject(objRaw.ToIDispatch()), nil
}

func (s *SWbemServices) ExecMethod(objectPath, methodName string, inParam *SWbemObject) (*SWbemObject, error) {
	s.m.Lock()
	defer s.m.Unlock()

	var objRaw *ole.VARIANT
	var err error
	if inParam != nil {
		objRaw, err = s.i.CallMethod("ExecMethod", objectPath, methodName, inParam.i)
	} else {
		objRaw, err = s.i.CallMethod("ExecMethod", objectPath, methodName)
	}

	if err != nil {
		return nil, err
	}
	return newSWbemObject(objRaw.ToIDispatch()), nil
}

type SWbemLocator struct {
	SWbemBase
}

func NewSWbemLocator() (*SWbemLocator, error) {
	// comshim.Add(1)

	var err error
	s := SWbemLocator{}

	unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		// comshim.Done()
		return nil, err
	}

	dispatch, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		// comshim.Done()
		return nil, err
	}

	s.init(dispatch, "SWbemLocator")

	return &s, nil
}

func (s *SWbemLocator) ConnectServer(server, namespace, user, password string) (*SWbemServices, error) {
	s.m.Lock()
	defer s.m.Unlock()

	serviceRaw, err := s.i.CallMethod("ConnectServer", server, namespace, user, password)
	if err != nil {
		return nil, err
	}
	// defer serviceRaw.Clear()

	service := serviceRaw.ToIDispatch()
	return newSWbemServices(service), nil
}

func (s *SWbemLocator) ConnectServerDefault() (*SWbemServices, error) {
	s.m.Lock()
	defer s.m.Unlock()

	serviceRaw, err := s.i.CallMethod("ConnectServer")
	if err != nil {
		return nil, err
	}
	// defer serviceRaw.Clear()

	service := serviceRaw.ToIDispatch()
	return newSWbemServices(service), nil
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
