package wmi

import (
	"fmt"
	"runtime"
	"sync"

	"errors"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
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

type Closer interface {
	Close()
}

func simplyClose(p Closer) {
	p.Close()
}

type SWbemProperty struct {
	m sync.Mutex
	i *ole.IDispatch

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

	p := &SWbemProperty{i: i, IsArray: isArray.Val > 0, IsLocal: isLocal.Val > 0, Name: name.ToString(), Origin: origin.ToString()}

	runtime.SetFinalizer(p, simplyClose)
	return p, nil
}

func (s *SWbemProperty) Close() {
	s.i.Release()
	// comshim.Done()
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

func (s *SWbemProperty) ValueSet(v interface{}) error {
	s.m.Lock()
	defer s.m.Unlock()

	r, err := s.i.PutProperty("Value", v)
	defer r.Clear()

	return err
}

type SWbemObject struct {
	m sync.Mutex
	i *ole.IDispatch
}

func newSWbemObject(i *ole.IDispatch) *SWbemObject {
	// comshim.Add(1)
	obj := &SWbemObject{i: i}
	runtime.SetFinalizer(obj, simplyClose)
	return obj
}

func (s *SWbemObject) Close() {
	s.i.Release()
	// comshim.Done()
}

func (s *SWbemObject) ExecMethod(methodName string, inParam *SWbemObject) (*SWbemObject, error) {
	s.m.Lock()
	defer s.m.Unlock()

	var objRaw *ole.VARIANT
	var err error
	if inParam != nil {
		objRaw, err = s.i.CallMethod("ExecMethod", methodName, inParam.i)
	} else {
		objRaw, err = s.i.CallMethod("ExecMethod", methodName)
	}

	if err != nil {
		return nil, err
	}
	return newSWbemObject(objRaw.ToIDispatch()), nil
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
	defer prop.Close()

	v, err := prop.ValueGet()
	if err != nil {
		panic(err)
	}
	return v
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
		return "err: " + err.Error()
	} else {
		return objText
	}
}

type SWbemObjectSet struct {
	m sync.Mutex
	i *ole.IDispatch
}

func newSWbemObjectSet(i *ole.IDispatch) *SWbemObjectSet {
	// comshim.Add(1)
	obj := &SWbemObjectSet{i: i}
	runtime.SetFinalizer(obj, simplyClose)
	return obj
}

func (s *SWbemObjectSet) Close() {
	s.i.Release()
	// comshim.Done()
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
	m sync.Mutex
	i *ole.IDispatch
}

func newSWbemServices(i *ole.IDispatch) *SWbemServices {
	// comshim.Add(1)
	obj := &SWbemServices{i: i}
	runtime.SetFinalizer(obj, simplyClose)
	return obj
}

func (s *SWbemServices) Close() {
	s.i.Release()
	// comshim.Done()
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
	m sync.Mutex
	i *ole.IDispatch
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

	s.i, err = unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		// comshim.Done()
		return nil, err
	}

	runtime.SetFinalizer(&s, simplyClose)
	return &s, nil
}

func (s *SWbemLocator) Close() {
	s.i.Release()
	// comshim.Done()
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
