package messages

import "unsafe"

const MSG_ID_SERVO_OUTPUT_RAW = 36

type ServoOutputRaw struct {
	TimeUsec   uint32 /*< [us] Timestamp (microseconds since system boot)*/
	Servo1Raw  uint16 /*< [us] Servo output 1 value, in microseconds*/
	Servo2Raw  uint16 /*< [us] Servo output 2 value, in microseconds*/
	Servo3Raw  uint16 /*< [us] Servo output 3 value, in microseconds*/
	Servo4Raw  uint16 /*< [us] Servo output 4 value, in microseconds*/
	Servo5Raw  uint16 /*< [us] Servo output 5 value, in microseconds*/
	Servo6Raw  uint16 /*< [us] Servo output 6 value, in microseconds*/
	Servo7Raw  uint16 /*< [us] Servo output 7 value, in microseconds*/
	Servo8Raw  uint16 /*< [us] Servo output 8 value, in microseconds*/
	Port       uint8  /*<  Servo output port (set of 8 outputs = 1 port). Most MAVs will just use one, but this allows to encode more than 8 Servos.*/
	Servo9Raw  uint16 /*< [us] Servo output 9 value, in microseconds*/
	Servo10Raw uint16 /*< [us] Servo output 10 value, in microseconds*/
	Servo11Raw uint16 /*< [us] Servo output 11 value, in microseconds*/
	Servo12Raw uint16 /*< [us] Servo output 12 value, in microseconds*/
	Servo13Raw uint16 /*< [us] Servo output 13 value, in microseconds*/
	Servo14Raw uint16 /*< [us] Servo output 14 value, in microseconds*/
	Servo15Raw uint16 /*< [us] Servo output 15 value, in microseconds*/
	Servo16Raw uint16 /*< [us] Servo output 16 value, in microseconds*/
}

func Byte2ServoOutputRaw(bts *[]byte) *ServoOutputRaw {
	sct := *(**ServoOutputRaw)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func ServoOutputRaw2Byte(sct *ServoOutputRaw) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
