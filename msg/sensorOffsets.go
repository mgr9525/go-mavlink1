package messages

import "unsafe"

const MSG_ID_SENSOR_OFFSETS = 150

type SensorOffsets struct {
	MagDeclination float32 /*< [rad] magnetic declination (radians)*/
	RawPress       int32   /*<  raw pressure from barometer*/
	RawTemp        int32   /*<  raw temperature from barometer*/
	GyroCalX       float32 /*<  gyro X calibration*/
	GyroCalY       float32 /*<  gyro Y calibration*/
	GyroCalZ       float32 /*<  gyro Z calibration*/
	AccelCalX      float32 /*<  accel X calibration*/
	AccelCalY      float32 /*<  accel Y calibration*/
	AccelCalZ      float32 /*<  accel Z calibration*/
	MagOfsX        int16   /*<  magnetometer X offset*/
	MagOfsY        int16   /*<  magnetometer Y offset*/
	MagOfsZ        int16   /*<  magnetometer Z offset*/
}

func Byte2SensorOffsets(bts *[]byte) *SensorOffsets {
	sct := *(**SensorOffsets)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func SensorOffsets2Byte(sct *SensorOffsets) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
