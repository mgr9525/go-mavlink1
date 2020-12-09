package messages

const MSG_ID_BATTERY_STATUS = 147

type BatteryStatus struct {
	CurrentConsumed  [4]byte  /* int32     /*< [mAh] Consumed charge, in milliampere hours (1 = 1 mAh), -1: autopilot does not provide mAh consumption estimate*/
	EnergyConsumed   [4]byte  /* int32     /*< [hJ] Consumed energy, in HectoJoules (intergrated U*I*dt)  (1 = 100 Joule), -1: autopilot does not provide energy consumption estimate*/
	Temperature      [2]byte  /* int16     /*< [cdegC] Temperature of the battery in centi-degrees celsius. INT16_MAX for unknown temperature.*/
	Voltages         [20]byte /* [10]int16 /*< [mV] Battery voltage of cells, in millivolts (1 = 1 millivolt). Cells above the valid cell count for this battery should have the UINT16_MAX value.*/
	CurrentBattery   [2]byte  /* int16     /*< [cA] Battery current, in 10*milliamperes (1 = 10 milliampere), -1: autopilot does not measure the current*/
	Id               byte     /* uint8     /*<  Battery ID*/
	BatteryFunction  byte     /* uint8     /*<  Function of the battery*/
	Typet            byte     /* uint8     /*<  Type (chemistry) of the battery*/
	BatteryRemaining byte     /* int8      /*< [%] Remaining battery energy: (0%: 0, 100%: 100), -1: autopilot does not estimate the remaining battery*/
	TimeRemaining    [4]byte  /* int32     /*< [s] Remaining battery time, in seconds (1 = 1s = 0% energy left), 0: autopilot does not provide remaining battery time estimate*/
	ChargeState      byte     /* uint8     /*<  State for extent of discharge, provided by autopilot for warning or external reactions*/
}
