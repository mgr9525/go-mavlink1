package messages

const MSG_ID_SYS_STATUS = 1

type SysStatus struct {
	OnboardControlSensorsPresent [4]byte /* uint32 /*<  Bitmask showing which onboard controllers and sensors are present. Value of 0: not present. Value of 1: present. Indices defined by ENUM MAV_SYS_STATUS_SENSOR*/
	OnboardControlSensorsEnabled [4]byte /* uint32 /*<  Bitmask showing which onboard controllers and sensors are enabled:  Value of 0: not enabled. Value of 1: enabled. Indices defined by ENUM MAV_SYS_STATUS_SENSOR*/
	OnboardControlSensorsHealth  [4]byte /* uint32 /*<  Bitmask showing which onboard controllers and sensors are operational or have an error:  Value of 0: not enabled. Value of 1: enabled. Indices defined by ENUM MAV_SYS_STATUS_SENSOR*/
	Load                         [2]byte /* uint16 /*< [d%] Maximum usage in percent of the mainloop time, (0%: 0, 100%: 1000) should be always below 1000*/
	VoltageBattery               [2]byte /* uint16 /*< [mV] Battery voltage, in millivolts (1 = 1 millivolt)*/
	CurrentBattery               [2]byte /* int16  /*< [cA] Battery current, in 10*milliamperes (1 = 10 milliampere), -1: autopilot does not measure the current*/
	DropRateComm                 [2]byte /* uint16 /*< [c%] Communication drops in percent, (0%: 0, 100%: 10'000), (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)*/
	ErrorsComm                   [2]byte /* uint16 /*<  Communication errors (UART, I2C, SPI, CAN), dropped packets on all links (packets that were corrupted on reception on the MAV)*/
	ErrorsCount1                 [2]byte /* uint16 /*<  Autopilot-specific errors*/
	ErrorsCount2                 [2]byte /* uint16 /*<  Autopilot-specific errors*/
	ErrorsCount3                 [2]byte /* uint16 /*<  Autopilot-specific errors*/
	ErrorsCount4                 [2]byte /* uint16 /*<  Autopilot-specific errors*/
	BatteryRemaining             byte    /* uint8  /*< [%] Remaining battery energy: (0%: 0, 100%: 100), -1: autopilot estimate the remaining battery*/
}
