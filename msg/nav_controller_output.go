package messages

import "unsafe"

const MSG_ID_NAV_CONTROLLER_OUTPUT = 62

type NavControllerOutput struct {
	NavRoll       float32 /*< [deg] Current desired roll in degrees*/
	NavPitch      float32 /*< [deg] Current desired pitch in degrees*/
	AltError      float32 /*< [m] Current altitude error in meters*/
	AspdError     float32 /*< [m/s] Current airspeed error in meters/second*/
	XtrackError   float32 /*< [m] Current crosstrack error on x-y plane in meters*/
	NavBearing    int16   /*< [deg] Current desired heading in degrees*/
	TargetBearing int16   /*< [deg] Bearing to current waypoint/target in degrees*/
	WpDist        uint16  /*< [m] Distance to active waypoint in meters*/
}

func Byte2NavControllerOutput(bts *[]byte) *NavControllerOutput {
	sct := *(**NavControllerOutput)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func NavControllerOutput2Byte(sct *NavControllerOutput) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
