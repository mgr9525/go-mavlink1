package messages

const MSG_ID_AHRS2 = 178

type Ahrs2 struct {
	Roll     [4]byte /* float32 /*< [rad] Roll angle (rad)*/
	Pitch    [4]byte /* float32 /*< [rad] Pitch angle (rad)*/
	Yaw      [4]byte /* float32 /*< [rad] Yaw angle (rad)*/
	Altitude [4]byte /* float32 /*< [m] Altitude (MSL)*/
	Lat      [4]byte /* int32   /*< [degE7] Latitude in degrees * 1E7*/
	Lng      [4]byte /* int32   /*< [degE7] Longitude in degrees * 1E7*/
}
