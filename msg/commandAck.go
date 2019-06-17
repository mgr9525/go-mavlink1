package messages

import "unsafe"

const MSG_ID_COMMAND_ACK = 77

type CommandAck struct {
	Command uint16 /*<  Command ID, as defined by MAV_CMD enum.*/
	Result  uint8  /*<  See MAV_RESULT enum*/
}

func Byte2CommandAck(bts *[]byte) *CommandAck {
	sct := *(**CommandAck)(unsafe.Pointer(bts))
	return sct
}
func CommandAck2Byte(sct *CommandAck) *[]byte {
	Len := unsafe.Sizeof(*sct)
	mock := &SliceMock{
		addr: uintptr(unsafe.Pointer(sct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(mock))
	return &data
}
