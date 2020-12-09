package messages

const MSG_ID_REQUEST_DATA_STREAM = 66

type RequestDataStream struct {
	ReqMessageRate  [2]byte /* uint16 /*< [Hz] The requested message rate*/
	TargetSystem    byte    /* uint8  /*<  The target requested to send the message stream.*/
	TargetComponent byte    /* uint8  /*<  The target requested to send the message stream.*/
	ReqStreamId     byte    /* uint8  /*<  The ID of the requested data stream*/
	StartStop       byte    /* uint8  /*<  1 to start sending, 0 to stop sending.*/
}
