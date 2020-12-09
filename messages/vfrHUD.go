package messages

const MSG_ID_VFR_HUD = 74

type VfrHud struct {
	VfrHudairspeed    [4]byte /* float32 /*< [m/s] Current airspeed in m/s*/
	VfrHudgroundspeed [4]byte /* float32 /*< [m/s] Current ground speed in m/s*/
	VfrHudalt         [4]byte /* float32 /*< [m] Current altitude (MSL), in meters*/
	VfrHudclimb       [4]byte /* float32 /*< [m/s] Current climb rate in meters/second*/
	Heading           [2]byte /* int16   /*< [deg] Current heading in degrees, in compass units (0..360, 0=north)*/
	Throttle          [2]byte /* uint16  /*< [%] Current throttle setting in integer percent, 0 to 100*/
}
