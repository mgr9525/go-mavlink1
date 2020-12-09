package messages

const MSG_ID_HEARTBEAT = 0

type Heartbeat struct {
	CustomMode     [4]byte /* uint32 /*<  A bitfield for use for autopilot-specific flags*/
	Type           byte    /* uint8  /*<  Type of the system (quadrotor, helicopter, etc.). Components use the same type as their associated system.*/
	autopilot      byte    /* uint8  /*<  Autopilot type / class.*/
	BaseMode       byte    /* uint8  /*<  System mode bitmap.*/
	SystemStatus   byte    /* uint8  /*<  System status flag.*/
	MavlinkVersion byte    /* uint8  /*<  MAVLink version, not writable by user, gets added by protocol because of magic data type: uint8_t_mavlink_version*/
}
