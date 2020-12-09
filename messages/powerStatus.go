package messages

const MSG_ID_POWER_STATUS = 125

type powerStatus struct {
	Vcc    [2]byte /* uint16 /*< [mV] 5V rail voltage in millivolts*/
	Vservo [2]byte /* uint16 /*< [mV] servo rail voltage in millivolts*/
	Flags  [2]byte /* uint16 /*<  power supply status flags (see MAV_POWER_STATUS enum)*/
}
