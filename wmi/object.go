package wmi

import (
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

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
