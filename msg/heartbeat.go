package messages

import "unsafe"

const MSG_ID_HEARTBEAT = 0

type Heartbeat struct {
	CustomMode     uint32 /*<  A bitfield for use for autopilot-specific flags*/
	Type           uint8  /*<  Type of the system (quadrotor, helicopter, etc.). Components use the same type as their associated system.*/
	autopilot      uint8  /*<  Autopilot type / class.*/
	BaseMode       uint8  /*<  System mode bitmap.*/
	SystemStatus   uint8  /*<  System status flag.*/
	MavlinkVersion uint8  /*<  MAVLink version, not writable by user, gets added by protocol because of magic data type: uint8_t_mavlink_version*/
}

func Byte2Heartbeat(bts *[]byte) *Heartbeat {
	sct := *(**Heartbeat)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func Heartbeat2Byte(sct *Heartbeat) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
