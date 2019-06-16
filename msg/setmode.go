package messages

import "unsafe"

const MSG_ID_SET_MODE = 11

type SetMode struct {
	CustomMode   uint32 /*<  The new autopilot-specific mode. This field can be ignored by an autopilot.*/
	TargetSystem uint8  /*<  The system setting the mode*/
	BaseMode     uint8  /*<  The new base mode*/
}

func Byte2SetMode(bts *[]byte) *SetMode {
	sct := *(**SetMode)(unsafe.Pointer(bts))
	return sct
}
func SetMode2Byte(sct *SetMode) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
