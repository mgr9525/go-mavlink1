package messages

import "unsafe"

const MSG_ID_GPS_RAW_INT = 24

type GpsRawInt struct {
	TimeUsec          uint64 /*< [us] Timestamp (microseconds since UNIX epoch or microseconds since system boot)*/
	Lat               int32  /*< [degE7] Latitude (WGS84, EGM96 ellipsoid), in degrees * 1E7*/
	Lon               int32  /*< [degE7] Longitude (WGS84, EGM96 ellipsoid), in degrees * 1E7*/
	Alt               int32  /*< [mm] Altitude (AMSL, NOT WGS84), in meters * 1000 (positive for up). Note that virtually all GPS modules provide the AMSL altitude in addition to the WGS84 altitude.*/
	Eph               uint16 /*<  GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX*/
	Epv               uint16 /*<  GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX*/
	Vel               uint16 /*< [cm/s] GPS ground speed (m/s * 100). If unknown, set to: UINT16_MAX*/
	Cog               uint16 /*< [cdeg] Course over ground (NOT heading, but direction of movement) in degrees * 100, 0.0..359.99 degrees. If unknown, set to: UINT16_MAX*/
	FixType           uint8  /*<  See the GPS_FIX_TYPE enum.*/
	SatellitesVisible uint8  /*<  Number of satellites visible. If unknown, set to 255*/
	AltEllipsoid      int32  /*< [mm] Altitude (above WGS84, EGM96 ellipsoid), in meters * 1000 (positive for up).*/
	HAcc              uint32 /*< [mm] Position uncertainty in meters * 1000 (positive for up).*/
	VAcc              uint32 /*< [mm] Altitude uncertainty in meters * 1000 (positive for up).*/
	VelAcc            uint32 /*< [mm] Speed uncertainty in meters * 1000 (positive for up).*/
	HdgAcc            uint32 /*< [degE5] Heading / track uncertainty in degrees * 1e5.*/
}

func Byte2GpsRawInt(bts *[]byte) *GpsRawInt {
	sct := *(**GpsRawInt)(unsafe.Pointer(bts))
	return sct
}
func GpsRawInt2Byte(sct *GpsRawInt) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
