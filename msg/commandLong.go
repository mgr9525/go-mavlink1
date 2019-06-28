package messages

import "unsafe"

const MSG_ID_COMMAND_LONG = 76

const (
	MAV_CMD_NAV_TAKEOFF               = 22
	MAV_CMD_PREFLIGHT_REBOOT_SHUTDOWN = 246
	MAV_CMD_COMPONENT_ARM_DISARM      = 400
	MAV_CMD_DO_SET_SERVO              = 183
	MAV_CMD_DO_PAUSE_CONTINUE         = 193
	MAV_CMD_CONDITION_YAW             = 115
)

type CommandLong struct {
	Param1          float32 /*<  Parameter 1 (for the specific command).*/
	Param2          float32 /*<  Parameter 2 (for the specific command).*/
	Param3          float32 /*<  Parameter 3 (for the specific command).*/
	Param4          float32 /*<  Parameter 4 (for the specific command).*/
	Param5          float32 /*<  Parameter 5 (for the specific command).*/
	Param6          float32 /*<  Parameter 6 (for the specific command).*/
	Param7          float32 /*<  Parameter 7 (for the specific command).*/
	Command         uint16  /*<  Command ID (of command to send).*/
	TargetSystem    uint8   /*<  System which should execute the command*/
	TargetComponent uint8   /*<  Component which should execute the command, 0 for all components*/
	Confirmation    uint8   /*<  0: First transmission of this command. 1-255: Confirmation transmissions (e.g. for kill command)*/
}

func Byte2CommandLong(bts *[]byte) *CommandLong {
	sct := *(**CommandLong)(unsafe.Pointer(bts))
	ret := *sct
	return &ret
}
func CommandLong2Byte(sct *CommandLong) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
