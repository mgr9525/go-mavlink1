package messages

import "unsafe"

const MSG_ID_POWER_STATUS = 125

type powerStatus struct {
	Vcc    uint16 /*< [mV] 5V rail voltage in millivolts*/
	Vservo uint16 /*< [mV] servo rail voltage in millivolts*/
	Flags  uint16 /*<  power supply status flags (see MAV_POWER_STATUS enum)*/
}

func Byte2powerStatus(bts *[]byte) *powerStatus {
	sct := *(**powerStatus)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func powerStatus2Byte(sct *powerStatus) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
