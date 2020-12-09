package messages

const MSG_ID_COMMAND_ACK = 77

type CommandAck struct {
	Command [2]byte /* uint16 /*<  Command ID, as defined by MAV_CMD enum.*/
	Result  byte    /* uint8  /*<  See MAV_RESULT enum*/
}
