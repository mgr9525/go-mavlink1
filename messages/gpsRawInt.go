package messages

const MSG_ID_GPS_RAW_INT = 24

type GpsRawInt struct {
	TimeUsec          [8]byte /* uint64 /*< [us] Timestamp (microseconds since UNIX epoch or microseconds since system boot)*/
	Lat               [4]byte /* int32  /*< [degE7] Latitude (WGS84, EGM96 ellipsoid), in degrees * 1E7*/
	Lon               [4]byte /* int32  /*< [degE7] Longitude (WGS84, EGM96 ellipsoid), in degrees * 1E7*/
	Alt               [4]byte /* int32  /*< [mm] Altitude (AMSL, NOT WGS84), in meters * 1000 (positive for up). Note that virtually all GPS modules provide the AMSL altitude in addition to the WGS84 altitude.*/
	Eph               [2]byte /* uint16 /*<  GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX*/
	Epv               [2]byte /* uint16 /*<  GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX*/
	Vel               [2]byte /* uint16 /*< [cm/s] GPS ground speed (m/s * 100). If unknown, set to: UINT16_MAX*/
	Cog               [2]byte /* uint16 /*< [cdeg] Course over ground (NOT heading, but direction of movement) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX*/
	FixType           byte    /* uint8  /*<  See the GPS_FIX_TYPE enum.*/
	SatellitesVisible byte    /* uint8  /*<  Number of satellites visible. If unknown, set to 255*/
	AltEllipsoid      [4]byte /* int32  /*< [mm] Altitude (above WGS84, EGM96 ellipsoid), in meters * 1000 (positive for up).*/
	HAcc              [4]byte /* uint32 /*< [mm] Position uncertainty in meters * 1000 (positive for up).*/
	VAcc              [4]byte /* uint32 /*< [mm] Altitude uncertainty in meters * 1000 (positive for up).*/
	VelAcc            [4]byte /* uint32 /*< [mm] Speed uncertainty in meters * 1000 (positive for up).*/
	HdgAcc            [4]byte /* uint32 /*< [degE5] Heading / track uncertainty in degrees * 1e5.*/
}
