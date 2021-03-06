package messages

const MSG_ID_SERVO_OUTPUT_RAW = 36

type ServoOutputRaw struct {
	TimeUsec   [4]byte /* uint32 /*< [us] Timestamp (microseconds since system boot)*/
	Servo1Raw  [2]byte /* uint16 /*< [us] Servo output 1 value, in microseconds*/
	Servo2Raw  [2]byte /* uint16 /*< [us] Servo output 2 value, in microseconds*/
	Servo3Raw  [2]byte /* uint16 /*< [us] Servo output 3 value, in microseconds*/
	Servo4Raw  [2]byte /* uint16 /*< [us] Servo output 4 value, in microseconds*/
	Servo5Raw  [2]byte /* uint16 /*< [us] Servo output 5 value, in microseconds*/
	Servo6Raw  [2]byte /* uint16 /*< [us] Servo output 6 value, in microseconds*/
	Servo7Raw  [2]byte /* uint16 /*< [us] Servo output 7 value, in microseconds*/
	Servo8Raw  [2]byte /* uint16 /*< [us] Servo output 8 value, in microseconds*/
	Port       byte    /* uint8  /*<  Servo output port (set of 8 outputs = 1 port). Most MAVs will just use one, but this allows to encode more than 8 Servos.*/
	Servo9Raw  [2]byte /* uint16 /*< [us] Servo output 9 value, in microseconds*/
	Servo10Raw [2]byte /* uint16 /*< [us] Servo output 10 value, in microseconds*/
	Servo11Raw [2]byte /* uint16 /*< [us] Servo output 11 value, in microseconds*/
	Servo12Raw [2]byte /* uint16 /*< [us] Servo output 12 value, in microseconds*/
	Servo13Raw [2]byte /* uint16 /*< [us] Servo output 13 value, in microseconds*/
	Servo14Raw [2]byte /* uint16 /*< [us] Servo output 14 value, in microseconds*/
	Servo15Raw [2]byte /* uint16 /*< [us] Servo output 15 value, in microseconds*/
	Servo16Raw [2]byte /* uint16 /*< [us] Servo output 16 value, in microseconds*/
}
