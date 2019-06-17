package messages

import "unsafe"

const MSG_ID_ATTITUDE = 30

type Attitude struct {
	TimeBootMs uint32  /*< [ms] Timestamp (time since system boot).*/
	Roll       float32 /*< [rad] Roll angle (-pi..+pi)*/
	Pitch      float32 /*< [rad] Pitch angle (-pi..+pi)*/
	Yaw        float32 /*< [rad] Yaw angle (-pi..+pi)*/
	Rollspeed  float32 /*< [rad/s] Roll angular speed*/
	Pitchspeed float32 /*< [rad/s] Pitch angular speed*/
	Yawspeed   float32 /*< [rad/s] Yaw angular speed*/
}

func Byte2Attitude(bts *[]byte) *Attitude {
	sct := *(**Attitude)(unsafe.Pointer(bts))
	return sct
}
func Attitude2Byte(sct *Attitude) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
