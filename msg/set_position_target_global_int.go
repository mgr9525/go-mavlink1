package messages

import "unsafe"

const MSG_ID_SET_POSITION_TARGET_GLOBAL_INT = 86

type SetPosTargetGlobalInt struct {
	TimeBootMs      uint32  /*< [ms] Timestamp in milliseconds since system boot. The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency.*/
	LatInt          uint32  /*< [degE7] X Position in WGS84 frame in 1e7 * degrees*/
	LonInt          uint32  /*< [degE7] Y Position in WGS84 frame in 1e7 * degrees*/
	Alt             float32 /*< [m] Altitude in meters in AMSL altitude, not WGS84 if absolute or relative, above terrain if GLOBAL_TERRAIN_ALT_INT*/
	Vx              float32 /*< [m/s] X velocity in NED frame in meter / s*/
	Vy              float32 /*< [m/s] Y velocity in NED frame in meter / s*/
	Vz              float32 /*< [m/s] Z velocity in NED frame in meter / s*/
	Afx             float32 /*< [m/s/s] X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N*/
	Afy             float32 /*< [m/s/s] Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N*/
	Afz             float32 /*< [m/s/s] Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N*/
	Yaw             float32 /*< [rad] yaw setpoint in rad*/
	YawRate         float32 /*< [rad/s] yaw rate setpoint in rad/s*/
	TypeMask        uint16  /*<  Bitmask to indicate which dimensions should be ignored by the vehicle: a value of 0b0000000000000000 or 0b0000001000000000 indicates that none of the setpoint dimensions should be ignored. If bit 10 is set the floats afx afy afz should be interpreted as force instead of acceleration. Mapping: bit 1: x, bit 2: y, bit 3: z, bit 4: vx, bit 5: vy, bit 6: vz, bit 7: ax, bit 8: ay, bit 9: az, bit 10: is force setpoint, bit 11: yaw, bit 12: yaw rate*/
	TargetSystem    uint8   /*<  System ID*/
	TargetComponent uint8   /*<  Component ID*/
	CoordinateFrame uint8   /*<  Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11*/
}

func Byte2SetPosTargetGlobalInt(bts *[]byte) *SetPosTargetGlobalInt {
	sct := *(**SetPosTargetGlobalInt)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func SetPosTargetGlobalInt2Byte(sct *SetPosTargetGlobalInt) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
