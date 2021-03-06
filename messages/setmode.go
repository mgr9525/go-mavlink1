package messages

const MSG_ID_SET_MODE = 11

const (
	MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE = 1   /* Eighth bit: 00000001 | */
	MAV_MODE_FLAG_DECODE_POSITION_TEST        = 2   /* Seventh bit: 00000010 | */
	MAV_MODE_FLAG_DECODE_POSITION_AUTO        = 4   /* Sixt bit:   00000100 | */
	MAV_MODE_FLAG_DECODE_POSITION_GUIDED      = 8   /* Fifth bit:  00001000 | */
	MAV_MODE_FLAG_DECODE_POSITION_STABILIZE   = 16  /* Fourth bit: 00010000 | */
	MAV_MODE_FLAG_DECODE_POSITION_HIL         = 32  /* Third bit:  00100000 | */
	MAV_MODE_FLAG_DECODE_POSITION_MANUAL      = 64  /* Second bit: 01000000 | */
	MAV_MODE_FLAG_DECODE_POSITION_SAFETY      = 128 /* First bit:  10000000 | */
	MAV_MODE_FLAG_DECODE_POSITION_ENUM_END    = 129 /*  | */
)

type SetMode struct {
	CustomMode   [4]byte /*<  The new autopilot-specific mode. This field can be ignored by an autopilot.*/
	TargetSystem byte    /*<  The system setting the mode*/
	BaseMode     byte    /*<  The new base mode*/
}
