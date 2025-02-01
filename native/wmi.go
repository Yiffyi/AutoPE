package native

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/yusufpapurcu/wmi"
)

type Win32_Volume struct {
	DeviceID    string
	DriveLetter string
	Name        string

	BootVolume   bool
	SystemVolume bool

	Capacity  uint64
	FreeSpace uint64

	FileSystem string
}

func WMIGetVolume(driveLetter string) (volume *Win32_Volume, err error) {
	s, err := wmi.InitializeSWbemServices(wmi.DefaultClient)
	if err != nil {
		panic(err)
	}

	var volumes []Win32_Volume
	q := wmi.CreateQuery(&volumes, "")

	err = s.Query(q, &volumes)
	if err != nil {
		panic(err)
	}

	d := filepath.VolumeName(driveLetter)

	for _, v := range volumes {
		if v.DriveLetter == d {
			return &v, nil
		}
	}
	return nil, errors.New("could not found target volume")
}

func SWbemGetVolumes() error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		oleCode := err.(*ole.OleError).Code()
		if oleCode != ole.S_OK {
			return fmt.Errorf("ole.CoInitializeEx error: %v", err)
		}
	}
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		return fmt.Errorf("CreateObject SWbemLocator error: %v", err)
	} else if unknown == nil {
		return errors.New("ErrNilCreateObject")
	}
	defer unknown.Release()

	dispatch, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return fmt.Errorf("SWbemLocator QueryInterface error: %v", err)
	}
	defer dispatch.Release()

	serviceRaw, err := oleutil.CallMethod(dispatch, "ConnectServer")
	if err != nil {
		return fmt.Errorf("SWbemLocator ConnectServer error: %v", err)
	}
	defer serviceRaw.Clear()
	service := serviceRaw.ToIDispatch()

	classRaw, err := service.CallMethod("Get", "Win32_Volume")
	if err != nil {
		return fmt.Errorf("SWbemServices Get error: %v", err)
	}
	defer classRaw.Clear()
	class := classRaw.ToIDispatch()

	methodsRaw, err := class.GetProperty("Methods_")
	if err != nil {
		return fmt.Errorf("SWbemObject GetProperty Methods_ error: %v", err)
	}
	defer methodsRaw.Clear()
	methods := methodsRaw.ToIDispatch()

	oleutil.ForEach(methods, func(v *ole.VARIANT) error {
		methodName, err := v.ToIDispatch().GetProperty("Name")
		fmt.Println(methodName.ToString(), err)
		return nil
	})

	formatRaw, err := methods.CallMethod("Item", "Format")
	if err != nil {
		return fmt.Errorf("SWbemMethodSet Item Format error: %v", err)
	}
	defer formatRaw.Clear()
	format := formatRaw.ToIDispatch()

	formatInParamRaw, err := format.GetProperty("InParameters")
	if err != nil {
		return fmt.Errorf("SWbemMethod InParameters error: %v", err)
	}
	defer formatInParamRaw.Clear()
	formatInParam := formatInParamRaw.ToIDispatch()

	instFormatInParamRaw, err := formatInParam.CallMethod("SpawnInstance_")
	if err != nil {
		return fmt.Errorf("SWbemMethod SpawnInstance_ error: %v", err)
	}
	defer instFormatInParamRaw.Clear()
	instFormatInParam := instFormatInParamRaw.ToIDispatch()

	str, err := instFormatInParam.CallMethod("GetObjectText_")
	if err != nil {
		return fmt.Errorf("SWbemMethod GetObjectText_ error: %v", err)
	}
	defer str.Clear()
	fmt.Println("instFormatInParam", str.ToString())
	return err
}
