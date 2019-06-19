package messages

import "unsafe"

const MSG_ID_VFR_HUD = 74

type VfrHud struct {
	VfrHudairspeed    float32 /*< [m/s] Current airspeed in m/s*/
	VfrHudgroundspeed float32 /*< [m/s] Current ground speed in m/s*/
	VfrHudalt         float32 /*< [m] Current altitude (MSL), in meters*/
	VfrHudclimb       float32 /*< [m/s] Current climb rate in meters/second*/
	Heading           int16   /*< [deg] Current heading in degrees, in compass units (0..360, 0=north)*/
	Throttle          uint16  /*< [%] Current throttle setting in integer percent, 0 to 100*/
}

func Byte2VfrHud(bts *[]byte) *VfrHud {
	sct := *(**VfrHud)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func VfrHud2Byte(sct *VfrHud) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
