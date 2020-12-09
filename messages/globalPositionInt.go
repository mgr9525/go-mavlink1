package messages

const MSG_ID_GLOBAL_POSITION_INT = 33

type GlobalPositionInt struct {
	TimeBootMs  [4]byte /* uint32 /*< [ms] Timestamp (milliseconds since system boot)*/
	Lat         [4]byte /* uint32 /*< [degE7] Latitude, expressed as degrees * 1E7*/
	Lon         [4]byte /* uint32 /*< [degE7] Longitude, expressed as degrees * 1E7*/
	Alt         [4]byte /* uint32 /*< [mm] Altitude in meters, expressed as * 1000 (millimeters), AMSL (not WGS84 - note that virtually all GPS modules provide the AMSL as well)*/
	RelativeAlt [4]byte /* uint32 /*< [mm] Altitude above ground in meters, expressed as * 1000 (millimeters)*/
	Vx          [2]byte /* int16  /*< [cm/s] Ground X Speed (Latitude, positive north), expressed as m/s * 100*/
	Vy          [2]byte /* int16  /*< [cm/s] Ground Y Speed (Longitude, positive east), expressed as m/s * 100*/
	Vz          [2]byte /* int16  /*< [cm/s] Ground Z Speed (Altitude, positive down), expressed as m/s * 100*/
	Hdg         [2]byte /* uint16 /*< [cdeg] Vehicle heading (yaw angle) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX*/
}
