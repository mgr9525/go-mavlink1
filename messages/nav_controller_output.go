package messages

const MSG_ID_NAV_CONTROLLER_OUTPUT = 62

type NavControllerOutput struct {
	NavRoll       [4]byte /* float32 /*< [deg] Current desired roll in degrees*/
	NavPitch      [4]byte /* float32 /*< [deg] Current desired pitch in degrees*/
	AltError      [4]byte /* float32 /*< [m] Current altitude error in meters*/
	AspdError     [4]byte /* float32 /*< [m/s] Current airspeed error in meters/second*/
	XtrackError   [4]byte /* float32 /*< [m] Current crosstrack error on x-y plane in meters*/
	NavBearing    [2]byte /* int16   /*< [deg] Current desired heading in degrees*/
	TargetBearing [2]byte /* int16   /*< [deg] Bearing to current waypoint/target in degrees*/
	WpDist        [2]byte /* uint16  /*< [m] Distance to active waypoint in meters*/
}
