package wmi

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

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
