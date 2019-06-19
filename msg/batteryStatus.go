package messages

import "unsafe"

const MSG_ID_BATTERY_STATUS = 147

type BatteryStatus struct {
	CurrentConsumed  int32     /*< [mAh] Consumed charge, in milliampere hours (1 = 1 mAh), -1: autopilot does not provide mAh consumption estimate*/
	EnergyConsumed   int32     /*< [hJ] Consumed energy, in HectoJoules (intergrated U*I*dt)  (1 = 100 Joule), -1: autopilot does not provide energy consumption estimate*/
	Temperature      int16     /*< [cdegC] Temperature of the battery in centi-degrees celsius. INT16_MAX for unknown temperature.*/
	Voltages         [10]int16 /*< [mV] Battery voltage of cells, in millivolts (1 = 1 millivolt). Cells above the valid cell count for this battery should have the UINT16_MAX value.*/
	CurrentBattery   int16     /*< [cA] Battery current, in 10*milliamperes (1 = 10 milliampere), -1: autopilot does not measure the current*/
	Id               uint8     /*<  Battery ID*/
	BatteryFunction  uint8     /*<  Function of the battery*/
	Typet            uint8     /*<  Type (chemistry) of the battery*/
	BatteryRemaining int8      /*< [%] Remaining battery energy: (0%: 0, 100%: 100), -1: autopilot does not estimate the remaining battery*/
	TimeRemaining    int32     /*< [s] Remaining battery time, in seconds (1 = 1s = 0% energy left), 0: autopilot does not provide remaining battery time estimate*/
	ChargeState      uint8     /*<  State for extent of discharge, provided by autopilot for warning or external reactions*/
}

func Byte2BatteryStatus(bts *[]byte) *BatteryStatus {
	sct := *(**BatteryStatus)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func BatteryStatus2Byte(sct *BatteryStatus) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
