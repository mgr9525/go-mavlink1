package messages

import "unsafe"

const MSG_ID_GLOBAL_POSITION_INT = 33

type GlobalPositionInt struct {
	TimeBootMs  uint32 /*< [ms] Timestamp (milliseconds since system boot)*/
	Lat         uint32 /*< [degE7] Latitude, expressed as degrees * 1E7*/
	Lon         uint32 /*< [degE7] Longitude, expressed as degrees * 1E7*/
	Alt         uint32 /*< [mm] Altitude in meters, expressed as * 1000 (millimeters), AMSL (not WGS84 - note that virtually all GPS modules provide the AMSL as well)*/
	RelativeAlt uint32 /*< [mm] Altitude above ground in meters, expressed as * 1000 (millimeters)*/
	Vx          int16  /*< [cm/s] Ground X Speed (Latitude, positive north), expressed as m/s * 100*/
	Vy          int16  /*< [cm/s] Ground Y Speed (Longitude, positive east), expressed as m/s * 100*/
	Vz          int16  /*< [cm/s] Ground Z Speed (Altitude, positive down), expressed as m/s * 100*/
	Hdg         uint16 /*< [cdeg] Vehicle heading (yaw angle) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX*/
}

func Byte2GlobalPositionInt(bts *[]byte) *GlobalPositionInt {
	sct := *(**GlobalPositionInt)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func GlobalPositionInt2Byte(sct *GlobalPositionInt) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
