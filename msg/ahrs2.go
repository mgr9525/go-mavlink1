package messages

import "unsafe"

const MSG_ID_AHRS2 = 178

type Ahrs2 struct {
	Roll     float32 /*< [rad] Roll angle (rad)*/
	Pitch    float32 /*< [rad] Pitch angle (rad)*/
	Yaw      float32 /*< [rad] Yaw angle (rad)*/
	Altitude float32 /*< [m] Altitude (MSL)*/
	Lat      int32   /*< [degE7] Latitude in degrees * 1E7*/
	Lng      int32   /*< [degE7] Longitude in degrees * 1E7*/
}

func Byte2Ahrs2(bts *[]byte) *Ahrs2 {
	sct := *(**Ahrs2)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func Ahrs22Byte(sct *Ahrs2) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
