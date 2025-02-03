package wmi

import "github.com/go-ole/go-ole"

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
